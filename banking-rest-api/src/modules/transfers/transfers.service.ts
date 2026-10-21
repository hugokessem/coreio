import { Injectable, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository } from 'typeorm';
import { Account, AccountStatus } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { AccountHold, HoldStatus } from '../../database/entities/account-hold.entity';
import {
  Transaction,
  TransactionType,
  TransactionStatus,
  TransactionChannel,
} from '../../database/entities/transaction.entity';
import { LedgerEntry, EntryType } from '../../database/entities/ledger-entry.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import {
  TransferVerifyDto,
  ExecuteTransferDto,
  TransferFromHoldDto,
  TransactionStatusDto,
  ReversalDto,
} from './dto/transfer.dto';
import {
  AccountNotFoundException,
  InsufficientFundsException,
  AccountRestrictedException,
  CurrencyMismatchException,
  HoldNotFoundException,
} from '../../common/exceptions/domain.exceptions';
import { FinancialMath } from '../../common/utils/financial-math.util';
import { IdempotencyService } from '../../common/services/idempotency.service';

@Injectable()
export class TransfersService {
  constructor(
    private readonly dataSource: DataSource,
    @InjectRepository(Account)
    private readonly accountRepository: Repository<Account>,
    @InjectRepository(Transaction)
    private readonly transactionRepository: Repository<Transaction>,
    @InjectRepository(ServiceEntity)
    private readonly serviceRepository: Repository<ServiceEntity>,
    private readonly idempotencyService: IdempotencyService,
  ) {}

  async verify(dto: TransferVerifyDto) {
    if (dto.sourceAccountNumber === dto.destinationAccountNumber) {
      throw new BadRequestException('Source and destination accounts must be different');
    }

    const source = await this.accountRepository.findOne({
      where: { accountNumber: dto.sourceAccountNumber },
      relations: ['balance', 'currency', 'customer'],
    });
    if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

    const destination = await this.accountRepository.findOne({
      where: { accountNumber: dto.destinationAccountNumber },
      relations: ['currency', 'customer', 'branch'],
    });
    if (!destination) throw new AccountNotFoundException(dto.destinationAccountNumber);

    if (source.currencyCode !== destination.currencyCode) {
      throw new CurrencyMismatchException(source.currencyCode, destination.currencyCode);
    }

    if (source.status !== AccountStatus.ACTIVE) {
      throw new AccountRestrictedException(source.accountNumber, `Source account status is ${source.status}`);
    }
    if (destination.status !== AccountStatus.ACTIVE) {
      throw new AccountRestrictedException(destination.accountNumber, `Destination account status is ${destination.status}`);
    }
    if (source.isDebitBlocked) {
      throw new AccountRestrictedException(source.accountNumber, 'Account is blocked for debit operations');
    }
    if (destination.isCreditBlocked) {
      throw new AccountRestrictedException(destination.accountNumber, 'Account is blocked for credit operations');
    }

    // Determine fee
    const service = await this.serviceRepository.findOne({ where: { serviceCode: 'FT_INT' } });
    const fee = service ? service.feeValue : '0.0000';
    const totalRequired = FinancialMath.add(dto.amount, fee);

    const available = source.balance?.availableBalance || '0.0000';
    const hasSufficientFunds = FinancialMath.isGreaterThanOrEqualTo(available, totalRequired);

    return {
      canExecute: hasSufficientFunds,
      sourceAccountNumber: source.accountNumber,
      sourceAccountName: source.accountName,
      destinationAccountNumber: destination.accountNumber,
      destinationAccountName: destination.accountName,
      destinationBranch: destination.branch?.branchName,
      amount: dto.amount,
      fee,
      totalDebit: totalRequired,
      currency: source.currencyCode,
      sourceAvailableBalance: available,
      sufficientFunds: hasSufficientFunds,
      reason: hasSufficientFunds ? 'Verification successful' : 'Insufficient funds in source account',
    };
  }

  async execute(dto: ExecuteTransferDto) {
    if (dto.sourceAccountNumber === dto.destinationAccountNumber) {
      throw new BadRequestException('Cannot transfer funds to the same account');
    }

    if (!FinancialMath.isPositive(dto.amount)) {
      throw new BadRequestException('Transfer amount must be strictly positive');
    }

    // Check Idempotency Key & Request Fingerprint Hash
    if (dto.idempotencyKey) {
      const check = await this.idempotencyService.checkOrRegister(
        dto.idempotencyKey,
        'FUND_TRANSFER',
        {
          source: dto.sourceAccountNumber,
          dest: dto.destinationAccountNumber,
          amount: dto.amount,
        },
      );
      if (check.isReplay && check.cachedResponse) {
        return check.cachedResponse;
      }
    }

    try {
      return await this.dataSource.transaction(async (manager) => {
        const source = await manager.findOne(Account, {
          where: { accountNumber: dto.sourceAccountNumber },
          relations: ['customer', 'currency'],
        });
        if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

        const destination = await manager.findOne(Account, {
          where: { accountNumber: dto.destinationAccountNumber },
          relations: ['customer', 'currency'],
        });
        if (!destination) throw new AccountNotFoundException(dto.destinationAccountNumber);

        if (source.currencyCode !== destination.currencyCode) {
          throw new CurrencyMismatchException(source.currencyCode, destination.currencyCode);
        }

        if (source.status !== AccountStatus.ACTIVE) {
          throw new AccountRestrictedException(source.accountNumber, `Source account status is ${source.status}`);
        }
        if (destination.status !== AccountStatus.ACTIVE) {
          throw new AccountRestrictedException(destination.accountNumber, `Destination account status is ${destination.status}`);
        }
        if (source.isDebitBlocked) {
          throw new AccountRestrictedException(source.accountNumber, 'Account is blocked for debit operations');
        }
        if (destination.isCreditBlocked) {
          throw new AccountRestrictedException(destination.accountNumber, 'Account is blocked for credit operations');
        }

        // Determine fee
        const service = await manager.findOne(ServiceEntity, { where: { serviceCode: 'FT_INT' } });
        const fee = service ? service.feeValue : '0.0000';
        const totalDebit = FinancialMath.add(dto.amount, fee);

        // Deadlock prevention: sort account IDs deterministically before acquiring locks
        const lockIds = [source.id, destination.id].sort();

        const balances = await manager
          .createQueryBuilder(AccountBalance, 'bal')
          .setLock('pessimistic_write')
          .where('bal.accountId IN (:...ids)', { ids: lockIds })
          .getMany();

        const sourceBalance = balances.find((b) => b.accountId === source.id);
        const destBalance = balances.find((b) => b.accountId === destination.id);

        if (!sourceBalance || !destBalance) {
          throw new BadRequestException('Failed to lock account balances');
        }

        if (FinancialMath.isLessThan(sourceBalance.availableBalance, totalDebit)) {
          throw new InsufficientFundsException(sourceBalance.availableBalance, totalDebit);
        }

        // Generate Reference FT26 + 8 random digits
        const randomSuffix = Math.floor(10000000 + Math.random() * 90000000).toString();
        const txRef = `FT26${randomSuffix}`;

        // Update source balance
        sourceBalance.ledgerBalance = FinancialMath.subtract(sourceBalance.ledgerBalance, totalDebit);
        sourceBalance.availableBalance = FinancialMath.subtract(sourceBalance.availableBalance, totalDebit);
        await manager.save(AccountBalance, sourceBalance);

        // Update destination balance
        destBalance.ledgerBalance = FinancialMath.add(destBalance.ledgerBalance, dto.amount);
        destBalance.availableBalance = FinancialMath.add(destBalance.availableBalance, dto.amount);
        await manager.save(AccountBalance, destBalance);

        // Create Transaction Record
        const tx = manager.create(Transaction, {
          transactionReference: txRef,
          idempotencyKey: dto.idempotencyKey || null,
          transactionType: TransactionType.INTERNAL_TRANSFER,
          channel: dto.channel || TransactionChannel.MOBILE,
          sourceAccountId: source.id,
          destinationAccountId: destination.id,
          amount: dto.amount,
          feeAmount: fee,
          currencyCode: source.currencyCode,
          exchangeRate: '1.000000',
          narration: dto.narration,
          status: TransactionStatus.COMPLETED,
          completedAt: new Date(),
        });
        const savedTx = await manager.save(Transaction, tx);

        // Post Debit Ledger Entry for source
        const debitEntry = manager.create(LedgerEntry, {
          transactionId: savedTx.id,
          accountId: source.id,
          entryType: EntryType.DEBIT,
          amount: totalDebit,
          currencyCode: source.currencyCode,
          balanceAfter: sourceBalance.ledgerBalance,
          narration: `TRF to ${destination.accountNumber} - ${dto.narration}`,
        });
        await manager.save(LedgerEntry, debitEntry);

        // Post Credit Ledger Entry for destination
        const creditEntry = manager.create(LedgerEntry, {
          transactionId: savedTx.id,
          accountId: destination.id,
          entryType: EntryType.CREDIT,
          amount: dto.amount,
          currencyCode: destination.currencyCode,
          balanceAfter: destBalance.ledgerBalance,
          narration: `TRF from ${source.accountNumber} - ${dto.narration}`,
        });
        await manager.save(LedgerEntry, creditEntry);

        savedTx.sourceAccount = source;
        savedTx.destinationAccount = destination;

        const response = this.mapToTransactionResponse(savedTx, 'Transfer completed successfully');

        // Record idempotency receipt
        await this.idempotencyService.recordSuccess(
          dto.idempotencyKey,
          'FUND_TRANSFER',
          {
            source: dto.sourceAccountNumber,
            dest: dto.destinationAccountNumber,
            amount: dto.amount,
          },
          savedTx.transactionReference,
          response,
          200,
          undefined,
          manager,
        );

        return response;
      });
    } catch (err: any) {
      if (dto.idempotencyKey && (err.code === '23505' || err.message?.includes('duplicate key') || err.message?.includes('unique constraint'))) {
        // Yield briefly for concurrent winner to commit
        await new Promise((resolve) => setTimeout(resolve, 100));
        const committed = await this.transactionRepository.findOne({
          where: { idempotencyKey: dto.idempotencyKey },
          relations: ['sourceAccount', 'destinationAccount'],
        });
        if (committed) {
          return this.mapToTransactionResponse(committed, 'Idempotent replay: previous successful transfer returned');
        }
      }
      throw err;
    }
  }

  async transferFromHold(dto: TransferFromHoldDto) {
    if (dto.idempotencyKey) {
      const existing = await this.transactionRepository.findOne({
        where: { idempotencyKey: dto.idempotencyKey },
        relations: ['sourceAccount', 'destinationAccount'],
      });
      if (existing) {
        return this.mapToTransactionResponse(existing, 'Idempotent replay: previous settlement returned');
      }
    }

    return await this.dataSource.transaction(async (manager) => {
      const hold = await manager.findOne(AccountHold, {
        where: { holdReference: dto.holdReference, status: HoldStatus.ACTIVE },
        relations: ['account'],
      });
      if (!hold) throw new HoldNotFoundException(dto.holdReference);

      const source = hold.account;
      const destination = await manager.findOne(Account, {
        where: { accountNumber: dto.destinationAccountNumber },
        relations: ['currency'],
      });
      if (!destination) throw new AccountNotFoundException(dto.destinationAccountNumber);

      if (source.id === destination.id) {
        throw new BadRequestException('Cannot settle hold to the same account');
      }

      const lockIds = [source.id, destination.id].sort();
      const balances = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId IN (:...ids)', { ids: lockIds })
        .getMany();

      const sourceBalance = balances.find((b) => b.accountId === source.id);
      const destBalance = balances.find((b) => b.accountId === destination.id);

      if (!sourceBalance || !destBalance) {
        throw new BadRequestException('Failed to lock account balances');
      }

      // Settle hold: available was already deducted when hold was placed.
      // Now deduct from ledgerBalance and lockedBalance.
      sourceBalance.ledgerBalance = FinancialMath.subtract(sourceBalance.ledgerBalance, hold.amount);
      sourceBalance.lockedBalance = FinancialMath.subtract(sourceBalance.lockedBalance, hold.amount);
      await manager.save(AccountBalance, sourceBalance);

      // Credit destination available and ledger
      destBalance.ledgerBalance = FinancialMath.add(destBalance.ledgerBalance, hold.amount);
      destBalance.availableBalance = FinancialMath.add(destBalance.availableBalance, hold.amount);
      await manager.save(AccountBalance, destBalance);

      hold.status = HoldStatus.SETTLED;
      hold.releasedAt = new Date();
      hold.releasedReason = `Settled to ${destination.accountNumber}`;
      await manager.save(AccountHold, hold);

      const randomSuffix = Math.floor(10000000 + Math.random() * 90000000).toString();
      const txRef = `FT26${randomSuffix}`;

      const tx = manager.create(Transaction, {
        transactionReference: txRef,
        idempotencyKey: dto.idempotencyKey || null,
        transactionType: TransactionType.TRANSFER_FROM_HOLD,
        channel: TransactionChannel.CORE,
        sourceAccountId: source.id,
        destinationAccountId: destination.id,
        amount: hold.amount,
        feeAmount: '0.0000',
        currencyCode: hold.currencyCode,
        narration: dto.narration || `Hold settlement ${hold.holdReference}`,
        status: TransactionStatus.COMPLETED,
        completedAt: new Date(),
      });
      const savedTx = await manager.save(Transaction, tx);

      // Ledger entries
      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: source.id,
        entryType: EntryType.DEBIT,
        amount: hold.amount,
        currencyCode: hold.currencyCode,
        balanceAfter: sourceBalance.ledgerBalance,
        narration: `Hold Settle to ${destination.accountNumber} [${hold.holdReference}]`,
      }));

      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: destination.id,
        entryType: EntryType.CREDIT,
        amount: hold.amount,
        currencyCode: hold.currencyCode,
        balanceAfter: destBalance.ledgerBalance,
        narration: `Hold Settle from ${source.accountNumber} [${hold.holdReference}]`,
      }));

      savedTx.sourceAccount = source;
      savedTx.destinationAccount = destination;

      return this.mapToTransactionResponse(savedTx, 'Hold settled and transferred successfully');
    });
  }

  async getStatus(dto: TransactionStatusDto) {
    return this.getByReference(dto.transactionReference);
  }

  async getByReference(transactionReference: string) {
    const tx = await this.transactionRepository.findOne({
      where: { transactionReference },
      relations: ['sourceAccount', 'destinationAccount', 'ledgerEntries'],
    });

    if (!tx) {
      throw new BadRequestException(`Transaction with reference '${transactionReference}' not found`);
    }

    return this.mapToTransactionResponse(tx);
  }

  async reverse(dto: ReversalDto) {
    return await this.dataSource.transaction(async (manager) => {
      const origTx = await manager.findOne(Transaction, {
        where: { transactionReference: dto.transactionReference },
        relations: ['sourceAccount', 'destinationAccount'],
      });

      if (!origTx) {
        throw new BadRequestException(`Transaction '${dto.transactionReference}' not found`);
      }

      if (origTx.status !== TransactionStatus.COMPLETED) {
        throw new BadRequestException(`Cannot reverse transaction with status '${origTx.status}'`);
      }

      if (origTx.reversalOfTransactionId) {
        throw new BadRequestException('This transaction is already a reversal');
      }

      const existingReversal = await manager.findOne(Transaction, {
        where: { reversalOfTransactionId: origTx.id },
      });
      if (existingReversal) {
        throw new BadRequestException(`Transaction has already been reversed by reference ${existingReversal.transactionReference}`);
      }

      if (!origTx.sourceAccountId || !origTx.destinationAccountId) {
        throw new BadRequestException('Cannot reverse single-legged cash transaction automatically');
      }

      const lockIds = [origTx.sourceAccountId, origTx.destinationAccountId].sort();
      const balances = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId IN (:...ids)', { ids: lockIds })
        .getMany();

      const sourceBal = balances.find((b) => b.accountId === origTx.sourceAccountId);
      const destBal = balances.find((b) => b.accountId === origTx.destinationAccountId);

      if (!sourceBal || !destBal) {
        throw new BadRequestException('Failed to lock accounts for reversal');
      }

      // Check if original destination still has sufficient funds to refund
      if (FinancialMath.isLessThan(destBal.availableBalance, origTx.amount)) {
        throw new InsufficientFundsException(
          destBal.availableBalance,
          origTx.amount,
        );
      }

      const totalToRefund = FinancialMath.add(origTx.amount, origTx.feeAmount);

      // Reverse balances: debit destination, credit original source
      destBal.ledgerBalance = FinancialMath.subtract(destBal.ledgerBalance, origTx.amount);
      destBal.availableBalance = FinancialMath.subtract(destBal.availableBalance, origTx.amount);
      await manager.save(AccountBalance, destBal);

      sourceBal.ledgerBalance = FinancialMath.add(sourceBal.ledgerBalance, totalToRefund);
      sourceBal.availableBalance = FinancialMath.add(sourceBal.availableBalance, totalToRefund);
      await manager.save(AccountBalance, sourceBal);

      // Mark original transaction as reversed
      origTx.status = TransactionStatus.REVERSED;
      await manager.save(Transaction, origTx);

      // Create Reversal Transaction
      const revRef = `REV-${origTx.transactionReference}`;
      const revTx = manager.create(Transaction, {
        transactionReference: revRef,
        transactionType: TransactionType.REVERSAL,
        channel: TransactionChannel.SYSTEM,
        sourceAccountId: origTx.destinationAccountId,
        destinationAccountId: origTx.sourceAccountId,
        amount: origTx.amount,
        feeAmount: origTx.feeAmount,
        currencyCode: origTx.currencyCode,
        narration: `Reversal of ${origTx.transactionReference}: ${dto.reason}`,
        status: TransactionStatus.COMPLETED,
        reversalOfTransactionId: origTx.id,
        completedAt: new Date(),
      });
      const savedRevTx = await manager.save(Transaction, revTx);

      // Double-entry postings for reversal
      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedRevTx.id,
        accountId: origTx.destinationAccountId,
        entryType: EntryType.DEBIT,
        amount: origTx.amount,
        currencyCode: origTx.currencyCode,
        balanceAfter: destBal.ledgerBalance,
        narration: `Reversal Debit: ${dto.reason}`,
      }));

      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedRevTx.id,
        accountId: origTx.sourceAccountId,
        entryType: EntryType.CREDIT,
        amount: totalToRefund,
        currencyCode: origTx.currencyCode,
        balanceAfter: sourceBal.ledgerBalance,
        narration: `Reversal Credit Refund: ${dto.reason}`,
      }));

      return {
        reversalReference: savedRevTx.transactionReference,
        originalReference: origTx.transactionReference,
        amount: origTx.amount,
        feeRefunded: origTx.feeAmount,
        status: savedRevTx.status,
        reason: dto.reason,
        completedAt: savedRevTx.completedAt,
      };
    });
  }

  private mapToTransactionResponse(tx: Transaction, message?: string) {
    return {
      transactionReference: tx.transactionReference,
      status: tx.status,
      transactionType: tx.transactionType,
      channel: tx.channel,
      sourceAccountNumber: tx.sourceAccount?.accountNumber || null,
      destinationAccountNumber: tx.destinationAccount?.accountNumber || null,
      amount: tx.amount,
      fee: tx.feeAmount,
      currency: tx.currencyCode,
      narration: tx.narration,
      idempotencyKey: tx.idempotencyKey,
      initiatedAt: tx.initiatedAt,
      completedAt: tx.completedAt,
      message: message || (tx.status === TransactionStatus.COMPLETED ? 'Transaction completed successfully' : 'Transaction failed'),
    };
  }
}
