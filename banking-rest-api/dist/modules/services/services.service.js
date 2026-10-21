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
exports.ServicesService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const service_entity_1 = require("../../database/entities/service.entity");
const customer_limit_entity_1 = require("../../database/entities/customer-limit.entity");
const customer_entity_1 = require("../../database/entities/customer.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
let ServicesService = class ServicesService {
    serviceRepository;
    customerLimitRepository;
    customerRepository;
    constructor(serviceRepository, customerLimitRepository, customerRepository) {
        this.serviceRepository = serviceRepository;
        this.customerLimitRepository = customerLimitRepository;
        this.customerRepository = customerRepository;
    }
    async getDetail(dto) {
        return this.getDetailByCode(dto.serviceCode);
    }
    async getDetailByCode(serviceCode) {
        const service = await this.serviceRepository.findOne({
            where: { serviceCode },
        });
        if (!service) {
            throw new common_1.NotFoundException(`Service '${serviceCode}' not found`);
        }
        return {
            serviceCode: service.serviceCode,
            serviceName: service.serviceName,
            category: service.category,
            status: service.status,
            minAmount: service.minAmount,
            maxAmount: service.maxAmount,
            dailyLimit: service.dailyLimit,
            feeType: service.feeType,
            feeValue: service.feeValue,
        };
    }
    async getServiceLimits() {
        const services = await this.serviceRepository.find({ order: { serviceCode: 'ASC' } });
        return {
            totalServices: services.length,
            services: services.map((s) => ({
                serviceCode: s.serviceCode,
                serviceName: s.serviceName,
                category: s.category,
                minAmount: s.minAmount,
                maxAmount: s.maxAmount,
                dailyLimit: s.dailyLimit,
                feeType: s.feeType,
                feeValue: s.feeValue,
            })),
        };
    }
    async getCustomerLimits(dto) {
        return this.getCustomerLimitsByNumber(dto.customerNumber, dto.serviceCode);
    }
    async getCustomerLimitsByNumber(customerNumber, serviceCode) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber },
        });
        if (!customer)
            throw new domain_exceptions_1.CustomerNotFoundException(customerNumber);
        const where = { customerId: customer.id };
        if (serviceCode) {
            where.serviceCode = serviceCode;
        }
        const limits = await this.customerLimitRepository.find({ where });
        return {
            customerNumber: customer.customerNumber,
            customerName: customer.fullName,
            customLimitsCount: limits.length,
            limits: limits.map((l) => ({
                serviceCode: l.serviceCode,
                dailyLimit: l.dailyLimit,
                perTransactionLimit: l.perTransactionLimit,
                dailyUsed: l.dailyUsed,
                lastResetDate: l.lastResetDate,
            })),
        };
    }
    async amendCustomerLimit(dto) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber: dto.customerNumber },
        });
        if (!customer)
            throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
        let limit = await this.customerLimitRepository.findOne({
            where: { customerId: customer.id, serviceCode: dto.serviceCode },
        });
        if (!limit) {
            limit = this.customerLimitRepository.create({
                customerId: customer.id,
                serviceCode: dto.serviceCode,
                dailyLimit: dto.dailyLimit || '500000.0000',
                perTransactionLimit: dto.perTransactionLimit || '100000.0000',
            });
        }
        else {
            if (dto.dailyLimit)
                limit.dailyLimit = dto.dailyLimit;
            if (dto.perTransactionLimit)
                limit.perTransactionLimit = dto.perTransactionLimit;
        }
        const saved = await this.customerLimitRepository.save(limit);
        return {
            customerNumber: customer.customerNumber,
            serviceCode: saved.serviceCode,
            dailyLimit: saved.dailyLimit,
            perTransactionLimit: saved.perTransactionLimit,
            dailyUsed: saved.dailyUsed,
            updatedAt: saved.updatedAt,
            message: 'Customer limit successfully amended',
        };
    }
    async listAll() {
        return this.serviceRepository.find({ order: { serviceCode: 'ASC' } });
    }
};
exports.ServicesService = ServicesService;
exports.ServicesService = ServicesService = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, typeorm_1.InjectRepository)(service_entity_1.ServiceEntity)),
    __param(1, (0, typeorm_1.InjectRepository)(customer_limit_entity_1.CustomerLimit)),
    __param(2, (0, typeorm_1.InjectRepository)(customer_entity_1.Customer)),
    __metadata("design:paramtypes", [typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], ServicesService);
//# sourceMappingURL=services.service.js.map