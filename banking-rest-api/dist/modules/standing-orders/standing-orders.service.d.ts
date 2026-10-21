import { DataSource, Repository } from 'typeorm';
import { StandingOrder, StandingOrderStatus, StandingOrderFrequency } from '../../database/entities/standing-order.entity';
import { StandingOrderExecution, ExecutionStatus } from '../../database/entities/standing-order-execution.entity';
import { Account } from '../../database/entities/account.entity';
import { CreateStandingOrderDto, ListStandingOrdersDto, CancelStandingOrderDto, UpdateStandingOrderDto } from './dto/standing-order.dto';
import { TransfersService } from '../transfers/transfers.service';
export declare class StandingOrdersService {
    private readonly dataSource;
    private readonly standingOrderRepo;
    private readonly executionRepo;
    private readonly accountRepo;
    private readonly transfersService;
    private readonly logger;
    constructor(dataSource: DataSource, standingOrderRepo: Repository<StandingOrder>, executionRepo: Repository<StandingOrderExecution>, accountRepo: Repository<Account>, transfersService: TransfersService);
    create(dto: CreateStandingOrderDto): Promise<{
        orderReference: string;
        sourceAccountNumber: string;
        beneficiaryName: string;
        beneficiaryAccount: string;
        amount: string;
        currency: string;
        frequency: StandingOrderFrequency;
        startDate: string;
        endDate: string | null;
        nextExecutionDate: string;
        status: StandingOrderStatus;
        createdAt: Date;
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
            frequency: StandingOrderFrequency;
            startDate: string;
            endDate: string | null;
            nextExecutionDate: string;
            status: StandingOrderStatus;
            failureCount: number;
            lastExecutedAt: Date | null;
        }[];
    }>;
    listByAccount(accountNumber: string): Promise<{
        accountNumber: string;
        totalOrders: number;
        orders: {
            orderReference: string;
            beneficiaryName: string;
            beneficiaryAccount: string;
            amount: string;
            currency: string;
            frequency: StandingOrderFrequency;
            startDate: string;
            endDate: string | null;
            nextExecutionDate: string;
            status: StandingOrderStatus;
            failureCount: number;
            lastExecutedAt: Date | null;
        }[];
    }>;
    getHistory(orderReference: string): Promise<{
        orderReference: string;
        sourceAccountNumber: string;
        beneficiaryAccount: string;
        beneficiaryName: string;
        status: StandingOrderStatus;
        frequency: StandingOrderFrequency;
        totalExecutions: number;
        executions: {
            id: string;
            executionDate: string;
            amount: string;
            status: ExecutionStatus;
            transactionReference: string | null;
            failureReason: string | null;
            executedAt: Date;
        }[];
    }>;
    update(orderReference: string, dto: UpdateStandingOrderDto): Promise<{
        orderReference: string;
        amount: string;
        frequency: StandingOrderFrequency;
        endDate: string | null;
        status: StandingOrderStatus;
        updatedAt: Date;
        message: string;
    }>;
    cancel(dto: CancelStandingOrderDto | string): Promise<{
        orderReference: string;
        status: StandingOrderStatus.CANCELLED;
        cancelledAt: string;
        message: string;
    }>;
    executeDueOrders(): Promise<{
        processed: number;
        succeeded: number;
        failed: number;
    }>;
    private calculateNextDate;
}
