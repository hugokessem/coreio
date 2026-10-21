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
exports.WalletController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const wallet_service_1 = require("./wallet.service");
const wallet_dto_1 = require("./dto/wallet.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
let WalletController = class WalletController {
    walletService;
    constructor(walletService) {
        this.walletService = walletService;
    }
    async customerLookup(dto) {
        return this.walletService.customerLookup(dto);
    }
    async transfer(dto) {
        return this.walletService.transferToWallet(dto);
    }
    async transferToWallet(dto) {
        return this.walletService.transferToWallet(dto);
    }
    async transferFromWallet(dto) {
        return this.walletService.transferFromWallet(dto);
    }
    async agentLookup(dto) {
        return this.walletService.agentLookup(dto);
    }
    async agentTransfer(dto) {
        return this.walletService.agentTransfer(dto);
    }
};
exports.WalletController = WalletController;
__decorate([
    (0, common_1.Post)('customer-lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Lookup customer profile and linked accounts by CBE Birr phone (Op 54: Wallet:CustomerAccountLookup)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [wallet_dto_1.WalletCustomerLookupDto]),
    __metadata("design:returntype", Promise)
], WalletController.prototype, "customerLookup", null);
__decorate([
    (0, common_1.Post)('transfer'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Transfer funds from bank account to CBE Birr mobile wallet - Canonical (Op 55: Wallet:CustomerFundTransfer)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [wallet_dto_1.AccountToWalletTransferDto]),
    __metadata("design:returntype", Promise)
], WalletController.prototype, "transfer", null);
__decorate([
    (0, common_1.Post)('transfer-to-wallet'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Transfer funds from bank account to CBE Birr mobile wallet - Legacy Alias (Op 55: Wallet:CustomerFundTransfer)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [wallet_dto_1.AccountToWalletTransferDto]),
    __metadata("design:returntype", Promise)
], WalletController.prototype, "transferToWallet", null);
__decorate([
    (0, common_1.Post)('transfer-from-wallet'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Transfer funds from CBE Birr mobile wallet to bank account' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [wallet_dto_1.WalletToAccountTransferDto]),
    __metadata("design:returntype", Promise)
], WalletController.prototype, "transferFromWallet", null);
__decorate([
    (0, common_1.Post)('agent-lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Lookup CBE Birr registered agent details and float status (Op 56: Wallet:AgentAccountLookup)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [wallet_dto_1.WalletAgentLookupDto]),
    __metadata("design:returntype", Promise)
], WalletController.prototype, "agentLookup", null);
__decorate([
    (0, common_1.Post)('agent-transfer'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Disburse liquidity float to CBE Birr Agent wallet (Op 57: Wallet:AgentFundTransfer)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [wallet_dto_1.WalletAgentTransferDto]),
    __metadata("design:returntype", Promise)
], WalletController.prototype, "agentTransfer", null);
exports.WalletController = WalletController = __decorate([
    (0, swagger_1.ApiTags)('CBE Birr Wallet'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard),
    (0, common_1.Controller)('api/v1/wallet'),
    __metadata("design:paramtypes", [wallet_service_1.WalletService])
], WalletController);
//# sourceMappingURL=wallet.controller.js.map