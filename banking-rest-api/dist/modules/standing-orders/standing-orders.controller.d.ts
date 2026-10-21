import { StandingOrdersService } from './standing-orders.service';
import { CreateStandingOrderDto, ListStandingOrdersDto, CancelStandingOrderDto, UpdateStandingOrderDto, ListStandingOrderHistoryDto } from './dto/standing-order.dto';
export declare class StandingOrdersController {
    private readonly standingOrdersService;
    constructor(standingOrdersService: StandingOrdersService);
    createCanonical(dto: CreateStandingOrderDto): Promise<{
        orderReference: string;
        sourceAccountNumber: string;
        beneficiaryName: string;
        beneficiaryAccount: string;
        amount: string;
        currency: string;
        frequency: import("../../database/entities").StandingOrderFrequency;
        startDate: string;
        endDate: string | null;
        nextExecutionDate: string;
        status: import("../../database/entities").StandingOrderStatus;
        createdAt: Date;
    }>;
    create(dto: CreateStandingOrderDto): Promise<{
        orderReference: string;
        sourceAccountNumber: string;
        beneficiaryName: string;
        beneficiaryAccount: string;
        amount: string;
        currency: string;
        frequency: import("../../database/entities").StandingOrderFrequency;
        startDate: string;
        endDate: string | null;
        nextExecutionDate: string;
        status: import("../../database/entities").StandingOrderStatus;
        createdAt: Date;
    }>;
    listQuery(accountNumber: string): Promise<{
        accountNumber: string;
        totalOrders: number;
        orders: {
            orderReference: string;
            beneficiaryName: string;
            beneficiaryAccount: string;
            amount: string;
            currency: string;
            frequency: import("../../database/entities").StandingOrderFrequency;
            startDate: string;
            endDate: string | null;
            nextExecutionDate: string;
            status: import("../../database/entities").StandingOrderStatus;
            failureCount: number;
            lastExecutedAt: Date | null;
        }[];
    }>;
    list(dto: ListStandingOrdersDto): Promise<{
        accountNumber: string;
        totalOrders: number;
        orders: {
            orderReference: string;
            beneficiaryName: string;
            beneficiaryAccount: string;
            amount: string;
            currency: string;
            frequency: import("../../database/entities").StandingOrderFrequency;
            startDate: string;
            endDate: string | null;
            nextExecutionDate: string;
            status: import("../../database/entities").StandingOrderStatus;
            failureCount: number;
            lastExecutedAt: Date | null;
        }[];
    }>;
    getHistory(orderReference: string): Promise<{
        orderReference: string;
        sourceAccountNumber: string;
        beneficiaryAccount: string;
        beneficiaryName: string;
        status: import("../../database/entities").StandingOrderStatus;
        frequency: import("../../database/entities").StandingOrderFrequency;
        totalExecutions: number;
        executions: {
            id: string;
            executionDate: string;
            amount: string;
            status: import("../../database/entities").ExecutionStatus;
            transactionReference: string | null;
            failureReason: string | null;
            executedAt: Date;
        }[];
    }>;
    getHistoryBody(dto: ListStandingOrderHistoryDto): Promise<{
        orderReference: string;
        sourceAccountNumber: string;
        beneficiaryAccount: string;
        beneficiaryName: string;
        status: import("../../database/entities").StandingOrderStatus;
        frequency: import("../../database/entities").StandingOrderFrequency;
        totalExecutions: number;
        executions: {
            id: string;
            executionDate: string;
            amount: string;
            status: import("../../database/entities").ExecutionStatus;
            transactionReference: string | null;
            failureReason: string | null;
            executedAt: Date;
        }[];
    }>;
    update(orderReference: string, dto: UpdateStandingOrderDto): Promise<{
        orderReference: string;
        amount: string;
        frequency: import("../../database/entities").StandingOrderFrequency;
        endDate: string | null;
        status: import("../../database/entities").StandingOrderStatus;
        updatedAt: Date;
        message: string;
    }>;
    deleteCanonical(orderReference: string): Promise<{
        orderReference: string;
        status: import("../../database/entities").StandingOrderStatus.CANCELLED;
        cancelledAt: string;
        message: string;
    }>;
    cancel(dto: CancelStandingOrderDto): Promise<{
        orderReference: string;
        status: import("../../database/entities").StandingOrderStatus.CANCELLED;
        cancelledAt: string;
        message: string;
    }>;
    executeDue(): Promise<{
        processed: number;
        succeeded: number;
        failed: number;
    }>;
}
