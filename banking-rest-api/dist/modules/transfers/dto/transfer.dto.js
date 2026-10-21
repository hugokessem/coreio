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
exports.ReversalDto = exports.TransactionStatusDto = exports.TransferFromHoldDto = exports.ExecuteTransferDto = exports.TransferVerifyDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
const transaction_entity_1 = require("../../../database/entities/transaction.entity");
class TransferVerifyDto {
    sourceAccountNumber;
    destinationAccountNumber;
    amount;
    currencyCode;
    channel;
}
exports.TransferVerifyDto = TransferVerifyDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Debiting account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransferVerifyDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000987654321', description: 'Crediting account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransferVerifyDto.prototype, "destinationAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1500.0000', description: 'Transfer amount' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], TransferVerifyDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'ETB', description: 'Transfer currency' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransferVerifyDto.prototype, "currencyCode", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ enum: transaction_entity_1.TransactionChannel, example: transaction_entity_1.TransactionChannel.MOBILE }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsEnum)(transaction_entity_1.TransactionChannel),
    __metadata("design:type", String)
], TransferVerifyDto.prototype, "channel", void 0);
class ExecuteTransferDto {
    sourceAccountNumber;
    destinationAccountNumber;
    amount;
    currencyCode;
    narration;
    idempotencyKey;
    channel;
}
exports.ExecuteTransferDto = ExecuteTransferDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Debiting account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ExecuteTransferDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000987654321', description: 'Crediting account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ExecuteTransferDto.prototype, "destinationAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1500.0000', description: 'Transfer amount' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], ExecuteTransferDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'ETB', description: 'Currency code' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ExecuteTransferDto.prototype, "currencyCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Invoice #8472 Payment', description: 'Payment narration/memo' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ExecuteTransferDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-UUID-98234-ABCD', description: 'Idempotency key for preventing duplicates' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ExecuteTransferDto.prototype, "idempotencyKey", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ enum: transaction_entity_1.TransactionChannel, example: transaction_entity_1.TransactionChannel.MOBILE }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsEnum)(transaction_entity_1.TransactionChannel),
    __metadata("design:type", String)
], ExecuteTransferDto.prototype, "channel", void 0);
class TransferFromHoldDto {
    holdReference;
    destinationAccountNumber;
    narration;
    idempotencyKey;
}
exports.TransferFromHoldDto = TransferFromHoldDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'HLD-001', description: 'Active Hold reference to settle' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransferFromHoldDto.prototype, "holdReference", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000987654321', description: 'Crediting destination account' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransferFromHoldDto.prototype, "destinationAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Settlement of judicial hold' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransferFromHoldDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-HOLD-SETTLE-001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransferFromHoldDto.prototype, "idempotencyKey", void 0);
class TransactionStatusDto {
    transactionReference;
}
exports.TransactionStatusDto = TransactionStatusDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'FT2629399881', description: 'Transaction reference number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], TransactionStatusDto.prototype, "transactionReference", void 0);
class ReversalDto {
    transactionReference;
    reason;
}
exports.ReversalDto = ReversalDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'FT2629399881', description: 'Transaction reference to reverse' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ReversalDto.prototype, "transactionReference", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Erroneous debit reversal requested by customer' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ReversalDto.prototype, "reason", void 0);
//# sourceMappingURL=transfer.dto.js.map