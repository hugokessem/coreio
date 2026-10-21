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
exports.StandingOrderExecution = exports.ExecutionStatus = void 0;
const typeorm_1 = require("typeorm");
const standing_order_entity_1 = require("./standing-order.entity");
var ExecutionStatus;
(function (ExecutionStatus) {
    ExecutionStatus["SUCCESS"] = "SUCCESS";
    ExecutionStatus["FAILED"] = "FAILED";
    ExecutionStatus["SKIPPED"] = "SKIPPED";
})(ExecutionStatus || (exports.ExecutionStatus = ExecutionStatus = {}));
let StandingOrderExecution = class StandingOrderExecution {
    id;
    standingOrderId;
    standingOrder;
    executionDate;
    amount;
    transactionReference;
    status;
    failureReason;
    createdAt;
};
exports.StandingOrderExecution = StandingOrderExecution;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], StandingOrderExecution.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'standing_order_id', type: 'uuid' }),
    __metadata("design:type", String)
], StandingOrderExecution.prototype, "standingOrderId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => standing_order_entity_1.StandingOrder, { onDelete: 'CASCADE' }),
    (0, typeorm_1.JoinColumn)({ name: 'standing_order_id' }),
    __metadata("design:type", standing_order_entity_1.StandingOrder)
], StandingOrderExecution.prototype, "standingOrder", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'execution_date', type: 'date' }),
    __metadata("design:type", String)
], StandingOrderExecution.prototype, "executionDate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'amount', type: 'numeric', precision: 18, scale: 4 }),
    __metadata("design:type", String)
], StandingOrderExecution.prototype, "amount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'transaction_reference', type: 'varchar', length: 50, nullable: true }),
    __metadata("design:type", Object)
], StandingOrderExecution.prototype, "transactionReference", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: ExecutionStatus,
        default: ExecutionStatus.SUCCESS,
    }),
    __metadata("design:type", String)
], StandingOrderExecution.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'failure_reason', type: 'varchar', length: 255, nullable: true }),
    __metadata("design:type", Object)
], StandingOrderExecution.prototype, "failureReason", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], StandingOrderExecution.prototype, "createdAt", void 0);
exports.StandingOrderExecution = StandingOrderExecution = __decorate([
    (0, typeorm_1.Entity)('standing_order_executions')
], StandingOrderExecution);
//# sourceMappingURL=standing-order-execution.entity.js.map