import { TransactionChannel } from '../../../database/entities/transaction.entity';
export declare class TransferVerifyDto {
    sourceAccountNumber: string;
    destinationAccountNumber: string;
    amount: string;
    currencyCode?: string;
    channel?: TransactionChannel;
}
export declare class ExecuteTransferDto {
    sourceAccountNumber: string;
    destinationAccountNumber: string;
    amount: string;
    currencyCode?: string;
    narration: string;
    idempotencyKey?: string;
    channel?: TransactionChannel;
}
export declare class TransferFromHoldDto {
    holdReference: string;
    destinationAccountNumber: string;
    narration?: string;
    idempotencyKey?: string;
}
export declare class TransactionStatusDto {
    transactionReference: string;
}
export declare class ReversalDto {
    transactionReference: string;
    reason: string;
}
