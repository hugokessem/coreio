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
exports.AmendCustomerLimitDto = exports.CustomerLimitsQueryDto = exports.ServiceDetailDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
class ServiceDetailDto {
    serviceCode;
}
exports.ServiceDetailDto = ServiceDetailDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'FT_INT', description: 'Service code (FT_INT, FT_IPS, CBE_BIRR, BILL_ETHIO)' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], ServiceDetailDto.prototype, "serviceCode", void 0);
class CustomerLimitsQueryDto {
    customerNumber;
    serviceCode;
}
exports.CustomerLimitsQueryDto = CustomerLimitsQueryDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '10000001' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerLimitsQueryDto.prototype, "customerNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FT_INT' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerLimitsQueryDto.prototype, "serviceCode", void 0);
class AmendCustomerLimitDto {
    customerNumber;
    serviceCode;
    dailyLimit;
    perTransactionLimit;
}
exports.AmendCustomerLimitDto = AmendCustomerLimitDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '10000001' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AmendCustomerLimitDto.prototype, "customerNumber", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'FT_INT' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AmendCustomerLimitDto.prototype, "serviceCode", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '750000.0000' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], AmendCustomerLimitDto.prototype, "dailyLimit", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '150000.0000' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], AmendCustomerLimitDto.prototype, "perTransactionLimit", void 0);
//# sourceMappingURL=service.dto.js.map