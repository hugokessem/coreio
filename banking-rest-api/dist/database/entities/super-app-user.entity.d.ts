import { Customer } from './customer.entity';
export declare enum SuperAppStatus {
    ACTIVE = "ACTIVE",
    SUSPENDED = "SUSPENDED",
    DEACTIVATED = "DEACTIVATED"
}
export declare class SuperAppUser {
    id: string;
    externalUserId: string;
    customerId: string;
    customer: Customer;
    partnerCode: string;
    status: SuperAppStatus;
    registeredAt: Date;
    updatedAt: Date;
}
