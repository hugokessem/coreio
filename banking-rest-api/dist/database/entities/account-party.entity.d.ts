import { Account } from './account.entity';
import { Customer } from './customer.entity';
export declare enum PartyRole {
    PRIMARY_OWNER = "PRIMARY_OWNER",
    JOINT_OWNER = "JOINT_OWNER",
    AUTHORIZED_SIGNATORY = "AUTHORIZED_SIGNATORY",
    GUARANTOR = "GUARANTOR",
    BENEFICIAL_OWNER = "BENEFICIAL_OWNER"
}
export declare class AccountParty {
    id: string;
    accountId: string;
    account: Account;
    customerId: string;
    customer: Customer;
    partyRole: PartyRole;
    authorizationLevel: string;
    status: string;
    createdAt: Date;
    updatedAt: Date;
}
