import { IpsService } from './ips.service';
import { IpsAccountLookupDto, IpsTransferDto, IpsQrPaymentDto, IpsStatusDto } from './dto/ips.dto';
export declare class IpsController {
    private readonly ipsService;
    constructor(ipsService: IpsService);
    accountLookup(dto: IpsAccountLookupDto): Promise<{
        accountNumber: string;
        accountName: string;
        currency: string;
        bankCode: string;
        bankName: string;
        branchName: string;
        status: import("../../database/entities").AccountStatus;
        isEligible: boolean;
    }>;
    transfer(dto: IpsTransferDto): Promise<{
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
        endToEndId: string;
        amount: string;
        fee: string;
        currency: string;
        sourceAccountNumber: string;
        destinationBankCode: string;
        destinationAccountNumber: string;
        beneficiaryName: string;
        remainingAvailableBalance: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    } | {
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
        amount: string;
        fee: string;
        currency: string;
        destinationBankCode: string;
        destinationAccountNumber: string;
        beneficiaryName: string;
        completedAt: Date | null;
        message: string;
    }>;
    qrPayment(dto: IpsQrPaymentDto): Promise<{
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
        merchantName: string;
        merchantAccount: string;
        amount: string;
        currency: string;
        remainingBalance: string;
        completedAt: Date | null;
        message: string;
    }>;
    status(dto: IpsStatusDto): Promise<{
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
        amount: string;
        fee: string;
        currency: string;
        destinationAccount: string | null;
        destinationBank: string | null;
        completedAt: Date | null;
    }>;
    statusCheck(dto: IpsStatusDto): Promise<{
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
        amount: string;
        fee: string;
        currency: string;
        destinationAccount: string | null;
        destinationBank: string | null;
        completedAt: Date | null;
    }>;
}
