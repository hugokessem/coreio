import { Customer } from './customer.entity';
import { Branch } from './branch.entity';
import { Currency } from './currency.entity';
import { AccountType } from './account-type.entity';
import { AccountBalance } from './account-balance.entity';
import { AccountHold } from './account-hold.entity';
export declare enum AccountStatus {
    ACTIVE = "ACTIVE",
    DORMANT = "DORMANT",
    FROZEN = "FROZEN",
    RESTRICTED = "RESTRICTED",
    CLOSED = "CLOSED"
}
export declare class Account {
    id: string;
    accountNumber: string;
    customerId: string;
    customer: Customer;
    accountTypeId: string;
    accountType: AccountType;
    branchId: string;
    branch: Branch;
    currencyCode: string;
    currency: Currency;
    accountName: string;
    status: AccountStatus;
    isPostingRestricted: boolean;
    postingRestrictionReason: string | null;
    isDebitBlocked: boolean;
    isCreditBlocked: boolean;
    openedAt: Date;
    closedAt: Date | null;
    lastActivityAt: Date | null;
    balance: AccountBalance;
    holds: AccountHold[];
    createdAt: Date;
    updatedAt: Date;
}
