import { Account } from './account.entity';
import { Customer } from './customer.entity';
import { CardType, CardBrand } from './card-request.entity';
export declare enum CardStatus {
    ACTIVE = "ACTIVE",
    BLOCKED = "BLOCKED",
    EXPIRED = "EXPIRED",
    CANCELLED = "CANCELLED"
}
export declare class Card {
    id: string;
    cardNumberMasked: string;
    cardHash: string;
    cardType: CardType;
    cardBrand: CardBrand;
    accountId: string;
    account: Account;
    customerId: string;
    customer: Customer;
    status: CardStatus;
    expiryDate: string;
    dailyLimit: string;
    createdAt: Date;
    updatedAt: Date;
}
