import { HttpException, HttpStatus } from '@nestjs/common';
export declare class BankingDomainException extends HttpException {
    readonly errorCode: string;
    constructor(message: string, statusCode?: HttpStatus, errorCode?: string);
}
export declare class AccountNotFoundException extends BankingDomainException {
    constructor(accountNumber: string);
}
export declare class CustomerNotFoundException extends BankingDomainException {
    constructor(identifier: string);
}
export declare class InsufficientFundsException extends BankingDomainException {
    constructor(available: string, required: string);
}
export declare class AccountRestrictedException extends BankingDomainException {
    constructor(accountNumber: string, reason: string);
}
export declare class CurrencyMismatchException extends BankingDomainException {
    constructor(expected: string, received: string);
}
export declare class IdempotencyConflictException extends BankingDomainException {
    constructor(key: string);
}
export declare class HoldNotFoundException extends BankingDomainException {
    constructor(holdReference: string);
}
export declare class DailyLimitExceededException extends BankingDomainException {
    constructor(limit: string, used: string);
}
