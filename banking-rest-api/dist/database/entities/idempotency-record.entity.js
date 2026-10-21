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
exports.IdempotencyRecord = exports.IdempotencyStatus = void 0;
const typeorm_1 = require("typeorm");
var IdempotencyStatus;
(function (IdempotencyStatus) {
    IdempotencyStatus["PENDING"] = "PENDING";
    IdempotencyStatus["COMPLETED"] = "COMPLETED";
    IdempotencyStatus["FAILED"] = "FAILED";
})(IdempotencyStatus || (exports.IdempotencyStatus = IdempotencyStatus = {}));
let IdempotencyRecord = class IdempotencyRecord {
    id;
    idempotencyKey;
    requestHash;
    operation;
    callerId;
    status;
    transactionReference;
    responsePayload;
    httpStatusCode;
    createdAt;
    updatedAt;
};
exports.IdempotencyRecord = IdempotencyRecord;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], IdempotencyRecord.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'idempotency_key', type: 'varchar', length: 100, unique: true }),
    __metadata("design:type", String)
], IdempotencyRecord.prototype, "idempotencyKey", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'request_hash', type: 'varchar', length: 64 }),
    __metadata("design:type", String)
], IdempotencyRecord.prototype, "requestHash", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'operation', type: 'varchar', length: 100 }),
    __metadata("design:type", String)
], IdempotencyRecord.prototype, "operation", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'caller_id', type: 'varchar', length: 100, nullable: true }),
    __metadata("design:type", Object)
], IdempotencyRecord.prototype, "callerId", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: IdempotencyStatus,
        default: IdempotencyStatus.PENDING,
    }),
    __metadata("design:type", String)
], IdempotencyRecord.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'transaction_reference', type: 'varchar', length: 50, nullable: true }),
    __metadata("design:type", Object)
], IdempotencyRecord.prototype, "transactionReference", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'response_payload', type: 'jsonb', nullable: true }),
    __metadata("design:type", Object)
], IdempotencyRecord.prototype, "responsePayload", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'http_status_code', type: 'int', nullable: true }),
    __metadata("design:type", Object)
], IdempotencyRecord.prototype, "httpStatusCode", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], IdempotencyRecord.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], IdempotencyRecord.prototype, "updatedAt", void 0);
exports.IdempotencyRecord = IdempotencyRecord = __decorate([
    (0, typeorm_1.Entity)('idempotency_records')
], IdempotencyRecord);
//# sourceMappingURL=idempotency-record.entity.js.map