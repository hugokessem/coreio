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
var StandingOrdersService_1;
Object.defineProperty(exports, "__esModule", { value: true });
exports.StandingOrdersService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const standing_order_entity_1 = require("../../database/entities/standing-order.entity");
const standing_order_execution_entity_1 = require("../../database/entities/standing-order-execution.entity");
const account_entity_1 = require("../../database/entities/account.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
const transfers_service_1 = require("../transfers/transfers.service");
let StandingOrdersService = StandingOrdersService_1 = class StandingOrdersService {
    dataSource;
    standingOrderRepo;
    executionRepo;
    accountRepo;
    transfersService;
    logger = new common_1.Logger(StandingOrdersService_1.name);
    constructor(dataSource, standingOrderRepo, executionRepo, accountRepo, transfersService) {
        this.dataSource = dataSource;
        this.standingOrderRepo = standingOrderRepo;
        this.executionRepo = executionRepo;
        this.accountRepo = accountRepo;
        this.transfersService = transfersService;
    }
    async create(dto) {
        const source = await this.accountRepo.findOne({
            where: { accountNumber: dto.sourceAccountNumber },
        });
        if (!source)
            throw new domain_exceptions_1.AccountNotFoundException(dto.sourceAccountNumber);
        const dest = await this.accountRepo.findOne({
            where: { accountNumber: dto.beneficiaryAccount },
        });
        const count = await this.standingOrderRepo.count();
        const orderReference = `SO-${source.accountNumber}-${(count + 1).toString().padStart(2, '0')}`;
        const so = this.standingOrderRepo.create({
            orderReference,
            sourceAccountId: source.id,
            destinationAccountId: dest ? dest.id : null,
            beneficiaryName: dto.beneficiaryName,
            beneficiaryAccount: dto.beneficiaryAccount,
            beneficiaryBankCode: dto.beneficiaryBankCode || null,
            amount: dto.amount,
            currencyCode: dto.currencyCode || source.currencyCode,
            frequency: dto.frequency,
            startDate: dto.startDate,
            endDate: dto.endDate || null,
            nextExecutionDate: dto.startDate,
            status: standing_order_entity_1.StandingOrderStatus.ACTIVE,
        });
        const saved = await this.standingOrderRepo.save(so);
        return {
            orderReference: saved.orderReference,
            sourceAccountNumber: source.accountNumber,
            beneficiaryName: saved.beneficiaryName,
            beneficiaryAccount: saved.beneficiaryAccount,
            amount: saved.amount,
            currency: saved.currencyCode,
            frequency: saved.frequency,
            startDate: saved.startDate,
            endDate: saved.endDate,
            nextExecutionDate: saved.nextExecutionDate,
            status: saved.status,
            createdAt: saved.createdAt,
        };
    }
    async list(dto) {
        return this.listByAccount(dto.accountNumber);
    }
    async listByAccount(accountNumber) {
        const source = await this.accountRepo.findOne({
            where: { accountNumber },
        });
        if (!source)
            throw new domain_exceptions_1.AccountNotFoundException(accountNumber);
        const orders = await this.standingOrderRepo.find({
            where: { sourceAccountId: source.id },
            order: { createdAt: 'DESC' },
        });
        return {
            accountNumber: source.accountNumber,
            totalOrders: orders.length,
            orders: orders.map((o) => ({
                orderReference: o.orderReference,
                beneficiaryName: o.beneficiaryName,
                beneficiaryAccount: o.beneficiaryAccount,
                amount: o.amount,
                currency: o.currencyCode,
                frequency: o.frequency,
                startDate: o.startDate,
                endDate: o.endDate,
                nextExecutionDate: o.nextExecutionDate,
                status: o.status,
                failureCount: o.failureCount,
                lastExecutedAt: o.lastExecutedAt,
            })),
        };
    }
    async getHistory(orderReference) {
        const order = await this.standingOrderRepo.findOne({
            where: { orderReference },
            relations: ['sourceAccount'],
        });
        if (!order) {
            throw new common_1.NotFoundException(`Standing order '${orderReference}' not found`);
        }
        const executions = await this.executionRepo.find({
            where: { standingOrderId: order.id },
            order: { createdAt: 'DESC' },
        });
        return {
            orderReference: order.orderReference,
            sourceAccountNumber: order.sourceAccount?.accountNumber,
            beneficiaryAccount: order.beneficiaryAccount,
            beneficiaryName: order.beneficiaryName,
            status: order.status,
            frequency: order.frequency,
            totalExecutions: executions.length,
            executions: executions.map((e) => ({
                id: e.id,
                executionDate: e.executionDate,
                amount: e.amount,
                status: e.status,
                transactionReference: e.transactionReference,
                failureReason: e.failureReason,
                executedAt: e.createdAt,
            })),
        };
    }
    async update(orderReference, dto) {
        const order = await this.standingOrderRepo.findOne({
            where: { orderReference },
        });
        if (!order) {
            throw new common_1.NotFoundException(`Standing order '${orderReference}' not found`);
        }
        if (order.status === standing_order_entity_1.StandingOrderStatus.CANCELLED) {
            throw new common_1.BadRequestException('Cannot update a cancelled standing order');
        }
        if (dto.amount !== undefined)
            order.amount = dto.amount;
        if (dto.frequency !== undefined)
            order.frequency = dto.frequency;
        if (dto.endDate !== undefined)
            order.endDate = dto.endDate;
        if (dto.status !== undefined)
            order.status = dto.status;
        const saved = await this.standingOrderRepo.save(order);
        return {
            orderReference: saved.orderReference,
            amount: saved.amount,
            frequency: saved.frequency,
            endDate: saved.endDate,
            status: saved.status,
            updatedAt: saved.updatedAt,
            message: 'Standing order updated successfully',
        };
    }
    async cancel(dto) {
        const orderReference = typeof dto === 'string' ? dto : dto.orderReference;
        const order = await this.standingOrderRepo.findOne({
            where: { orderReference },
        });
        if (!order) {
            throw new common_1.NotFoundException(`Standing order '${orderReference}' not found`);
        }
        order.status = standing_order_entity_1.StandingOrderStatus.CANCELLED;
        await this.standingOrderRepo.save(order);
        return {
            orderReference: order.orderReference,
            status: order.status,
            cancelledAt: new Date().toISOString(),
            message: 'Standing order cancelled successfully',
        };
    }
    async executeDueOrders() {
        const today = new Date().toISOString().split('T')[0];
        let succeeded = 0;
        let failed = 0;
        return await this.dataSource.transaction(async (manager) => {
            const dueOrders = await manager
                .createQueryBuilder(standing_order_entity_1.StandingOrder, 'so')
                .where('so.status = :status', { status: standing_order_entity_1.StandingOrderStatus.ACTIVE })
                .andWhere('so.nextExecutionDate <= :today', { today })
                .setLock('pessimistic_write')
                .setOnLocked('skip_locked')
                .limit(50)
                .getMany();
            for (const order of dueOrders) {
                const sourceAccount = await manager.findOne(account_entity_1.Account, {
                    where: { id: order.sourceAccountId },
                });
                if (!sourceAccount)
                    continue;
                const idempotencyKey = `SO_EXEC_${order.orderReference}_${order.nextExecutionDate}`;
                try {
                    const result = await this.transfersService.execute({
                        sourceAccountNumber: sourceAccount.accountNumber,
                        destinationAccountNumber: order.beneficiaryAccount,
                        amount: order.amount,
                        narration: `Standing Order ${order.orderReference}`,
                        idempotencyKey,
                    });
                    const exec = manager.create(standing_order_execution_entity_1.StandingOrderExecution, {
                        standingOrderId: order.id,
                        executionDate: order.nextExecutionDate,
                        amount: order.amount,
                        transactionReference: result.transactionReference,
                        status: standing_order_execution_entity_1.ExecutionStatus.SUCCESS,
                    });
                    await manager.save(standing_order_execution_entity_1.StandingOrderExecution, exec);
                    order.lastExecutedAt = new Date();
                    order.failureCount = 0;
                    order.nextExecutionDate = this.calculateNextDate(order.nextExecutionDate, order.frequency);
                    if (order.endDate && order.nextExecutionDate > order.endDate) {
                        order.status = standing_order_entity_1.StandingOrderStatus.COMPLETED;
                    }
                    await manager.save(standing_order_entity_1.StandingOrder, order);
                    succeeded++;
                }
                catch (err) {
                    this.logger.error(`Failed executing standing order ${order.orderReference}: ${err.message}`);
                    failed++;
                    const exec = manager.create(standing_order_execution_entity_1.StandingOrderExecution, {
                        standingOrderId: order.id,
                        executionDate: order.nextExecutionDate,
                        amount: order.amount,
                        status: standing_order_execution_entity_1.ExecutionStatus.FAILED,
                        failureReason: err.message,
                    });
                    await manager.save(standing_order_execution_entity_1.StandingOrderExecution, exec);
                    order.failureCount += 1;
                    if (order.failureCount >= 3) {
                        order.status = standing_order_entity_1.StandingOrderStatus.PAUSED;
                    }
                    await manager.save(standing_order_entity_1.StandingOrder, order);
                }
            }
            return {
                processed: dueOrders.length,
                succeeded,
                failed,
            };
        });
    }
    calculateNextDate(currentDateStr, freq) {
        const d = new Date(currentDateStr);
        switch (freq) {
            case standing_order_entity_1.StandingOrderFrequency.DAILY:
                d.setDate(d.getDate() + 1);
                break;
            case standing_order_entity_1.StandingOrderFrequency.WEEKLY:
                d.setDate(d.getDate() + 7);
                break;
            case standing_order_entity_1.StandingOrderFrequency.MONTHLY:
                d.setMonth(d.getMonth() + 1);
                break;
            case standing_order_entity_1.StandingOrderFrequency.QUARTERLY:
                d.setMonth(d.getMonth() + 3);
                break;
            case standing_order_entity_1.StandingOrderFrequency.ANNUALLY:
                d.setFullYear(d.getFullYear() + 1);
                break;
            default:
                d.setMonth(d.getMonth() + 1);
        }
        return d.toISOString().split('T')[0];
    }
};
exports.StandingOrdersService = StandingOrdersService;
exports.StandingOrdersService = StandingOrdersService = StandingOrdersService_1 = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(standing_order_entity_1.StandingOrder)),
    __param(2, (0, typeorm_1.InjectRepository)(standing_order_execution_entity_1.StandingOrderExecution)),
    __param(3, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        transfers_service_1.TransfersService])
], StandingOrdersService);
//# sourceMappingURL=standing-orders.service.js.map