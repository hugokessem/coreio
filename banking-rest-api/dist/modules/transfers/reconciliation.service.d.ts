import { DataSource, Repository } from 'typeorm';
import { Account } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { LedgerEntry } from '../../database/entities/ledger-entry.entity';
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
export declare class ReconciliationService {
    private readonly dataSource;
    private readonly accountRepo;
    private readonly balanceRepo;
    private readonly ledgerRepo;
    private readonly logger;
    constructor(dataSource: DataSource, accountRepo: Repository<Account>, balanceRepo: Repository<AccountBalance>, ledgerRepo: Repository<LedgerEntry>);
    reconcileAllAccounts(): Promise<SystemReconciliationReport>;
    reconcileAccount(accountNumber: string): Promise<AccountReconciliationResult>;
}
