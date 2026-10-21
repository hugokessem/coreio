import { Repository } from 'typeorm';
import { CardRequest, CardRequestStatus, CardType, CardBrand } from '../../database/entities/card-request.entity';
import { Card, CardStatus } from '../../database/entities/card.entity';
import { Account } from '../../database/entities/account.entity';
import { Branch } from '../../database/entities/branch.entity';
import { Customer } from '../../database/entities/customer.entity';
import { RequestCardDto, ReplaceCardDto, ListCardsDto } from './dto/card.dto';
export declare class CardsService {
    private readonly cardRequestRepo;
    private readonly cardRepo;
    private readonly accountRepo;
    private readonly branchRepo;
    private readonly customerRepo;
    constructor(cardRequestRepo: Repository<CardRequest>, cardRepo: Repository<Card>, accountRepo: Repository<Account>, branchRepo: Repository<Branch>, customerRepo: Repository<Customer>);
    requestCard(dto: RequestCardDto): Promise<{
        requestReference: string;
        accountNumber: string;
        customerName: string;
        cardType: CardType;
        cardBrand: CardBrand;
        collectionBranch: string;
        status: CardRequestStatus;
        createdAt: Date;
        message: string;
    }>;
    replaceCard(dto: ReplaceCardDto): Promise<{
        requestReference: string;
        previousCardMasked: string;
        previousCardStatus: CardStatus.BLOCKED;
        replacementStatus: CardRequestStatus;
        collectionBranch: string;
        createdAt: Date;
        message: string;
    }>;
    listCards(dto: ListCardsDto): Promise<{
        count: number;
        cards: {
            id: string;
            cardNumberMasked: string;
            cardType: CardType;
            cardBrand: CardBrand;
            accountNumber: string;
            status: CardStatus;
            expiryDate: string;
            dailyLimit: string;
            createdAt: Date;
        }[];
    }>;
}
