import { DataSource, Repository } from 'typeorm';
import { Account } from '../../database/entities/account.entity';
import { Customer } from '../../database/entities/customer.entity';
import { Transaction, TransactionStatus } from '../../database/entities/transaction.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import { WalletCustomerLookupDto, AccountToWalletTransferDto, WalletToAccountTransferDto, WalletAgentLookupDto, WalletAgentTransferDto } from './dto/wallet.dto';
export declare class WalletService {
    private readonly dataSource;
    private readonly accountRepo;
    private readonly customerRepo;
    private readonly transactionRepo;
    private readonly serviceRepo;
    constructor(dataSource: DataSource, accountRepo: Repository<Account>, customerRepo: Repository<Customer>, transactionRepo: Repository<Transaction>, serviceRepo: Repository<ServiceEntity>);
    customerLookup(dto: WalletCustomerLookupDto): Promise<{
        phoneNumber: string | null;
        customerNumber: string;
        customerName: string;
        status: import("../../database/entities/customer.entity").CustomerStatus;
        linkedAccounts: string[];
        walletTier: string;
        isAllowedToTransact: boolean;
    }>;
    agentLookup(dto: WalletAgentLookupDto): Promise<{
        agentIdentifier: string;
        agentName: string;
        phoneNumber: string | null;
        agentType: string;
        status: import("../../database/entities/customer.entity").CustomerStatus;
        canReceiveFloat: boolean;
        settlementAccount: string | null;
        simulation: boolean;
        message: string;
    }>;
    transferToWallet(dto: AccountToWalletTransferDto): Promise<{
        transactionReference: string;
        status: TransactionStatus;
        walletPhoneNumber: string;
        amount: string;
        fee: string;
        currency: string;
        remainingBalance: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    } | {
        transactionReference: string;
        status: TransactionStatus;
        amount: string;
        currency: string;
        walletPhoneNumber: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    }>;
    transferFromWallet(dto: WalletToAccountTransferDto): Promise<{
        transactionReference: string;
        status: TransactionStatus;
        walletPhoneNumber: string;
        destinationAccountNumber: string;
        amount: string;
        currency: string;
        newAvailableBalance: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    } | {
        transactionReference: string;
        status: TransactionStatus;
        amount: string;
        currency: string;
        destinationAccountNumber: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    }>;
    agentTransfer(dto: WalletAgentTransferDto): Promise<{
        transactionReference: string;
        status: TransactionStatus;
        walletPhoneNumber: string;
        amount: string;
        fee: string;
        currency: string;
        remainingBalance: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    } | {
        transactionReference: string;
        status: TransactionStatus;
        amount: string;
        currency: string;
        walletPhoneNumber: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    }>;
}
