import { Injectable, BadRequestException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository, Between, LessThanOrEqual, MoreThanOrEqual } from 'typeorm';
import { Account, AccountStatus } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { AccountHold, HoldStatus } from '../../database/entities/account-hold.entity';
import { Customer } from '../../database/entities/customer.entity';
import { Branch } from '../../database/entities/branch.entity';
import { Currency } from '../../database/entities/currency.entity';
import { AccountType } from '../../database/entities/account-type.entity';
import { Transaction, TransactionType, TransactionStatus, TransactionChannel } from '../../database/entities/transaction.entity';
import { LedgerEntry, EntryType } from '../../database/entities/ledger-entry.entity';
import { AccountParty, PartyRole } from '../../database/entities/account-party.entity';
import {
  AccountLookupDto,
  AccountNameLookupDto,
  CreateAccountDto,
  AccountListDto,
  AccountEligibilityDto,
  CreateHoldDto,
  ReleaseHoldDto,
  AccountStatementDto,
  AccountMiniStatementDto,
} from './dto/account.dto';
import {
  AccountNotFoundException,
  CustomerNotFoundException,
  InsufficientFundsException,
  AccountRestrictedException,
  HoldNotFoundException,
} from '../../common/exceptions/domain.exceptions';
import { FinancialMath } from '../../common/utils/financial-math.util';

@Injectable()
export class AccountsService {
  constructor(
    private readonly dataSource: DataSource,
    @InjectRepository(Account)
    private readonly accountRepository: Repository<Account>,
    @InjectRepository(AccountBalance)
    private readonly balanceRepository: Repository<AccountBalance>,
    @InjectRepository(AccountHold)
    private readonly holdRepository: Repository<AccountHold>,
    @InjectRepository(Customer)
    private readonly customerRepository: Repository<Customer>,
    @InjectRepository(Branch)
    private readonly branchRepository: Repository<Branch>,
    @InjectRepository(Currency)
    private readonly currencyRepository: Repository<Currency>,
    @InjectRepository(AccountType)
    private readonly accountTypeRepository: Repository<AccountType>,
    @InjectRepository(LedgerEntry)
    private readonly ledgerRepository: Repository<LedgerEntry>,
  ) {}

  async lookup(dto: AccountLookupDto) {
    const account = await this.accountRepository.findOne({
      where: { accountNumber: dto.accountNumber },
      relations: ['customer', 'branch', 'currency', 'accountType', 'balance', 'holds'],
    });

    if (!account) {
      throw new AccountNotFoundException(dto.accountNumber);
    }

    return this.mapToAccountResponse(account);
  }

  async nameLookup(dto: AccountNameLookupDto) {
    const account = await this.accountRepository.findOne({
      where: { accountNumber: dto.accountNumber },
      relations: ['branch', 'currency'],
    });

    if (!account) {
      throw new AccountNotFoundException(dto.accountNumber);
    }

    return {
      accountNumber: account.accountNumber,
      accountName: account.accountName,
      currency: account.currencyCode,
      branchCode: account.branch?.branchCode,
      status: account.status,
      isAllowedToTransact: account.status === AccountStatus.ACTIVE && !account.isPostingRestricted,
    };
  }

  async list(dto: AccountListDto) {
    const customer = await this.customerRepository.findOne({
      where: { customerNumber: dto.customerNumber },
    });

    if (!customer) {
      throw new CustomerNotFoundException(dto.customerNumber);
    }

    const accounts = await this.accountRepository.find({
      where: { customerId: customer.id },
      relations: ['branch', 'currency', 'accountType', 'balance'],
      order: { createdAt: 'ASC' },
    });

    return {
      customerNumber: customer.customerNumber,
      customerName: customer.fullName,
      totalAccounts: accounts.length,
      accounts: accounts.map((a) => this.mapToAccountResponse(a)),
    };
  }

  async create(dto: CreateAccountDto) {
    const customer = await this.customerRepository.findOne({
      where: { customerNumber: dto.customerNumber },
    });
    if (!customer) {
      throw new CustomerNotFoundException(dto.customerNumber);
    }

    const branch = await this.branchRepository.findOne({
      where: { branchCode: dto.branchCode },
    });
    if (!branch) {
      throw new BadRequestException(`Invalid branch code '${dto.branchCode}'`);
    }

    const currency = await this.currencyRepository.findOne({
      where: { code: dto.currencyCode },
    });
    if (!currency) {
      throw new BadRequestException(`Invalid currency code '${dto.currencyCode}'`);
    }

    const accountType = await this.accountTypeRepository.findOne({
      where: { code: dto.accountTypeCode },
    });
    if (!accountType) {
      throw new BadRequestException(`Invalid account type code '${dto.accountTypeCode}'`);
    }

    return await this.dataSource.transaction(async (manager) => {
      // Generate 13-digit account number (1000 + random 9 digits)
      const count = await manager.count(Account);
      const accountNumber = `1000${(100000000 + count + 1).toString()}`;

      const accountName = dto.accountName || customer.fullName;

      const account = manager.create(Account, {
        accountNumber,
        customerId: customer.id,
        branchId: branch.id,
        currencyCode: currency.code,
        accountTypeId: accountType.id,
        accountName,
        status: AccountStatus.ACTIVE,
      });

      const savedAccount = await manager.save(Account, account);

      const initialDepositStr = dto.initialDeposit ? FinancialMath.add(dto.initialDeposit, '0') : '0.0000';

      const balance = manager.create(AccountBalance, {
        accountId: savedAccount.id,
        ledgerBalance: initialDepositStr,
        availableBalance: initialDepositStr,
        lockedBalance: '0.0000',
        currencyCode: currency.code,
      });
      await manager.save(AccountBalance, balance);

      if (FinancialMath.isPositive(initialDepositStr)) {
        const tx = manager.create(Transaction, {
          transactionReference: `INIT-${accountNumber}`,
          transactionType: TransactionType.CASH_DEPOSIT,
          channel: TransactionChannel.SYSTEM,
          destinationAccountId: savedAccount.id,
          amount: initialDepositStr,
          currencyCode: currency.code,
          narration: 'Account Opening Deposit',
          status: TransactionStatus.COMPLETED,
          completedAt: new Date(),
        });
        await manager.save(Transaction, tx);

        const entry = manager.create(LedgerEntry, {
          transactionId: tx.id,
          accountId: savedAccount.id,
          entryType: EntryType.CREDIT,
          amount: initialDepositStr,
          currencyCode: currency.code,
          balanceAfter: initialDepositStr,
          narration: 'Initial Deposit',
        });
        await manager.save(LedgerEntry, entry);
      }

      savedAccount.balance = balance;
      savedAccount.branch = branch;
      savedAccount.currency = currency;
      savedAccount.accountType = accountType;
      savedAccount.customer = customer;

      // Create primary AccountParty record
      const party = manager.create(AccountParty, {
        accountId: savedAccount.id,
        customerId: customer.id,
        partyRole: PartyRole.PRIMARY_OWNER,
        authorizationLevel: 'SOLE_SIGNATORY',
        status: 'ACTIVE',
      });
      await manager.save(AccountParty, party);

      return this.mapToAccountResponse(savedAccount);
    });
  }

  async checkEligibility(dto: AccountEligibilityDto) {
    const account = await this.accountRepository.findOne({
      where: { accountNumber: dto.accountNumber },
      relations: ['balance'],
    });

    if (!account) {
      throw new AccountNotFoundException(dto.accountNumber);
    }

    const errors: string[] = [];
    if (account.status !== AccountStatus.ACTIVE) {
      errors.push(`Account status is ${account.status}`);
    }
    if (account.isPostingRestricted) {
      errors.push(`Account posting restricted: ${account.postingRestrictionReason || 'No reason specified'}`);
    }
    if (account.isDebitBlocked && dto.transactionType === 'DEBIT') {
      errors.push('Account is blocked for debit transactions');
    }
    if (account.isCreditBlocked && dto.transactionType === 'CREDIT') {
      errors.push('Account is blocked for credit transactions');
    }

    const available = account.balance ? account.balance.availableBalance : '0.0000';
    const hasSufficientFunds = FinancialMath.isGreaterThanOrEqualTo(available, dto.amount);

    if (dto.transactionType === 'DEBIT' && !hasSufficientFunds) {
      errors.push(`Insufficient funds. Available: ${available}, Requested: ${dto.amount}`);
    }

    return {
      accountNumber: account.accountNumber,
      accountName: account.accountName,
      eligible: errors.length === 0,
      currency: account.currencyCode,
      availableBalance: available,
      requestedAmount: dto.amount,
      reasons: errors,
    };
  }

  async createHold(dto: CreateHoldDto) {
    return await this.dataSource.transaction(async (manager) => {
      const account = await manager.findOne(Account, {
        where: { accountNumber: dto.accountNumber },
      });

      if (!account) {
        throw new AccountNotFoundException(dto.accountNumber);
      }

      // Lock balance row using SELECT FOR UPDATE
      const balance = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId = :accountId', { accountId: account.id })
        .getOne();

      if (!balance) {
        throw new BadRequestException('Balance record missing for account');
      }

      if (FinancialMath.isLessThan(balance.availableBalance, dto.amount)) {
        throw new InsufficientFundsException(balance.availableBalance, dto.amount);
      }

      const holdCount = await manager.count(AccountHold);
      const holdReference = `HLD-${new Date().toISOString().slice(0, 10).replace(/-/g, '')}-${(1000 + holdCount + 1).toString()}`;

      const hold = manager.create(AccountHold, {
        holdReference,
        accountId: account.id,
        amount: dto.amount,
        remainingAmount: dto.amount,
        currencyCode: account.currencyCode,
        reason: dto.reason,
        lockEvent: dto.lockEvent || 'HOLD',
        status: HoldStatus.ACTIVE,
        fromDate: dto.fromDate || new Date().toISOString().slice(0, 10),
        toDate: dto.toDate || null,
        createdBy: 'SYSTEM_API',
      });
      await manager.save(AccountHold, hold);

      // Adjust balances
      balance.lockedBalance = FinancialMath.add(balance.lockedBalance, dto.amount);
      balance.availableBalance = FinancialMath.subtract(balance.availableBalance, dto.amount);
      await manager.save(AccountBalance, balance);

      return {
        holdReference: hold.holdReference,
        accountNumber: account.accountNumber,
        amount: hold.amount,
        currency: hold.currencyCode,
        reason: hold.reason,
        lockEvent: hold.lockEvent,
        status: hold.status,
        newAvailableBalance: balance.availableBalance,
        newLockedBalance: balance.lockedBalance,
        createdAt: hold.createdAt,
      };
    });
  }

  async releaseHold(dto: ReleaseHoldDto) {
    return await this.dataSource.transaction(async (manager) => {
      const hold = await manager.findOne(AccountHold, {
        where: { holdReference: dto.holdReference, status: HoldStatus.ACTIVE },
        relations: ['account'],
      });

      if (!hold) {
        throw new HoldNotFoundException(dto.holdReference);
      }

      const balance = await manager
        .createQueryBuilder(AccountBalance, 'bal')
        .setLock('pessimistic_write')
        .where('bal.accountId = :accountId', { accountId: hold.accountId })
        .getOne();

      if (!balance) {
        throw new BadRequestException('Balance record not found');
      }

      hold.status = HoldStatus.RELEASED;
      hold.releasedAt = new Date();
      hold.releasedReason = dto.reason || 'Administrative Release';
      await manager.save(AccountHold, hold);

      // Restore available balance
      balance.lockedBalance = FinancialMath.subtract(balance.lockedBalance, hold.amount);
      balance.availableBalance = FinancialMath.add(balance.availableBalance, hold.amount);
      await manager.save(AccountBalance, balance);

      return {
        holdReference: hold.holdReference,
        accountNumber: hold.account?.accountNumber,
        amount: hold.amount,
        status: hold.status,
        releasedAt: hold.releasedAt,
        releasedReason: hold.releasedReason,
        newAvailableBalance: balance.availableBalance,
        newLockedBalance: balance.lockedBalance,
      };
    });
  }

  async listHolds(accountNumber: string) {
    const account = await this.accountRepository.findOne({
      where: { accountNumber },
    });

    if (!account) {
      throw new AccountNotFoundException(accountNumber);
    }

    const holds = await this.holdRepository.find({
      where: { accountId: account.id, status: HoldStatus.ACTIVE },
      order: { createdAt: 'DESC' },
    });

    return {
      accountNumber,
      activeHoldsCount: holds.length,
      holds: holds.map((h) => ({
        holdReference: h.holdReference,
        amount: h.amount,
        currency: h.currencyCode,
        reason: h.reason,
        lockEvent: h.lockEvent,
        status: h.status,
        fromDate: h.fromDate,
        toDate: h.toDate,
        createdAt: h.createdAt,
      })),
    };
  }

  async getStatement(dto: AccountStatementDto) {
    const account = await this.accountRepository.findOne({
      where: { accountNumber: dto.accountNumber },
      relations: ['balance'],
    });

    if (!account) {
      throw new AccountNotFoundException(dto.accountNumber);
    }

    const query = this.ledgerRepository
      .createQueryBuilder('entry')
      .where('entry.accountId = :accountId', { accountId: account.id })
      .orderBy('entry.postedAt', 'DESC')
      .take(dto.limit || 50);

    if (dto.startDate) {
      query.andWhere('entry.valueDate >= :startDate', { startDate: dto.startDate });
    }
    if (dto.endDate) {
      query.andWhere('entry.valueDate <= :endDate', { endDate: dto.endDate });
    }

    const entries = await query.getMany();
    const mappedEntries = entries.map((e) => ({
      id: e.id,
      transactionId: e.transactionId,
      entryType: e.entryType,
      amount: e.amount,
      currency: e.currencyCode,
      balanceAfter: e.balanceAfter,
      narration: e.narration,
      valueDate: e.valueDate,
      postedAt: e.postedAt,
    }));

    const currentLedger = account.balance?.ledgerBalance || '0.0000';
    const openingBal = entries.length > 0 ? entries[entries.length - 1].balanceAfter : currentLedger;

    return {
      accountNumber: account.accountNumber,
      accountName: account.accountName,
      currency: account.currencyCode,
      currentLedgerBalance: currentLedger,
      currentAvailableBalance: account.balance?.availableBalance || '0.0000',
      openingBalance: openingBal,
      closingBalance: currentLedger,
      recordCount: entries.length,
      entries: mappedEntries,
      transactions: mappedEntries,
    };
  }

  async getMiniStatement(dto: AccountMiniStatementDto) {
    return this.getStatement({
      accountNumber: dto.accountNumber,
      limit: dto.limit || dto.count || 5,
    });
  }

  private mapToAccountResponse(a: Account) {
    const activeHolds = a.holds ? a.holds.filter((h) => h.status === HoldStatus.ACTIVE) : [];
    const lockedSum = activeHolds.reduce((acc, h) => FinancialMath.add(acc, h.amount), '0.0000');

    return {
      accountNumber: a.accountNumber,
      accountName: a.accountName,
      customerNumber: a.customer?.customerNumber,
      customerName: a.customer?.fullName,
      accountTypeCode: a.accountType?.code,
      accountTypeName: a.accountType?.name,
      branchCode: a.branch?.branchCode,
      branchName: a.branch?.branchName,
      currency: a.currencyCode,
      status: a.status,
      isPostingRestricted: a.isPostingRestricted,
      postingRestrictionReason: a.postingRestrictionReason,
      isDebitBlocked: a.isDebitBlocked,
      isCreditBlocked: a.isCreditBlocked,
      ledgerBalance: a.balance?.ledgerBalance || '0.0000',
      availableBalance: a.balance?.availableBalance || '0.0000',
      lockedBalance: a.balance?.lockedBalance || lockedSum,
      openedAt: a.openedAt,
      closedAt: a.closedAt,
      createdAt: a.createdAt,
    };
  }
}
