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
exports.IpsStatusDto = exports.IpsQrPaymentDto = exports.IpsTransferDto = exports.IpsAccountLookupDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
class IpsAccountLookupDto {
    accountNumber;
    bankCode;
}
exports.IpsAccountLookupDto = IpsAccountLookupDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Account number to lookup via IPS' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsAccountLookupDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'CBETETAA', description: 'Bank identifier / BIC' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsAccountLookupDto.prototype, "bankCode", void 0);
class IpsTransferDto {
    sourceAccountNumber;
    destinationBankCode;
    destinationAccountNumber;
    beneficiaryName;
    amount;
    narration;
    endToEndId;
    idempotencyKey;
}
exports.IpsTransferDto = IpsTransferDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'Local source account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'AWINETAA', description: 'Recipient bank BIC/code' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "destinationBankCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '0132049281900', description: 'Recipient external account number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "destinationAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Dawit Yohannes', description: 'Beneficiary name' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "beneficiaryName", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '2500.0000', description: 'Transfer amount' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IPS P2P Interbank Transfer' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IPS-TX-20261020-0091', description: 'ISO 20022 EndToEndId' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "endToEndId", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-IPS-98124' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsTransferDto.prototype, "idempotencyKey", void 0);
class IpsQrPaymentDto {
    sourceAccountNumber;
    qrPayload;
    merchantAccount;
    merchantName;
    amount;
    narration;
    idempotencyKey;
}
exports.IpsQrPaymentDto = IpsQrPaymentDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsQrPaymentDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'ETH.IPS.MCH.8841029', description: 'Payload decoded from QR standard' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsQrPaymentDto.prototype, "qrPayload", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000987654321', description: 'Merchant receiving account' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsQrPaymentDto.prototype, "merchantAccount", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Fresh Supermarket Bole', description: 'Merchant legal trade name' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsQrPaymentDto.prototype, "merchantName", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '450.0000' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], IpsQrPaymentDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Retail Grocery Checkout' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsQrPaymentDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-QR-8821' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsQrPaymentDto.prototype, "idempotencyKey", void 0);
class IpsStatusDto {
    transactionReference;
}
exports.IpsStatusDto = IpsStatusDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'IPS2699182341' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], IpsStatusDto.prototype, "transactionReference", void 0);
//# sourceMappingURL=ips.dto.js.map