import { HttpException, HttpStatus } from '@nestjs/common';

export class BankingDomainException extends HttpException {
  constructor(message: string, statusCode = HttpStatus.BAD_REQUEST, public readonly errorCode = 'BANKING_DOMAIN_ERROR') {
    super({ message, errorCode, statusCode }, statusCode);
  }
}

export class AccountNotFoundException extends BankingDomainException {
  constructor(accountNumber: string) {
    super(`Account with number '${accountNumber}' was not found`, HttpStatus.NOT_FOUND, 'ACCOUNT_NOT_FOUND');
  }
}

export class CustomerNotFoundException extends BankingDomainException {
  constructor(identifier: string) {
    super(`Customer '${identifier}' was not found`, HttpStatus.NOT_FOUND, 'CUSTOMER_NOT_FOUND');
  }
}

export class InsufficientFundsException extends BankingDomainException {
  constructor(available: string, required: string) {
    super(`Insufficient available funds. Available: ${available}, Required: ${required}`, HttpStatus.UNPROCESSABLE_ENTITY, 'INSUFFICIENT_FUNDS');
  }
}

export class AccountRestrictedException extends BankingDomainException {
  constructor(accountNumber: string, reason: string) {
    super(`Account '${accountNumber}' is restricted: ${reason}`, HttpStatus.FORBIDDEN, 'ACCOUNT_RESTRICTED');
  }
}

export class CurrencyMismatchException extends BankingDomainException {
  constructor(expected: string, received: string) {
    super(`Currency mismatch: Expected ${expected}, received ${received}`, HttpStatus.BAD_REQUEST, 'CURRENCY_MISMATCH');
  }
}

export class IdempotencyConflictException extends BankingDomainException {
  constructor(key: string) {
    super(`Duplicate request detected for idempotency key '${key}'`, HttpStatus.CONFLICT, 'IDEMPOTENCY_CONFLICT');
  }
}

export class HoldNotFoundException extends BankingDomainException {
  constructor(holdReference: string) {
    super(`Account hold '${holdReference}' was not found or is already released`, HttpStatus.NOT_FOUND, 'HOLD_NOT_FOUND');
  }
}

export class DailyLimitExceededException extends BankingDomainException {
  constructor(limit: string, used: string) {
    super(`Daily limit exceeded. Limit: ${limit}, Used today: ${used}`, HttpStatus.UNPROCESSABLE_ENTITY, 'DAILY_LIMIT_EXCEEDED');
  }
}
