import { Account } from './account.entity';
export declare class AccountBalance {
    id: string;
    accountId: string;
    account: Account;
    ledgerBalance: string;
    availableBalance: string;
    lockedBalance: string;
    currencyCode: string;
    version: number;
    updatedAt: Date;
}
