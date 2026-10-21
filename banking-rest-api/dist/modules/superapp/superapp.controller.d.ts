import { SuperAppService } from './superapp.service';
import { SuperAppSubscribeDto, SuperAppUnsubscribeDto, SuperAppStatusDto, SuperAppStatusChangeDto } from './dto/superapp.dto';
export declare class SuperAppController {
    private readonly superAppService;
    constructor(superAppService: SuperAppService);
    subscribe(dto: SuperAppSubscribeDto): Promise<{
        externalUserId: string;
        customerNumber: string;
        customerName: string;
        status: import("../../database/entities").SuperAppStatus;
        partnerCode: string;
        registeredAt: Date;
        message: string;
    }>;
    unsubscribe(dto: SuperAppUnsubscribeDto): Promise<{
        externalUserId: string;
        status: import("../../database/entities").SuperAppStatus.DEACTIVATED;
        message: string;
    }>;
    changeStatusPut(dto: SuperAppStatusChangeDto): Promise<{
        externalUserId: string;
        customerNumber: string;
        status: import("../../database/entities").SuperAppStatus;
        reason: string;
        updatedAt: Date;
        message: string;
    }>;
    changeStatusPost(dto: SuperAppStatusChangeDto): Promise<{
        externalUserId: string;
        customerNumber: string;
        status: import("../../database/entities").SuperAppStatus;
        reason: string;
        updatedAt: Date;
        message: string;
    }>;
    status(dto: SuperAppStatusDto): Promise<{
        externalUserId: string;
        customerNumber: string;
        customerName: string;
        status: import("../../database/entities").SuperAppStatus;
        partnerCode: string;
        registeredAt: Date;
        updatedAt: Date;
    }>;
    getStatusParam(externalUserId: string): Promise<{
        externalUserId: string;
        customerNumber: string;
        customerName: string;
        status: import("../../database/entities").SuperAppStatus;
        partnerCode: string;
        registeredAt: Date;
        updatedAt: Date;
    }>;
}
