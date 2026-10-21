import { CustomerType } from '../../../database/entities/customer.entity';
export declare class CustomerLookupDto {
    customerNumber?: string;
    phoneNumber?: string;
    accountNumber?: string;
    nationalId?: string;
}
export declare class CustomerDetailDto {
    customerNumber: string;
}
export declare class CreateCustomerDto {
    customerType: CustomerType;
    firstName?: string;
    middleName?: string;
    lastName?: string;
    legalName?: string;
    fullName?: string;
    dateOfBirth?: string;
    gender?: string;
    phoneNumber: string;
    email?: string;
    address?: string;
    addressLine1?: string;
    city?: string;
    nationalId?: string;
    tin?: string;
}
export declare class CustomerPhoneLookupDto {
    phoneNumber: string;
}
export declare class AmlVerifyDto {
    customerNumber: string;
}
export declare class CheckExistenceDto {
    phoneNumber?: string;
    nationalId?: string;
}
export declare class CustomerEligibilityDto {
    customerNumber?: string;
    phoneNumber?: string;
    nationalId?: string;
    dateOfBirth?: string;
    channel?: string;
}
export declare class FaydaVerifyDto {
    nationalId?: string;
    faydaId?: string;
    fullName?: string;
    dateOfBirth?: string;
    phoneNumber?: string;
}
export declare class CompositeOnboardDto {
    customerType: CustomerType;
    firstName?: string;
    middleName?: string;
    lastName?: string;
    phoneNumber: string;
    nationalId?: string;
    accountTypeCode: string;
    branchCode: string;
    currencyCode: string;
    initialDeposit?: string;
}
