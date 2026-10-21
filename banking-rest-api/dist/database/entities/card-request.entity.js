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
exports.CardRequest = exports.CardRequestStatus = exports.CardRequestType = exports.CardBrand = exports.CardType = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
const customer_entity_1 = require("./customer.entity");
const branch_entity_1 = require("./branch.entity");
var CardType;
(function (CardType) {
    CardType["DEBIT"] = "DEBIT";
    CardType["PREPAID"] = "PREPAID";
    CardType["CREDIT"] = "CREDIT";
})(CardType || (exports.CardType = CardType = {}));
var CardBrand;
(function (CardBrand) {
    CardBrand["MASTERCARD"] = "MASTERCARD";
    CardBrand["VISA"] = "VISA";
    CardBrand["LOCAL"] = "LOCAL";
})(CardBrand || (exports.CardBrand = CardBrand = {}));
var CardRequestType;
(function (CardRequestType) {
    CardRequestType["NEW"] = "NEW";
    CardRequestType["REPLACEMENT"] = "REPLACEMENT";
    CardRequestType["RENEWAL"] = "RENEWAL";
})(CardRequestType || (exports.CardRequestType = CardRequestType = {}));
var CardRequestStatus;
(function (CardRequestStatus) {
    CardRequestStatus["PENDING"] = "PENDING";
    CardRequestStatus["APPROVED"] = "APPROVED";
    CardRequestStatus["PRINTING"] = "PRINTING";
    CardRequestStatus["DISPATCHED"] = "DISPATCHED";
    CardRequestStatus["DELIVERED"] = "DELIVERED";
    CardRequestStatus["REJECTED"] = "REJECTED";
})(CardRequestStatus || (exports.CardRequestStatus = CardRequestStatus = {}));
let CardRequest = class CardRequest {
    id;
    requestReference;
    accountId;
    account;
    customerId;
    customer;
    cardType;
    cardBrand;
    requestType;
    status;
    targetBranchId;
    targetBranch;
    rejectionReason;
    createdAt;
    updatedAt;
};
exports.CardRequest = CardRequest;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], CardRequest.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'request_reference', type: 'varchar', length: 50, unique: true }),
    __metadata("design:type", String)
], CardRequest.prototype, "requestReference", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'account_id', type: 'uuid' }),
    __metadata("design:type", String)
], CardRequest.prototype, "accountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account),
    (0, typeorm_1.JoinColumn)({ name: 'account_id' }),
    __metadata("design:type", account_entity_1.Account)
], CardRequest.prototype, "account", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'customer_id', type: 'uuid' }),
    __metadata("design:type", String)
], CardRequest.prototype, "customerId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => customer_entity_1.Customer),
    (0, typeorm_1.JoinColumn)({ name: 'customer_id' }),
    __metadata("design:type", customer_entity_1.Customer)
], CardRequest.prototype, "customer", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'card_type',
        type: 'enum',
        enum: CardType,
        default: CardType.DEBIT,
    }),
    __metadata("design:type", String)
], CardRequest.prototype, "cardType", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'card_brand',
        type: 'enum',
        enum: CardBrand,
        default: CardBrand.MASTERCARD,
    }),
    __metadata("design:type", String)
], CardRequest.prototype, "cardBrand", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'request_type',
        type: 'enum',
        enum: CardRequestType,
        default: CardRequestType.NEW,
    }),
    __metadata("design:type", String)
], CardRequest.prototype, "requestType", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: CardRequestStatus,
        default: CardRequestStatus.PENDING,
    }),
    __metadata("design:type", String)
], CardRequest.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'target_branch_id', type: 'uuid' }),
    __metadata("design:type", String)
], CardRequest.prototype, "targetBranchId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => branch_entity_1.Branch),
    (0, typeorm_1.JoinColumn)({ name: 'target_branch_id' }),
    __metadata("design:type", branch_entity_1.Branch)
], CardRequest.prototype, "targetBranch", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'rejection_reason', type: 'varchar', length: 255, nullable: true }),
    __metadata("design:type", Object)
], CardRequest.prototype, "rejectionReason", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], CardRequest.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], CardRequest.prototype, "updatedAt", void 0);
exports.CardRequest = CardRequest = __decorate([
    (0, typeorm_1.Entity)('card_requests')
], CardRequest);
//# sourceMappingURL=card-request.entity.js.map