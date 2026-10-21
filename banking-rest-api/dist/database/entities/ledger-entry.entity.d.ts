import { Transaction } from './transaction.entity';
import { Account } from './account.entity';
export declare enum EntryType {
    DEBIT = "DEBIT",
    CREDIT = "CREDIT"
}
export declare class LedgerEntry {
    id: string;
    transactionId: string;
    transaction: Transaction;
    accountId: string;
    account: Account;
    entryType: EntryType;
    amount: string;
    currencyCode: string;
    balanceAfter: string;
    narration: string;
    valueDate: string;
    postedAt: Date;
}
