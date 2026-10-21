import { CardType, CardBrand } from '../../../database/entities/card-request.entity';
export declare class RequestCardDto {
    accountNumber: string;
    cardType?: CardType;
    cardBrand?: CardBrand;
    targetBranchCode: string;
}
export declare class ReplaceCardDto {
    cardNumberMasked: string;
    reason: string;
    targetBranchCode: string;
}
export declare class ListCardsDto {
    accountNumber?: string;
    customerNumber?: string;
}
