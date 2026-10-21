import { Injectable, BadRequestException, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository } from 'typeorm';
import { Account, AccountStatus } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { Customer } from '../../database/entities/customer.entity';
import {
  Transaction,
  TransactionType,
  TransactionStatus,
  TransactionChannel,
} from '../../database/entities/transaction.entity';
import { LedgerEntry, EntryType } from '../../database/entities/ledger-entry.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import {
  WalletCustomerLookupDto,
  AccountToWalletTransferDto,
  WalletToAccountTransferDto,
  WalletAgentLookupDto,
  WalletAgentTransferDto,
} from './dto/wallet.dto';
import {
  AccountNotFoundException,
  CustomerNotFoundException,
  InsufficientFundsException,
  AccountRestrictedException,
} from '../../common/exceptions/domain.exceptions';
import { FinancialMath } from '../../common/utils/financial-math.util';

@Injectable()
export class WalletService {
  constructor(
    private readonly dataSource: DataSource,
    @InjectRepository(Account)
    private readonly accountRepo: Repository<Account>,
    @InjectRepository(Customer)
    private readonly customerRepo: Repository<Customer>,
    @InjectRepository(Transaction)
    private readonly transactionRepo: Repository<Transaction>,
    @InjectRepository(ServiceEntity)
    private readonly serviceRepo: Repository<ServiceEntity>,
  ) {}

  async customerLookup(dto: WalletCustomerLookupDto) {
    const customer = await this.customerRepo.findOne({
      where: { phoneNumber: dto.phoneNumber },
      relations: ['accounts'],
    });

    if (!customer) {
      throw new CustomerNotFoundException(dto.phoneNumber);
    }

    return {
      phoneNumber: customer.phoneNumber,
      customerNumber: customer.customerNumber,
      customerName: customer.fullName,
      status: customer.status,
      linkedAccounts: customer.accounts ? customer.accounts.map((a) => a.accountNumber) : [],
      walletTier: 'LEVEL_3_FULL_KYC',
      isAllowedToTransact: customer.status === 'ACTIVE',
    };
  }

  async agentLookup(dto: WalletAgentLookupDto) {
    const customer = await this.customerRepo.findOne({
      where: [
        { phoneNumber: dto.agentIdentifier },
        { customerNumber: dto.agentIdentifier },
      ],
      relations: ['accounts'],
    });

    if (!customer) {
      throw new NotFoundException(`CBE Birr Agent '${dto.agentIdentifier}' not found`);
    }

    return {
      agentIdentifier: dto.agentIdentifier,
      agentName: customer.fullName,
      phoneNumber: customer.phoneNumber,
      agentType: 'SUPER_AGENT',
      status: customer.status,
      canReceiveFloat: customer.status === 'ACTIVE',
      settlementAccount: customer.accounts?.[0]?.accountNumber || null,
      simulation: true,
      message: 'Agent profile retrieved from local database',
    };
  }

  async transferToWallet(dto: AccountToWalletTransferDto) {
    if (dto.idempotencyKey) {
      const existing = await this.transactionRepo.findOne({
        where: { idempotencyKey: dto.idempotencyKey },
      });
      if (existing) {
        return {
          transactionReference: existing.transactionReference,
          status: existing.status,
          amount: existing.amount,
          currency: existing.currencyCode,
          walletPhoneNumber: dto.walletPhoneNumber,
          completedAt: existing.completedAt,
          simulation: true,
          externalSettlementConfirmed: false,
          message: 'Idempotent replay: previous wallet transfer returned',
        };
      }
    }

    return await this.dataSource.transaction(async (manager) => {
      const source = await manager.findOne(Account, {
        where: { accountNumber: dto.sourceAccountNumber },
      });
      if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

      if (source.status !== AccountStatus.ACTIVE || source.isDebitBlocked) {
        throw new AccountRestrictedException(source.accountNumber, 'Account not eligible for debit');
      }

      const service = await this.serviceRepo.findOne({ where: { serviceCode: 'CBE_BIRR' } });
      const fee = service ? service.feeValue : '0.0000';
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

      const txRef = `WLT26${Math.floor(10000000 + Math.random() * 90000000)}`;

      const tx = manager.create(Transaction, {
        transactionReference: txRef,
        idempotencyKey: dto.idempotencyKey || null,
        transactionType: TransactionType.WALLET_TRANSFER,
        channel: TransactionChannel.CBE_BIRR,
        sourceAccountId: source.id,
        externalAccount: dto.walletPhoneNumber,
        amount: dto.amount,
        feeAmount: fee,
        currencyCode: source.currencyCode,
        narration: dto.narration || `Transfer to CBE Birr: ${dto.walletPhoneNumber}`,
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
        narration: `Top-up CBE Birr Wallet [${dto.walletPhoneNumber}]`,
      }));

      return {
        transactionReference: savedTx.transactionReference,
        status: savedTx.status,
        walletPhoneNumber: dto.walletPhoneNumber,
        amount: dto.amount,
        fee,
        currency: source.currencyCode,
        remainingBalance: balance.availableBalance,
        completedAt: savedTx.completedAt,
        simulation: true,
        externalSettlementConfirmed: false,
        message: 'Debit posted to bank ledger; CBE Birr wallet credit simulated in local sandbox',
      };
    });
  }

  async transferFromWallet(dto: WalletToAccountTransferDto) {
    if (dto.idempotencyKey) {
      const existing = await this.transactionRepo.findOne({
        where: { idempotencyKey: dto.idempotencyKey },
      });
      if (existing) {
        return {
          transactionReference: existing.transactionReference,
          status: existing.status,
          amount: existing.amount,
          currency: existing.currencyCode,
          destinationAccountNumber: dto.destinationAccountNumber,
          completedAt: existing.completedAt,
          simulation: true,
          externalSettlementConfirmed: false,
          message: 'Idempotent replay: previous transfer returned',
        };
      }
    }

    return await this.dataSource.transaction(async (manager) => {
      const dest = await manager.findOne(Account, {
        where: { accountNumber: dto.destinationAccountNumber },
      });
      if (!dest) throw new AccountNotFoundException(dto.destinationAccountNumber);

      if (dest.status !== AccountStatus.ACTIVE || dest.isCreditBlocked) {
        throw new AccountRestrictedException(dest.accountNumber, 'Account not eligible for credit');
      }

      const balance = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId = :accountId', { accountId: dest.id })
        .getOne();

      if (!balance) throw new BadRequestException('Balance record missing');

      balance.ledgerBalance = FinancialMath.add(balance.ledgerBalance, dto.amount);
      balance.availableBalance = FinancialMath.add(balance.availableBalance, dto.amount);
      await manager.save(AccountBalance, balance);

      const txRef = `WLT26${Math.floor(10000000 + Math.random() * 90000000)}`;

      const tx = manager.create(Transaction, {
        transactionReference: txRef,
        idempotencyKey: dto.idempotencyKey || null,
        transactionType: TransactionType.WALLET_TRANSFER,
        channel: TransactionChannel.CBE_BIRR,
        destinationAccountId: dest.id,
        externalAccount: dto.walletPhoneNumber,
        amount: dto.amount,
        feeAmount: '0.0000',
        currencyCode: dest.currencyCode,
        narration: dto.narration || `Transfer from CBE Birr: ${dto.walletPhoneNumber}`,
        status: TransactionStatus.COMPLETED,
        completedAt: new Date(),
      });
      const savedTx = await manager.save(Transaction, tx);

      await manager.save(LedgerEntry, manager.create(LedgerEntry, {
        transactionId: savedTx.id,
        accountId: dest.id,
        entryType: EntryType.CREDIT,
        amount: dto.amount,
        currencyCode: dest.currencyCode,
        balanceAfter: balance.ledgerBalance,
        narration: `Cash-in from CBE Birr [${dto.walletPhoneNumber}]`,
      }));

      return {
        transactionReference: savedTx.transactionReference,
        status: savedTx.status,
        walletPhoneNumber: dto.walletPhoneNumber,
        destinationAccountNumber: dest.accountNumber,
        amount: dto.amount,
        currency: dest.currencyCode,
        newAvailableBalance: balance.availableBalance,
        completedAt: savedTx.completedAt,
        simulation: true,
        externalSettlementConfirmed: false,
        message: 'Funds received from CBE Birr wallet (local simulation)',
      };
    });
  }

  async agentTransfer(dto: WalletAgentTransferDto) {
    return this.transferToWallet({
      sourceAccountNumber: dto.sourceAccountNumber,
      walletPhoneNumber: dto.agentIdentifier,
      amount: dto.amount,
      narration: dto.narration || `Agent Float Transfer: ${dto.agentIdentifier}`,
      idempotencyKey: dto.idempotencyKey,
    });
  }
}
