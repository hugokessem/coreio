export declare enum IdempotencyStatus {
    PENDING = "PENDING",
    COMPLETED = "COMPLETED",
    FAILED = "FAILED"
}
export declare class IdempotencyRecord {
    id: string;
    idempotencyKey: string;
    requestHash: string;
    operation: string;
    callerId: string | null;
    status: IdempotencyStatus;
    transactionReference: string | null;
    responsePayload: Record<string, any> | null;
    httpStatusCode: number | null;
    createdAt: Date;
    updatedAt: Date;
}
