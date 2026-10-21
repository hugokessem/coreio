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
exports.TransfersService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const account_hold_entity_1 = require("../../database/entities/account-hold.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const service_entity_1 = require("../../database/entities/service.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
const financial_math_util_1 = require("../../common/utils/financial-math.util");
const idempotency_service_1 = require("../../common/services/idempotency.service");
let TransfersService = class TransfersService {
    dataSource;
    accountRepository;
    transactionRepository;
    serviceRepository;
    idempotencyService;
    constructor(dataSource, accountRepository, transactionRepository, serviceRepository, idempotencyService) {
        this.dataSource = dataSource;
        this.accountRepository = accountRepository;
        this.transactionRepository = transactionRepository;
        this.serviceRepository = serviceRepository;
        this.idempotencyService = idempotencyService;
    }
    async verify(dto) {
        if (dto.sourceAccountNumber === dto.destinationAccountNumber) {
            throw new common_1.BadRequestException('Source and destination accounts must be different');
        }
        const source = await this.accountRepository.findOne({
            where: { accountNumber: dto.sourceAccountNumber },
            relations: ['balance', 'currency', 'customer'],
        });
        if (!source)
            throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
        const destination = await this.accountRepository.findOne({
            where: { accountNumber: dto.destinationAccountNumber },
            relations: ['currency', 'customer', 'branch'],
        });
        if (!destination)
            throw new domain_exceptions_1.AccountNotFoundException(dto.destinationAccountNumber);
        if (source.currencyCode !== destination.currencyCode) {
            throw new domain_exceptions_1.CurrencyMismatchException(source.currencyCode, destination.currencyCode);
        }
        if (source.status !== account_entity_1.AccountStatus.ACTIVE) {
            throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, `Source account status is ${source.status}`);
        }
        if (destination.status !== account_entity_1.AccountStatus.ACTIVE) {
            throw new domain_exceptions_1.AccountRestrictedException(destination.accountNumber, `Destination account status is ${destination.status}`);
        }
        if (source.isDebitBlocked) {
            throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, 'Account is blocked for debit operations');
        }
        if (destination.isCreditBlocked) {
            throw new domain_exceptions_1.AccountRestrictedException(destination.accountNumber, 'Account is blocked for credit operations');
        }
        const service = await this.serviceRepository.findOne({ where: { serviceCode: 'FT_INT' } });
        const fee = service ? service.feeValue : '0.0000';
        const totalRequired = financial_math_util_1.FinancialMath.add(dto.amount, fee);
        const available = source.balance?.availableBalance || '0.0000';
        const hasSufficientFunds = financial_math_util_1.FinancialMath.isGreaterThanOrEqualTo(available, totalRequired);
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
    async execute(dto) {
        if (dto.sourceAccountNumber === dto.destinationAccountNumber) {
            throw new common_1.BadRequestException('Cannot transfer funds to the same account');
        }
        if (!financial_math_util_1.FinancialMath.isPositive(dto.amount)) {
            throw new common_1.BadRequestException('Transfer amount must be strictly positive');
        }
        if (dto.idempotencyKey) {
            const check = await this.idempotencyService.checkOrRegister(dto.idempotencyKey, 'FUND_TRANSFER', {
                source: dto.sourceAccountNumber,
                dest: dto.destinationAccountNumber,
                amount: dto.amount,
            });
            if (check.isReplay && check.cachedResponse) {
                return check.cachedResponse;
            }
        }
        try {
            return await this.dataSource.transaction(async (manager) => {
                const source = await manager.findOne(account_entity_1.Account, {
                    where: { accountNumber: dto.sourceAccountNumber },
                    relations: ['customer', 'currency'],
                });
                if (!source)
                    throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
                const destination = await manager.findOne(account_entity_1.Account, {
                    where: { accountNumber: dto.destinationAccountNumber },
                    relations: ['customer', 'currency'],
                });
                if (!destination)
                    throw new domain_exceptions_1.AccountNotFoundException(dto.destinationAccountNumber);
                if (source.currencyCode !== destination.currencyCode) {
                    throw new domain_exceptions_1.CurrencyMismatchException(source.currencyCode, destination.currencyCode);
                }
                if (source.status !== account_entity_1.AccountStatus.ACTIVE) {
                    throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, `Source account status is ${source.status}`);
                }
                if (destination.status !== account_entity_1.AccountStatus.ACTIVE) {
                    throw new domain_exceptions_1.AccountRestrictedException(destination.accountNumber, `Destination account status is ${destination.status}`);
                }
                if (source.isDebitBlocked) {
                    throw new domain_exceptions_1.AccountRestrictedException(source.accountNumber, 'Account is blocked for debit operations');
                }
                if (destination.isCreditBlocked) {
                    throw new domain_exceptions_1.AccountRestrictedException(destination.accountNumber, 'Account is blocked for credit operations');
                }
                const service = await manager.findOne(service_entity_1.ServiceEntity, { where: { serviceCode: 'FT_INT' } });
                const fee = service ? service.feeValue : '0.0000';
                const totalDebit = financial_math_util_1.FinancialMath.add(dto.amount, fee);
                const lockIds = [source.id, destination.id].sort();
                const balances = await manager
                    .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                    .setLock('pessimistic_write')
                    .where('bal.accountId IN (:...ids)', { ids: lockIds })
                    .getMany();
                const sourceBalance = balances.find((b) => b.accountId === source.id);
                const destBalance = balances.find((b) => b.accountId === destination.id);
                if (!sourceBalance || !destBalance) {
                    throw new common_1.BadRequestException('Failed to lock account balances');
                }
                if (financial_math_util_1.FinancialMath.isLessThan(sourceBalance.availableBalance, totalDebit)) {
                    throw new domain_exceptions_1.InsufficientFundsException(sourceBalance.availableBalance, totalDebit);
                }
                const randomSuffix = Math.floor(10000000 + Math.random() * 90000000).toString();
                const txRef = `FT26${randomSuffix}`;
                sourceBalance.ledgerBalance = financial_math_util_1.FinancialMath.subtract(sourceBalance.ledgerBalance, totalDebit);
                sourceBalance.availableBalance = financial_math_util_1.FinancialMath.subtract(sourceBalance.availableBalance, totalDebit);
                await manager.save(account_balance_entity_1.AccountBalance, sourceBalance);
                destBalance.ledgerBalance = financial_math_util_1.FinancialMath.add(destBalance.ledgerBalance, dto.amount);
                destBalance.availableBalance = financial_math_util_1.FinancialMath.add(destBalance.availableBalance, dto.amount);
                await manager.save(account_balance_entity_1.AccountBalance, destBalance);
                const tx = manager.create(transaction_entity_1.Transaction, {
                    transactionReference: txRef,
                    idempotencyKey: dto.idempotencyKey || null,
                    transactionType: transaction_entity_1.TransactionType.INTERNAL_TRANSFER,
                    channel: dto.channel || transaction_entity_1.TransactionChannel.MOBILE,
                    sourceAccountId: source.id,
                    destinationAccountId: destination.id,
                    amount: dto.amount,
                    feeAmount: fee,
                    currencyCode: source.currencyCode,
                    exchangeRate: '1.000000',
                    narration: dto.narration,
                    status: transaction_entity_1.TransactionStatus.COMPLETED,
                    completedAt: new Date(),
                });
                const savedTx = await manager.save(transaction_entity_1.Transaction, tx);
                const debitEntry = manager.create(ledger_entry_entity_1.LedgerEntry, {
                    transactionId: savedTx.id,
                    accountId: source.id,
                    entryType: ledger_entry_entity_1.EntryType.DEBIT,
                    amount: totalDebit,
                    currencyCode: source.currencyCode,
                    balanceAfter: sourceBalance.ledgerBalance,
                    narration: `TRF to ${destination.accountNumber} - ${dto.narration}`,
                });
                await manager.save(ledger_entry_entity_1.LedgerEntry, debitEntry);
                const creditEntry = manager.create(ledger_entry_entity_1.LedgerEntry, {
                    transactionId: savedTx.id,
                    accountId: destination.id,
                    entryType: ledger_entry_entity_1.EntryType.CREDIT,
                    amount: dto.amount,
                    currencyCode: destination.currencyCode,
                    balanceAfter: destBalance.ledgerBalance,
                    narration: `TRF from ${source.accountNumber} - ${dto.narration}`,
                });
                await manager.save(ledger_entry_entity_1.LedgerEntry, creditEntry);
                savedTx.sourceAccount = source;
                savedTx.destinationAccount = destination;
                const response = this.mapToTransactionResponse(savedTx, 'Transfer completed successfully');
                await this.idempotencyService.recordSuccess(dto.idempotencyKey, 'FUND_TRANSFER', {
                    source: dto.sourceAccountNumber,
                    dest: dto.destinationAccountNumber,
                    amount: dto.amount,
                }, savedTx.transactionReference, response, 200, undefined, manager);
                return response;
            });
        }
        catch (err) {
            if (dto.idempotencyKey && (err.code === '23505' || err.message?.includes('duplicate key') || err.message?.includes('unique constraint'))) {
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
    async transferFromHold(dto) {
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
            const hold = await manager.findOne(account_hold_entity_1.AccountHold, {
                where: { holdReference: dto.holdReference, status: account_hold_entity_1.HoldStatus.ACTIVE },
                relations: ['account'],
            });
            if (!hold)
                throw new domain_exceptions_1.HoldNotFoundException(dto.holdReference);
            const source = hold.account;
            const destination = await manager.findOne(account_entity_1.Account, {
                where: { accountNumber: dto.destinationAccountNumber },
                relations: ['currency'],
            });
            if (!destination)
                throw new domain_exceptions_1.AccountNotFoundException(dto.destinationAccountNumber);
            if (source.id === destination.id) {
                throw new common_1.BadRequestException('Cannot settle hold to the same account');
            }
            const lockIds = [source.id, destination.id].sort();
            const balances = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId IN (:...ids)', { ids: lockIds })
                .getMany();
            const sourceBalance = balances.find((b) => b.accountId === source.id);
            const destBalance = balances.find((b) => b.accountId === destination.id);
            if (!sourceBalance || !destBalance) {
                throw new common_1.BadRequestException('Failed to lock account balances');
            }
            sourceBalance.ledgerBalance = financial_math_util_1.FinancialMath.subtract(sourceBalance.ledgerBalance, hold.amount);
            sourceBalance.lockedBalance = financial_math_util_1.FinancialMath.subtract(sourceBalance.lockedBalance, hold.amount);
            await manager.save(account_balance_entity_1.AccountBalance, sourceBalance);
            destBalance.ledgerBalance = financial_math_util_1.FinancialMath.add(destBalance.ledgerBalance, hold.amount);
            destBalance.availableBalance = financial_math_util_1.FinancialMath.add(destBalance.availableBalance, hold.amount);
            await manager.save(account_balance_entity_1.AccountBalance, destBalance);
            hold.status = account_hold_entity_1.HoldStatus.SETTLED;
            hold.releasedAt = new Date();
            hold.releasedReason = `Settled to ${destination.accountNumber}`;
            await manager.save(account_hold_entity_1.AccountHold, hold);
            const randomSuffix = Math.floor(10000000 + Math.random() * 90000000).toString();
            const txRef = `FT26${randomSuffix}`;
            const tx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: txRef,
                idempotencyKey: dto.idempotencyKey || null,
                transactionType: transaction_entity_1.TransactionType.TRANSFER_FROM_HOLD,
                channel: transaction_entity_1.TransactionChannel.CORE,
                sourceAccountId: source.id,
                destinationAccountId: destination.id,
                amount: hold.amount,
                feeAmount: '0.0000',
                currencyCode: hold.currencyCode,
                narration: dto.narration || `Hold settlement ${hold.holdReference}`,
                status: transaction_entity_1.TransactionStatus.COMPLETED,
                completedAt: new Date(),
            });
            const savedTx = await manager.save(transaction_entity_1.Transaction, tx);
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedTx.id,
                accountId: source.id,
                entryType: ledger_entry_entity_1.EntryType.DEBIT,
                amount: hold.amount,
                currencyCode: hold.currencyCode,
                balanceAfter: sourceBalance.ledgerBalance,
                narration: `Hold Settle to ${destination.accountNumber} [${hold.holdReference}]`,
            }));
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedTx.id,
                accountId: destination.id,
                entryType: ledger_entry_entity_1.EntryType.CREDIT,
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
    async getStatus(dto) {
        return this.getByReference(dto.transactionReference);
    }
    async getByReference(transactionReference) {
        const tx = await this.transactionRepository.findOne({
            where: { transactionReference },
            relations: ['sourceAccount', 'destinationAccount', 'ledgerEntries'],
        });
        if (!tx) {
            throw new common_1.BadRequestException(`Transaction with reference '${transactionReference}' not found`);
        }
        return this.mapToTransactionResponse(tx);
    }
    async reverse(dto) {
        return await this.dataSource.transaction(async (manager) => {
            const origTx = await manager.findOne(transaction_entity_1.Transaction, {
                where: { transactionReference: dto.transactionReference },
                relations: ['sourceAccount', 'destinationAccount'],
            });
            if (!origTx) {
                throw new common_1.BadRequestException(`Transaction '${dto.transactionReference}' not found`);
            }
            if (origTx.status !== transaction_entity_1.TransactionStatus.COMPLETED) {
                throw new common_1.BadRequestException(`Cannot reverse transaction with status '${origTx.status}'`);
            }
            if (origTx.reversalOfTransactionId) {
                throw new common_1.BadRequestException('This transaction is already a reversal');
            }
            const existingReversal = await manager.findOne(transaction_entity_1.Transaction, {
                where: { reversalOfTransactionId: origTx.id },
            });
            if (existingReversal) {
                throw new common_1.BadRequestException(`Transaction has already been reversed by reference ${existingReversal.transactionReference}`);
            }
            if (!origTx.sourceAccountId || !origTx.destinationAccountId) {
                throw new common_1.BadRequestException('Cannot reverse single-legged cash transaction automatically');
            }
            const lockIds = [origTx.sourceAccountId, origTx.destinationAccountId].sort();
            const balances = await manager
                .createQueryBuilder(account_balance_entity_1.AccountBalance, 'bal')
                .setLock('pessimistic_write')
                .where('bal.accountId IN (:...ids)', { ids: lockIds })
                .getMany();
            const sourceBal = balances.find((b) => b.accountId === origTx.sourceAccountId);
            const destBal = balances.find((b) => b.accountId === origTx.destinationAccountId);
            if (!sourceBal || !destBal) {
                throw new common_1.BadRequestException('Failed to lock accounts for reversal');
            }
            if (financial_math_util_1.FinancialMath.isLessThan(destBal.availableBalance, origTx.amount)) {
                throw new domain_exceptions_1.InsufficientFundsException(destBal.availableBalance, origTx.amount);
            }
            const totalToRefund = financial_math_util_1.FinancialMath.add(origTx.amount, origTx.feeAmount);
            destBal.ledgerBalance = financial_math_util_1.FinancialMath.subtract(destBal.ledgerBalance, origTx.amount);
            destBal.availableBalance = financial_math_util_1.FinancialMath.subtract(destBal.availableBalance, origTx.amount);
            await manager.save(account_balance_entity_1.AccountBalance, destBal);
            sourceBal.ledgerBalance = financial_math_util_1.FinancialMath.add(sourceBal.ledgerBalance, totalToRefund);
            sourceBal.availableBalance = financial_math_util_1.FinancialMath.add(sourceBal.availableBalance, totalToRefund);
            await manager.save(account_balance_entity_1.AccountBalance, sourceBal);
            origTx.status = transaction_entity_1.TransactionStatus.REVERSED;
            await manager.save(transaction_entity_1.Transaction, origTx);
            const revRef = `REV-${origTx.transactionReference}`;
            const revTx = manager.create(transaction_entity_1.Transaction, {
                transactionReference: revRef,
                transactionType: transaction_entity_1.TransactionType.REVERSAL,
                channel: transaction_entity_1.TransactionChannel.SYSTEM,
                sourceAccountId: origTx.destinationAccountId,
                destinationAccountId: origTx.sourceAccountId,
                amount: origTx.amount,
                feeAmount: origTx.feeAmount,
                currencyCode: origTx.currencyCode,
                narration: `Reversal of ${origTx.transactionReference}: ${dto.reason}`,
                status: transaction_entity_1.TransactionStatus.COMPLETED,
                reversalOfTransactionId: origTx.id,
                completedAt: new Date(),
            });
            const savedRevTx = await manager.save(transaction_entity_1.Transaction, revTx);
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedRevTx.id,
                accountId: origTx.destinationAccountId,
                entryType: ledger_entry_entity_1.EntryType.DEBIT,
                amount: origTx.amount,
                currencyCode: origTx.currencyCode,
                balanceAfter: destBal.ledgerBalance,
                narration: `Reversal Debit: ${dto.reason}`,
            }));
            await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                transactionId: savedRevTx.id,
                accountId: origTx.sourceAccountId,
                entryType: ledger_entry_entity_1.EntryType.CREDIT,
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
    mapToTransactionResponse(tx, message) {
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
            message: message || (tx.status === transaction_entity_1.TransactionStatus.COMPLETED ? 'Transaction completed successfully' : 'Transaction failed'),
        };
    }
};
exports.TransfersService = TransfersService;
exports.TransfersService = TransfersService = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(2, (0, typeorm_1.InjectRepository)(transaction_entity_1.Transaction)),
    __param(3, (0, typeorm_1.InjectRepository)(service_entity_1.ServiceEntity)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        idempotency_service_1.IdempotencyService])
], TransfersService);
//# sourceMappingURL=transfers.service.js.map