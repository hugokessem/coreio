import { PaymentsService } from './payments.service';
import { BillPaymentDto, BillPaymentVerifyDto, SplitPaymentDto, RtgsPaymentDto } from './dto/payment.dto';
export declare class PaymentsController {
    private readonly paymentsService;
    constructor(paymentsService: PaymentsService);
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
        status: import("../../database/entities").TransactionStatus;
        billerCode: string;
        billReference: string;
        amount: string;
        currency: string;
        remainingAvailableBalance: string;
        completedAt: Date | null;
        message: string;
    } | {
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
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
        status: import("../../database/entities").TransactionStatus;
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
