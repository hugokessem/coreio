import { Injectable, BadRequestException, NotFoundException, Logger } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { DataSource, Repository } from 'typeorm';
import {
  StandingOrder,
  StandingOrderStatus,
  StandingOrderFrequency,
} from '../../database/entities/standing-order.entity';
import {
  StandingOrderExecution,
  ExecutionStatus,
} from '../../database/entities/standing-order-execution.entity';
import { Account } from '../../database/entities/account.entity';
import {
  CreateStandingOrderDto,
  ListStandingOrdersDto,
  CancelStandingOrderDto,
  UpdateStandingOrderDto,
} from './dto/standing-order.dto';
import { AccountNotFoundException } from '../../common/exceptions/domain.exceptions';
import { TransfersService } from '../transfers/transfers.service';

@Injectable()
export class StandingOrdersService {
  private readonly logger = new Logger(StandingOrdersService.name);

  constructor(
    private readonly dataSource: DataSource,
    @InjectRepository(StandingOrder)
    private readonly standingOrderRepo: Repository<StandingOrder>,
    @InjectRepository(StandingOrderExecution)
    private readonly executionRepo: Repository<StandingOrderExecution>,
    @InjectRepository(Account)
    private readonly accountRepo: Repository<Account>,
    private readonly transfersService: TransfersService,
  ) {}

  async create(dto: CreateStandingOrderDto) {
    const source = await this.accountRepo.findOne({
      where: { accountNumber: dto.sourceAccountNumber },
    });
    if (!source) throw new AccountNotFoundException(dto.sourceAccountNumber);

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
      status: StandingOrderStatus.ACTIVE,
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

  async list(dto: ListStandingOrdersDto) {
    return this.listByAccount(dto.accountNumber);
  }

  async listByAccount(accountNumber: string) {
    const source = await this.accountRepo.findOne({
      where: { accountNumber },
    });
    if (!source) throw new AccountNotFoundException(accountNumber);

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

  async getHistory(orderReference: string) {
    const order = await this.standingOrderRepo.findOne({
      where: { orderReference },
      relations: ['sourceAccount'],
    });
    if (!order) {
      throw new NotFoundException(`Standing order '${orderReference}' not found`);
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

  async update(orderReference: string, dto: UpdateStandingOrderDto) {
    const order = await this.standingOrderRepo.findOne({
      where: { orderReference },
    });
    if (!order) {
      throw new NotFoundException(`Standing order '${orderReference}' not found`);
    }

    if (order.status === StandingOrderStatus.CANCELLED) {
      throw new BadRequestException('Cannot update a cancelled standing order');
    }

    if (dto.amount !== undefined) order.amount = dto.amount;
    if (dto.frequency !== undefined) order.frequency = dto.frequency;
    if (dto.endDate !== undefined) order.endDate = dto.endDate;
    if (dto.status !== undefined) order.status = dto.status;

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

  async cancel(dto: CancelStandingOrderDto | string) {
    const orderReference = typeof dto === 'string' ? dto : dto.orderReference;
    const order = await this.standingOrderRepo.findOne({
      where: { orderReference },
    });
    if (!order) {
      throw new NotFoundException(`Standing order '${orderReference}' not found`);
    }

    order.status = StandingOrderStatus.CANCELLED;
    await this.standingOrderRepo.save(order);

    return {
      orderReference: order.orderReference,
      status: order.status,
      cancelledAt: new Date().toISOString(),
      message: 'Standing order cancelled successfully',
    };
  }

  /**
   * Durable scheduler worker: acquires locks via FOR UPDATE SKIP LOCKED and processes due orders
   */
  async executeDueOrders(): Promise<{ processed: number; succeeded: number; failed: number }> {
    const today = new Date().toISOString().split('T')[0];
    let succeeded = 0;
    let failed = 0;

    return await this.dataSource.transaction(async (manager) => {
      const dueOrders = await manager
        .createQueryBuilder(StandingOrder, 'so')
        .where('so.status = :status', { status: StandingOrderStatus.ACTIVE })
        .andWhere('so.nextExecutionDate <= :today', { today })
        .setLock('pessimistic_write')
        .setOnLocked('skip_locked')
        .limit(50)
        .getMany();

      for (const order of dueOrders) {
        const sourceAccount = await manager.findOne(Account, {
          where: { id: order.sourceAccountId },
        });
        if (!sourceAccount) continue;

        const idempotencyKey = `SO_EXEC_${order.orderReference}_${order.nextExecutionDate}`;
        try {
          const result = await this.transfersService.execute({
            sourceAccountNumber: sourceAccount.accountNumber,
            destinationAccountNumber: order.beneficiaryAccount,
            amount: order.amount,
            narration: `Standing Order ${order.orderReference}`,
            idempotencyKey,
          });

          // Record execution success
          const exec = manager.create(StandingOrderExecution, {
            standingOrderId: order.id,
            executionDate: order.nextExecutionDate,
            amount: order.amount,
            transactionReference: result.transactionReference,
            status: ExecutionStatus.SUCCESS,
          });
          await manager.save(StandingOrderExecution, exec);

          // Advance nextExecutionDate
          order.lastExecutedAt = new Date();
          order.failureCount = 0;
          order.nextExecutionDate = this.calculateNextDate(order.nextExecutionDate, order.frequency);

          // Check if passed end date
          if (order.endDate && order.nextExecutionDate > order.endDate) {
            order.status = StandingOrderStatus.COMPLETED;
          }

          await manager.save(StandingOrder, order);
          succeeded++;
        } catch (err: any) {
          this.logger.error(`Failed executing standing order ${order.orderReference}: ${err.message}`);
          failed++;

          const exec = manager.create(StandingOrderExecution, {
            standingOrderId: order.id,
            executionDate: order.nextExecutionDate,
            amount: order.amount,
            status: ExecutionStatus.FAILED,
            failureReason: err.message,
          });
          await manager.save(StandingOrderExecution, exec);

          order.failureCount += 1;
          if (order.failureCount >= 3) {
            order.status = StandingOrderStatus.PAUSED;
          }
          await manager.save(StandingOrder, order);
        }
      }

      return {
        processed: dueOrders.length,
        succeeded,
        failed,
      };
    });
  }

  private calculateNextDate(currentDateStr: string, freq: StandingOrderFrequency): string {
    const d = new Date(currentDateStr);
    switch (freq) {
      case StandingOrderFrequency.DAILY:
        d.setDate(d.getDate() + 1);
        break;
      case StandingOrderFrequency.WEEKLY:
        d.setDate(d.getDate() + 7);
        break;
      case StandingOrderFrequency.MONTHLY:
        d.setMonth(d.getMonth() + 1);
        break;
      case StandingOrderFrequency.QUARTERLY:
        d.setMonth(d.getMonth() + 3);
        break;
      case StandingOrderFrequency.ANNUALLY:
        d.setFullYear(d.getFullYear() + 1);
        break;
      default:
        d.setMonth(d.getMonth() + 1);
    }
    return d.toISOString().split('T')[0];
  }
}
