export declare class AccountLookupDto {
    accountNumber: string;
}
export declare class AccountNameLookupDto {
    accountNumber: string;
}
export declare class CreateAccountDto {
    customerNumber: string;
    accountTypeCode: string;
    branchCode: string;
    currencyCode: string;
    accountName?: string;
    initialDeposit?: string;
}
export declare class AccountListDto {
    customerNumber: string;
}
export declare class AccountEligibilityDto {
    accountNumber: string;
    amount: string;
    transactionType?: string;
}
export declare class CreateHoldDto {
    accountNumber: string;
    amount: string;
    reason: string;
    lockEvent?: string;
    fromDate?: string;
    toDate?: string;
}
export declare class ReleaseHoldDto {
    holdReference: string;
    reason?: string;
}
export declare class AccountStatementDto {
    accountNumber: string;
    startDate?: string;
    endDate?: string;
    limit?: number;
}
export declare class AccountMiniStatementDto {
    accountNumber: string;
    count?: number;
    limit?: number;
}
