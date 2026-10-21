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
exports.RatesController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const rates_service_1 = require("./rates.service");
const rate_dto_1 = require("./dto/rate.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
let RatesController = class RatesController {
    ratesService;
    constructor(ratesService) {
        this.ratesService = ratesService;
    }
    async lookup(dto) {
        return this.ratesService.lookup(dto);
    }
    async getAllRates(fromCurrency, toCurrency) {
        return this.ratesService.lookup({ fromCurrency, toCurrency });
    }
    async getExchangeRates(fromCurrency, toCurrency) {
        return this.ratesService.lookup({ fromCurrency, toCurrency });
    }
    async postExchangeRates(dto) {
        return this.ratesService.lookup(dto);
    }
};
exports.RatesController = RatesController;
__decorate([
    (0, common_1.Post)('lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Lookup FX exchange rates by currency pair or date (Op 35: ExchangeRates)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [rate_dto_1.RateLookupDto]),
    __metadata("design:returntype", Promise)
], RatesController.prototype, "lookup", null);
__decorate([
    (0, common_1.Get)(),
    (0, swagger_1.ApiOperation)({ summary: 'GET all active FX exchange rates - REST Canonical (Op 35: ExchangeRates)' }),
    __param(0, (0, common_1.Query)('fromCurrency')),
    __param(1, (0, common_1.Query)('toCurrency')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String, String]),
    __metadata("design:returntype", Promise)
], RatesController.prototype, "getAllRates", null);
__decorate([
    (0, common_1.Get)('exchange-rates'),
    (0, swagger_1.ApiOperation)({ summary: 'GET FX exchange rates - CoreIO Route Alias (Op 35: ExchangeRates)' }),
    __param(0, (0, common_1.Query)('fromCurrency')),
    __param(1, (0, common_1.Query)('toCurrency')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String, String]),
    __metadata("design:returntype", Promise)
], RatesController.prototype, "getExchangeRates", null);
__decorate([
    (0, common_1.Post)('exchange-rates'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'POST FX exchange rates lookup - CoreIO Route Alias (Op 35: ExchangeRates)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [rate_dto_1.RateLookupDto]),
    __metadata("design:returntype", Promise)
], RatesController.prototype, "postExchangeRates", null);
exports.RatesController = RatesController = __decorate([
    (0, swagger_1.ApiTags)('Exchange Rates'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard),
    (0, common_1.Controller)('api/v1/rates'),
    __metadata("design:paramtypes", [rates_service_1.RatesService])
], RatesController);
//# sourceMappingURL=rates.controller.js.map