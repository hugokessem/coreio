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
exports.AccountBalance = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
let AccountBalance = class AccountBalance {
    id;
    accountId;
    account;
    ledgerBalance;
    availableBalance;
    lockedBalance;
    currencyCode;
    version;
    updatedAt;
};
exports.AccountBalance = AccountBalance;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], AccountBalance.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'account_id', type: 'uuid', unique: true }),
    __metadata("design:type", String)
], AccountBalance.prototype, "accountId", void 0);
__decorate([
    (0, typeorm_1.OneToOne)(() => account_entity_1.Account, (account) => account.balance, { onDelete: 'CASCADE' }),
    (0, typeorm_1.JoinColumn)({ name: 'account_id' }),
    __metadata("design:type", account_entity_1.Account)
], AccountBalance.prototype, "account", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'ledger_balance',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], AccountBalance.prototype, "ledgerBalance", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'available_balance',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], AccountBalance.prototype, "availableBalance", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'locked_balance',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], AccountBalance.prototype, "lockedBalance", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'currency_code', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], AccountBalance.prototype, "currencyCode", void 0);
__decorate([
    (0, typeorm_1.VersionColumn)({ default: 1 }),
    __metadata("design:type", Number)
], AccountBalance.prototype, "version", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], AccountBalance.prototype, "updatedAt", void 0);
exports.AccountBalance = AccountBalance = __decorate([
    (0, typeorm_1.Entity)('account_balances')
], AccountBalance);
//# sourceMappingURL=account-balance.entity.js.map