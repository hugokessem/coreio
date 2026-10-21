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
exports.SuperAppController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const superapp_service_1 = require("./superapp.service");
const superapp_dto_1 = require("./dto/superapp.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
let SuperAppController = class SuperAppController {
    superAppService;
    constructor(superAppService) {
        this.superAppService = superAppService;
    }
    async subscribe(dto) {
        return this.superAppService.subscribe(dto);
    }
    async unsubscribe(dto) {
        return this.superAppService.unsubscribe(dto);
    }
    async changeStatusPut(dto) {
        return this.superAppService.changeStatus(dto);
    }
    async changeStatusPost(dto) {
        return this.superAppService.changeStatus(dto);
    }
    async status(dto) {
        return this.superAppService.getStatus(dto);
    }
    async getStatusParam(externalUserId) {
        return this.superAppService.getByExternalUserId(externalUserId);
    }
};
exports.SuperAppController = SuperAppController;
__decorate([
    (0, common_1.Post)('subscribe'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Link/subscribe external partner user to bank customer profile (Op 47: SuperAppSubscribe)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [superapp_dto_1.SuperAppSubscribeDto]),
    __metadata("design:returntype", Promise)
], SuperAppController.prototype, "subscribe", null);
__decorate([
    (0, common_1.Post)('unsubscribe'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Deactivate external partner subscription link (Op 48: SuperAppUnsubscribe)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [superapp_dto_1.SuperAppUnsubscribeDto]),
    __metadata("design:returntype", Promise)
], SuperAppController.prototype, "unsubscribe", null);
__decorate([
    (0, common_1.Put)('status'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Update status of external partner subscription - Canonical REST (Op 49: SuperAppStatusChange)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [superapp_dto_1.SuperAppStatusChangeDto]),
    __metadata("design:returntype", Promise)
], SuperAppController.prototype, "changeStatusPut", null);
__decorate([
    (0, common_1.Post)('status/change'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Update status of external partner subscription - Alias (Op 49: SuperAppStatusChange)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [superapp_dto_1.SuperAppStatusChangeDto]),
    __metadata("design:returntype", Promise)
], SuperAppController.prototype, "changeStatusPost", null);
__decorate([
    (0, common_1.Post)('status'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Inquire status of external partner user link - Body Query' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [superapp_dto_1.SuperAppStatusDto]),
    __metadata("design:returntype", Promise)
], SuperAppController.prototype, "status", null);
__decorate([
    (0, common_1.Get)('status/:externalUserId'),
    (0, swagger_1.ApiOperation)({ summary: 'Inquire status of external partner user link - Route Param' }),
    __param(0, (0, common_1.Param)('externalUserId')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], SuperAppController.prototype, "getStatusParam", null);
exports.SuperAppController = SuperAppController = __decorate([
    (0, swagger_1.ApiTags)('SuperApp'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard),
    (0, common_1.Controller)('api/v1/superapp'),
    __metadata("design:paramtypes", [superapp_service_1.SuperAppService])
], SuperAppController);
//# sourceMappingURL=superapp.controller.js.map