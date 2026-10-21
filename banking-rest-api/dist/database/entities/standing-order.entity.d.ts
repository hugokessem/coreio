import { Account } from './account.entity';
export declare enum StandingOrderFrequency {
    DAILY = "DAILY",
    WEEKLY = "WEEKLY",
    MONTHLY = "MONTHLY",
    QUARTERLY = "QUARTERLY",
    ANNUALLY = "ANNUALLY"
}
export declare enum StandingOrderStatus {
    ACTIVE = "ACTIVE",
    PAUSED = "PAUSED",
    CANCELLED = "CANCELLED",
    COMPLETED = "COMPLETED"
}
export declare class StandingOrder {
    id: string;
    orderReference: string;
    sourceAccountId: string;
    sourceAccount: Account;
    destinationAccountId: string | null;
    destinationAccount: Account | null;
    beneficiaryName: string;
    beneficiaryAccount: string;
    beneficiaryBankCode: string | null;
    amount: string;
    currencyCode: string;
    frequency: StandingOrderFrequency;
    startDate: string;
    endDate: string | null;
    nextExecutionDate: string;
    status: StandingOrderStatus;
    failureCount: number;
    lastExecutedAt: Date | null;
    createdAt: Date;
    updatedAt: Date;
}
