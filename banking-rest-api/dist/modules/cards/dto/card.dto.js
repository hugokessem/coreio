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
exports.ListCardsDto = exports.ReplaceCardDto = exports.RequestCardDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
const card_request_entity_1 = require("../../../database/entities/card-request.entity");
class RequestCardDto {
    accountNumber;
    cardType;
    cardBrand;
    targetBranchCode;
}
exports.RequestCardDto = RequestCardDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '1000123456789' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RequestCardDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ enum: card_request_entity_1.CardType, default: card_request_entity_1.CardType.DEBIT }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsEnum)(card_request_entity_1.CardType),
    __metadata("design:type", String)
], RequestCardDto.prototype, "cardType", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ enum: card_request_entity_1.CardBrand, default: card_request_entity_1.CardBrand.MASTERCARD }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsEnum)(card_request_entity_1.CardBrand),
    __metadata("design:type", String)
], RequestCardDto.prototype, "cardBrand", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '001', description: 'Branch code for collection' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], RequestCardDto.prototype, "targetBranchCode", void 0);
class ReplaceCardDto {
    cardNumberMasked;
    reason;
    targetBranchCode;
}
exports.ReplaceCardDto = ReplaceCardDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '541234******7890', description: 'Masked card number or card hash' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ReplaceCardDto.prototype, "cardNumberMasked", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'Card damaged / magnetic stripe unreadable' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ReplaceCardDto.prototype, "reason", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '100', description: 'Collection branch code' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ReplaceCardDto.prototype, "targetBranchCode", void 0);
class ListCardsDto {
    accountNumber;
    customerNumber;
}
exports.ListCardsDto = ListCardsDto;
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '1000123456789' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ListCardsDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '10000001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ListCardsDto.prototype, "customerNumber", void 0);
//# sourceMappingURL=card.dto.js.map