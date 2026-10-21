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
exports.ListStandingOrderHistoryDto = exports.UpdateStandingOrderDto = exports.CancelStandingOrderDto = exports.ListStandingOrdersDto = exports.CreateStandingOrderDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
const standing_order_entity_1 = require("../../../database/entities/standing-order.entity");
class CreateStandingOrderDto {
    sourceAccountNumber;
    beneficiaryAccount;
    beneficiaryName;
    beneficiaryBankCode;
    amount;
    currencyCode;
    frequency;
    startDate;
    endDate;
}
exports.CreateStandingOrderDto = CreateStandingOrderDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "sourceAccountNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000987654321' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "beneficiaryAccount", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Almaz Tadesse Mengesha' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "beneficiaryName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'CBETETAA' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "beneficiaryBankCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '2500.0000' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'ETB' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "currencyCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ enum: standing_order_entity_1.StandingOrderFrequency, example: standing_order_entity_1.StandingOrderFrequency.MONTHLY }),
    (0, class_validator_1.IsEnum)(standing_order_entity_1.StandingOrderFrequency),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "frequency", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '2026-11-01' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "startDate", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '2027-11-01' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateStandingOrderDto.prototype, "endDate", void 0);
class ListStandingOrdersDto {
    accountNumber;
}
exports.ListStandingOrdersDto = ListStandingOrdersDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ListStandingOrdersDto.prototype, "accountNumber", void 0);
class CancelStandingOrderDto {
    orderReference;
}
exports.CancelStandingOrderDto = CancelStandingOrderDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'SO-1000123456789-01' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CancelStandingOrderDto.prototype, "orderReference", void 0);
class UpdateStandingOrderDto {
    amount;
    frequency;
    endDate;
    status;
}
exports.UpdateStandingOrderDto = UpdateStandingOrderDto;
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '3000.0000' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], UpdateStandingOrderDto.prototype, "amount", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ enum: standing_order_entity_1.StandingOrderFrequency, example: standing_order_entity_1.StandingOrderFrequency.WEEKLY }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsEnum)(standing_order_entity_1.StandingOrderFrequency),
    __metadata("design:type", String)
], UpdateStandingOrderDto.prototype, "frequency", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '2027-12-31' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], UpdateStandingOrderDto.prototype, "endDate", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ enum: standing_order_entity_1.StandingOrderStatus, example: standing_order_entity_1.StandingOrderStatus.PAUSED }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsEnum)(standing_order_entity_1.StandingOrderStatus),
    __metadata("design:type", String)
], UpdateStandingOrderDto.prototype, "status", void 0);
class ListStandingOrderHistoryDto {
    orderReference;
}
exports.ListStandingOrderHistoryDto = ListStandingOrderHistoryDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'SO-1000123456789-01' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ListStandingOrderHistoryDto.prototype, "orderReference", void 0);
//# sourceMappingURL=standing-order.dto.js.map