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
exports.IpsController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const ips_service_1 = require("./ips.service");
const ips_dto_1 = require("./dto/ips.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
let IpsController = class IpsController {
    ipsService;
    constructor(ipsService) {
        this.ipsService = ipsService;
    }
    async accountLookup(dto) {
        return this.ipsService.accountLookup(dto);
    }
    async transfer(dto) {
        return this.ipsService.executeTransfer(dto);
    }
    async qrPayment(dto) {
        return this.ipsService.qrPayment(dto);
    }
    async status(dto) {
        return this.ipsService.getStatus(dto);
    }
    async statusCheck(dto) {
        return this.ipsService.getStatus(dto);
    }
};
exports.IpsController = IpsController;
__decorate([
    (0, common_1.Post)('account-lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'IPS Interbank ISO 20022 account & title lookup (Op 50: IPS:AccountLookup)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [ips_dto_1.IpsAccountLookupDto]),
    __metadata("design:returntype", Promise)
], IpsController.prototype, "accountLookup", null);
__decorate([
    (0, common_1.Post)('transfer'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Execute outgoing/incoming ISO 20022 IPS interbank fund transfer (Op 51: IPS:FundTransfer)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [ips_dto_1.IpsTransferDto]),
    __metadata("design:returntype", Promise)
], IpsController.prototype, "transfer", null);
__decorate([
    (0, common_1.Post)('qr-payment'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Process merchant QR payment via IPS switch (Op 53: IPS:QRPayment)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [ips_dto_1.IpsQrPaymentDto]),
    __metadata("design:returntype", Promise)
], IpsController.prototype, "qrPayment", null);
__decorate([
    (0, common_1.Post)('status'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Inquire clearing status of an IPS transaction (Op 52: IPS:StatusCheck)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [ips_dto_1.IpsStatusDto]),
    __metadata("design:returntype", Promise)
], IpsController.prototype, "status", null);
__decorate([
    (0, common_1.Post)('status-check'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Inquire clearing status of an IPS transaction - CoreIO Route Alias (Op 52: IPS:StatusCheck)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [ips_dto_1.IpsStatusDto]),
    __metadata("design:returntype", Promise)
], IpsController.prototype, "statusCheck", null);
exports.IpsController = IpsController = __decorate([
    (0, swagger_1.ApiTags)('IPS (EthSwitch Interbank)'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard),
    (0, common_1.Controller)('api/v1/ips'),
    __metadata("design:paramtypes", [ips_service_1.IpsService])
], IpsController);
//# sourceMappingURL=ips.controller.js.map