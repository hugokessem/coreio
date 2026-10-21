import { Account } from './account.entity';
export declare class Branch {
    id: string;
    branchCode: string;
    branchName: string;
    address: string | null;
    city: string;
    swiftCode: string | null;
    status: string;
    accounts: Account[];
    createdAt: Date;
    updatedAt: Date;
}
