export declare class ServiceDetailDto {
    serviceCode: string;
}
export declare class CustomerLimitsQueryDto {
    customerNumber: string;
    serviceCode?: string;
}
export declare class AmendCustomerLimitDto {
    customerNumber: string;
    serviceCode: string;
    dailyLimit?: string;
    perTransactionLimit?: string;
}
