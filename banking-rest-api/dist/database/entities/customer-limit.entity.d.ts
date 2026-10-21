import { Customer } from './customer.entity';
export declare class CustomerLimit {
    id: string;
    customerId: string;
    customer: Customer;
    serviceCode: string;
    dailyLimit: string;
    perTransactionLimit: string;
    dailyUsed: string;
    lastResetDate: string;
    createdAt: Date;
    updatedAt: Date;
}
