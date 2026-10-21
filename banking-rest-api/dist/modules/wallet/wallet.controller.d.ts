import { WalletService } from './wallet.service';
import { WalletCustomerLookupDto, AccountToWalletTransferDto, WalletToAccountTransferDto, WalletAgentLookupDto, WalletAgentTransferDto } from './dto/wallet.dto';
export declare class WalletController {
    private readonly walletService;
    constructor(walletService: WalletService);
    customerLookup(dto: WalletCustomerLookupDto): Promise<{
        phoneNumber: string | null;
        customerNumber: string;
        customerName: string;
        status: import("../../database/entities").CustomerStatus;
        linkedAccounts: string[];
        walletTier: string;
        isAllowedToTransact: boolean;
    }>;
    transfer(dto: AccountToWalletTransferDto): Promise<{
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
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
        status: import("../../database/entities").TransactionStatus;
        amount: string;
        currency: string;
        walletPhoneNumber: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    }>;
    transferToWallet(dto: AccountToWalletTransferDto): Promise<{
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
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
        status: import("../../database/entities").TransactionStatus;
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
        status: import("../../database/entities").TransactionStatus;
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
        status: import("../../database/entities").TransactionStatus;
        amount: string;
        currency: string;
        destinationAccountNumber: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    }>;
    agentLookup(dto: WalletAgentLookupDto): Promise<{
        agentIdentifier: string;
        agentName: string;
        phoneNumber: string | null;
        agentType: string;
        status: import("../../database/entities").CustomerStatus;
        canReceiveFloat: boolean;
        settlementAccount: string | null;
        simulation: boolean;
        message: string;
    }>;
    agentTransfer(dto: WalletAgentTransferDto): Promise<{
        transactionReference: string;
        status: import("../../database/entities").TransactionStatus;
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
        status: import("../../database/entities").TransactionStatus;
        amount: string;
        currency: string;
        walletPhoneNumber: string;
        completedAt: Date | null;
        simulation: boolean;
        externalSettlementConfirmed: boolean;
        message: string;
    }>;
}
