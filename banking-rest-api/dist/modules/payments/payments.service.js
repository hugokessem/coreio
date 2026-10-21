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
exports.PaymentsService = void 0;
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
let PaymentsService = class PaymentsService {
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
    async verifyBill(dto) {
        const source = await this.accountRepo.findOne({
            where: { accountNumber: dto.sourceAccountNumber },
            relations: ['balance', 'currency'],
        });
        if (!source)
            throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
        if (source.status !== account_entity_1.AccountStatus.ACTIVE) {
            throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, `Source account status is ${source.status}`);
        }
        if (source.isDebitBlocked) {
            throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, 'Account is blocked for debit operations');
        }
        const biller = await this.serviceRepo.findOne({
            where: { serviceCode: dto.billerCode },
        });
        const fee = biller ? biller.feeValue : '0.0000';
        const totalRequired = financial_math_util_1.FinancialMath.add(dto.amount, fee);
        const available = source.balance?.availableBalance || '0.0000';
        const hasSufficientFunds = financial_math_util_1.FinancialMath.isGreaterThanOrEqualTo(available, totalRequired);
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
    async payBill(dto) {
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
            const source = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.sourceAccountNumber },
            });
            if (!source)
                throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
            if (source.status !== account_entity_1.AccountStatus.ACTIVE || source.isDebitBlocked) {
                throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, 'Account not eligible for bill payment debit');
            }
            const balance = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId = :accountId', { accountId: source.id })
                .getOne();
            if (!balance)
                throw new common_1.BadRequestException('Balance record not found');
            if (financial_math_util_1.FinancialMath.isLessThan(balance.availableBalance, dto.amount)) {
                throw new domain_exceptions_1.InsufficientFundsException(balance.availableBalance, dto.amount);
            }
            balance.ledgerBalance = financial_math_util_1.FinancialMath.subtract(balance.ledgerBalance, dto.amount);
            balance.availableBalance = financial_math_util_1.FinancialMath.subtract(balance.availableBalance, dto.amount);
            await manager.save(account_balance_entity_1.AccountBalance, balance);
            const randomSuffix = Math.floor(10000000 + Math.random() * 90000000).toString();
            const txRef = `BP26${randomSuffix}`;
            const tx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: txRef,
                idempotencyKey: dto.idempotencyKey || null,
                transactionType: transaction_entity_1.TransactionType.BILL_PAYMENT,
                channel: transaction_entity_1.TransactionChannel.MOBILE,
                sourceAccountId: source.id,
                externalAccount: dto.billReference,
                amount: dto.amount,
                feeAmount: '0.0000',
                currencyCode: source.currencyCode,
                narration: dto.narration || `Bill Payment: ${dto.billerCode} - ${dto.billReference}`,
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
    async splitPayment(dto) {
        if (!dto.splits || dto.splits.length === 0) {
            throw new common_1.BadRequestException('At least one split recipient must be specified');
        }
        const totalAmount = dto.splits.reduce((acc, s) => financial_math_util_1.FinancialMath.add(acc, s.amount), '0.0000');
        return await this.dataSource.transaction(async (manager) => {
            const source = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.sourceAccountNumber },
            });
            if (!source)
                throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
            const destAccounts = [];
            for (const split of dto.splits) {
                const dest = await manager.findOne(account_entity_1.Account, {
                    where: { accountNumber: split.destinationAccountNumber },
                });
                if (!dest)
                    throw new domain_exceptions_1.AccountNotFoundException(split.destinationAccountNumber);
                destAccounts.push(dest);
            }
            const allAccountIds = [source.id, ...destAccounts.map((d) => d.id)].sort();
            const balances = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId IN (:...ids)', { ids: allAccountIds })
                .getMany();
            const sourceBal = balances.find((b) => b.accountId === source.id);
            if (!sourceBal)
                throw new common_1.BadRequestException('Source balance not found');
            if (financial_math_util_1.FinancialMath.isLessThan(sourceBal.availableBalance, totalAmount)) {
                throw new domain_exceptions_1.InsufficientFundsException(sourceBal.availableBalance, totalAmount);
            }
            sourceBal.ledgerBalance = financial_math_util_1.FinancialMath.subtract(sourceBal.ledgerBalance, totalAmount);
            sourceBal.availableBalance = financial_math_util_1.FinancialMath.subtract(sourceBal.availableBalance, totalAmount);
            await manager.save(account_balance_entity_1.AccountBalance, sourceBal);
            const splitResults = [];
            const batchRef = `SP26${Math.floor(10000000 + Math.random() * 90000000)}`;
            for (let i = 0; i < dto.splits.length; i++) {
                const split = dto.splits[i];
                const destAcc = destAccounts[i];
                const destBal = balances.find((b) => b.accountId === destAcc.id);
                if (!destBal)
                    continue;
                destBal.ledgerBalance = financial_math_util_1.FinancialMath.add(destBal.ledgerBalance, split.amount);
                destBal.availableBalance = financial_math_util_1.FinancialMath.add(destBal.availableBalance, split.amount);
                await manager.save(account_balance_entity_1.AccountBalance, destBal);
                const subTxRef = `${batchRef}-${(i + 1).toString().padStart(2, '0')}`;
                const tx = manager.create(transaction_entity_1.Transaction, {
                    transactionReference: subTxRef,
                    transactionType: transaction_entity_1.TransactionType.SPLIT_PAYMENT,
                    channel: transaction_entity_1.TransactionChannel.CORE,
                    sourceAccountId: source.id,
                    destinationAccountId: destAcc.id,
                    amount: split.amount,
                    feeAmount: '0.0000',
                    currencyCode: source.currencyCode,
                    narration: split.narration || dto.batchNarration || 'Split payment distribution',
                    status: transaction_entity_1.TransactionStatus.COMPLETED,
                    completedAt: new Date(),
                });
                const savedTx = await manager.save(transaction_entity_1.Transaction, tx);
                await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                    transactionId: savedTx.id,
                    accountId: source.id,
                    entryType: ledger_entry_entity_1.EntryType.DEBIT,
                    amount: split.amount,
                    currencyCode: source.currencyCode,
                    balanceAfter: sourceBal.ledgerBalance,
                    narration: `Split Batch Debit -> ${destAcc.accountNumber}`,
                }));
                await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                    transactionId: savedTx.id,
                    accountId: destAcc.id,
                    entryType: ledger_entry_entity_1.EntryType.CREDIT,
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
    async rtgsPayment(dto) {
        return await this.dataSource.transaction(async (manager) => {
            const source = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.sourceAccountNumber },
            });
            if (!source)
                throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
            const fee = '100.0000';
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
            const txRef = `RTGS26${Math.floor(10000000 + Math.random() * 90000000)}`;
            const tx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: txRef,
                idempotencyKey: dto.idempotencyKey || null,
                transactionType: transaction_entity_1.TransactionType.RTGS_TRANSFER,
                channel: transaction_entity_1.TransactionChannel.CORE,
                sourceAccountId: source.id,
                externalAccount: dto.beneficiaryAccountNumber,
                externalBankCode: dto.beneficiaryBankSwift,
                amount: dto.amount,
                feeAmount: fee,
                currencyCode: source.currencyCode,
                narration: dto.narration || `RTGS to ${dto.beneficiaryBankSwift}: ${dto.beneficiaryName}`,
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
};
exports.PaymentsService = PaymentsService;
exports.PaymentsService = PaymentsService = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(2, (0, typeorm_1.InjectRepository)(transaction_entity_1.Transaction)),
    __param(3, (0, typeorm_1.InjectRepository)(service_entity_1.ServiceEntity)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], PaymentsService);
//# sourceMappingURL=payments.service.js.map