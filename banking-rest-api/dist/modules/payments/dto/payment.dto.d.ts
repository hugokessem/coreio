export declare class BillPaymentDto {
    sourceAccountNumber: string;
    billerCode: string;
    billReference: string;
    amount: string;
    narration?: string;
    idempotencyKey?: string;
}
export declare class BillPaymentVerifyDto {
    sourceAccountNumber: string;
    billerCode: string;
    billReference: string;
    amount: string;
}
export declare class SplitDestinationDto {
    destinationAccountNumber: string;
    amount: string;
    narration?: string;
}
export declare class SplitPaymentDto {
    sourceAccountNumber: string;
    splits: SplitDestinationDto[];
    batchNarration?: string;
    idempotencyKey?: string;
}
export declare class RtgsPaymentDto {
    sourceAccountNumber: string;
    beneficiaryBankSwift: string;
    beneficiaryAccountNumber: string;
    beneficiaryName: string;
    amount: string;
    narration?: string;
    idempotencyKey?: string;
}
