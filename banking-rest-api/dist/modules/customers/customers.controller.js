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
exports.CustomersController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const customers_service_1 = require("./customers.service");
const customer_dto_1 = require("./dto/customer.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
const roles_guard_1 = require("../../common/guards/roles.guard");
const roles_decorator_1 = require("../../common/decorators/roles.decorator");
const user_entity_1 = require("../../database/entities/user.entity");
let CustomersController = class CustomersController {
    customersService;
    constructor(customersService) {
        this.customersService = customersService;
    }
    async lookup(dto) {
        return this.customersService.lookup(dto);
    }
    async detail(dto) {
        return this.customersService.getDetail(dto);
    }
    async phoneLookup(dto) {
        return this.customersService.phoneLookup(dto);
    }
    async create(dto, req) {
        const creator = req.user?.username || 'SYSTEM';
        return this.customersService.create(dto, creator);
    }
    async fetch(body) {
        return this.customersService.fetchAll(body.page || 1, body.limit || 50);
    }
    async deletePost(body) {
        return this.customersService.deleteCustomer(body.customerNumber);
    }
    async amlVerify(dto, paramNumber) {
        const customerNumber = paramNumber || dto.customerNumber;
        return this.customersService.amlVerify({ customerNumber });
    }
    async authorizeCustomer(customerNumber, body, req) {
        const authorizer = body?.authorizerUsername || req.user?.username || 'ADMIN';
        return this.customersService.authorizeCustomer(customerNumber, authorizer);
    }
    async checkExistence(dto) {
        return this.customersService.checkExistence(dto);
    }
    async checkEligibility(dto) {
        return this.customersService.checkCustomerEligibility(dto);
    }
    async verifyFayda(dto) {
        return this.customersService.verifyFayda(dto);
    }
    async compositeOnboard(dto, req) {
        const creator = req.user?.username || 'SYSTEM';
        return this.customersService.compositeOnboard(dto, creator);
    }
    async amlLookupByAccount(accountNumber) {
        return this.customersService.amlLookupByAccount(accountNumber);
    }
    async getAll(page = '1', limit = '50') {
        return this.customersService.fetchAll(parseInt(page, 10), parseInt(limit, 10));
    }
    async getOne(customerNumber) {
        return this.customersService.getDetail({ customerNumber });
    }
    async deleteOne(customerNumber) {
        return this.customersService.deleteCustomer(customerNumber);
    }
};
exports.CustomersController = CustomersController;
__decorate([
    (0, common_1.Post)('customers/lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Lookup customer by customer number, phone, NID, or account number' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Customer details retrieved' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.CustomerLookupDto]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "lookup", null);
__decorate([
    (0, common_1.Post)('customers/detail'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Get full customer detail by customer number' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.CustomerDetailDto]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "detail", null);
__decorate([
    (0, common_1.Post)('customers/phone-lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Lookup customer by registered phone number' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.CustomerPhoneLookupDto]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "phoneLookup", null);
__decorate([
    (0, common_1.Post)(['customers/create', 'customers']),
    (0, common_1.HttpCode)(common_1.HttpStatus.CREATED),
    (0, swagger_1.ApiOperation)({ summary: 'Register/create new customer record' }),
    __param(0, (0, common_1.Body)()),
    __param(1, (0, common_1.Req)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.CreateCustomerDto, Object]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "create", null);
__decorate([
    (0, common_1.Post)('customers/fetch'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Fetch paginated list of customers' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "fetch", null);
__decorate([
    (0, roles_decorator_1.Roles)(user_entity_1.UserRole.ADMIN),
    (0, common_1.Post)('customers/delete'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Delete customer record (Requires ADMIN role)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "deletePost", null);
__decorate([
    (0, common_1.Post)(['customers/aml-verify', 'customers/:customerNumber/verify-aml']),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Verify customer AML compliance' }),
    __param(0, (0, common_1.Body)()),
    __param(1, (0, common_1.Param)('customerNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.AmlVerifyDto, String]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "amlVerify", null);
__decorate([
    (0, roles_decorator_1.Roles)(user_entity_1.UserRole.TELLER, user_entity_1.UserRole.ADMIN),
    (0, common_1.Post)('customers/:customerNumber/authorize'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Maker-checker approval of a customer record (Enforces dual control)' }),
    __param(0, (0, common_1.Param)('customerNumber')),
    __param(1, (0, common_1.Body)()),
    __param(2, (0, common_1.Req)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String, Object, Object]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "authorizeCustomer", null);
__decorate([
    (0, common_1.Post)('customers/check-existence'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Pre-onboarding existence check: validates if phone or national ID is registered' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.CheckExistenceDto]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "checkExistence", null);
__decorate([
    (0, common_1.Post)('customers/eligibility'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Check customer eligibility for digital channels and services' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.CustomerEligibilityDto]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "checkEligibility", null);
__decorate([
    (0, common_1.Post)('identity/fayda/verify'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Verify national citizen identity against local Fayda identity persistence store' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.FaydaVerifyDto]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "verifyFayda", null);
__decorate([
    (0, common_1.Post)('accounts/onboard'),
    (0, common_1.HttpCode)(common_1.HttpStatus.CREATED),
    (0, swagger_1.ApiOperation)({ summary: 'Composite onboarding workflow: Customer registration, AML check, and account opening' }),
    __param(0, (0, common_1.Body)()),
    __param(1, (0, common_1.Req)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [customer_dto_1.CompositeOnboardDto, Object]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "compositeOnboard", null);
__decorate([
    (0, common_1.Post)('accounts/:accountNumber/aml-lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Check AML status associated with an account holder' }),
    __param(0, (0, common_1.Param)('accountNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "amlLookupByAccount", null);
__decorate([
    (0, common_1.Get)('customers'),
    (0, swagger_1.ApiOperation)({ summary: 'GET all customers (REST convention)' }),
    __param(0, (0, common_1.Query)('page')),
    __param(1, (0, common_1.Query)('limit')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object, Object]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "getAll", null);
__decorate([
    (0, common_1.Get)('customers/:customerNumber'),
    (0, swagger_1.ApiOperation)({ summary: 'GET customer by customer number (REST convention)' }),
    __param(0, (0, common_1.Param)('customerNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "getOne", null);
__decorate([
    (0, roles_decorator_1.Roles)(user_entity_1.UserRole.ADMIN),
    (0, common_1.Delete)('customers/:customerNumber'),
    (0, swagger_1.ApiOperation)({ summary: 'DELETE customer by customer number (Requires ADMIN role)' }),
    __param(0, (0, common_1.Param)('customerNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], CustomersController.prototype, "deleteOne", null);
exports.CustomersController = CustomersController = __decorate([
    (0, swagger_1.ApiTags)('Customers & Identity'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard, roles_guard_1.RolesGuard),
    (0, common_1.Controller)('api/v1'),
    __metadata("design:paramtypes", [customers_service_1.CustomersService])
], CustomersController);
//# sourceMappingURL=customers.controller.js.map