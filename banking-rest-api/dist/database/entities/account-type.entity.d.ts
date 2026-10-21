export declare enum AccountCategory {
    SAVINGS = "SAVINGS",
    CURRENT = "CURRENT",
    FIXED_DEPOSIT = "FIXED_DEPOSIT",
    OVERDRAFT = "OVERDRAFT",
    LOAN = "LOAN"
}
export declare class AccountType {
    id: string;
    code: string;
    name: string;
    category: AccountCategory;
    minimumBalance: string;
    isInterestBearing: boolean;
    status: string;
    createdAt: Date;
    updatedAt: Date;
}
