import { Repository } from 'typeorm';
import { SuperAppUser, SuperAppStatus } from '../../database/entities/super-app-user.entity';
import { Customer } from '../../database/entities/customer.entity';
import { SuperAppSubscribeDto, SuperAppUnsubscribeDto, SuperAppStatusDto, SuperAppStatusChangeDto } from './dto/superapp.dto';
export declare class SuperAppService {
    private readonly superAppUserRepo;
    private readonly customerRepo;
    constructor(superAppUserRepo: Repository<SuperAppUser>, customerRepo: Repository<Customer>);
    subscribe(dto: SuperAppSubscribeDto): Promise<{
        externalUserId: string;
        customerNumber: string;
        customerName: string;
        status: SuperAppStatus;
        partnerCode: string;
        registeredAt: Date;
        message: string;
    }>;
    unsubscribe(dto: SuperAppUnsubscribeDto): Promise<{
        externalUserId: string;
        status: SuperAppStatus.DEACTIVATED;
        message: string;
    }>;
    getStatus(dto: SuperAppStatusDto): Promise<{
        externalUserId: string;
        customerNumber: string;
        customerName: string;
        status: SuperAppStatus;
        partnerCode: string;
        registeredAt: Date;
        updatedAt: Date;
    }>;
    getByExternalUserId(externalUserId: string): Promise<{
        externalUserId: string;
        customerNumber: string;
        customerName: string;
        status: SuperAppStatus;
        partnerCode: string;
        registeredAt: Date;
        updatedAt: Date;
    }>;
    changeStatus(dto: SuperAppStatusChangeDto): Promise<{
        externalUserId: string;
        customerNumber: string;
        status: SuperAppStatus;
        reason: string;
        updatedAt: Date;
        message: string;
    }>;
}
