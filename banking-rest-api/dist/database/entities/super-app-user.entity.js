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
exports.SuperAppUser = exports.SuperAppStatus = void 0;
const typeorm_1 = require("typeorm");
const customer_entity_1 = require("./customer.entity");
var SuperAppStatus;
(function (SuperAppStatus) {
    SuperAppStatus["ACTIVE"] = "ACTIVE";
    SuperAppStatus["SUSPENDED"] = "SUSPENDED";
    SuperAppStatus["DEACTIVATED"] = "DEACTIVATED";
})(SuperAppStatus || (exports.SuperAppStatus = SuperAppStatus = {}));
let SuperAppUser = class SuperAppUser {
    id;
    externalUserId;
    customerId;
    customer;
    partnerCode;
    status;
    registeredAt;
    updatedAt;
};
exports.SuperAppUser = SuperAppUser;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], SuperAppUser.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'external_user_id', type: 'varchar', length: 100, unique: true }),
    __metadata("design:type", String)
], SuperAppUser.prototype, "externalUserId", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'customer_id', type: 'uuid' }),
    __metadata("design:type", String)
], SuperAppUser.prototype, "customerId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => customer_entity_1.Customer, { onDelete: 'CASCADE' }),
    (0, typeorm_1.JoinColumn)({ name: 'customer_id' }),
    __metadata("design:type", customer_entity_1.Customer)
], SuperAppUser.prototype, "customer", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'partner_code', type: 'varchar', length: 50, default: 'SUPERAPP_DEFAULT' }),
    __metadata("design:type", String)
], SuperAppUser.prototype, "partnerCode", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: SuperAppStatus,
        default: SuperAppStatus.ACTIVE,
    }),
    __metadata("design:type", String)
], SuperAppUser.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'registered_at' }),
    __metadata("design:type", Date)
], SuperAppUser.prototype, "registeredAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], SuperAppUser.prototype, "updatedAt", void 0);
exports.SuperAppUser = SuperAppUser = __decorate([
    (0, typeorm_1.Entity)('super_app_users')
], SuperAppUser);
//# sourceMappingURL=super-app-user.entity.js.map