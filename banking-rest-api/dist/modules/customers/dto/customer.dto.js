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
exports.CompositeOnboardDto = exports.FaydaVerifyDto = exports.CustomerEligibilityDto = exports.CheckExistenceDto = exports.AmlVerifyDto = exports.CustomerPhoneLookupDto = exports.CreateCustomerDto = exports.CustomerDetailDto = exports.CustomerLookupDto = void 0;
const swagger_1 = require("@nestjs/swagger");
const class_validator_1 = require("class-validator");
const customer_entity_1 = require("../../../database/entities/customer.entity");
class CustomerLookupDto {
    customerNumber;
    phoneNumber;
    accountNumber;
    nationalId;
}
exports.CustomerLookupDto = CustomerLookupDto;
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '10000001', description: 'Customer number' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerLookupDto.prototype, "customerNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '+251911223344', description: 'Customer registered phone number' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerLookupDto.prototype, "phoneNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '1000123456789', description: 'Account number associated with customer' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerLookupDto.prototype, "accountNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FAYDA-9876543210', description: 'National ID / Fayda number' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerLookupDto.prototype, "nationalId", void 0);
class CustomerDetailDto {
    customerNumber;
}
exports.CustomerDetailDto = CustomerDetailDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '10000001', description: 'Unique customer number' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerDetailDto.prototype, "customerNumber", void 0);
class CreateCustomerDto {
    customerType;
    firstName;
    middleName;
    lastName;
    legalName;
    fullName;
    dateOfBirth;
    gender;
    phoneNumber;
    email;
    address;
    addressLine1;
    city;
    nationalId;
    tin;
}
exports.CreateCustomerDto = CreateCustomerDto;
__decorate([
    (0, swagger_1.ApiProperty)({ enum: customer_entity_1.CustomerType, example: customer_entity_1.CustomerType.INDIVIDUAL }),
    (0, class_validator_1.IsEnum)(customer_entity_1.CustomerType),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "customerType", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Abebe' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "firstName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Kebede' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "middleName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Bikila' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "lastName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Abebe Kebede Bikila' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "legalName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Abebe Kebede Bikila' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "fullName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '1985-04-12' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "dateOfBirth", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'MALE' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "gender", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911223344' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "phoneNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'abebe.kebede@example.com' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsEmail)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "email", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Kirkos Sub-City, Addis Ababa' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "address", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Bole Medhanialem' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "addressLine1", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Addis Ababa' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "city", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FAYDA-9876543210' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "nationalId", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '1234567890' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CreateCustomerDto.prototype, "tin", void 0);
class CustomerPhoneLookupDto {
    phoneNumber;
}
exports.CustomerPhoneLookupDto = CustomerPhoneLookupDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911223344', description: 'Phone number to lookup' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerPhoneLookupDto.prototype, "phoneNumber", void 0);
class AmlVerifyDto {
    customerNumber;
}
exports.AmlVerifyDto = AmlVerifyDto;
__decorate([
    (0, swagger_1.ApiProperty)({ example: '10000001' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], AmlVerifyDto.prototype, "customerNumber", void 0);
class CheckExistenceDto {
    phoneNumber;
    nationalId;
}
exports.CheckExistenceDto = CheckExistenceDto;
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '+251911223344' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CheckExistenceDto.prototype, "phoneNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FAYDA-9876543210' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CheckExistenceDto.prototype, "nationalId", void 0);
class CustomerEligibilityDto {
    customerNumber;
    phoneNumber;
    nationalId;
    dateOfBirth;
    channel;
}
exports.CustomerEligibilityDto = CustomerEligibilityDto;
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '10000001' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerEligibilityDto.prototype, "customerNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '+251911223344' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerEligibilityDto.prototype, "phoneNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FAYDA-9876543210' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerEligibilityDto.prototype, "nationalId", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '1995-05-15' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerEligibilityDto.prototype, "dateOfBirth", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'MOBILE_BANKING' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CustomerEligibilityDto.prototype, "channel", void 0);
class FaydaVerifyDto {
    nationalId;
    faydaId;
    fullName;
    dateOfBirth;
    phoneNumber;
}
exports.FaydaVerifyDto = FaydaVerifyDto;
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FAYDA-9876543210', description: 'Ethiopian Fayda National Identity string' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], FaydaVerifyDto.prototype, "nationalId", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FIN-ETH-99281928' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], FaydaVerifyDto.prototype, "faydaId", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Abebe Bikila' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], FaydaVerifyDto.prototype, "fullName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '1985-04-12' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], FaydaVerifyDto.prototype, "dateOfBirth", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '+251911223344' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], FaydaVerifyDto.prototype, "phoneNumber", void 0);
class CompositeOnboardDto {
    customerType;
    firstName;
    middleName;
    lastName;
    phoneNumber;
    nationalId;
    accountTypeCode;
    branchCode;
    currencyCode;
    initialDeposit;
}
exports.CompositeOnboardDto = CompositeOnboardDto;
__decorate([
    (0, swagger_1.ApiProperty)({ enum: customer_entity_1.CustomerType, example: customer_entity_1.CustomerType.INDIVIDUAL }),
    (0, class_validator_1.IsEnum)(customer_entity_1.CustomerType),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "customerType", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Tewodros' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "firstName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Kassahun' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "middleName", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'Bekele' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "lastName", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '+251911889900' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "phoneNumber", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: 'FAYDA-7788990011' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "nationalId", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '6501', description: 'Account product category code (6501 Savings)' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "accountTypeCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: '001', description: 'Branch code (001 Addis Ababa Main)' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "branchCode", void 0);
__decorate([
    (0, swagger_1.ApiProperty)({ example: 'ETB' }),
    (0, class_validator_1.IsNotEmpty)(),
    (0, class_validator_1.IsString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "currencyCode", void 0);
__decorate([
    (0, swagger_1.ApiPropertyOptional)({ example: '500.0000', description: 'Initial deposit amount' }),
    (0, class_validator_1.IsOptional)(),
    (0, class_validator_1.IsNumberString)(),
    __metadata("design:type", String)
], CompositeOnboardDto.prototype, "initialDeposit", void 0);
//# sourceMappingURL=customer.dto.js.map