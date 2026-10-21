export declare class IpsAccountLookupDto {
    accountNumber: string;
    bankCode?: string;
}
export declare class IpsTransferDto {
    sourceAccountNumber: string;
    destinationBankCode: string;
    destinationAccountNumber: string;
    beneficiaryName: string;
    amount: string;
    narration?: string;
    endToEndId?: string;
    idempotencyKey?: string;
}
export declare class IpsQrPaymentDto {
    sourceAccountNumber: string;
    qrPayload: string;
    merchantAccount: string;
    merchantName: string;
    amount: string;
    narration?: string;
    idempotencyKey?: string;
}
export declare class IpsStatusDto {
    transactionReference: string;
}
