import { CardsService } from './cards.service';
import { RequestCardDto, ReplaceCardDto, ListCardsDto } from './dto/card.dto';
export declare class CardsController {
    private readonly cardsService;
    constructor(cardsService: CardsService);
    requestCard(dto: RequestCardDto): Promise<{
        requestReference: string;
        accountNumber: string;
        customerName: string;
        cardType: import("../../database/entities").CardType;
        cardBrand: import("../../database/entities").CardBrand;
        collectionBranch: string;
        status: import("../../database/entities").CardRequestStatus;
        createdAt: Date;
        message: string;
    }>;
    replaceCard(dto: ReplaceCardDto): Promise<{
        requestReference: string;
        previousCardMasked: string;
        previousCardStatus: import("../../database/entities").CardStatus.BLOCKED;
        replacementStatus: import("../../database/entities").CardRequestStatus;
        collectionBranch: string;
        createdAt: Date;
        message: string;
    }>;
    listCards(dto: ListCardsDto): Promise<{
        count: number;
        cards: {
            id: string;
            cardNumberMasked: string;
            cardType: import("../../database/entities").CardType;
            cardBrand: import("../../database/entities").CardBrand;
            accountNumber: string;
            status: import("../../database/entities").CardStatus;
            expiryDate: string;
            dailyLimit: string;
            createdAt: Date;
        }[];
    }>;
    getByAccount(accountNumber: string): Promise<{
        count: number;
        cards: {
            id: string;
            cardNumberMasked: string;
            cardType: import("../../database/entities").CardType;
            cardBrand: import("../../database/entities").CardBrand;
            accountNumber: string;
            status: import("../../database/entities").CardStatus;
            expiryDate: string;
            dailyLimit: string;
            createdAt: Date;
        }[];
    }>;
}
