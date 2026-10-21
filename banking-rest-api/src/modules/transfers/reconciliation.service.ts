import { Injectable, Logger } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository } from 'typeorm';
import { Account } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { LedgerEntry, EntryType } from '../../database/entities/ledger-entry.entity';
import { FinancialMath } from '../../common/utils/financial-math.util';

export interface AccountReconciliationResult {
  accountNumber: string;
  currency: string;
  projectedLedgerBalance: string;
  calculatedFromLedgerEntries: string;
  isBalanced: boolean;
  discrepancy: string;
  totalDebits: string;
  totalCredits: string;
  totalEntryCount: number;
}

export interface SystemReconciliationReport {
  timestamp: string;
  accountsAudited: number;
  balancedAccounts: number;
  imbalancedAccounts: number;
  totalLedgerDebits: string;
  totalLedgerCredits: string;
  systemBalanced: boolean;
  details: AccountReconciliationResult[];
}

@Injectable()
export class ReconciliationService {
  private readonly logger = new Logger(ReconciliationService.name);

  constructor(
    private readonly dataSource: DataSource,
    @InjectRepository(Account)
    private readonly accountRepo: Repository<Account>,
    @InjectRepository(AccountBalance)
    private readonly balanceRepo: Repository<AccountBalance>,
    @InjectRepository(LedgerEntry)
    private readonly ledgerRepo: Repository<LedgerEntry>,
  ) {}

  /**
   * Performs an authoritative reconciliation audit comparing account_balances projection
   * against the cumulative sum of immutable ledger entries.
   */
  async reconcileAllAccounts(): Promise<SystemReconciliationReport> {
    const accounts = await this.accountRepo.find({
      relations: ['balance'],
      order: { accountNumber: 'ASC' },
    });

    const results: AccountReconciliationResult[] = [];
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
        if (entry.entryType === EntryType.CREDIT) {
          totalCredits = FinancialMath.add(totalCredits, entry.amount);
        } else if (entry.entryType === EntryType.DEBIT) {
          totalDebits = FinancialMath.add(totalDebits, entry.amount);
        }
      }

      // Authoritative calculated balance from ledger
      const calculatedBalance = FinancialMath.subtract(totalCredits, totalDebits);
      const projectedBalance = acc.balance?.ledgerBalance || '0.0000';

      const discrepancy = FinancialMath.subtract(projectedBalance, calculatedBalance);
      const isBalanced = FinancialMath.isEqualTo(projectedBalance, calculatedBalance);

      if (!isBalanced) {
        imbalancedCount++;
        this.logger.error(
          `RECONCILIATION MISMATCH: Account ${acc.accountNumber} projected=${projectedBalance}, calculated=${calculatedBalance}`,
        );
      }

      grandTotalDebits = FinancialMath.add(grandTotalDebits, totalDebits);
      grandTotalCredits = FinancialMath.add(grandTotalCredits, totalCredits);

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

  /**
   * Reconcile a single account by its account number
   */
  async reconcileAccount(accountNumber: string): Promise<AccountReconciliationResult> {
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
      if (entry.entryType === EntryType.CREDIT) {
        totalCredits = FinancialMath.add(totalCredits, entry.amount);
      } else {
        totalDebits = FinancialMath.add(totalDebits, entry.amount);
      }
    }

    const calculatedBalance = FinancialMath.subtract(totalCredits, totalDebits);
    const projectedBalance = acc.balance?.ledgerBalance || '0.0000';
    const discrepancy = FinancialMath.subtract(projectedBalance, calculatedBalance);

    return {
      accountNumber: acc.accountNumber,
      currency: acc.currencyCode,
      projectedLedgerBalance: projectedBalance,
      calculatedFromLedgerEntries: calculatedBalance,
      isBalanced: FinancialMath.isEqualTo(projectedBalance, calculatedBalance),
      discrepancy,
      totalDebits,
      totalCredits,
      totalEntryCount: entries.length,
    };
  }
}
