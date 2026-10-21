import { StandingOrder } from './standing-order.entity';
export declare enum ExecutionStatus {
    SUCCESS = "SUCCESS",
    FAILED = "FAILED",
    SKIPPED = "SKIPPED"
}
export declare class StandingOrderExecution {
    id: string;
    standingOrderId: string;
    standingOrder: StandingOrder;
    executionDate: string;
    amount: string;
    transactionReference: string | null;
    status: ExecutionStatus;
    failureReason: string | null;
    createdAt: Date;
}
