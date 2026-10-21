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
exports.Branch = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
let Branch = class Branch {
    id;
    branchCode;
    branchName;
    address;
    city;
    swiftCode;
    status;
    accounts;
    createdAt;
    updatedAt;
};
exports.Branch = Branch;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], Branch.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'branch_code', type: 'varchar', length: 10, unique: true }),
    __metadata("design:type", String)
], Branch.prototype, "branchCode", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'branch_name', type: 'varchar', length: 150 }),
    __metadata("design:type", String)
], Branch.prototype, "branchName", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'address', type: 'text', nullable: true }),
    __metadata("design:type", Object)
], Branch.prototype, "address", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'city', type: 'varchar', length: 100, default: 'Addis Ababa' }),
    __metadata("design:type", String)
], Branch.prototype, "city", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'swift_code', type: 'varchar', length: 20, nullable: true }),
    __metadata("design:type", Object)
], Branch.prototype, "swiftCode", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' }),
    __metadata("design:type", String)
], Branch.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.OneToMany)(() => account_entity_1.Account, (account) => account.branch),
    __metadata("design:type", Array)
], Branch.prototype, "accounts", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], Branch.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], Branch.prototype, "updatedAt", void 0);
exports.Branch = Branch = __decorate([
    (0, typeorm_1.Entity)('branches')
], Branch);
//# sourceMappingURL=branch.entity.js.map