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
exports.ServiceEntity = exports.FeeType = void 0;
const typeorm_1 = require("typeorm");
var FeeType;
(function (FeeType) {
    FeeType["FLAT"] = "FLAT";
    FeeType["PERCENTAGE"] = "PERCENTAGE";
})(FeeType || (exports.FeeType = FeeType = {}));
let ServiceEntity = class ServiceEntity {
    id;
    serviceCode;
    serviceName;
    category;
    status;
    minAmount;
    maxAmount;
    dailyLimit;
    feeType;
    feeValue;
    createdAt;
    updatedAt;
};
exports.ServiceEntity = ServiceEntity;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], ServiceEntity.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'service_code', type: 'varchar', length: 50, unique: true }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "serviceCode", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'service_name', type: 'varchar', length: 100 }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "serviceName", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'category', type: 'varchar', length: 50 }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "category", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'min_amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '1.0000',
    }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "minAmount", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'max_amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '1000000.0000',
    }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "maxAmount", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'daily_limit',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '5000000.0000',
    }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "dailyLimit", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'fee_type',
        type: 'enum',
        enum: FeeType,
        default: FeeType.FLAT,
    }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "feeType", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'fee_value',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], ServiceEntity.prototype, "feeValue", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], ServiceEntity.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], ServiceEntity.prototype, "updatedAt", void 0);
exports.ServiceEntity = ServiceEntity = __decorate([
    (0, typeorm_1.Entity)('services')
], ServiceEntity);
//# sourceMappingURL=service.entity.js.map