"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.AccountsService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const account_hold_entity_1 = require("../../database/entities/account-hold.entity");
const customer_entity_1 = require("../../database/entities/customer.entity");
const branch_entity_1 = require("../../database/entities/branch.entity");
const currency_entity_1 = require("../../database/entities/currency.entity");
const account_type_entity_1 = require("../../database/entities/account-type.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const account_party_entity_1 = require("../../database/entities/account-party.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
const financial_math_util_1 = require("../../common/utils/financial-math.util");
let AccountsService = class AccountsService {
    dataSource;
    accountRepository;
    balanceRepository;
    holdRepository;
    customerRepository;
    branchRepository;
    currencyRepository;
    accountTypeRepository;
    ledgerRepository;
    constructor(dataSource, accountRepository, balanceRepository, holdRepository, customerRepository, branchRepository, currencyRepository, accountTypeRepository, ledgerRepository) {
        this.dataSource = dataSource;
        this.accountRepository = accountRepository;
        this.balanceRepository = balanceRepository;
        this.holdRepository = holdRepository;
        this.customerRepository = customerRepository;
        this.branchRepository = branchRepository;
        this.currencyRepository = currencyRepository;
        this.accountTypeRepository = accountTypeRepository;
        this.ledgerRepository = ledgerRepository;
    }
    async lookup(dto) {
        const account = await this.accountRepository.findOne({
            where: { accountNumber: dto.accountNumber },
            relations: ['customer', 'branch', 'currency', 'accountType', 'balance', 'holds'],
        });
        if (!account) {
            throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
        }
        return this.mapToAccountResponse(account);
    }
    async nameLookup(dto) {
        const account = await this.accountRepository.findOne({
            where: { accountNumber: dto.accountNumber },
            relations: ['branch', 'currency'],
        });
        if (!account) {
            throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
        }
        return {
            accountNumber: account.accountNumber,
            accountName: account.accountName,
            currency: account.currencyCode,
            branchCode: account.branch?.branchCode,
            status: account.status,
            isAllowedToTransact: account.status === account_entity_1.AccountStatus.ACTIVE && !account.isPostingRestricted,
        };
    }
    async list(dto) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber: dto.customerNumber },
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
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
    async create(dto) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber: dto.customerNumber },
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
        }
        const branch = await this.branchRepository.findOne({
            where: { branchCode: dto.branchCode },
        });
        if (!branch) {
            throw new common_1.BadRequestException(`Invalid branch code '${dto.branchCode}'`);
        }
        const currency = await this.currencyRepository.findOne({
            where: { code: dto.currencyCode },
        });
        if (!currency) {
            throw new common_1.BadRequestException(`Invalid currency code '${dto.currencyCode}'`);
        }
        const accountType = await this.accountTypeRepository.findOne({
            where: { code: dto.accountTypeCode },
        });
        if (!accountType) {
            throw new common_1.BadRequestException(`Invalid account type code '${dto.accountTypeCode}'`);
        }
        return await this.dataSource.transaction(async (manager) => {
            const count = await manager.count(account_entity_1.Account);
            const accountNumber = `1000${(100000000 + count + 1).toString()}`;
            const accountName = dto.accountName || customer.fullName;
            const account = manager.create(account_entity_1.Account, {
                accountNumber,
                customerId: customer.id,
                branchId: branch.id,
                currencyCode: currency.code,
                accountTypeId: accountType.id,
                accountName,
                status: account_entity_1.AccountStatus.ACTIVE,
            });
            const savedAccount = await manager.save(account_entity_1.Account, account);
            const initialDepositStr = dto.initialDeposit ? financial_math_util_1.FinancialMath.add(dto.initialDeposit, '0') : '0.0000';
            const balance = manager.create(account_balance_entity_1.AccountBalance, {
                accountId: savedAccount.id,
                ledgerBalance: initialDepositStr,
                availableBalance: initialDepositStr,
                lockedBalance: '0.0000',
                currencyCode: currency.code,
            });
            await manager.save(account_balance_entity_1.AccountBalance, balance);
            if (financial_math_util_1.FinancialMath.isPositive(initialDepositStr)) {
                const tx = manager.create(transaction_entity_1.Transaction, {
                    transactionReference: `INIT-${accountNumber}`,
                    transactionType: transaction_entity_1.TransactionType.CASH_DEPOSIT,
                    channel: transaction_entity_1.TransactionChannel.SYSTEM,
                    destinationAccountId: savedAccount.id,
                    amount: initialDepositStr,
                    currencyCode: currency.code,
                    narration: 'Account Opening Deposit',
                    status: transaction_entity_1.TransactionStatus.COMPLETED,
                    completedAt: new Date(),
                });
                await manager.save(transaction_entity_1.Transaction, tx);
                const entry = manager.create(ledger_entry_entity_1.LedgerEntry, {
                    transactionId: tx.id,
                    accountId: savedAccount.id,
                    entryType: ledger_entry_entity_1.EntryType.CREDIT,
                    amount: initialDepositStr,
                    currencyCode: currency.code,
                    balanceAfter: initialDepositStr,
                    narration: 'Initial Deposit',
                });
                await manager.save(ledger_entry_entity_1.LedgerEntry, entry);
            }
            savedAccount.balance = balance;
            savedAccount.branch = branch;
            savedAccount.currency = currency;
            savedAccount.accountType = accountType;
            savedAccount.customer = customer;
            const party = manager.create(account_party_entity_1.AccountParty, {
                accountId: savedAccount.id,
                customerId: customer.id,
                partyRole: account_party_entity_1.PartyRole.PRIMARY_OWNER,
                authorizationLevel: 'SOLE_SIGNATORY',
                status: 'ACTIVE',
            });
            await manager.save(account_party_entity_1.AccountParty, party);
            return this.mapToAccountResponse(savedAccount);
        });
    }
    async checkEligibility(dto) {
        const account = await this.accountRepository.findOne({
            where: { accountNumber: dto.accountNumber },
            relations: ['balance'],
        });
        if (!account) {
            throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
        }
        const errors = [];
        if (account.status !== account_entity_1.AccountStatus.ACTIVE) {
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
        const hasSufficientFunds = financial_math_util_1.FinancialMath.isGreaterThanOrEqualTo(available, dto.amount);
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
    async createHold(dto) {
        return await this.dataSource.transaction(async (manager) => {
            const account = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.accountNumber },
            });
            if (!account) {
                throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
            }
            const balance = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId = :accountId', { accountId: account.id })
                .getOne();
            if (!balance) {
                throw new common_1.BadRequestException('Balance record missing for account');
            }
            if (financial_math_util_1.FinancialMath.isLessThan(balance.availableBalance, dto.amount)) {
                throw new domain_exceptions_1.InsufficientFundsException(balance.availableBalance, dto.amount);
            }
            const holdCount = await manager.count(account_hold_entity_1.AccountHold);
            const holdReference = `HLD-${new Date().toISOString().slice(0, 10).replace(/-/g, '')}-${(1000 + holdCount + 1).toString()}`;
            const hold = manager.create(account_hold_entity_1.AccountHold, {
                holdReference,
                accountId: account.id,
                amount: dto.amount,
                remainingAmount: dto.amount,
                currencyCode: account.currencyCode,
                reason: dto.reason,
                lockEvent: dto.lockEvent || 'HOLD',
                status: account_hold_entity_1.HoldStatus.ACTIVE,
                fromDate: dto.fromDate || new Date().toISOString().slice(0, 10),
                toDate: dto.toDate || null,
                createdBy: 'SYSTEM_API',
            });
            await manager.save(account_hold_entity_1.AccountHold, hold);
            balance.lockedBalance = financial_math_util_1.FinancialMath.add(balance.lockedBalance, dto.amount);
            balance.availableBalance = financial_math_util_1.FinancialMath.subtract(balance.availableBalance, dto.amount);
            await manager.save(account_balance_entity_1.AccountBalance, balance);
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
    async releaseHold(dto) {
        return await this.dataSource.transaction(async (manager) => {
            const hold = await manager.findOne(account_hold_entity_1.AccountHold, {
                where: { holdReference: dto.holdReference, status: account_hold_entity_1.HoldStatus.ACTIVE },
                relations: ['account'],
            });
            if (!hold) {
                throw new domain_exceptions_1.HoldNotFoundException(dto.holdReference);
            }
            const balance = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId = :accountId', { accountId: hold.accountId })
                .getOne();
            if (!balance) {
                throw new common_1.BadRequestException('Balance record not found');
            }
            hold.status = account_hold_entity_1.HoldStatus.RELEASED;
            hold.releasedAt = new Date();
            hold.releasedReason = dto.reason || 'Administrative Release';
            await manager.save(account_hold_entity_1.AccountHold, hold);
            balance.lockedBalance = financial_math_util_1.FinancialMath.subtract(balance.lockedBalance, hold.amount);
            balance.availableBalance = financial_math_util_1.FinancialMath.add(balance.availableBalance, hold.amount);
            await manager.save(account_balance_entity_1.AccountBalance, balance);
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
    async listHolds(accountNumber) {
        const account = await this.accountRepository.findOne({
            where: { accountNumber },
        });
        if (!account) {
            throw new domain_exceptions_1.AccountNotFoundException(accountNumber);
        }
        const holds = await this.holdRepository.find({
            where: { accountId: account.id, status: account_hold_entity_1.HoldStatus.ACTIVE },
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
    async getStatement(dto) {
        const account = await this.accountRepository.findOne({
            where: { accountNumber: dto.accountNumber },
            relations: ['balance'],
        });
        if (!account) {
            throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
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
    async getMiniStatement(dto) {
        return this.getStatement({
            accountNumber: dto.accountNumber,
            limit: dto.limit || dto.count || 5,
        });
    }
    mapToAccountResponse(a) {
        const activeHolds = a.holds ? a.holds.filter((h) => h.status === account_hold_entity_1.HoldStatus.ACTIVE) : [];
        const lockedSum = activeHolds.reduce((acc, h) => financial_math_util_1.FinancialMath.add(acc, h.amount), '0.0000');
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
};
exports.AccountsService = AccountsService;
exports.AccountsService = AccountsService = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(2, (0, typeorm_1.InjectRepository)(account_balance_entity_1.AccountBalance)),
    __param(3, (0, typeorm_1.InjectRepository)(account_hold_entity_1.AccountHold)),
    __param(4, (0, typeorm_1.InjectRepository)(customer_entity_1.Customer)),
    __param(5, (0, typeorm_1.InjectRepository)(branch_entity_1.Branch)),
    __param(6, (0, typeorm_1.InjectRepository)(currency_entity_1.Currency)),
    __param(7, (0, typeorm_1.InjectRepository)(account_type_entity_1.AccountType)),
    __param(8, (0, typeorm_1.InjectRepository)(ledger_entry_entity_1.LedgerEntry)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], AccountsService);
//# sourceMappingURL=accounts.service.js.map