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
exports.IpsService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const service_entity_1 = require("../../database/entities/service.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
const financial_math_util_1 = require("../../common/utils/financial-math.util");
let IpsService = class IpsService {
    dataSource;
    accountRepo;
    transactionRepo;
    serviceRepo;
    constructor(dataSource, accountRepo, transactionRepo, serviceRepo) {
        this.dataSource = dataSource;
        this.accountRepo = accountRepo;
        this.transactionRepo = transactionRepo;
        this.serviceRepo = serviceRepo;
    }
    async accountLookup(dto) {
        const account = await this.accountRepo.findOne({
            where: { accountNumber: dto.accountNumber },
            relations: ['customer', 'currency', 'branch'],
        });
        if (!account) {
            throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
        }
        return {
            accountNumber: account.accountNumber,
            accountName: account.accountName,
            currency: account.currencyCode,
            bankCode: dto.bankCode || 'CBETETAA',
            bankName: 'Commercial Bank of Ethiopia',
            branchName: account.branch?.branchName,
            status: account.status,
            isEligible: account.status === account_entity_1.AccountStatus.ACTIVE && !account.isCreditBlocked,
        };
    }
    async executeTransfer(dto) {
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
            const source = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.sourceAccountNumber },
            });
            if (!source)
                throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
            if (source.status !== account_entity_1.AccountStatus.ACTIVE || source.isDebitBlocked) {
                throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, 'Account not permitted for IPS transfer debit');
            }
            const service = await manager.findOne(service_entity_1.ServiceEntity, { where: { serviceCode: 'FT_IPS' } });
            const fee = service ? service.feeValue : '5.0000';
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
            const txRef = `IPS26${Math.floor(10000000 + Math.random() * 90000000)}`;
            const tx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: txRef,
                idempotencyKey: dto.idempotencyKey || null,
                transactionType: transaction_entity_1.TransactionType.IPS_TRANSFER,
                channel: transaction_entity_1.TransactionChannel.IPS,
                sourceAccountId: source.id,
                externalAccount: dto.destinationAccountNumber,
                externalBankCode: dto.destinationBankCode,
                amount: dto.amount,
                feeAmount: fee,
                currencyCode: source.currencyCode,
                narration: dto.narration || `IPS to ${dto.destinationBankCode}: ${dto.beneficiaryName}`,
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
    async qrPayment(dto) {
        return await this.dataSource.transaction(async (manager) => {
            const source = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.sourceAccountNumber },
            });
            if (!source)
                throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
            const merchant = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.merchantAccount },
            });
            if (!merchant)
                throw new domain_exceptions_1.AccountNotFoundException(dto.merchantAccount);
            const lockIds = [source.id, merchant.id].sort();
            const balances = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId IN (:...ids)', { ids: lockIds })
                .getMany();
            const sourceBal = balances.find((b) => b.accountId === source.id);
            const merchantBal = balances.find((b) => b.accountId === merchant.id);
            if (!sourceBal || !merchantBal)
                throw new common_1.BadRequestException('Failed to lock accounts');
            if (financial_math_util_1.FinancialMath.isLessThan(sourceBal.availableBalance, dto.amount)) {
                throw new domain_exceptions_1.InsufficientFundsException(sourceBal.availableBalance, dto.amount);
            }
            sourceBal.ledgerBalance = financial_math_util_1.FinancialMath.subtract(sourceBal.ledgerBalance, dto.amount);
            sourceBal.availableBalance = financial_math_util_1.FinancialMath.subtract(sourceBal.availableBalance, dto.amount);
            await manager.save(account_balance_entity_1.AccountBalance, sourceBal);
            merchantBal.ledgerBalance = financial_math_util_1.FinancialMath.add(merchantBal.ledgerBalance, dto.amount);
            merchantBal.availableBalance = financial_math_util_1.FinancialMath.add(merchantBal.availableBalance, dto.amount);
            await manager.save(account_balance_entity_1.AccountBalance, merchantBal);
            const txRef = `QR26${Math.floor(10000000 + Math.random() * 90000000)}`;
            const tx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: txRef,
                idempotencyKey: dto.idempotencyKey || null,
                transactionType: transaction_entity_1.TransactionType.IPS_QR_PAYMENT,
                channel: transaction_entity_1.TransactionChannel.IPS,
                sourceAccountId: source.id,
                destinationAccountId: merchant.id,
                amount: dto.amount,
                feeAmount: '0.0000',
                currencyCode: source.currencyCode,
                narration: dto.narration || `IPS QR Merchant Payment: ${dto.merchantName}`,
                status: transaction_entity_1.TransactionStatus.COMPLETED,
                completedAt: new Date(),
            });
            const savedTx = await manager.save(transaction_entity_1.Transaction, tx);
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedTx.id,
                accountId: source.id,
                entryType: ledger_entry_entity_1.EntryType.DEBIT,
                amount: dto.amount,
                currencyCode: source.currencyCode,
                balanceAfter: sourceBal.ledgerBalance,
                narration: `QR Pay -> ${dto.merchantName}`,
            }));
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedTx.id,
                accountId: merchant.id,
                entryType: ledger_entry_entity_1.EntryType.CREDIT,
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
    async getStatus(dto) {
        const tx = await this.transactionRepo.findOne({
            where: { transactionReference: dto.transactionReference },
        });
        if (!tx)
            throw new common_1.BadRequestException(`IPS transaction '${dto.transactionReference}' not found`);
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
};
exports.IpsService = IpsService;
exports.IpsService = IpsService = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(2, (0, typeorm_1.InjectRepository)(transaction_entity_1.Transaction)),
    __param(3, (0, typeorm_1.InjectRepository)(service_entity_1.ServiceEntity)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], IpsService);
//# sourceMappingURL=ips.service.js.map