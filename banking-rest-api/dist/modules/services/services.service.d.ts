import { Repository } from 'typeorm';
import { ServiceEntity } from '../../database/entities/service.entity';
import { CustomerLimit } from '../../database/entities/customer-limit.entity';
import { Customer } from '../../database/entities/customer.entity';
import { ServiceDetailDto, CustomerLimitsQueryDto, AmendCustomerLimitDto } from './dto/service.dto';
export declare class ServicesService {
    private readonly serviceRepository;
    private readonly customerLimitRepository;
    private readonly customerRepository;
    constructor(serviceRepository: Repository<ServiceEntity>, customerLimitRepository: Repository<CustomerLimit>, customerRepository: Repository<Customer>);
    getDetail(dto: ServiceDetailDto): Promise<{
        serviceCode: string;
        serviceName: string;
        category: string;
        status: string;
        minAmount: string;
        maxAmount: string;
        dailyLimit: string;
        feeType: import("../../database/entities/service.entity").FeeType;
        feeValue: string;
    }>;
    getDetailByCode(serviceCode: string): Promise<{
        serviceCode: string;
        serviceName: string;
        category: string;
        status: string;
        minAmount: string;
        maxAmount: string;
        dailyLimit: string;
        feeType: import("../../database/entities/service.entity").FeeType;
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
            feeType: import("../../database/entities/service.entity").FeeType;
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
    getCustomerLimitsByNumber(customerNumber: string, serviceCode?: string): Promise<{
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
    amendCustomerLimit(dto: AmendCustomerLimitDto): Promise<{
        customerNumber: string;
        serviceCode: string;
        dailyLimit: string;
        perTransactionLimit: string;
        dailyUsed: string;
        updatedAt: Date;
        message: string;
    }>;
    listAll(): Promise<ServiceEntity[]>;
}
