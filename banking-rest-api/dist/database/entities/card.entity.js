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
exports.Card = exports.CardStatus = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
const customer_entity_1 = require("./customer.entity");
const card_request_entity_1 = require("./card-request.entity");
var CardStatus;
(function (CardStatus) {
    CardStatus["ACTIVE"] = "ACTIVE";
    CardStatus["BLOCKED"] = "BLOCKED";
    CardStatus["EXPIRED"] = "EXPIRED";
    CardStatus["CANCELLED"] = "CANCELLED";
})(CardStatus || (exports.CardStatus = CardStatus = {}));
let Card = class Card {
    id;
    cardNumberMasked;
    cardHash;
    cardType;
    cardBrand;
    accountId;
    account;
    customerId;
    customer;
    status;
    expiryDate;
    dailyLimit;
    createdAt;
    updatedAt;
};
exports.Card = Card;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], Card.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'card_number_masked', type: 'varchar', length: 30 }),
    __metadata("design:type", String)
], Card.prototype, "cardNumberMasked", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'card_hash', type: 'varchar', length: 64, unique: true }),
    __metadata("design:type", String)
], Card.prototype, "cardHash", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'card_type',
        type: 'enum',
        enum: card_request_entity_1.CardType,
        default: card_request_entity_1.CardType.DEBIT,
    }),
    __metadata("design:type", String)
], Card.prototype, "cardType", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'card_brand',
        type: 'enum',
        enum: card_request_entity_1.CardBrand,
        default: card_request_entity_1.CardBrand.MASTERCARD,
    }),
    __metadata("design:type", String)
], Card.prototype, "cardBrand", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'account_id', type: 'uuid' }),
    __metadata("design:type", String)
], Card.prototype, "accountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account),
    (0, typeorm_1.JoinColumn)({ name: 'account_id' }),
    __metadata("design:type", account_entity_1.Account)
], Card.prototype, "account", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'customer_id', type: 'uuid' }),
    __metadata("design:type", String)
], Card.prototype, "customerId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => customer_entity_1.Customer),
    (0, typeorm_1.JoinColumn)({ name: 'customer_id' }),
    __metadata("design:type", customer_entity_1.Customer)
], Card.prototype, "customer", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: CardStatus,
        default: CardStatus.ACTIVE,
    }),
    __metadata("design:type", String)
], Card.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'expiry_date', type: 'varchar', length: 5 }),
    __metadata("design:type", String)
], Card.prototype, "expiryDate", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'daily_limit',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '50000.0000',
    }),
    __metadata("design:type", String)
], Card.prototype, "dailyLimit", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], Card.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], Card.prototype, "updatedAt", void 0);
exports.Card = Card = __decorate([
    (0, typeorm_1.Entity)('cards')
], Card);
//# sourceMappingURL=card.entity.js.map