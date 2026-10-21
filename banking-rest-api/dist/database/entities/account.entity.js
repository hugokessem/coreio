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
exports.Account = exports.AccountStatus = void 0;
const typeorm_1 = require("typeorm");
const customer_entity_1 = require("./customer.entity");
const branch_entity_1 = require("./branch.entity");
const currency_entity_1 = require("./currency.entity");
const account_type_entity_1 = require("./account-type.entity");
const account_balance_entity_1 = require("./account-balance.entity");
const account_hold_entity_1 = require("./account-hold.entity");
var AccountStatus;
(function (AccountStatus) {
    AccountStatus["ACTIVE"] = "ACTIVE";
    AccountStatus["DORMANT"] = "DORMANT";
    AccountStatus["FROZEN"] = "FROZEN";
    AccountStatus["RESTRICTED"] = "RESTRICTED";
    AccountStatus["CLOSED"] = "CLOSED";
})(AccountStatus || (exports.AccountStatus = AccountStatus = {}));
let Account = class Account {
    id;
    accountNumber;
    customerId;
    customer;
    accountTypeId;
    accountType;
    branchId;
    branch;
    currencyCode;
    currency;
    accountName;
    status;
    isPostingRestricted;
    postingRestrictionReason;
    isDebitBlocked;
    isCreditBlocked;
    openedAt;
    closedAt;
    lastActivityAt;
    balance;
    holds;
    createdAt;
    updatedAt;
};
exports.Account = Account;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], Account.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'account_number', type: 'varchar', length: 30, unique: true }),
    __metadata("design:type", String)
], Account.prototype, "accountNumber", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'customer_id', type: 'uuid' }),
    __metadata("design:type", String)
], Account.prototype, "customerId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => customer_entity_1.Customer, (customer) => customer.accounts),
    (0, typeorm_1.JoinColumn)({ name: 'customer_id' }),
    __metadata("design:type", customer_entity_1.Customer)
], Account.prototype, "customer", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'account_type_id', type: 'uuid' }),
    __metadata("design:type", String)
], Account.prototype, "accountTypeId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_type_entity_1.AccountType),
    (0, typeorm_1.JoinColumn)({ name: 'account_type_id' }),
    __metadata("design:type", account_type_entity_1.AccountType)
], Account.prototype, "accountType", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'branch_id', type: 'uuid' }),
    __metadata("design:type", String)
], Account.prototype, "branchId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => branch_entity_1.Branch, (branch) => branch.accounts),
    (0, typeorm_1.JoinColumn)({ name: 'branch_id' }),
    __metadata("design:type", branch_entity_1.Branch)
], Account.prototype, "branch", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'currency_code', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], Account.prototype, "currencyCode", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => currency_entity_1.Currency),
    (0, typeorm_1.JoinColumn)({ name: 'currency_code' }),
    __metadata("design:type", currency_entity_1.Currency)
], Account.prototype, "currency", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'account_name', type: 'varchar', length: 255 }),
    __metadata("design:type", String)
], Account.prototype, "accountName", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: AccountStatus,
        default: AccountStatus.ACTIVE,
    }),
    __metadata("design:type", String)
], Account.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'is_posting_restricted', type: 'boolean', default: false }),
    __metadata("design:type", Boolean)
], Account.prototype, "isPostingRestricted", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'posting_restriction_reason', type: 'varchar', length: 255, nullable: true }),
    __metadata("design:type", Object)
], Account.prototype, "postingRestrictionReason", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'is_debit_blocked', type: 'boolean', default: false }),
    __metadata("design:type", Boolean)
], Account.prototype, "isDebitBlocked", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'is_credit_blocked', type: 'boolean', default: false }),
    __metadata("design:type", Boolean)
], Account.prototype, "isCreditBlocked", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'opened_at', type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' }),
    __metadata("design:type", Date)
], Account.prototype, "openedAt", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'closed_at', type: 'timestamp', nullable: true }),
    __metadata("design:type", Object)
], Account.prototype, "closedAt", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'last_activity_at', type: 'timestamp', nullable: true }),
    __metadata("design:type", Object)
], Account.prototype, "lastActivityAt", void 0);
__decorate([
    (0, typeorm_1.OneToOne)(() => account_balance_entity_1.AccountBalance, (balance) => balance.account, { cascade: true }),
    __metadata("design:type", account_balance_entity_1.AccountBalance)
], Account.prototype, "balance", void 0);
__decorate([
    (0, typeorm_1.OneToMany)(() => account_hold_entity_1.AccountHold, (hold) => hold.account),
    __metadata("design:type", Array)
], Account.prototype, "holds", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], Account.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], Account.prototype, "updatedAt", void 0);
exports.Account = Account = __decorate([
    (0, typeorm_1.Entity)('accounts')
], Account);
//# sourceMappingURL=account.entity.js.map