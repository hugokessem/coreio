import { SuperAppStatus } from '../../../database/entities/super-app-user.entity';
export declare class SuperAppSubscribeDto {
    externalUserId: string;
    customerNumber: string;
    partnerCode?: string;
}
export declare class SuperAppUnsubscribeDto {
    externalUserId: string;
}
export declare class SuperAppStatusDto {
    externalUserId: string;
}
export declare class SuperAppStatusChangeDto {
    externalUserId: string;
    status: SuperAppStatus;
    reason?: string;
}
