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
exports.TransfersController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const transfers_service_1 = require("./transfers.service");
const reconciliation_service_1 = require("./reconciliation.service");
const transfer_dto_1 = require("./dto/transfer.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
const roles_guard_1 = require("../../common/guards/roles.guard");
const roles_decorator_1 = require("../../common/decorators/roles.decorator");
const user_role_enum_1 = require("../../common/enums/user-role.enum");
let TransfersController = class TransfersController {
    transfersService;
    reconciliationService;
    constructor(transfersService, reconciliationService) {
        this.transfersService = transfersService;
        this.reconciliationService = reconciliationService;
    }
    async verify(dto) {
        return this.transfersService.verify(dto);
    }
    async executeCanonical(dto) {
        return this.transfersService.execute(dto);
    }
    async execute(dto) {
        return this.transfersService.execute(dto);
    }
    async getByReference(transactionReference) {
        return this.transfersService.getByReference(transactionReference);
    }
    async getStatusByParam(transactionReference) {
        return this.transfersService.getByReference(transactionReference);
    }
    async status(dto) {
        return this.transfersService.getStatus(dto);
    }
    async transferFromHold(dto) {
        return this.transfersService.transferFromHold(dto);
    }
    async reverse(dto) {
        return this.transfersService.reverse(dto);
    }
    async revert(dto) {
        return this.transfersService.reverse(dto);
    }
    async reconcile() {
        return this.reconciliationService.reconcileAllAccounts();
    }
};
exports.TransfersController = TransfersController;
__decorate([
    (0, common_1.Post)('verify'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Pre-transfer verification, fee calculation, and balance sufficiency check (Op 23: FundTransferVerify)' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Transfer verified' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [transfer_dto_1.TransferVerifyDto]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "verify", null);
__decorate([
    (0, common_1.Post)(),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Execute funds transfer - Canonical REST route (Op 24: FundTransfer)' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Transfer successfully posted' }),
    (0, swagger_1.ApiResponse)({ status: 422, description: 'Insufficient funds' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [transfer_dto_1.ExecuteTransferDto]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "executeCanonical", null);
__decorate([
    (0, common_1.Post)('execute'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Execute funds transfer with double-entry ledger, row locking, and idempotency (Op 24: FundTransfer alias)' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Transfer successfully posted' }),
    (0, swagger_1.ApiResponse)({ status: 422, description: 'Insufficient funds' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [transfer_dto_1.ExecuteTransferDto]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "execute", null);
__decorate([
    (0, common_1.Get)(':transactionReference'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Query transfer details and verification by reference (Op 25: FundTransferCheck)' }),
    __param(0, (0, common_1.Param)('transactionReference')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "getByReference", null);
__decorate([
    (0, common_1.Get)('status/:transactionReference'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Query transaction status by transaction reference (Op 26: StatusCheck GET)' }),
    __param(0, (0, common_1.Param)('transactionReference')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "getStatusByParam", null);
__decorate([
    (0, common_1.Post)('status'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Query transaction status by body reference (Op 26: StatusCheck POST)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [transfer_dto_1.TransactionStatusDto]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "status", null);
__decorate([
    (0, common_1.Post)('from-hold'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Settle and execute transfer from an active pre-allocated hold (Op 44: LockedAmountFT)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [transfer_dto_1.TransferFromHoldDto]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "transferFromHold", null);
__decorate([
    (0, common_1.Post)('reverse'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Perform compensating reversal of a completed transfer (Op 27: RevertFundTransfer)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [transfer_dto_1.ReversalDto]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "reverse", null);
__decorate([
    (0, common_1.Post)('revert'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Perform compensating reversal of a completed transfer - Alias (Op 27: RevertFundTransfer)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [transfer_dto_1.ReversalDto]),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "revert", null);
__decorate([
    (0, common_1.Post)('reconcile'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, roles_decorator_1.Roles)(user_role_enum_1.UserRole.ADMIN),
    (0, swagger_1.ApiOperation)({ summary: 'Run system-wide double-entry ledger vs account balance reconciliation audit' }),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], TransfersController.prototype, "reconcile", null);
exports.TransfersController = TransfersController = __decorate([
    (0, swagger_1.ApiTags)('Transfers'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard, roles_guard_1.RolesGuard),
    (0, common_1.Controller)('api/v1/transfers'),
    __metadata("design:paramtypes", [transfers_service_1.TransfersService,
        reconciliation_service_1.ReconciliationService])
], TransfersController);
//# sourceMappingURL=transfers.controller.js.map