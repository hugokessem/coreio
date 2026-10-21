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
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.IdempotencyService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const crypto = require("crypto");
const idempotency_record_entity_1 = require("../../database/entities/idempotency-record.entity");
const domain_exceptions_1 = require("../exceptions/domain.exceptions");
let IdempotencyService = class IdempotencyService {
    idempotencyRepo;
    dataSource;
    constructor(idempotencyRepo, dataSource) {
        this.idempotencyRepo = idempotencyRepo;
        this.dataSource = dataSource;
    }
    hashPayload(payload) {
        const canonical = JSON.stringify(payload, Object.keys(payload || {}).sort());
        return crypto.createHash('sha256').update(canonical).digest('hex');
    }
    async checkOrRegister(key, operation, payload, callerId) {
        if (!key) {
            return { isReplay: false };
        }
        const requestHash = this.hashPayload(payload);
        const existing = await this.idempotencyRepo.findOne({
            where: { idempotencyKey: key },
        });
        if (existing) {
            if (existing.requestHash !== requestHash) {
                throw new domain_exceptions_1.IdempotencyConflictException(`Idempotency key '${key}' was previously used with a different request payload. Duplicate execution rejected.`);
            }
            if (existing.status === idempotency_record_entity_1.IdempotencyStatus.COMPLETED && existing.responsePayload) {
                return {
                    isReplay: true,
                    cachedResponse: {
                        ...existing.responsePayload,
                        message: `Idempotent replay: previously completed ${operation} returned`,
                    },
                };
            }
        }
        return { isReplay: false };
    }
    async recordSuccess(key, operation, payload, transactionReference, response, httpStatus = 200, callerId, entityManager) {
        if (!key)
            return;
        const manager = entityManager || this.dataSource.manager;
        const requestHash = this.hashPayload(payload);
        let record = await manager.findOne(idempotency_record_entity_1.IdempotencyRecord, {
            where: { idempotencyKey: key },
        });
        if (record) {
            record.status = idempotency_record_entity_1.IdempotencyStatus.COMPLETED;
            record.transactionReference = transactionReference;
            record.responsePayload = response;
            record.httpStatusCode = httpStatus;
            await manager.save(idempotency_record_entity_1.IdempotencyRecord, record);
        }
        else {
            record = manager.create(idempotency_record_entity_1.IdempotencyRecord, {
                idempotencyKey: key,
                requestHash,
                operation,
                callerId: callerId || null,
                status: idempotency_record_entity_1.IdempotencyStatus.COMPLETED,
                transactionReference,
                responsePayload: response,
                httpStatusCode: httpStatus,
            });
            await manager.save(idempotency_record_entity_1.IdempotencyRecord, record);
        }
    }
};
exports.IdempotencyService = IdempotencyService;
exports.IdempotencyService = IdempotencyService = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, typeorm_1.InjectRepository)(idempotency_record_entity_1.IdempotencyRecord)),
    __metadata("design:paramtypes", [typeorm_2.Repository,
        typeorm_2.DataSource])
], IdempotencyService);
//# sourceMappingURL=idempotency.service.js.map