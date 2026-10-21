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
exports.WalletAgentTransferDto = exports.WalletAgentLookupDto = exports.WalletToAccountTransferDto = exports.AccountToWalletTransferDto = exports.WalletCustomerLookupDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
class WalletCustomerLookupDto {
    phoneNumber;
}
exports.WalletCustomerLookupDto = WalletCustomerLookupDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911223344', description: 'CBE Birr registered mobile phone number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletCustomerLookupDto.prototype, "phoneNumber", void 0);
class AccountToWalletTransferDto {
    sourceAccountNumber;
    walletPhoneNumber;
    amount;
    narration;
    idempotencyKey;
}
exports.AccountToWalletTransferDto = AccountToWalletTransferDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'CBE Bank Account Number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountToWalletTransferDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911223344', description: 'CBE Birr Wallet Phone Number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountToWalletTransferDto.prototype, "walletPhoneNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '500.0000', description: 'Amount in ETB' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], AccountToWalletTransferDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Top-up CBE Birr Wallet' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountToWalletTransferDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-WALLET-TOPUP-001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AccountToWalletTransferDto.prototype, "idempotencyKey", void 0);
class WalletToAccountTransferDto {
    walletPhoneNumber;
    destinationAccountNumber;
    amount;
    narration;
    idempotencyKey;
}
exports.WalletToAccountTransferDto = WalletToAccountTransferDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911223344', description: 'CBE Birr Wallet Phone Number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletToAccountTransferDto.prototype, "walletPhoneNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'CBE Bank Account Number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletToAccountTransferDto.prototype, "destinationAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '500.0000', description: 'Amount in ETB' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], WalletToAccountTransferDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Transfer from CBE Birr to Savings' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletToAccountTransferDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-WALLET-CASHOUT-001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletToAccountTransferDto.prototype, "idempotencyKey", void 0);
class WalletAgentLookupDto {
    agentIdentifier;
}
exports.WalletAgentLookupDto = WalletAgentLookupDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911223344', description: 'Agent phone number or unique agent identifier' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletAgentLookupDto.prototype, "agentIdentifier", void 0);
class WalletAgentTransferDto {
    sourceAccountNumber;
    agentIdentifier;
    amount;
    narration;
    idempotencyKey;
}
exports.WalletAgentTransferDto = WalletAgentTransferDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789', description: 'CBE Bank Account Number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletAgentTransferDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911223344', description: 'Agent mobile phone or agent wallet ID' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletAgentTransferDto.prototype, "agentIdentifier", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000.0000', description: 'Amount in ETB' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], WalletAgentTransferDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Float disbursement to agent' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletAgentTransferDto.prototype, "narration", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'IDEMP-AGENT-TRF-001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], WalletAgentTransferDto.prototype, "idempotencyKey", void 0);
//# sourceMappingURL=wallet.dto.js.map