"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.AccountMiniStatementDto = exports.AccountStatementDto = exports.ReleaseHoldDto = exports.CreateHoldDto = exports.AccountEligibilityDto = exports.AccountListDto = exports.CreateAccountDto = exports.AccountNameLookupDto = exports.AccountLookupDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
class AccountLookupDto {
    accountNumber;
}
exports.AccountLookupDto = AccountLookupDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Account number to lookup' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountLookupDto.prototype, "accountNumber", void 0);
class AccountNameLookupDto {
    accountNumber;
}
exports.AccountNameLookupDto = AccountNameLookupDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Account number for beneficiary/name lookup' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountNameLookupDto.prototype, "accountNumber", void 0);
class CreateAccountDto {
    customerNumber;
    accountTypeCode;
    branchCode;
    currencyCode;
    accountName;
    initialDeposit;
}
exports.CreateAccountDto = CreateAccountDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '10000001', description: 'Customer number owning the account' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateAccountDto.prototype, "customerNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '6501', description: 'Account category code (6501 Savings, 6001 Current)' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateAccountDto.prototype, "accountTypeCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '001', description: 'Branch code (001 Addis Ababa Main, 100 Bole)' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateAccountDto.prototype, "branchCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'ETB', description: 'Currency code (ETB, USD, EUR, GBP)' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateAccountDto.prototype, "currencyCode", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Abebe Kebede Personal Account' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateAccountDto.prototype, "accountName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '1000.0000', description: 'Initial deposit amount' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], CreateAccountDto.prototype, "initialDeposit", void 0);
class AccountListDto {
    customerNumber;
}
exports.AccountListDto = AccountListDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '10000001', description: 'Customer number to list all accounts for' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountListDto.prototype, "customerNumber", void 0);
class AccountEligibilityDto {
    accountNumber;
    amount;
    transactionType;
}
exports.AccountEligibilityDto = AccountEligibilityDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountEligibilityDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '500.0000', description: 'Amount intended for transaction' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], AccountEligibilityDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'DEBIT', description: 'DEBIT or CREDIT' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountEligibilityDto.prototype, "transactionType", void 0);
class CreateHoldDto {
    accountNumber;
    amount;
    reason;
    lockEvent;
    fromDate;
    toDate;
}
exports.CreateHoldDto = CreateHoldDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateHoldDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '2500.0000' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], CreateHoldDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Court Order Freeze' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateHoldDto.prototype, "reason", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'COURT_ORDER' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateHoldDto.prototype, "lockEvent", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '2026-10-20' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateHoldDto.prototype, "fromDate", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '2026-12-31' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateHoldDto.prototype, "toDate", void 0);
class ReleaseHoldDto {
    holdReference;
    reason;
}
exports.ReleaseHoldDto = ReleaseHoldDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'HLD-001', description: 'Hold / Lock reference to release' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ReleaseHoldDto.prototype, "holdReference", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Court order lifted' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ReleaseHoldDto.prototype, "reason", void 0);
class AccountStatementDto {
    accountNumber;
    startDate;
    endDate;
    limit;
}
exports.AccountStatementDto = AccountStatementDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountStatementDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '2026-01-01' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountStatementDto.prototype, "startDate", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '2026-12-31' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountStatementDto.prototype, "endDate", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 50, default: 50 }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsInt)(),
    (0, class_validator_1.Min)(1),
    (0, class_validator_1.Max)(200),
    __metadata("design:type", Number)
], AccountStatementDto.prototype, "limit", void 0);
class AccountMiniStatementDto {
    accountNumber;
    count;
    limit;
}
exports.AccountMiniStatementDto = AccountMiniStatementDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountMiniStatementDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 5, default: 5 }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsInt)(),
    (0, class_validator_1.Min)(1),
    (0, class_validator_1.Max)(50),
    __metadata("design:type", Number)
], AccountMiniStatementDto.prototype, "count", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 5, default: 5 }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsInt)(),
    (0, class_validator_1.Min)(1),
    (0, class_validator_1.Max)(50),
    __metadata("design:type", Number)
], AccountMiniStatementDto.prototype, "limit", void 0);
//# sourceMappingURL=account.dto.js.map