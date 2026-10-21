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
exports.PaymentsController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const payments_service_1 = require("./payments.service");
const payment_dto_1 = require("./dto/payment.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
let PaymentsController = class PaymentsController {
    paymentsService;
    constructor(paymentsService) {
        this.paymentsService = paymentsService;
    }
    async verifyBill(dto) {
        return this.paymentsService.verifyBill(dto);
    }
    async payBill(dto) {
        return this.paymentsService.payBill(dto);
    }
    async splitPayment(dto) {
        return this.paymentsService.splitPayment(dto);
    }
    async rtgsPayment(dto) {
        return this.paymentsService.rtgsPayment(dto);
    }
};
exports.PaymentsController = PaymentsController;
__decorate([
    (0, common_1.Post)('bill/verify'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Pre-verify bill payment parameters, fee tariff, and customer balance (Op 41: BillPaymentVerify)' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Bill payment verified' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [payment_dto_1.BillPaymentVerifyDto]),
    __metadata("design:returntype", Promise)
], PaymentsController.prototype, "verifyBill", null);
__decorate([
    (0, common_1.Post)('bill'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Process bill payment to utility/merchant (Op 40: BillPayment)' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Bill payment posted' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [payment_dto_1.BillPaymentDto]),
    __metadata("design:returntype", Promise)
], PaymentsController.prototype, "payBill", null);
__decorate([
    (0, common_1.Post)('split'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Atomic split payment distribution from 1 source to multiple destinations (Op 42: SplitPayment)' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'All split transfer legs posted atomically' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [payment_dto_1.SplitPaymentDto]),
    __metadata("design:returntype", Promise)
], PaymentsController.prototype, "splitPayment", null);
__decorate([
    (0, common_1.Post)('rtgs'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Interbank RTGS payment instruction (Op 43: RTGS - Simulated Sandbox Settlement)' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'RTGS debit posted and clearing simulated' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [payment_dto_1.RtgsPaymentDto]),
    __metadata("design:returntype", Promise)
], PaymentsController.prototype, "rtgsPayment", null);
exports.PaymentsController = PaymentsController = __decorate([
    (0, swagger_1.ApiTags)('Payments'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard),
    (0, common_1.Controller)('api/v1/payments'),
    __metadata("design:paramtypes", [payments_service_1.PaymentsService])
], PaymentsController);
//# sourceMappingURL=payments.controller.js.map