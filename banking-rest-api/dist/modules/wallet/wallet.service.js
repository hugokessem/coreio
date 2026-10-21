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
exports.WalletService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const customer_entity_1 = require("../../database/entities/customer.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const service_entity_1 = require("../../database/entities/service.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
const financial_math_util_1 = require("../../common/utils/financial-math.util");
let WalletService = class WalletService {
    dataSource;
    accountRepo;
    customerRepo;
    transactionRepo;
    serviceRepo;
    constructor(dataSource, accountRepo, customerRepo, transactionRepo, serviceRepo) {
        this.dataSource = dataSource;
        this.accountRepo = accountRepo;
        this.customerRepo = customerRepo;
        this.transactionRepo = transactionRepo;
        this.serviceRepo = serviceRepo;
    }
    async customerLookup(dto) {
        const customer = await this.customerRepo.findOne({
            where: { phoneNumber: dto.phoneNumber },
            relations: ['accounts'],
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(dto.phoneNumber);
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
    async agentLookup(dto) {
        const customer = await this.customerRepo.findOne({
            where: [
                { phoneNumber: dto.agentIdentifier },
                { customerNumber: dto.agentIdentifier },
            ],
            relations: ['accounts'],
        });
        if (!customer) {
            throw new common_1.NotFoundException(`CBE Birr Agent '${dto.agentIdentifier}' not found`);
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
    async transferToWallet(dto) {
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
            const source = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.sourceAccountNumber },
            });
            if (!source)
                throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
            if (source.status !== account_entity_1.AccountStatus.ACTIVE || source.isDebitBlocked) {
                throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, 'Account not eligible for debit');
            }
            const service = await this.serviceRepo.findOne({ where: { serviceCode: 'CBE_BIRR' } });
            const fee = service ? service.feeValue : '0.0000';
            const totalDebit = financial_math_util_1.FinancialMath.add(dto.amount, fee);
            const balance = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId = :accountId', { accountId: source.id })
                .getOne();
            if (!balance)
                throw new common_1.BadRequestException('Balance record missing');
            if (financial_math_util_1.FinancialMath.isLessThan(balance.availableBalance, totalDebit)) {
                throw new domain_exceptions_1.InsufficientFundsException(balance.availableBalance, totalDebit);
            }
            balance.ledgerBalance = financial_math_util_1.FinancialMath.subtract(balance.ledgerBalance, totalDebit);
            balance.availableBalance = financial_math_util_1.FinancialMath.subtract(balance.availableBalance, totalDebit);
            await manager.save(account_balance_entity_1.AccountBalance, balance);
            const txRef = `WLT26${Math.floor(10000000 + Math.random() * 90000000)}`;
            const tx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: txRef,
                idempotencyKey: dto.idempotencyKey || null,
                transactionType: transaction_entity_1.TransactionType.WALLET_TRANSFER,
                channel: transaction_entity_1.TransactionChannel.CBE_BIRR,
                sourceAccountId: source.id,
                externalAccount: dto.walletPhoneNumber,
                amount: dto.amount,
                feeAmount: fee,
                currencyCode: source.currencyCode,
                narration: dto.narration || `Transfer to CBE Birr: ${dto.walletPhoneNumber}`,
                status: transaction_entity_1.TransactionStatus.COMPLETED,
                completedAt: new Date(),
            });
            const savedTx = await manager.save(transaction_entity_1.Transaction, tx);
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedTx.id,
                accountId: source.id,
                entryType: ledger_entry_entity_1.EntryType.DEBIT,
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
    async transferFromWallet(dto) {
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
            const dest = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.destinationAccountNumber },
            });
            if (!dest)
                throw new domain_exceptions_1.AccountNotFoundException(dto.destinationAccountNumber);
            if (dest.status !== account_entity_1.AccountStatus.ACTIVE || dest.isCreditBlocked) {
                throw new domain_exceptions_1.AccountRestrictedException(dest.accountNumber, 'Account not eligible for credit');
            }
            const balance = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId = :accountId', { accountId: dest.id })
                .getOne();
            if (!balance)
                throw new common_1.BadRequestException('Balance record missing');
            balance.ledgerBalance = financial_math_util_1.FinancialMath.add(balance.ledgerBalance, dto.amount);
            balance.availableBalance = financial_math_util_1.FinancialMath.add(balance.availableBalance, dto.amount);
            await manager.save(account_balance_entity_1.AccountBalance, balance);
            const txRef = `WLT26${Math.floor(10000000 + Math.random() * 90000000)}`;
            const tx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: txRef,
                idempotencyKey: dto.idempotencyKey || null,
                transactionType: transaction_entity_1.TransactionType.WALLET_TRANSFER,
                channel: transaction_entity_1.TransactionChannel.CBE_BIRR,
                destinationAccountId: dest.id,
                externalAccount: dto.walletPhoneNumber,
                amount: dto.amount,
                feeAmount: '0.0000',
                currencyCode: dest.currencyCode,
                narration: dto.narration || `Transfer from CBE Birr: ${dto.walletPhoneNumber}`,
                status: transaction_entity_1.TransactionStatus.COMPLETED,
                completedAt: new Date(),
            });
            const savedTx = await manager.save(transaction_entity_1.Transaction, tx);
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedTx.id,
                accountId: dest.id,
                entryType: ledger_entry_entity_1.EntryType.CREDIT,
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
    async agentTransfer(dto) {
        return this.transferToWallet({
            sourceAccountNumber: dto.sourceAccountNumber,
            walletPhoneNumber: dto.agentIdentifier,
            amount: dto.amount,
            narration: dto.narration || `Agent Float Transfer: ${dto.agentIdentifier}`,
            idempotencyKey: dto.idempotencyKey,
        });
    }
};
exports.WalletService = WalletService;
exports.WalletService = WalletService = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(2, (0, typeorm_1.InjectRepository)(customer_entity_1.Customer)),
    __param(3, (0, typeorm_1.InjectRepository)(transaction_entity_1.Transaction)),
    __param(4, (0, typeorm_1.InjectRepository)(service_entity_1.ServiceEntity)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], WalletService);
//# sourceMappingURL=wallet.service.js.map