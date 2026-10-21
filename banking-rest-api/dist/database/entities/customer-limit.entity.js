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
exports.CustomerLimit = void 0;
const typeorm_1 = require("typeorm");
const customer_entity_1 = require("./customer.entity");
let CustomerLimit = class CustomerLimit {
    id;
    customerId;
    customer;
    serviceCode;
    dailyLimit;
    perTransactionLimit;
    dailyUsed;
    lastResetDate;
    createdAt;
    updatedAt;
};
exports.CustomerLimit = CustomerLimit;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], CustomerLimit.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'customer_id', type: 'uuid' }),
    __metadata("design:type", String)
], CustomerLimit.prototype, "customerId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => customer_entity_1.Customer, (c) => c.limits, { onDelete: 'CASCADE' }),
    (0, typeorm_1.JoinColumn)({ name: 'customer_id' }),
    __metadata("design:type", customer_entity_1.Customer)
], CustomerLimit.prototype, "customer", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'service_code', type: 'varchar', length: 50 }),
    __metadata("design:type", String)
], CustomerLimit.prototype, "serviceCode", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'daily_limit',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '500000.0000',
    }),
    __metadata("design:type", String)
], CustomerLimit.prototype, "dailyLimit", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'per_transaction_limit',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '100000.0000',
    }),
    __metadata("design:type", String)
], CustomerLimit.prototype, "perTransactionLimit", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'daily_used',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], CustomerLimit.prototype, "dailyUsed", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'last_reset_date', type: 'date', default: () => 'CURRENT_DATE' }),
    __metadata("design:type", String)
], CustomerLimit.prototype, "lastResetDate", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], CustomerLimit.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], CustomerLimit.prototype, "updatedAt", void 0);
exports.CustomerLimit = CustomerLimit = __decorate([
    (0, typeorm_1.Entity)('customer_limits'),
    (0, typeorm_1.Index)(['customerId', 'serviceCode'], { unique: true })
], CustomerLimit);
//# sourceMappingURL=customer-limit.entity.js.map