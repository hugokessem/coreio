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
var ReconciliationService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReconciliationService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const financial_math_util_1 = require("../../common/utils/financial-math.util");
let ReconciliationService = ReconciliationService_1 = class ReconciliationService {
    dataSource;
    accountRepo;
    balanceRepo;
    ledgerRepo;
    logger = new common_1.Logger(ReconciliationService_1.name);
    constructor(dataSource, accountRepo, balanceRepo, ledgerRepo) {
        this.dataSource = dataSource;
        this.accountRepo = accountRepo;
        this.balanceRepo = balanceRepo;
        this.ledgerRepo = ledgerRepo;
    }
    async reconcileAllAccounts() {
        const accounts = await this.accountRepo.find({
            relations: ['balance'],
            order: { accountNumber: 'ASC' },
        });
        const results = [];
        let grandTotalDebits = '0.0000';
        let grandTotalCredits = '0.0000';
        let imbalancedCount = 0;
        for (const acc of accounts) {
            const entries = await this.ledgerRepo.find({
                where: { accountId: acc.id },
            });
            let totalCredits = '0.0000';
            let totalDebits = '0.0000';
            for (const entry of entries) {
                if (entry.entryType === ledger_entry_entity_1.EntryType.CREDIT) {
                    totalCredits = financial_math_util_1.FinancialMath.add(totalCredits, entry.amount);
                }
                else if (entry.entryType === ledger_entry_entity_1.EntryType.DEBIT) {
                    totalDebits = financial_math_util_1.FinancialMath.add(totalDebits, entry.amount);
                }
            }
            const calculatedBalance = financial_math_util_1.FinancialMath.subtract(totalCredits, totalDebits);
            const projectedBalance = acc.balance?.ledgerBalance || '0.0000';
            const discrepancy = financial_math_util_1.FinancialMath.subtract(projectedBalance, calculatedBalance);
            const isBalanced = financial_math_util_1.FinancialMath.isEqualTo(projectedBalance, calculatedBalance);
            if (!isBalanced) {
                imbalancedCount++;
                this.logger.error(`RECONCILIATION MISMATCH: Account ${acc.accountNumber} projected=${projectedBalance}, calculated=${calculatedBalance}`);
            }
            grandTotalDebits = financial_math_util_1.FinancialMath.add(grandTotalDebits, totalDebits);
            grandTotalCredits = financial_math_util_1.FinancialMath.add(grandTotalCredits, totalCredits);
            results.push({
                accountNumber: acc.accountNumber,
                currency: acc.currencyCode,
                projectedLedgerBalance: projectedBalance,
                calculatedFromLedgerEntries: calculatedBalance,
                isBalanced,
                discrepancy,
                totalDebits,
                totalCredits,
                totalEntryCount: entries.length,
            });
        }
        return {
            timestamp: new Date().toISOString(),
            accountsAudited: accounts.length,
            balancedAccounts: accounts.length - imbalancedCount,
            imbalancedAccounts: imbalancedCount,
            totalLedgerDebits: grandTotalDebits,
            totalLedgerCredits: grandTotalCredits,
            systemBalanced: imbalancedCount === 0,
            details: results,
        };
    }
    async reconcileAccount(accountNumber) {
        const acc = await this.accountRepo.findOne({
            where: { accountNumber },
            relations: ['balance'],
        });
        if (!acc) {
            throw new Error(`Account ${accountNumber} not found`);
        }
        const entries = await this.ledgerRepo.find({
            where: { accountId: acc.id },
        });
        let totalCredits = '0.0000';
        let totalDebits = '0.0000';
        for (const entry of entries) {
            if (entry.entryType === ledger_entry_entity_1.EntryType.CREDIT) {
                totalCredits = financial_math_util_1.FinancialMath.add(totalCredits, entry.amount);
            }
            else {
                totalDebits = financial_math_util_1.FinancialMath.add(totalDebits, entry.amount);
            }
        }
        const calculatedBalance = financial_math_util_1.FinancialMath.subtract(totalCredits, totalDebits);
        const projectedBalance = acc.balance?.ledgerBalance || '0.0000';
        const discrepancy = financial_math_util_1.FinancialMath.subtract(projectedBalance, calculatedBalance);
        return {
            accountNumber: acc.accountNumber,
            currency: acc.currencyCode,
            projectedLedgerBalance: projectedBalance,
            calculatedFromLedgerEntries: calculatedBalance,
            isBalanced: financial_math_util_1.FinancialMath.isEqualTo(projectedBalance, calculatedBalance),
            discrepancy,
            totalDebits,
            totalCredits,
            totalEntryCount: entries.length,
        };
    }
};
exports.ReconciliationService = ReconciliationService;
exports.ReconciliationService = ReconciliationService = ReconciliationService_1 = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(2, (0, typeorm_1.InjectRepository)(account_balance_entity_1.AccountBalance)),
    __param(3, (0, typeorm_1.InjectRepository)(ledger_entry_entity_1.LedgerEntry)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], ReconciliationService);
//# sourceMappingURL=reconciliation.service.js.map