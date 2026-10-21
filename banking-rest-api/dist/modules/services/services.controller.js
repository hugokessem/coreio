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
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.ServicesController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const services_service_1 = require("./services.service");
const service_dto_1 = require("./dto/service.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
const roles_guard_1 = require("../../common/guards/roles.guard");
const roles_decorator_1 = require("../../common/decorators/roles.decorator");
const user_role_enum_1 = require("../../common/enums/user-role.enum");
let ServicesController = class ServicesController {
    servicesService;
    constructor(servicesService) {
        this.servicesService = servicesService;
    }
    async listAll() {
        return this.servicesService.listAll();
    }
    async getDetailByCode(serviceCode) {
        return this.servicesService.getDetailByCode(serviceCode);
    }
    async getDetail(dto) {
        return this.servicesService.getDetail(dto);
    }
    async getServiceLimits() {
        return this.servicesService.getServiceLimits();
    }
    async getCustomerLimits(dto) {
        return this.servicesService.getCustomerLimits(dto);
    }
    async getCustomerLimitsCanonical(customerNumber, serviceCode) {
        return this.servicesService.getCustomerLimitsByNumber(customerNumber, serviceCode);
    }
    async amendCustomerLimitPut(dto) {
        return this.servicesService.amendCustomerLimit(dto);
    }
    async amendCustomerLimitPost(dto) {
        return this.servicesService.amendCustomerLimit(dto);
    }
};
exports.ServicesController = ServicesController;
__decorate([
    (0, common_1.Get)(),
    (0, swagger_1.ApiOperation)({ summary: 'List all supported banking services (REST convention)' }),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "listAll", null);
__decorate([
    (0, common_1.Get)('detail/:serviceCode'),
    (0, swagger_1.ApiOperation)({ summary: 'Get service tariff rules, channel limits, and configuration (Op 36: ServiceDetail Canonical)' }),
    __param(0, (0, common_1.Param)('serviceCode')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "getDetailByCode", null);
__decorate([
    (0, common_1.Post)('detail'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Get service tariff rules, channel limits, and configuration (Op 36: ServiceDetail Legacy)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [service_dto_1.ServiceDetailDto]),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "getDetail", null);
__decorate([
    (0, common_1.Get)('limits'),
    (0, swagger_1.ApiOperation)({ summary: 'Query tariff limits across all services (Op 37: ServiceLimit)' }),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "getServiceLimits", null);
__decorate([
    (0, common_1.Post)('limits'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Query custom per-customer transaction limits (Op 38: CustomerLimitFetchByCIF)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [service_dto_1.CustomerLimitsQueryDto]),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "getCustomerLimits", null);
__decorate([
    (0, common_1.Get)('limits/customer/:customerNumber'),
    (0, swagger_1.ApiOperation)({ summary: 'Query custom per-customer transaction limits - Canonical REST (Op 38: CustomerLimitFetchByCIF)' }),
    __param(0, (0, common_1.Param)('customerNumber')),
    __param(1, (0, common_1.Query)('serviceCode')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String, String]),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "getCustomerLimitsCanonical", null);
__decorate([
    (0, common_1.Put)('limits/amend'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, roles_decorator_1.Roles)(user_role_enum_1.UserRole.ADMIN),
    (0, swagger_1.ApiOperation)({ summary: 'Amend per-customer transaction limits - Admin Dual Control (Op 39: CustomerLimitAmendByCIF)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [service_dto_1.AmendCustomerLimitDto]),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "amendCustomerLimitPut", null);
__decorate([
    (0, common_1.Post)('limits/amend'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, roles_decorator_1.Roles)(user_role_enum_1.UserRole.ADMIN),
    (0, swagger_1.ApiOperation)({ summary: 'Amend per-customer transaction limits - Legacy Alias (Op 39: CustomerLimitAmendByCIF)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [service_dto_1.AmendCustomerLimitDto]),
    __metadata("design:returntype", Promise)
], ServicesController.prototype, "amendCustomerLimitPost", null);
exports.ServicesController = ServicesController = __decorate([
    (0, swagger_1.ApiTags)('Services & Limits'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard, roles_guard_1.RolesGuard),
    (0, common_1.Controller)('api/v1/services'),
    __metadata("design:paramtypes", [services_service_1.ServicesService])
], ServicesController);
//# sourceMappingURL=services.controller.js.map