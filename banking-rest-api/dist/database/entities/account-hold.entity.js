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
exports.AccountHold = exports.HoldStatus = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
var HoldStatus;
(function (HoldStatus) {
    HoldStatus["ACTIVE"] = "ACTIVE";
    HoldStatus["RELEASED"] = "RELEASED";
    HoldStatus["SETTLED"] = "SETTLED";
    HoldStatus["EXPIRED"] = "EXPIRED";
})(HoldStatus || (exports.HoldStatus = HoldStatus = {}));
let AccountHold = class AccountHold {
    id;
    holdReference;
    accountId;
    account;
    amount;
    remainingAmount;
    currencyCode;
    reason;
    lockEvent;
    status;
    fromDate;
    toDate;
    releasedAt;
    releasedReason;
    createdBy;
    createdAt;
    updatedAt;
};
exports.AccountHold = AccountHold;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], AccountHold.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'hold_reference', type: 'varchar', length: 50, unique: true }),
    __metadata("design:type", String)
], AccountHold.prototype, "holdReference", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'account_id', type: 'uuid' }),
    __metadata("design:type", String)
], AccountHold.prototype, "accountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account, (account) => account.holds, { onDelete: 'CASCADE' }),
    (0, typeorm_1.JoinColumn)({ name: 'account_id' }),
    __metadata("design:type", account_entity_1.Account)
], AccountHold.prototype, "account", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
    }),
    __metadata("design:type", String)
], AccountHold.prototype, "amount", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'remaining_amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], AccountHold.prototype, "remainingAmount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'currency_code', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], AccountHold.prototype, "currencyCode", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'reason', type: 'varchar', length: 255 }),
    __metadata("design:type", String)
], AccountHold.prototype, "reason", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'lock_event', type: 'varchar', length: 50, default: 'HOLD' }),
    __metadata("design:type", String)
], AccountHold.prototype, "lockEvent", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: HoldStatus,
        default: HoldStatus.ACTIVE,
    }),
    __metadata("design:type", String)
], AccountHold.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'from_date', type: 'date', nullable: true }),
    __metadata("design:type", Object)
], AccountHold.prototype, "fromDate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'to_date', type: 'date', nullable: true }),
    __metadata("design:type", Object)
], AccountHold.prototype, "toDate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'released_at', type: 'timestamp', nullable: true }),
    __metadata("design:type", Object)
], AccountHold.prototype, "releasedAt", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'released_reason', type: 'varchar', length: 255, nullable: true }),
    __metadata("design:type", Object)
], AccountHold.prototype, "releasedReason", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'created_by', type: 'varchar', length: 100, default: 'SYSTEM' }),
    __metadata("design:type", String)
], AccountHold.prototype, "createdBy", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], AccountHold.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], AccountHold.prototype, "updatedAt", void 0);
exports.AccountHold = AccountHold = __decorate([
    (0, typeorm_1.Entity)('account_holds')
], AccountHold);
//# sourceMappingURL=account-hold.entity.js.map