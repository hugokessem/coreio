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
exports.Currency = void 0;
const typeorm_1 = require("typeorm");
let Currency = class Currency {
    code;
    name;
    numericCode;
    decimalPlaces;
    isBaseCurrency;
    status;
    createdAt;
    updatedAt;
};
exports.Currency = Currency;
__decorate([
    (0, typeorm_1.PrimaryColumn)({ type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], Currency.prototype, "code", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'name', type: 'varchar', length: 50 }),
    __metadata("design:type", String)
], Currency.prototype, "name", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'numeric_code', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], Currency.prototype, "numericCode", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'decimal_places', type: 'int', default: 2 }),
    __metadata("design:type", Number)
], Currency.prototype, "decimalPlaces", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'is_base_currency', type: 'boolean', default: false }),
    __metadata("design:type", Boolean)
], Currency.prototype, "isBaseCurrency", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' }),
    __metadata("design:type", String)
], Currency.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], Currency.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], Currency.prototype, "updatedAt", void 0);
exports.Currency = Currency = __decorate([
    (0, typeorm_1.Entity)('currencies')
], Currency);
//# sourceMappingURL=currency.entity.js.map