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
exports.Customer = exports.AMLStatus = exports.CustomerStatus = exports.CustomerType = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
const customer_limit_entity_1 = require("./customer-limit.entity");
var CustomerType;
(function (CustomerType) {
    CustomerType["INDIVIDUAL"] = "INDIVIDUAL";
    CustomerType["CORPORATE"] = "CORPORATE";
})(CustomerType || (exports.CustomerType = CustomerType = {}));
var CustomerStatus;
(function (CustomerStatus) {
    CustomerStatus["ACTIVE"] = "ACTIVE";
    CustomerStatus["INACTIVE"] = "INACTIVE";
    CustomerStatus["SUSPENDED"] = "SUSPENDED";
    CustomerStatus["DECEASED"] = "DECEASED";
})(CustomerStatus || (exports.CustomerStatus = CustomerStatus = {}));
var AMLStatus;
(function (AMLStatus) {
    AMLStatus["CLEAR"] = "CLEAR";
    AMLStatus["SUSPECT"] = "SUSPECT";
    AMLStatus["BLOCKED"] = "BLOCKED";
})(AMLStatus || (exports.AMLStatus = AMLStatus = {}));
let Customer = class Customer {
    id;
    customerNumber;
    customerType;
    firstName;
    middleName;
    lastName;
    legalName;
    dateOfBirth;
    gender;
    phoneNumber;
    email;
    address;
    nationalId;
    tin;
    kycLevel;
    amlStatus;
    status;
    isAuthorized;
    createdBy;
    authorizedBy;
    authorizedAt;
    accounts;
    limits;
    createdAt;
    updatedAt;
    get fullName() {
        if (this.legalName)
            return this.legalName;
        return [this.firstName, this.middleName, this.lastName].filter(Boolean).join(' ');
    }
};
exports.Customer = Customer;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], Customer.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'customer_number', type: 'varchar', length: 20, unique: true }),
    __metadata("design:type", String)
], Customer.prototype, "customerNumber", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'customer_type',
        type: 'enum',
        enum: CustomerType,
        default: CustomerType.INDIVIDUAL,
    }),
    __metadata("design:type", String)
], Customer.prototype, "customerType", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'first_name', type: 'varchar', length: 100, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "firstName", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'middle_name', type: 'varchar', length: 100, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "middleName", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'last_name', type: 'varchar', length: 100, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "lastName", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'legal_name', type: 'varchar', length: 255, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "legalName", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'date_of_birth', type: 'date', nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "dateOfBirth", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'gender', type: 'varchar', length: 10, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "gender", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'phone_number', type: 'varchar', length: 20, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "phoneNumber", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'email', type: 'varchar', length: 100, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "email", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'address', type: 'text', nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "address", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'national_id', type: 'varchar', length: 50, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "nationalId", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'tin', type: 'varchar', length: 30, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "tin", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'kyc_level', type: 'varchar', length: 20, default: 'LEVEL_2' }),
    __metadata("design:type", String)
], Customer.prototype, "kycLevel", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'aml_status',
        type: 'enum',
        enum: AMLStatus,
        default: AMLStatus.CLEAR,
    }),
    __metadata("design:type", String)
], Customer.prototype, "amlStatus", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: CustomerStatus,
        default: CustomerStatus.ACTIVE,
    }),
    __metadata("design:type", String)
], Customer.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'is_authorized', type: 'boolean', default: true }),
    __metadata("design:type", Boolean)
], Customer.prototype, "isAuthorized", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'created_by', type: 'varchar', length: 100, default: 'SYSTEM' }),
    __metadata("design:type", String)
], Customer.prototype, "createdBy", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'authorized_by', type: 'varchar', length: 100, nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "authorizedBy", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'authorized_at', type: 'timestamp', nullable: true }),
    __metadata("design:type", Object)
], Customer.prototype, "authorizedAt", void 0);
__decorate([
    (0, typeorm_1.OneToMany)(() => account_entity_1.Account, (account) => account.customer),
    __metadata("design:type", Array)
], Customer.prototype, "accounts", void 0);
__decorate([
    (0, typeorm_1.OneToMany)(() => customer_limit_entity_1.CustomerLimit, (limit) => limit.customer),
    __metadata("design:type", Array)
], Customer.prototype, "limits", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], Customer.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], Customer.prototype, "updatedAt", void 0);
exports.Customer = Customer = __decorate([
    (0, typeorm_1.Entity)('customers')
], Customer);
//# sourceMappingURL=customer.entity.js.map