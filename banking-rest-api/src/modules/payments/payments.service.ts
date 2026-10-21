import { Injectable, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository } from 'typeorm';
import { Account, AccountStatus } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import {
  Transaction,
  TransactionType,
  TransactionStatus,
  TransactionChannel,
} from '../../database/entities/transaction.entity';
import { LedgerEntry, EntryType } from '../../database/entities/ledger-entry.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import { BillPaymentDto, BillPaymentVerifyDto, SplitPaymentDto, RtgsPaymentDto } from './dto/payment.dto';
import {
  AccountNotFoundException,
  InsufficientFundsException,
  AccountRestrictedException,
} from '../../common/exceptions/domain.exceptions';
import { FinancialMath } from '../../common/utils/financial-math.util';

@Injectable()
export class PaymentsService {
  constructor(
    private readonly dataSource: DataSource,
    @InjectRepository(Account)
    private readonly accountRepo: Repository<Account>,
    @InjectRepository(Transaction)
    private readonly transactionRepo: Repository<Transaction>,
    @InjectRepository(ServiceEntity)
    private readonly serviceRepo: Repository<ServiceEntity>,
  ) {}

  async verifyBill(dto: BillPaymentVerifyDto) {
    const source = await this.accountRepo.findOne({
      where: { accountNumber: dto.sourceAccountNumber },
      relations: ['balance', 'currency'],
    });
    if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

    if (source.status !== AccountStatus.ACTIVE) {
      throw new AccountRestrictedException(source.accountNumber, `Source account status is ${source.status}`);
    }
    if (source.isDebitBlocked) {
      throw new AccountRestrictedException(source.accountNumber, 'Account is blocked for debit operations');
    }

    const biller = await this.serviceRepo.findOne({
      where: { serviceCode: dto.billerCode },
    });
    const fee = biller ? biller.feeValue : '0.0000';
    const totalRequired = FinancialMath.add(dto.amount, fee);

    const available = source.balance?.availableBalance || '0.0000';
    const hasSufficientFunds = FinancialMath.isGreaterThanOrEqualTo(available, totalRequired);

    return {
      canExecute: hasSufficientFunds,
      sourceAccountNumber: source.accountNumber,
      sourceAccountName: source.accountName,
      billerCode: dto.billerCode,
      billerName: biller?.serviceName || dto.billerCode,
      billReference: dto.billReference,
      amount: dto.amount,
      fee,
      totalDebit: totalRequired,
      currency: source.currencyCode,
      sourceAvailableBalance: available,
      sufficientFunds: hasSufficientFunds,
      status: 'VERIFIED',
      reason: hasSufficientFunds ? 'Bill verified successfully' : 'Insufficient funds in source account',
    };
  }

  async payBill(dto: BillPaymentDto) {
    if (dto.idempotencyKey) {
      const existing = await this.transactionRepo.findOne({
        where: { idempotencyKey: dto.idempotencyKey },
        relations: ['sourceAccount'],
      });
      if (existing) {
        return {
          transactionReference: existing.transactionReference,
          status: existing.status,
          billerCode: dto.billerCode,
          billReference: dto.billReference,
          amount: existing.amount,
          narration: existing.narration,
          completedAt: existing.completedAt,
          message: 'Idempotent replay: previous bill payment returned',
        };
      }
    }

    return await this.dataSource.transaction(async (manager) => {
      const source = await manager.findOne(Account, {
        where: { accountNumber: dto.sourceAccountNumber },
      });
      if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

      if (source.status !== AccountStatus.ACTIVE || source.isDebitBlocked) {
        throw new AccountRestrictedException(source.accountNumber, 'Account not eligible for bill payment debit');
      }

      const balance = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId = :accountId', { accountId: source.id })
        .getOne();

      if (!balance) throw new BadRequestException('Balance record not found');

      if (FinancialMath.isLessThan(balance.availableBalance, dto.amount)) {
        throw new InsufficientFundsException(balance.availableBalance, dto.amount);
      }

      // Deduct balance
      balance.ledgerBalance = FinancialMath.subtract(balance.ledgerBalance, dto.amount);
      balance.availableBalance = FinancialMath.subtract(balance.availableBalance, dto.amount);
      await manager.save(AccountBalance, balance);

      const randomSuffix = Math.floor(10000000 + Math.random() * 90000000).toString();
      const txRef = `BP26${randomSuffix}`;

      const tx = manager.create(Transaction, {
        transactionReference: txRef,
        idempotencyKey: dto.idempotencyKey || null,
        transactionType: TransactionType.BILL_PAYMENT,
        channel: TransactionChannel.MOBILE,
        sourceAccountId: source.id,
        externalAccount: dto.billReference,
        amount: dto.amount,
        feeAmount: '0.0000',
        currencyCode: source.currencyCode,
        narration: dto.narration || `Bill Payment: ${dto.billerCode} - ${dto.billReference}`,
        status: TransactionStatus.COMPLETED,
        completedAt: new Date(),
      });
      const savedTx = await manager.save(Transaction, tx);

      // Debit ledger entry
      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: source.id,
        entryType: EntryType.DEBIT,
        amount: dto.amount,
        currencyCode: source.currencyCode,
        balanceAfter: balance.ledgerBalance,
        narration: `Bill Payment ${dto.billerCode} [${dto.billReference}]`,
      }));

      return {
        transactionReference: savedTx.transactionReference,
        status: savedTx.status,
        billerCode: dto.billerCode,
        billReference: dto.billReference,
        amount: savedTx.amount,
        currency: savedTx.currencyCode,
        remainingAvailableBalance: balance.availableBalance,
        completedAt: savedTx.completedAt,
        message: 'Bill payment processed successfully',
      };
    });
  }

  async splitPayment(dto: SplitPaymentDto) {
    if (!dto.splits || dto.splits.length === 0) {
      throw new BadRequestException('At least one split recipient must be specified');
    }

    const totalAmount = dto.splits.reduce((acc, s) => FinancialMath.add(acc, s.amount), '0.0000');

    return await this.dataSource.transaction(async (manager) => {
      const source = await manager.findOne(Account, {
        where: { accountNumber: dto.sourceAccountNumber },
      });
      if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

      // Find all destination accounts
      const destAccounts: Account[] = [];
      for (const split of dto.splits) {
        const dest = await manager.findOne(Account, {
          where: { accountNumber: split.destinationAccountNumber },
        });
        if (!dest) throw new AccountNotFoundException(split.destinationAccountNumber);
        destAccounts.push(dest);
      }

      // Lock all accounts deterministically
      const allAccountIds = [source.id, ...destAccounts.map((d) => d.id)].sort();
      const balances = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId IN (:...ids)', { ids: allAccountIds })
        .getMany();

      const sourceBal = balances.find((b) => b.accountId === source.id);
      if (!sourceBal) throw new BadRequestException('Source balance not found');

      if (FinancialMath.isLessThan(sourceBal.availableBalance, totalAmount)) {
        throw new InsufficientFundsException(sourceBal.availableBalance, totalAmount);
      }

      // Deduct total from source
      sourceBal.ledgerBalance = FinancialMath.subtract(sourceBal.ledgerBalance, totalAmount);
      sourceBal.availableBalance = FinancialMath.subtract(sourceBal.availableBalance, totalAmount);
      await manager.save(AccountBalance, sourceBal);

      const splitResults: any[] = [];
      const batchRef = `SP26${Math.floor(10000000 + Math.random() * 90000000)}`;

      for (let i = 0; i < dto.splits.length; i++) {
        const split = dto.splits[i];
        const destAcc = destAccounts[i];
        const destBal = balances.find((b) => b.accountId === destAcc.id);
        if (!destBal) continue;

        destBal.ledgerBalance = FinancialMath.add(destBal.ledgerBalance, split.amount);
        destBal.availableBalance = FinancialMath.add(destBal.availableBalance, split.amount);
        await manager.save(AccountBalance, destBal);

        const subTxRef = `${batchRef}-${(i + 1).toString().padStart(2, '0')}`;
        const tx = manager.create(Transaction, {
          transactionReference: subTxRef,
          transactionType: TransactionType.SPLIT_PAYMENT,
          channel: TransactionChannel.CORE,
          sourceAccountId: source.id,
          destinationAccountId: destAcc.id,
          amount: split.amount,
          feeAmount: '0.0000',
          currencyCode: source.currencyCode,
          narration: split.narration || dto.batchNarration || 'Split payment distribution',
          status: TransactionStatus.COMPLETED,
          completedAt: new Date(),
        });
        const savedTx = await manager.save(Transaction, tx);

        // Ledger postings
        await manager.save(LedgerEntry, manager.create(LedgerEntry, {
          transactionId: savedTx.id,
          accountId: source.id,
          entryType: EntryType.DEBIT,
          amount: split.amount,
          currencyCode: source.currencyCode,
          balanceAfter: sourceBal.ledgerBalance,
          narration: `Split Batch Debit -> ${destAcc.accountNumber}`,
        }));

        await manager.save(LedgerEntry, manager.create(LedgerEntry, {
          transactionId: savedTx.id,
          accountId: destAcc.id,
          entryType: EntryType.CREDIT,
          amount: split.amount,
          currencyCode: destAcc.currencyCode,
          balanceAfter: destBal.ledgerBalance,
          narration: `Split Batch Credit <- ${source.accountNumber}`,
        }));

        splitResults.push({
          subReference: subTxRef,
          destinationAccount: destAcc.accountNumber,
          amount: split.amount,
          status: 'COMPLETED',
        });
      }

      return {
        batchReference: batchRef,
        sourceAccountNumber: source.accountNumber,
        totalDistributed: totalAmount,
        splitCount: splitResults.length,
        splits: splitResults,
        remainingAvailableBalance: sourceBal.availableBalance,
        completedAt: new Date().toISOString(),
        message: 'Split payment batch executed successfully',
      };
    });
  }

  async rtgsPayment(dto: RtgsPaymentDto) {
    return await this.dataSource.transaction(async (manager) => {
      const source = await manager.findOne(Account, {
        where: { accountNumber: dto.sourceAccountNumber },
      });
      if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

      const fee = '100.0000'; // Standard RTGS processing tariff
      const totalDebit = FinancialMath.add(dto.amount, fee);

      const balance = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId = :accountId', { accountId: source.id })
        .getOne();

      if (!balance) throw new BadRequestException('Balance record missing');

      if (FinancialMath.isLessThan(balance.availableBalance, totalDebit)) {
        throw new InsufficientFundsException(balance.availableBalance, totalDebit);
      }

      balance.ledgerBalance = FinancialMath.subtract(balance.ledgerBalance, totalDebit);
      balance.availableBalance = FinancialMath.subtract(balance.availableBalance, totalDebit);
      await manager.save(AccountBalance, balance);

      const txRef = `RTGS26${Math.floor(10000000 + Math.random() * 90000000)}`;

      const tx = manager.create(Transaction, {
        transactionReference: txRef,
        idempotencyKey: dto.idempotencyKey || null,
        transactionType: TransactionType.RTGS_TRANSFER,
        channel: TransactionChannel.CORE,
        sourceAccountId: source.id,
        externalAccount: dto.beneficiaryAccountNumber,
        externalBankCode: dto.beneficiaryBankSwift,
        amount: dto.amount,
        feeAmount: fee,
        currencyCode: source.currencyCode,
        narration: dto.narration || `RTGS to ${dto.beneficiaryBankSwift}: ${dto.beneficiaryName}`,
        status: TransactionStatus.COMPLETED,
        completedAt: new Date(),
      });
      const savedTx = await manager.save(Transaction, tx);

      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: source.id,
        entryType: EntryType.DEBIT,
        amount: totalDebit,
        currencyCode: source.currencyCode,
        balanceAfter: balance.ledgerBalance,
        narration: `RTGS to ${dto.beneficiaryBankSwift} [${dto.beneficiaryAccountNumber}]`,
      }));

      return {
        transactionReference: savedTx.transactionReference,
        status: savedTx.status,
        beneficiaryBankSwift: dto.beneficiaryBankSwift,
        beneficiaryAccountNumber: dto.beneficiaryAccountNumber,
        beneficiaryName: dto.beneficiaryName,
        amount: dto.amount,
        fee,
        currency: source.currencyCode,
        remainingBalance: balance.availableBalance,
        completedAt: savedTx.completedAt,
        simulation: true,
        externalSettlementConfirmed: false,
        message: 'RTGS debit posted to internal ledger; interbank clearing is simulated in local sandbox',
      };
    });
  }
}
