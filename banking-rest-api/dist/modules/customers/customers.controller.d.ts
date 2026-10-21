import { CustomersService } from './customers.service';
import { CustomerLookupDto, CustomerDetailDto, CreateCustomerDto, CustomerPhoneLookupDto, AmlVerifyDto, CheckExistenceDto, CustomerEligibilityDto, FaydaVerifyDto, CompositeOnboardDto } from './dto/customer.dto';
export declare class CustomersController {
    private readonly customersService;
    constructor(customersService: CustomersService);
    lookup(dto: CustomerLookupDto): Promise<{
        customerNumber: string;
        customerType: import("../../database/entities").CustomerType;
        fullName: string;
        firstName: string | null;
        middleName: string | null;
        lastName: string | null;
        legalName: string | null;
        phoneNumber: string | null;
        email: string | null;
        address: string | null;
        dateOfBirth: string | null;
        gender: string | null;
        nationalId: string | null;
        tin: string | null;
        kycLevel: string;
        amlStatus: import("../../database/entities").AMLStatus;
        status: import("../../database/entities").CustomerStatus;
        isAuthorized: boolean;
        authorizedBy: string | null;
        authorizedAt: Date | null;
        accounts: {
            accountNumber: string;
            accountName: string;
            currency: string;
            status: import("../../database/entities").AccountStatus;
            branchCode: string | null;
        }[];
        createdAt: Date;
        updatedAt: Date;
    }>;
    detail(dto: CustomerDetailDto): Promise<{
        customerNumber: string;
        customerType: import("../../database/entities").CustomerType;
        fullName: string;
        firstName: string | null;
        middleName: string | null;
        lastName: string | null;
        legalName: string | null;
        phoneNumber: string | null;
        email: string | null;
        address: string | null;
        dateOfBirth: string | null;
        gender: string | null;
        nationalId: string | null;
        tin: string | null;
        kycLevel: string;
        amlStatus: import("../../database/entities").AMLStatus;
        status: import("../../database/entities").CustomerStatus;
        isAuthorized: boolean;
        authorizedBy: string | null;
        authorizedAt: Date | null;
        accounts: {
            accountNumber: string;
            accountName: string;
            currency: string;
            status: import("../../database/entities").AccountStatus;
            branchCode: string | null;
        }[];
        createdAt: Date;
        updatedAt: Date;
    }>;
    phoneLookup(dto: CustomerPhoneLookupDto): Promise<{
        customerNumber: string;
        customerType: import("../../database/entities").CustomerType;
        fullName: string;
        firstName: string | null;
        middleName: string | null;
        lastName: string | null;
        legalName: string | null;
        phoneNumber: string | null;
        email: string | null;
        address: string | null;
        dateOfBirth: string | null;
        gender: string | null;
        nationalId: string | null;
        tin: string | null;
        kycLevel: string;
        amlStatus: import("../../database/entities").AMLStatus;
        status: import("../../database/entities").CustomerStatus;
        isAuthorized: boolean;
        authorizedBy: string | null;
        authorizedAt: Date | null;
        accounts: {
            accountNumber: string;
            accountName: string;
            currency: string;
            status: import("../../database/entities").AccountStatus;
            branchCode: string | null;
        }[];
        createdAt: Date;
        updatedAt: Date;
    }>;
    create(dto: CreateCustomerDto, req: any): Promise<{
        customerNumber: string;
        customerType: import("../../database/entities").CustomerType;
        fullName: string;
        firstName: string | null;
        middleName: string | null;
        lastName: string | null;
        legalName: string | null;
        phoneNumber: string | null;
        email: string | null;
        address: string | null;
        dateOfBirth: string | null;
        gender: string | null;
        nationalId: string | null;
        tin: string | null;
        kycLevel: string;
        amlStatus: import("../../database/entities").AMLStatus;
        status: import("../../database/entities").CustomerStatus;
        isAuthorized: boolean;
        authorizedBy: string | null;
        authorizedAt: Date | null;
        accounts: {
            accountNumber: string;
            accountName: string;
            currency: string;
            status: import("../../database/entities").AccountStatus;
            branchCode: string | null;
        }[];
        createdAt: Date;
        updatedAt: Date;
    }>;
    fetch(body: {
        page?: number;
        limit?: number;
    }): Promise<{
        data: {
            customerNumber: string;
            customerType: import("../../database/entities").CustomerType;
            fullName: string;
            firstName: string | null;
            middleName: string | null;
            lastName: string | null;
            legalName: string | null;
            phoneNumber: string | null;
            email: string | null;
            address: string | null;
            dateOfBirth: string | null;
            gender: string | null;
            nationalId: string | null;
            tin: string | null;
            kycLevel: string;
            amlStatus: import("../../database/entities").AMLStatus;
            status: import("../../database/entities").CustomerStatus;
            isAuthorized: boolean;
            authorizedBy: string | null;
            authorizedAt: Date | null;
            accounts: {
                accountNumber: string;
                accountName: string;
                currency: string;
                status: import("../../database/entities").AccountStatus;
                branchCode: string | null;
            }[];
            createdAt: Date;
            updatedAt: Date;
        }[];
        meta: {
            page: number;
            limit: number;
            total: number;
            totalPages: number;
        };
    }>;
    deletePost(body: {
        customerNumber: string;
    }): Promise<{
        success: boolean;
        message: string;
    }>;
    amlVerify(dto: AmlVerifyDto, paramNumber?: string): Promise<{
        customerNumber: string;
        customerName: string;
        amlStatus: import("../../database/entities").AMLStatus;
        isCleared: boolean;
        screenedAt: string;
        flags: string[];
    }>;
    authorizeCustomer(customerNumber: string, body: {
        authorizerUsername?: string;
    }, req: any): Promise<{
        customerNumber: string;
        isAuthorized: boolean;
        authorizedBy: string;
        authorizedAt: Date;
        message: string;
    }>;
    checkExistence(dto: CheckExistenceDto): Promise<{
        exists: boolean;
        matchType: string | null;
        customerNumber: string | null;
        phoneNumber: string | null;
        nationalId: string | null;
        canProceedWithOnboarding: boolean;
    }>;
    checkEligibility(dto: CustomerEligibilityDto): Promise<{
        customerNumber: string;
        customerName: string;
        eligible: boolean;
        isEligible: boolean;
        amlStatus: import("../../database/entities").AMLStatus;
        kycLevel: string;
        channel: string;
        reasons: string[];
    } | {
        customerNumber: null;
        customerName: null;
        eligible: boolean;
        isEligible: boolean;
        amlStatus: string;
        channel: string;
        reasons: never[];
        kycLevel?: undefined;
    }>;
    verifyFayda(dto: FaydaVerifyDto): Promise<{
        nationalId: string | null;
        customerId: string | null;
        customerNumber: string | null;
        customerName: string;
        dateOfBirth: string | null;
        isVerified: boolean;
        verified: boolean;
        customerFlag: string;
        simulation: boolean;
        source: string;
        message: string;
    }>;
    compositeOnboard(dto: CompositeOnboardDto, req: any): Promise<{
        customerNumber: string;
        customerName: string | null;
        accountNumber: string;
        currency: string;
        branchCode: string;
        accountTypeCode: string;
        initialBalance: string;
        status: string;
        timestamp: string;
    }>;
    amlLookupByAccount(accountNumber: string): Promise<{
        accountNumber: string;
        customerNumber: string;
        customerName: string;
        amlStatus: import("../../database/entities").AMLStatus;
        isCleared: boolean;
        passedScreening: boolean;
        screenedAt: string;
        evaluation: string;
    }>;
    getAll(page?: string, limit?: string): Promise<{
        data: {
            customerNumber: string;
            customerType: import("../../database/entities").CustomerType;
            fullName: string;
            firstName: string | null;
            middleName: string | null;
            lastName: string | null;
            legalName: string | null;
            phoneNumber: string | null;
            email: string | null;
            address: string | null;
            dateOfBirth: string | null;
            gender: string | null;
            nationalId: string | null;
            tin: string | null;
            kycLevel: string;
            amlStatus: import("../../database/entities").AMLStatus;
            status: import("../../database/entities").CustomerStatus;
            isAuthorized: boolean;
            authorizedBy: string | null;
            authorizedAt: Date | null;
            accounts: {
                accountNumber: string;
                accountName: string;
                currency: string;
                status: import("../../database/entities").AccountStatus;
                branchCode: string | null;
            }[];
            createdAt: Date;
            updatedAt: Date;
        }[];
        meta: {
            page: number;
            limit: number;
            total: number;
            totalPages: number;
        };
    }>;
    getOne(customerNumber: string): Promise<{
        customerNumber: string;
        customerType: import("../../database/entities").CustomerType;
        fullName: string;
        firstName: string | null;
        middleName: string | null;
        lastName: string | null;
        legalName: string | null;
        phoneNumber: string | null;
        email: string | null;
        address: string | null;
        dateOfBirth: string | null;
        gender: string | null;
        nationalId: string | null;
        tin: string | null;
        kycLevel: string;
        amlStatus: import("../../database/entities").AMLStatus;
        status: import("../../database/entities").CustomerStatus;
        isAuthorized: boolean;
        authorizedBy: string | null;
        authorizedAt: Date | null;
        accounts: {
            accountNumber: string;
            accountName: string;
            currency: string;
            status: import("../../database/entities").AccountStatus;
            branchCode: string | null;
        }[];
        createdAt: Date;
        updatedAt: Date;
    }>;
    deleteOne(customerNumber: string): Promise<{
        success: boolean;
        message: string;
    }>;
}
