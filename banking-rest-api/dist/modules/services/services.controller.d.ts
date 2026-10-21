import { ServicesService } from './services.service';
import { ServiceDetailDto, CustomerLimitsQueryDto, AmendCustomerLimitDto } from './dto/service.dto';
export declare class ServicesController {
    private readonly servicesService;
    constructor(servicesService: ServicesService);
    listAll(): Promise<import("../../database/entities").ServiceEntity[]>;
    getDetailByCode(serviceCode: string): Promise<{
        serviceCode: string;
        serviceName: string;
        category: string;
        status: string;
        minAmount: string;
        maxAmount: string;
        dailyLimit: string;
        feeType: import("../../database/entities").FeeType;
        feeValue: string;
    }>;
    getDetail(dto: ServiceDetailDto): Promise<{
        serviceCode: string;
        serviceName: string;
        category: string;
        status: string;
        minAmount: string;
        maxAmount: string;
        dailyLimit: string;
        feeType: import("../../database/entities").FeeType;
        feeValue: string;
    }>;
    getServiceLimits(): Promise<{
        totalServices: number;
        services: {
            serviceCode: string;
            serviceName: string;
            category: string;
            minAmount: string;
            maxAmount: string;
            dailyLimit: string;
            feeType: import("../../database/entities").FeeType;
            feeValue: string;
        }[];
    }>;
    getCustomerLimits(dto: CustomerLimitsQueryDto): Promise<{
        customerNumber: string;
        customerName: string;
        customLimitsCount: number;
        limits: {
            serviceCode: string;
            dailyLimit: string;
            perTransactionLimit: string;
            dailyUsed: string;
            lastResetDate: string;
        }[];
    }>;
    getCustomerLimitsCanonical(customerNumber: string, serviceCode?: string): Promise<{
        customerNumber: string;
        customerName: string;
        customLimitsCount: number;
        limits: {
            serviceCode: string;
            dailyLimit: string;
            perTransactionLimit: string;
            dailyUsed: string;
            lastResetDate: string;
        }[];
    }>;
    amendCustomerLimitPut(dto: AmendCustomerLimitDto): Promise<{
        customerNumber: string;
        serviceCode: string;
        dailyLimit: string;
        perTransactionLimit: string;
        dailyUsed: string;
        updatedAt: Date;
        message: string;
    }>;
    amendCustomerLimitPost(dto: AmendCustomerLimitDto): Promise<{
        customerNumber: string;
        serviceCode: string;
        dailyLimit: string;
        perTransactionLimit: string;
        dailyUsed: string;
        updatedAt: Date;
        message: string;
    }>;
}
