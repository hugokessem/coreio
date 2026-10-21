import { StandingOrderFrequency, StandingOrderStatus } from '../../../database/entities/standing-order.entity';
export declare class CreateStandingOrderDto {
    sourceAccountNumber: string;
    beneficiaryAccount: string;
    beneficiaryName: string;
    beneficiaryBankCode?: string;
    amount: string;
    currencyCode?: string;
    frequency: StandingOrderFrequency;
    startDate: string;
    endDate?: string;
}
export declare class ListStandingOrdersDto {
    accountNumber: string;
}
export declare class CancelStandingOrderDto {
    orderReference: string;
}
export declare class UpdateStandingOrderDto {
    amount?: string;
    frequency?: StandingOrderFrequency;
    endDate?: string;
    status?: StandingOrderStatus;
}
export declare class ListStandingOrderHistoryDto {
    orderReference: string;
}
