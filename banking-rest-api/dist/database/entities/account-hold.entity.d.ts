import { Account } from './account.entity';
export declare enum HoldStatus {
    ACTIVE = "ACTIVE",
    RELEASED = "RELEASED",
    SETTLED = "SETTLED",
    EXPIRED = "EXPIRED"
}
export declare class AccountHold {
    id: string;
    holdReference: string;
    accountId: string;
    account: Account;
    amount: string;
    remainingAmount: string;
    currencyCode: string;
    reason: string;
    lockEvent: string;
    status: HoldStatus;
    fromDate: string | null;
    toDate: string | null;
    releasedAt: Date | null;
    releasedReason: string | null;
    createdBy: string;
    createdAt: Date;
    updatedAt: Date;
}
