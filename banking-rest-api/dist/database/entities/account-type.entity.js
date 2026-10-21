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
exports.AccountType = exports.AccountCategory = void 0;
const typeorm_1 = require("typeorm");
var AccountCategory;
(function (AccountCategory) {
    AccountCategory["SAVINGS"] = "SAVINGS";
    AccountCategory["CURRENT"] = "CURRENT";
    AccountCategory["FIXED_DEPOSIT"] = "FIXED_DEPOSIT";
    AccountCategory["OVERDRAFT"] = "OVERDRAFT";
    AccountCategory["LOAN"] = "LOAN";
})(AccountCategory || (exports.AccountCategory = AccountCategory = {}));
let AccountType = class AccountType {
    id;
    code;
    name;
    category;
    minimumBalance;
    isInterestBearing;
    status;
    createdAt;
    updatedAt;
};
exports.AccountType = AccountType;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], AccountType.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'code', type: 'varchar', length: 20, unique: true }),
    __metadata("design:type", String)
], AccountType.prototype, "code", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'name', type: 'varchar', length: 100 }),
    __metadata("design:type", String)
], AccountType.prototype, "name", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'category',
        type: 'enum',
        enum: AccountCategory,
        default: AccountCategory.SAVINGS,
    }),
    __metadata("design:type", String)
], AccountType.prototype, "category", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'minimum_balance',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], AccountType.prototype, "minimumBalance", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'is_interest_bearing', type: 'boolean', default: false }),
    __metadata("design:type", Boolean)
], AccountType.prototype, "isInterestBearing", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' }),
    __metadata("design:type", String)
], AccountType.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], AccountType.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], AccountType.prototype, "updatedAt", void 0);
exports.AccountType = AccountType = __decorate([
    (0, typeorm_1.Entity)('account_types')
], AccountType);
//# sourceMappingURL=account-type.entity.js.map