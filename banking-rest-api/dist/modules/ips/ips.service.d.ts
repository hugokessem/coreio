import { DataSource, Repository } from 'typeorm';
import { Account, AccountStatus } from '../../database/entities/account.entity';
import { Transaction, TransactionStatus } from '../../database/entities/transaction.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import { IpsAccountLookupDto, IpsTransferDto, IpsQrPaymentDto, IpsStatusDto } from './dto/ips.dto';
export declare class IpsService {
    private readonly dataSource;
    private readonly accountRepo;
    private readonly transactionRepo;
    private readonly serviceRepo;
    constructor(dataSource: DataSource, accountRepo: Repository<Account>, transactionRepo: Repository<Transaction>, serviceRepo: Repository<ServiceEntity>);
    accountLookup(dto: IpsAccountLookupDto): Promise<{
        accountNumber: string;
        accountName: string;
        currency: string;
        bankCode: string;
        bankName: string;
        branchName: string;
        status: AccountStatus;
        isEligible: boolean;
    }>;
    executeTransfer(dto: IpsTransferDto): Promise<{
        transactionReference: string;
        status: TransactionStatus;
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
        status: TransactionStatus;
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
        status: TransactionStatus;
        merchantName: string;
        merchantAccount: string;
        amount: string;
        currency: string;
        remainingBalance: string;
        completedAt: Date | null;
        message: string;
    }>;
    getStatus(dto: IpsStatusDto): Promise<{
        transactionReference: string;
        status: TransactionStatus;
        amount: string;
        fee: string;
        currency: string;
        destinationAccount: string | null;
        destinationBank: string | null;
        completedAt: Date | null;
    }>;
}
