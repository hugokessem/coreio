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
import {
  IpsAccountLookupDto,
  IpsTransferDto,
  IpsQrPaymentDto,
  IpsStatusDto,
} from './dto/ips.dto';
import {
  AccountNotFoundException,
  InsufficientFundsException,
  AccountRestrictedException,
} from '../../common/exceptions/domain.exceptions';
import { FinancialMath } from '../../common/utils/financial-math.util';

@Injectable()
export class IpsService {
  constructor(
    private readonly dataSource: DataSource,
    @InjectRepository(Account)
    private readonly accountRepo: Repository<Account>,
    @InjectRepository(Transaction)
    private readonly transactionRepo: Repository<Transaction>,
    @InjectRepository(ServiceEntity)
    private readonly serviceRepo: Repository<ServiceEntity>,
  ) {}

  async accountLookup(dto: IpsAccountLookupDto) {
    const account = await this.accountRepo.findOne({
      where: { accountNumber: dto.accountNumber },
      relations: ['customer', 'currency', 'branch'],
    });

    if (!account) {
      throw new AccountNotFoundException(dto.accountNumber);
    }

    return {
      accountNumber: account.accountNumber,
      accountName: account.accountName,
      currency: account.currencyCode,
      bankCode: dto.bankCode || 'CBETETAA',
      bankName: 'Commercial Bank of Ethiopia',
      branchName: account.branch?.branchName,
      status: account.status,
      isEligible: account.status === AccountStatus.ACTIVE && !account.isCreditBlocked,
    };
  }

  async executeTransfer(dto: IpsTransferDto) {
    if (dto.idempotencyKey) {
      const existing = await this.transactionRepo.findOne({
        where: { idempotencyKey: dto.idempotencyKey },
      });
      if (existing) {
        return {
          transactionReference: existing.transactionReference,
          status: existing.status,
          amount: existing.amount,
          fee: existing.feeAmount,
          currency: existing.currencyCode,
          destinationBankCode: dto.destinationBankCode,
          destinationAccountNumber: dto.destinationAccountNumber,
          beneficiaryName: dto.beneficiaryName,
          completedAt: existing.completedAt,
          message: 'Idempotent replay: previous IPS transfer returned',
        };
      }
    }

    return await this.dataSource.transaction(async (manager) => {
      const source = await manager.findOne(Account, {
        where: { accountNumber: dto.sourceAccountNumber },
      });
      if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

      if (source.status !== AccountStatus.ACTIVE || source.isDebitBlocked) {
        throw new AccountRestrictedException(source.accountNumber, 'Account not permitted for IPS transfer debit');
      }

      // Interbank IPS tariff (5 ETB standard)
      const service = await manager.findOne(ServiceEntity, { where: { serviceCode: 'FT_IPS' } });
      const fee = service ? service.feeValue : '5.0000';
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

      const txRef = `IPS26${Math.floor(10000000 + Math.random() * 90000000)}`;

      const tx = manager.create(Transaction, {
        transactionReference: txRef,
        idempotencyKey: dto.idempotencyKey || null,
        transactionType: TransactionType.IPS_TRANSFER,
        channel: TransactionChannel.IPS,
        sourceAccountId: source.id,
        externalAccount: dto.destinationAccountNumber,
        externalBankCode: dto.destinationBankCode,
        amount: dto.amount,
        feeAmount: fee,
        currencyCode: source.currencyCode,
        narration: dto.narration || `IPS to ${dto.destinationBankCode}: ${dto.beneficiaryName}`,
        status: TransactionStatus.COMPLETED,
        completedAt: new Date(),
      });
      const savedTx = await manager.save(Transaction, tx);

      // Debit source ledger entry
      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: source.id,
        entryType: EntryType.DEBIT,
        amount: totalDebit,
        currencyCode: source.currencyCode,
        balanceAfter: balance.ledgerBalance,
        narration: `IPS Interbank Transfer to ${dto.destinationAccountNumber} [${dto.destinationBankCode}]`,
      }));

      return {
        transactionReference: savedTx.transactionReference,
        status: savedTx.status,
        endToEndId: dto.endToEndId || `E2E-${savedTx.transactionReference}`,
        amount: dto.amount,
        fee,
        currency: source.currencyCode,
        sourceAccountNumber: source.accountNumber,
        destinationBankCode: dto.destinationBankCode,
        destinationAccountNumber: dto.destinationAccountNumber,
        beneficiaryName: dto.beneficiaryName,
        remainingAvailableBalance: balance.availableBalance,
        completedAt: savedTx.completedAt,
        simulation: true,
        externalSettlementConfirmed: false,
        message: 'IPS debit posted to internal ledger; EthSwitch interbank clearing is simulated in local sandbox',
      };
    });
  }

  async qrPayment(dto: IpsQrPaymentDto) {
    return await this.dataSource.transaction(async (manager) => {
      const source = await manager.findOne(Account, {
        where: { accountNumber: dto.sourceAccountNumber },
      });
      if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

      const merchant = await manager.findOne(Account, {
        where: { accountNumber: dto.merchantAccount },
      });
      if (!merchant) throw new AccountNotFoundException(dto.merchantAccount);

      const lockIds = [source.id, merchant.id].sort();
      const balances = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId IN (:...ids)', { ids: lockIds })
        .getMany();

      const sourceBal = balances.find((b) => b.accountId === source.id);
      const merchantBal = balances.find((b) => b.accountId === merchant.id);

      if (!sourceBal || !merchantBal) throw new BadRequestException('Failed to lock accounts');

      if (FinancialMath.isLessThan(sourceBal.availableBalance, dto.amount)) {
        throw new InsufficientFundsException(sourceBal.availableBalance, dto.amount);
      }

      sourceBal.ledgerBalance = FinancialMath.subtract(sourceBal.ledgerBalance, dto.amount);
      sourceBal.availableBalance = FinancialMath.subtract(sourceBal.availableBalance, dto.amount);
      await manager.save(AccountBalance, sourceBal);

      merchantBal.ledgerBalance = FinancialMath.add(merchantBal.ledgerBalance, dto.amount);
      merchantBal.availableBalance = FinancialMath.add(merchantBal.availableBalance, dto.amount);
      await manager.save(AccountBalance, merchantBal);

      const txRef = `QR26${Math.floor(10000000 + Math.random() * 90000000)}`;

      const tx = manager.create(Transaction, {
        transactionReference: txRef,
        idempotencyKey: dto.idempotencyKey || null,
        transactionType: TransactionType.IPS_QR_PAYMENT,
        channel: TransactionChannel.IPS,
        sourceAccountId: source.id,
        destinationAccountId: merchant.id,
        amount: dto.amount,
        feeAmount: '0.0000',
        currencyCode: source.currencyCode,
        narration: dto.narration || `IPS QR Merchant Payment: ${dto.merchantName}`,
        status: TransactionStatus.COMPLETED,
        completedAt: new Date(),
      });
      const savedTx = await manager.save(Transaction, tx);

      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: source.id,
        entryType: EntryType.DEBIT,
        amount: dto.amount,
        currencyCode: source.currencyCode,
        balanceAfter: sourceBal.ledgerBalance,
        narration: `QR Pay -> ${dto.merchantName}`,
      }));

      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: merchant.id,
        entryType: EntryType.CREDIT,
        amount: dto.amount,
        currencyCode: merchant.currencyCode,
        balanceAfter: merchantBal.ledgerBalance,
        narration: `QR Receive <- ${source.accountNumber}`,
      }));

      return {
        transactionReference: savedTx.transactionReference,
        status: savedTx.status,
        merchantName: dto.merchantName,
        merchantAccount: merchant.accountNumber,
        amount: dto.amount,
        currency: source.currencyCode,
        remainingBalance: sourceBal.availableBalance,
        completedAt: savedTx.completedAt,
        message: 'IPS QR payment successful',
      };
    });
  }

  async getStatus(dto: IpsStatusDto) {
    const tx = await this.transactionRepo.findOne({
      where: { transactionReference: dto.transactionReference },
    });
    if (!tx) throw new BadRequestException(`IPS transaction '${dto.transactionReference}' not found`);

    return {
      transactionReference: tx.transactionReference,
      status: tx.status,
      amount: tx.amount,
      fee: tx.feeAmount,
      currency: tx.currencyCode,
      destinationAccount: tx.externalAccount,
      destinationBank: tx.externalBankCode,
      completedAt: tx.completedAt,
    };
  }
}
