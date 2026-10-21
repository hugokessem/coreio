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
exports.StandingOrder = exports.StandingOrderStatus = exports.StandingOrderFrequency = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
var StandingOrderFrequency;
(function (StandingOrderFrequency) {
    StandingOrderFrequency["DAILY"] = "DAILY";
    StandingOrderFrequency["WEEKLY"] = "WEEKLY";
    StandingOrderFrequency["MONTHLY"] = "MONTHLY";
    StandingOrderFrequency["QUARTERLY"] = "QUARTERLY";
    StandingOrderFrequency["ANNUALLY"] = "ANNUALLY";
})(StandingOrderFrequency || (exports.StandingOrderFrequency = StandingOrderFrequency = {}));
var StandingOrderStatus;
(function (StandingOrderStatus) {
    StandingOrderStatus["ACTIVE"] = "ACTIVE";
    StandingOrderStatus["PAUSED"] = "PAUSED";
    StandingOrderStatus["CANCELLED"] = "CANCELLED";
    StandingOrderStatus["COMPLETED"] = "COMPLETED";
})(StandingOrderStatus || (exports.StandingOrderStatus = StandingOrderStatus = {}));
let StandingOrder = class StandingOrder {
    id;
    orderReference;
    sourceAccountId;
    sourceAccount;
    destinationAccountId;
    destinationAccount;
    beneficiaryName;
    beneficiaryAccount;
    beneficiaryBankCode;
    amount;
    currencyCode;
    frequency;
    startDate;
    endDate;
    nextExecutionDate;
    status;
    failureCount;
    lastExecutedAt;
    createdAt;
    updatedAt;
};
exports.StandingOrder = StandingOrder;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], StandingOrder.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'order_reference', type: 'varchar', length: 50, unique: true }),
    __metadata("design:type", String)
], StandingOrder.prototype, "orderReference", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'source_account_id', type: 'uuid' }),
    __metadata("design:type", String)
], StandingOrder.prototype, "sourceAccountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account, { onDelete: 'RESTRICT' }),
    (0, typeorm_1.JoinColumn)({ name: 'source_account_id' }),
    __metadata("design:type", account_entity_1.Account)
], StandingOrder.prototype, "sourceAccount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'destination_account_id', type: 'uuid', nullable: true }),
    __metadata("design:type", Object)
], StandingOrder.prototype, "destinationAccountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account, { nullable: true, onDelete: 'SET NULL' }),
    (0, typeorm_1.JoinColumn)({ name: 'destination_account_id' }),
    __metadata("design:type", Object)
], StandingOrder.prototype, "destinationAccount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'beneficiary_name', type: 'varchar', length: 150 }),
    __metadata("design:type", String)
], StandingOrder.prototype, "beneficiaryName", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'beneficiary_account', type: 'varchar', length: 50 }),
    __metadata("design:type", String)
], StandingOrder.prototype, "beneficiaryAccount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'beneficiary_bank_code', type: 'varchar', length: 30, nullable: true }),
    __metadata("design:type", Object)
], StandingOrder.prototype, "beneficiaryBankCode", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
    }),
    __metadata("design:type", String)
], StandingOrder.prototype, "amount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'currency_code', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], StandingOrder.prototype, "currencyCode", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'frequency',
        type: 'enum',
        enum: StandingOrderFrequency,
        default: StandingOrderFrequency.MONTHLY,
    }),
    __metadata("design:type", String)
], StandingOrder.prototype, "frequency", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'start_date', type: 'date' }),
    __metadata("design:type", String)
], StandingOrder.prototype, "startDate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'end_date', type: 'date', nullable: true }),
    __metadata("design:type", Object)
], StandingOrder.prototype, "endDate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'next_execution_date', type: 'date' }),
    __metadata("design:type", String)
], StandingOrder.prototype, "nextExecutionDate", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: StandingOrderStatus,
        default: StandingOrderStatus.ACTIVE,
    }),
    __metadata("design:type", String)
], StandingOrder.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'failure_count', type: 'int', default: 0 }),
    __metadata("design:type", Number)
], StandingOrder.prototype, "failureCount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'last_executed_at', type: 'timestamp', nullable: true }),
    __metadata("design:type", Object)
], StandingOrder.prototype, "lastExecutedAt", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], StandingOrder.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], StandingOrder.prototype, "updatedAt", void 0);
exports.StandingOrder = StandingOrder = __decorate([
    (0, typeorm_1.Entity)('standing_orders')
], StandingOrder);
//# sourceMappingURL=standing-order.entity.js.map