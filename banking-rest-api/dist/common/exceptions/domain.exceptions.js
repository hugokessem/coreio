"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DailyLimitExceededException = exports.HoldNotFoundException = exports.IdempotencyConflictException = exports.CurrencyMismatchException = exports.AccountRestrictedException = exports.InsufficientFundsException = exports.CustomerNotFoundException = exports.AccountNotFoundException = exports.BankingDomainException = void 0;
const common_1 = require("@nestjs/common");
class BankingDomainException extends common_1.HttpException {
    errorCode;
    constructor(message, statusCode = common_1.HttpStatus.BAD_REQUEST, errorCode = 'BANKING_DOMAIN_ERROR') {
        super({ message, errorCode, statusCode }, statusCode);
        this.errorCode = errorCode;
    }
}
exports.BankingDomainException = BankingDomainException;
class AccountNotFoundException extends BankingDomainException {
    constructor(accountNumber) {
        super(`Account with number '${accountNumber}' was not found`, common_1.HttpStatus.NOT_FOUND, 'ACCOUNT_NOT_FOUND');
    }
}
exports.AccountNotFoundException = AccountNotFoundException;
class CustomerNotFoundException extends BankingDomainException {
    constructor(identifier) {
        super(`Customer '${identifier}' was not found`, common_1.HttpStatus.NOT_FOUND, 'CUSTOMER_NOT_FOUND');
    }
}
exports.CustomerNotFoundException = CustomerNotFoundException;
class InsufficientFundsException extends BankingDomainException {
    constructor(available, required) {
        super(`Insufficient available funds. Available: ${available}, Required: ${required}`, common_1.HttpStatus.UNPROCESSABLE_ENTITY, 'INSUFFICIENT_FUNDS');
    }
}
exports.InsufficientFundsException = InsufficientFundsException;
class AccountRestrictedException extends BankingDomainException {
    constructor(accountNumber, reason) {
        super(`Account '${accountNumber}' is restricted: ${reason}`, common_1.HttpStatus.FORBIDDEN, 'ACCOUNT_RESTRICTED');
    }
}
exports.AccountRestrictedException = AccountRestrictedException;
class CurrencyMismatchException extends BankingDomainException {
    constructor(expected, received) {
        super(`Currency mismatch: Expected ${expected}, received ${received}`, common_1.HttpStatus.BAD_REQUEST, 'CURRENCY_MISMATCH');
    }
}
exports.CurrencyMismatchException = CurrencyMismatchException;
class IdempotencyConflictException extends BankingDomainException {
    constructor(key) {
        super(`Duplicate request detected for idempotency key '${key}'`, common_1.HttpStatus.CONFLICT, 'IDEMPOTENCY_CONFLICT');
    }
}
exports.IdempotencyConflictException = IdempotencyConflictException;
class HoldNotFoundException extends BankingDomainException {
    constructor(holdReference) {
        super(`Account hold '${holdReference}' was not found or is already released`, common_1.HttpStatus.NOT_FOUND, 'HOLD_NOT_FOUND');
    }
}
exports.HoldNotFoundException = HoldNotFoundException;
class DailyLimitExceededException extends BankingDomainException {
    constructor(limit, used) {
        super(`Daily limit exceeded. Limit: ${limit}, Used today: ${used}`, common_1.HttpStatus.UNPROCESSABLE_ENTITY, 'DAILY_LIMIT_EXCEEDED');
    }
}
exports.DailyLimitExceededException = DailyLimitExceededException;
//# sourceMappingURL=domain.exceptions.js.map