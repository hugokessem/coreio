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
exports.RtgsPaymentDto = exports.SplitPaymentDto = exports.SplitDestinationDto = exports.BillPaymentVerifyDto = exports.BillPaymentDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
const class_transformer_1 = require("class-transformer");
class BillPaymentDto {
    sourceAccountNumber;
    billerCode;
    billReference;
    amount;
    narration;
    idempotencyKey;
}
exports.BillPaymentDto = BillPaymentDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Customer debit account' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'BILL_ETHIO', description: 'Biller service code' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentDto.prototype, "billerCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '0911223344', description: 'Customer utility/phone/bill account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentDto.prototype, "billReference", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '350.0000', description: 'Payment amount' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], BillPaymentDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Ethio Telecom Monthly Postpaid Bill' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-BILL-0099' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentDto.prototype, "idempotencyKey", void 0);
class BillPaymentVerifyDto {
    sourceAccountNumber;
    billerCode;
    billReference;
    amount;
}
exports.BillPaymentVerifyDto = BillPaymentVerifyDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Customer debit account' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentVerifyDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'BILL_ETHIO', description: 'Biller service code' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentVerifyDto.prototype, "billerCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '0911223344', description: 'Customer utility/phone/bill account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], BillPaymentVerifyDto.prototype, "billReference", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '350.0000', description: 'Payment amount' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], BillPaymentVerifyDto.prototype, "amount", void 0);
class SplitDestinationDto {
    destinationAccountNumber;
    amount;
    narration;
}
exports.SplitDestinationDto = SplitDestinationDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000987654321' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SplitDestinationDto.prototype, "destinationAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '500.0000' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], SplitDestinationDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Employee Bonus Portion' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SplitDestinationDto.prototype, "narration", void 0);
class SplitPaymentDto {
    sourceAccountNumber;
    splits;
    batchNarration;
    idempotencyKey;
}
exports.SplitPaymentDto = SplitPaymentDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '2000123456789', description: 'Source corporate/master account' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SplitPaymentDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ type: [SplitDestinationDto], description: 'List of recipients and split amounts' }),
    (0, class_validator_1.IsArray)(),
    (0, class_validator_1.ValidateNested)({ each: true }),
    (0, class_transformer_1.Type)(() => SplitDestinationDto),
    __metadata("design:type", Array)
], SplitPaymentDto.prototype, "splits", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Vendor Bulk Settlement' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SplitPaymentDto.prototype, "batchNarration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-SPLIT-BATCH-001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SplitPaymentDto.prototype, "idempotencyKey", void 0);
class RtgsPaymentDto {
    sourceAccountNumber;
    beneficiaryBankSwift;
    beneficiaryAccountNumber;
    beneficiaryName;
    amount;
    narration;
    idempotencyKey;
}
exports.RtgsPaymentDto = RtgsPaymentDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RtgsPaymentDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'AWINETAA', description: 'Destination Bank SWIFT/BIC code' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RtgsPaymentDto.prototype, "beneficiaryBankSwift", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '0132049281900', description: 'Beneficiary account number at destination bank' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RtgsPaymentDto.prototype, "beneficiaryAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Awash Bank Commercial Customer PLC' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RtgsPaymentDto.prototype, "beneficiaryName", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '250000.0000' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], RtgsPaymentDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Import LC Settlement' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RtgsPaymentDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-RTGS-001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RtgsPaymentDto.prototype, "idempotencyKey", void 0);
//# sourceMappingURL=payment.dto.js.map