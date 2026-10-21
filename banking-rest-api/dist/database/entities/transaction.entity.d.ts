import { Account } from './account.entity';
import { LedgerEntry } from './ledger-entry.entity';
export declare enum TransactionType {
    INTERNAL_TRANSFER = "INTERNAL_TRANSFER",
    TRANSFER_FROM_HOLD = "TRANSFER_FROM_HOLD",
    IPS_TRANSFER = "IPS_TRANSFER",
    IPS_QR_PAYMENT = "IPS_QR_PAYMENT",
    WALLET_TRANSFER = "WALLET_TRANSFER",
    BILL_PAYMENT = "BILL_PAYMENT",
    SPLIT_PAYMENT = "SPLIT_PAYMENT",
    RTGS_TRANSFER = "RTGS_TRANSFER",
    CASH_DEPOSIT = "CASH_DEPOSIT",
    CASH_WITHDRAWAL = "CASH_WITHDRAWAL",
    FEE = "FEE",
    REVERSAL = "REVERSAL"
}
export declare enum TransactionStatus {
    PENDING = "PENDING",
    COMPLETED = "COMPLETED",
    FAILED = "FAILED",
    REVERSED = "REVERSED"
}
export declare enum TransactionChannel {
    CORE = "CORE",
    WEB = "WEB",
    MOBILE = "MOBILE",
    IPS = "IPS",
    CBE_BIRR = "CBE_BIRR",
    ATM = "ATM",
    POS = "POS",
    SYSTEM = "SYSTEM"
}
export declare class Transaction {
    id: string;
    transactionReference: string;
    idempotencyKey: string | null;
    transactionType: TransactionType;
    channel: TransactionChannel;
    sourceAccountId: string | null;
    sourceAccount: Account | null;
    destinationAccountId: string | null;
    destinationAccount: Account | null;
    externalAccount: string | null;
    externalBankCode: string | null;
    amount: string;
    feeAmount: string;
    currencyCode: string;
    exchangeRate: string;
    narration: string;
    status: TransactionStatus;
    failureReason: string | null;
    reversalOfTransactionId: string | null;
    initiatedAt: Date;
    completedAt: Date | null;
    ledgerEntries: LedgerEntry[];
    createdAt: Date;
    updatedAt: Date;
}
