export declare class WalletCustomerLookupDto {
    phoneNumber: string;
}
export declare class AccountToWalletTransferDto {
    sourceAccountNumber: string;
    walletPhoneNumber: string;
    amount: string;
    narration?: string;
    idempotencyKey?: string;
}
export declare class WalletToAccountTransferDto {
    walletPhoneNumber: string;
    destinationAccountNumber: string;
    amount: string;
    narration?: string;
    idempotencyKey?: string;
}
export declare class WalletAgentLookupDto {
    agentIdentifier: string;
}
export declare class WalletAgentTransferDto {
    sourceAccountNumber: string;
    agentIdentifier: string;
    amount: string;
    narration?: string;
    idempotencyKey?: string;
}
