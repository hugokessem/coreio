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
exports.LedgerEntry = exports.EntryType = void 0;
const typeorm_1 = require("typeorm");
const transaction_entity_1 = require("./transaction.entity");
const account_entity_1 = require("./account.entity");
var EntryType;
(function (EntryType) {
    EntryType["DEBIT"] = "DEBIT";
    EntryType["CREDIT"] = "CREDIT";
})(EntryType || (exports.EntryType = EntryType = {}));
let LedgerEntry = class LedgerEntry {
    id;
    transactionId;
    transaction;
    accountId;
    account;
    entryType;
    amount;
    currencyCode;
    balanceAfter;
    narration;
    valueDate;
    postedAt;
};
exports.LedgerEntry = LedgerEntry;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], LedgerEntry.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'transaction_id', type: 'uuid' }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "transactionId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => transaction_entity_1.Transaction, (tx) => tx.ledgerEntries, { onDelete: 'CASCADE' }),
    (0, typeorm_1.JoinColumn)({ name: 'transaction_id' }),
    __metadata("design:type", transaction_entity_1.Transaction)
], LedgerEntry.prototype, "transaction", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'account_id', type: 'uuid' }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "accountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account, { onDelete: 'RESTRICT' }),
    (0, typeorm_1.JoinColumn)({ name: 'account_id' }),
    __metadata("design:type", account_entity_1.Account)
], LedgerEntry.prototype, "account", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'entry_type',
        type: 'enum',
        enum: EntryType,
    }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "entryType", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
    }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "amount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'currency_code', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "currencyCode", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'balance_after',
        type: 'numeric',
        precision: 18,
        scale: 4,
    }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "balanceAfter", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'narration', type: 'varchar', length: 255 }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "narration", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'value_date', type: 'date', default: () => 'CURRENT_DATE' }),
    __metadata("design:type", String)
], LedgerEntry.prototype, "valueDate", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'posted_at' }),
    __metadata("design:type", Date)
], LedgerEntry.prototype, "postedAt", void 0);
exports.LedgerEntry = LedgerEntry = __decorate([
    (0, typeorm_1.Entity)('ledger_entries')
], LedgerEntry);
//# sourceMappingURL=ledger-entry.entity.js.map