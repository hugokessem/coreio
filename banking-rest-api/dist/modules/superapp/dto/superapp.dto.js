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
exports.SuperAppStatusChangeDto = exports.SuperAppStatusDto = exports.SuperAppUnsubscribeDto = exports.SuperAppSubscribeDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
const super_app_user_entity_1 = require("../../../database/entities/super-app-user.entity");
class SuperAppSubscribeDto {
    externalUserId;
    customerNumber;
    partnerCode;
}
exports.SuperAppSubscribeDto = SuperAppSubscribeDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'EXT-USER-TELCO-9921', description: 'External ecosystem partner user ID' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SuperAppSubscribeDto.prototype, "externalUserId", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '10000001', description: 'Bank customer number to link' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SuperAppSubscribeDto.prototype, "customerNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'TELEBIRR_MINIAPP', description: 'Partner platform identifier' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SuperAppSubscribeDto.prototype, "partnerCode", void 0);
class SuperAppUnsubscribeDto {
    externalUserId;
}
exports.SuperAppUnsubscribeDto = SuperAppUnsubscribeDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'EXT-USER-TELCO-9921' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SuperAppUnsubscribeDto.prototype, "externalUserId", void 0);
class SuperAppStatusDto {
    externalUserId;
}
exports.SuperAppStatusDto = SuperAppStatusDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'EXT-USER-TELCO-9921' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SuperAppStatusDto.prototype, "externalUserId", void 0);
class SuperAppStatusChangeDto {
    externalUserId;
    status;
    reason;
}
exports.SuperAppStatusChangeDto = SuperAppStatusChangeDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'EXT-USER-TELCO-9921' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SuperAppStatusChangeDto.prototype, "externalUserId", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ enum: super_app_user_entity_1.SuperAppStatus, example: super_app_user_entity_1.SuperAppStatus.SUSPENDED }),
    (0, class_validator_1.IsEnum)(super_app_user_entity_1.SuperAppStatus),
    __metadata("design:type", String)
], SuperAppStatusChangeDto.prototype, "status", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Temporary suspension requested by customer' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], SuperAppStatusChangeDto.prototype, "reason", void 0);
//# sourceMappingURL=superapp.dto.js.map