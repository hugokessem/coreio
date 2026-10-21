import { Account } from './account.entity';
import { Customer } from './customer.entity';
import { Branch } from './branch.entity';
export declare enum CardType {
    DEBIT = "DEBIT",
    PREPAID = "PREPAID",
    CREDIT = "CREDIT"
}
export declare enum CardBrand {
    MASTERCARD = "MASTERCARD",
    VISA = "VISA",
    LOCAL = "LOCAL"
}
export declare enum CardRequestType {
    NEW = "NEW",
    REPLACEMENT = "REPLACEMENT",
    RENEWAL = "RENEWAL"
}
export declare enum CardRequestStatus {
    PENDING = "PENDING",
    APPROVED = "APPROVED",
    PRINTING = "PRINTING",
    DISPATCHED = "DISPATCHED",
    DELIVERED = "DELIVERED",
    REJECTED = "REJECTED"
}
export declare class CardRequest {
    id: string;
    requestReference: string;
    accountId: string;
    account: Account;
    customerId: string;
    customer: Customer;
    cardType: CardType;
    cardBrand: CardBrand;
    requestType: CardRequestType;
    status: CardRequestStatus;
    targetBranchId: string;
    targetBranch: Branch;
    rejectionReason: string | null;
    createdAt: Date;
    updatedAt: Date;
}
