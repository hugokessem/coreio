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
exports.ExchangeRate = void 0;
const typeorm_1 = require("typeorm");
let ExchangeRate = class ExchangeRate {
    id;
    fromCurrency;
    toCurrency;
    buyRate;
    sellRate;
    midRate;
    effectiveDate;
    createdAt;
    updatedAt;
};
exports.ExchangeRate = ExchangeRate;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], ExchangeRate.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'from_currency', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], ExchangeRate.prototype, "fromCurrency", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'to_currency', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], ExchangeRate.prototype, "toCurrency", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'buy_rate', type: 'numeric', precision: 18, scale: 6 }),
    __metadata("design:type", String)
], ExchangeRate.prototype, "buyRate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'sell_rate', type: 'numeric', precision: 18, scale: 6 }),
    __metadata("design:type", String)
], ExchangeRate.prototype, "sellRate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'mid_rate', type: 'numeric', precision: 18, scale: 6 }),
    __metadata("design:type", String)
], ExchangeRate.prototype, "midRate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'effective_date', type: 'date' }),
    __metadata("design:type", String)
], ExchangeRate.prototype, "effectiveDate", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], ExchangeRate.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], ExchangeRate.prototype, "updatedAt", void 0);
exports.ExchangeRate = ExchangeRate = __decorate([
    (0, typeorm_1.Entity)('exchange_rates'),
    (0, typeorm_1.Index)(['fromCurrency', 'toCurrency', 'effectiveDate'], { unique: true })
], ExchangeRate);
//# sourceMappingURL=exchange-rate.entity.js.map