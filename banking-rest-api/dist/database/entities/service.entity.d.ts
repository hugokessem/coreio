export declare enum FeeType {
    FLAT = "FLAT",
    PERCENTAGE = "PERCENTAGE"
}
export declare class ServiceEntity {
    id: string;
    serviceCode: string;
    serviceName: string;
    category: string;
    status: string;
    minAmount: string;
    maxAmount: string;
    dailyLimit: string;
    feeType: FeeType;
    feeValue: string;
    createdAt: Date;
    updatedAt: Date;
}
