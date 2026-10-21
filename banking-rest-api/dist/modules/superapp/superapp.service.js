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
exports.SuperAppService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const super_app_user_entity_1 = require("../../database/entities/super-app-user.entity");
const customer_entity_1 = require("../../database/entities/customer.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
let SuperAppService = class SuperAppService {
    superAppUserRepo;
    customerRepo;
    constructor(superAppUserRepo, customerRepo) {
        this.superAppUserRepo = superAppUserRepo;
        this.customerRepo = customerRepo;
    }
    async subscribe(dto) {
        const customer = await this.customerRepo.findOne({
            where: { customerNumber: dto.customerNumber },
        });
        if (!customer)
            throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
        let user = await this.superAppUserRepo.findOne({
            where: { externalUserId: dto.externalUserId },
        });
        if (user) {
            user.status = super_app_user_entity_1.SuperAppStatus.ACTIVE;
            user.customerId = customer.id;
            user.partnerCode = dto.partnerCode || user.partnerCode;
            await this.superAppUserRepo.save(user);
        }
        else {
            user = this.superAppUserRepo.create({
                externalUserId: dto.externalUserId,
                customerId: customer.id,
                partnerCode: dto.partnerCode || 'SUPERAPP_DEFAULT',
                status: super_app_user_entity_1.SuperAppStatus.ACTIVE,
            });
            await this.superAppUserRepo.save(user);
        }
        return {
            externalUserId: user.externalUserId,
            customerNumber: customer.customerNumber,
            customerName: customer.fullName,
            status: user.status,
            partnerCode: user.partnerCode,
            registeredAt: user.registeredAt,
            message: 'SuperApp customer binding established successfully',
        };
    }
    async unsubscribe(dto) {
        const user = await this.superAppUserRepo.findOne({
            where: { externalUserId: dto.externalUserId },
        });
        if (!user) {
            throw new common_1.NotFoundException(`SuperApp user '${dto.externalUserId}' not found`);
        }
        user.status = super_app_user_entity_1.SuperAppStatus.DEACTIVATED;
        await this.superAppUserRepo.save(user);
        return {
            externalUserId: user.externalUserId,
            status: user.status,
            message: 'SuperApp subscription deactivated',
        };
    }
    async getStatus(dto) {
        return this.getByExternalUserId(dto.externalUserId);
    }
    async getByExternalUserId(externalUserId) {
        const user = await this.superAppUserRepo.findOne({
            where: { externalUserId },
            relations: ['customer'],
        });
        if (!user) {
            throw new common_1.NotFoundException(`SuperApp user '${externalUserId}' not found`);
        }
        return {
            externalUserId: user.externalUserId,
            customerNumber: user.customer?.customerNumber,
            customerName: user.customer?.fullName,
            status: user.status,
            partnerCode: user.partnerCode,
            registeredAt: user.registeredAt,
            updatedAt: user.updatedAt,
        };
    }
    async changeStatus(dto) {
        const user = await this.superAppUserRepo.findOne({
            where: { externalUserId: dto.externalUserId },
            relations: ['customer'],
        });
        if (!user) {
            throw new common_1.NotFoundException(`SuperApp user '${dto.externalUserId}' not found`);
        }
        user.status = dto.status;
        const saved = await this.superAppUserRepo.save(user);
        return {
            externalUserId: saved.externalUserId,
            customerNumber: saved.customer?.customerNumber,
            status: saved.status,
            reason: dto.reason || 'Status updated',
            updatedAt: saved.updatedAt,
            message: `SuperApp subscription status changed to ${saved.status}`,
        };
    }
};
exports.SuperAppService = SuperAppService;
exports.SuperAppService = SuperAppService = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, typeorm_1.InjectRepository)(super_app_user_entity_1.SuperAppUser)),
    __param(1, (0, typeorm_1.InjectRepository)(customer_entity_1.Customer)),
    __metadata("design:paramtypes", [typeorm_2.Repository,
        typeorm_2.Repository])
], SuperAppService);
//# sourceMappingURL=superapp.service.js.map