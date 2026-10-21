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
exports.StandingOrdersController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const standing_orders_service_1 = require("./standing-orders.service");
const standing_order_dto_1 = require("./dto/standing-order.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
const roles_guard_1 = require("../../common/guards/roles.guard");
const roles_decorator_1 = require("../../common/decorators/roles.decorator");
const user_role_enum_1 = require("../../common/enums/user-role.enum");
let StandingOrdersController = class StandingOrdersController {
    standingOrdersService;
    constructor(standingOrdersService) {
        this.standingOrdersService = standingOrdersService;
    }
    async createCanonical(dto) {
        return this.standingOrdersService.create(dto);
    }
    async create(dto) {
        return this.standingOrdersService.create(dto);
    }
    async listQuery(accountNumber) {
        return this.standingOrdersService.listByAccount(accountNumber);
    }
    async list(dto) {
        return this.standingOrdersService.list(dto);
    }
    async getHistory(orderReference) {
        return this.standingOrdersService.getHistory(orderReference);
    }
    async getHistoryBody(dto) {
        return this.standingOrdersService.getHistory(dto.orderReference);
    }
    async update(orderReference, dto) {
        return this.standingOrdersService.update(orderReference, dto);
    }
    async deleteCanonical(orderReference) {
        return this.standingOrdersService.cancel(orderReference);
    }
    async cancel(dto) {
        return this.standingOrdersService.cancel(dto);
    }
    async executeDue() {
        return this.standingOrdersService.executeDueOrders();
    }
};
exports.StandingOrdersController = StandingOrdersController;
__decorate([
    (0, common_1.Post)(),
    (0, common_1.HttpCode)(common_1.HttpStatus.CREATED),
    (0, swagger_1.ApiOperation)({ summary: 'Setup recurring standing order instruction - Canonical (Op 30: CreateStandingOrder)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [standing_order_dto_1.CreateStandingOrderDto]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "createCanonical", null);
__decorate([
    (0, common_1.Post)('create'),
    (0, common_1.HttpCode)(common_1.HttpStatus.CREATED),
    (0, swagger_1.ApiOperation)({ summary: 'Setup recurring standing order instruction - Legacy Alias (Op 30: CreateStandingOrder)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [standing_order_dto_1.CreateStandingOrderDto]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "create", null);
__decorate([
    (0, common_1.Get)(),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'List all standing orders for an account - Canonical (Op 31: ListStandingOrder)' }),
    __param(0, (0, common_1.Query)('accountNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "listQuery", null);
__decorate([
    (0, common_1.Post)('list'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'List all standing orders for an account - Legacy Alias (Op 31: ListStandingOrder)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [standing_order_dto_1.ListStandingOrdersDto]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "list", null);
__decorate([
    (0, common_1.Get)(':orderReference/history'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Query execution history of a standing order (Op 32: ListStandingOrderHistory)' }),
    __param(0, (0, common_1.Param)('orderReference')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "getHistory", null);
__decorate([
    (0, common_1.Post)('history'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Query execution history of a standing order - Body Alias (Op 32: ListStandingOrderHistory)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [standing_order_dto_1.ListStandingOrderHistoryDto]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "getHistoryBody", null);
__decorate([
    (0, common_1.Put)(':orderReference'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Update parameters of an active standing order (Op 33: UpdateStandingOrder)' }),
    __param(0, (0, common_1.Param)('orderReference')),
    __param(1, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String, standing_order_dto_1.UpdateStandingOrderDto]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "update", null);
__decorate([
    (0, common_1.Delete)(':orderReference'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Cancel an active standing order - Canonical (Op 34: CancleStandingOrder)' }),
    __param(0, (0, common_1.Param)('orderReference')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "deleteCanonical", null);
__decorate([
    (0, common_1.Post)('cancel'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Cancel an active standing order instruction - Legacy Alias (Op 34: CancleStandingOrder)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [standing_order_dto_1.CancelStandingOrderDto]),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "cancel", null);
__decorate([
    (0, common_1.Post)('execute-due'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, roles_decorator_1.Roles)(user_role_enum_1.UserRole.ADMIN),
    (0, swagger_1.ApiOperation)({ summary: 'Execute all due standing orders with FOR UPDATE SKIP LOCKED concurrency' }),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], StandingOrdersController.prototype, "executeDue", null);
exports.StandingOrdersController = StandingOrdersController = __decorate([
    (0, swagger_1.ApiTags)('Standing Orders'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard, roles_guard_1.RolesGuard),
    (0, common_1.Controller)('api/v1/standing-orders'),
    __metadata("design:paramtypes", [standing_orders_service_1.StandingOrdersService])
], StandingOrdersController);
//# sourceMappingURL=standing-orders.controller.js.map