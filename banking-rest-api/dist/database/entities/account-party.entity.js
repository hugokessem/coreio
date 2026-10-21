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
exports.AccountParty = exports.PartyRole = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
const customer_entity_1 = require("./customer.entity");
var PartyRole;
(function (PartyRole) {
    PartyRole["PRIMARY_OWNER"] = "PRIMARY_OWNER";
    PartyRole["JOINT_OWNER"] = "JOINT_OWNER";
    PartyRole["AUTHORIZED_SIGNATORY"] = "AUTHORIZED_SIGNATORY";
    PartyRole["GUARANTOR"] = "GUARANTOR";
    PartyRole["BENEFICIAL_OWNER"] = "BENEFICIAL_OWNER";
})(PartyRole || (exports.PartyRole = PartyRole = {}));
let AccountParty = class AccountParty {
    id;
    accountId;
    account;
    customerId;
    customer;
    partyRole;
    authorizationLevel;
    status;
    createdAt;
    updatedAt;
};
exports.AccountParty = AccountParty;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], AccountParty.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'account_id', type: 'uuid' }),
    __metadata("design:type", String)
], AccountParty.prototype, "accountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account, { onDelete: 'CASCADE' }),
    (0, typeorm_1.JoinColumn)({ name: 'account_id' }),
    __metadata("design:type", account_entity_1.Account)
], AccountParty.prototype, "account", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'customer_id', type: 'uuid' }),
    __metadata("design:type", String)
], AccountParty.prototype, "customerId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => customer_entity_1.Customer, { onDelete: 'RESTRICT' }),
    (0, typeorm_1.JoinColumn)({ name: 'customer_id' }),
    __metadata("design:type", customer_entity_1.Customer)
], AccountParty.prototype, "customer", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'party_role',
        type: 'enum',
        enum: PartyRole,
        default: PartyRole.PRIMARY_OWNER,
    }),
    __metadata("design:type", String)
], AccountParty.prototype, "partyRole", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'authorization_level', type: 'varchar', length: 50, default: 'SOLE_SIGNATORY' }),
    __metadata("design:type", String)
], AccountParty.prototype, "authorizationLevel", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' }),
    __metadata("design:type", String)
], AccountParty.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], AccountParty.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], AccountParty.prototype, "updatedAt", void 0);
exports.AccountParty = AccountParty = __decorate([
    (0, typeorm_1.Entity)('account_parties'),
    (0, typeorm_1.Index)(['accountId', 'customerId', 'partyRole'], { unique: true })
], AccountParty);
//# sourceMappingURL=account-party.entity.js.map