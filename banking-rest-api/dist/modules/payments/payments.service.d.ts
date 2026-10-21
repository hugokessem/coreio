import { DataSource, Repository } from 'typeorm';
import { Account } from '../../database/entities/account.entity';
import { Transaction, TransactionStatus } from '../../database/entities/transaction.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import { BillPaymentDto, BillPaymentVerifyDto, SplitPaymentDto, RtgsPaymentDto } from './dto/payment.dto';
export declare class PaymentsService {
    private readonly dataSource;
    private readonly accountRepo;
    private readonly transactionRepo;
    private readonly serviceRepo;
    constructor(dataSource: DataSource, accountRepo: Repository<Account>, transactionRepo: Repository<Transaction>, serviceRepo: Repository<ServiceEntity>);
    verifyBill(dto: BillPaymentVerifyDto): Promise<{
        canExecute: boolean;
        sourceAccountNumber: string;
        sourceAccountName: string;
        billerCode: string;
        billerName: string;
        billReference: string;
        amount: string;
        fee: string;
        totalDebit: string;
        currency: string;
        sourceAvailableBalance: string;
        sufficientFunds: boolean;
        status: string;
        reason: string;
    }>;
    payBill(dto: BillPaymentDto): Promise<{
        transactionReference: string;
        status: TransactionStatus;
        billerCode: string;
        billReference: string;
        amount: string;
        currency: string;
        remainingAvailableBalance: string;
        completedAt: Date | null;
        message: string;
    } | {
        transactionReference: string;
        status: TransactionStatus;
        billerCode: string;
        billReference: string;
        amount: string;
        narration: string;
        completedAt: Date | null;
        message: string;
    }>;
    splitPayment(dto: SplitPaymentDto): Promise<{
        batchReference: string;
        sourceAccountNumber: string;
        totalDistributed: string;
        splitCount: number;
        splits: any[];
        remainingAvailableBalance: string;
        completedAt: string;
        message: string;
    }>;
    rtgsPayment(dto: RtgsPaymentDto): Promise<{
        transactionReference: string;
        status: TransactionStatus;
        beneficiaryBankSwift: string;
        beneficiaryAccountNumber: string;
        beneficiaryName: string;
        amount: string;
        fee: string;
        currency: string;
        remainingBalance: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    }>;
}
