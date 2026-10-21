import { Account } from './account.entity';
import { CustomerLimit } from './customer-limit.entity';
export declare enum CustomerType {
    INDIVIDUAL = "INDIVIDUAL",
    CORPORATE = "CORPORATE"
}
export declare enum CustomerStatus {
    ACTIVE = "ACTIVE",
    INACTIVE = "INACTIVE",
    SUSPENDED = "SUSPENDED",
    DECEASED = "DECEASED"
}
export declare enum AMLStatus {
    CLEAR = "CLEAR",
    SUSPECT = "SUSPECT",
    BLOCKED = "BLOCKED"
}
export declare class Customer {
    id: string;
    customerNumber: string;
    customerType: CustomerType;
    firstName: string | null;
    middleName: string | null;
    lastName: string | null;
    legalName: string | null;
    dateOfBirth: string | null;
    gender: string | null;
    phoneNumber: string | null;
    email: string | null;
    address: string | null;
    nationalId: string | null;
    tin: string | null;
    kycLevel: string;
    amlStatus: AMLStatus;
    status: CustomerStatus;
    isAuthorized: boolean;
    createdBy: string;
    authorizedBy: string | null;
    authorizedAt: Date | null;
    accounts: Account[];
    limits: CustomerLimit[];
    createdAt: Date;
    updatedAt: Date;
    get fullName(): string;
}
