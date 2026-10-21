import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { SuperAppUser, SuperAppStatus } from '../../database/entities/super-app-user.entity';
import { Customer } from '../../database/entities/customer.entity';
import {
  SuperAppSubscribeDto,
  SuperAppUnsubscribeDto,
  SuperAppStatusDto,
  SuperAppStatusChangeDto,
} from './dto/superapp.dto';
import { CustomerNotFoundException } from '../../common/exceptions/domain.exceptions';

@Injectable()
export class SuperAppService {
  constructor(
    @InjectRepository(SuperAppUser)
    private readonly superAppUserRepo: Repository<SuperAppUser>,
    @InjectRepository(Customer)
    private readonly customerRepo: Repository<Customer>,
  ) {}

  async subscribe(dto: SuperAppSubscribeDto) {
    const customer = await this.customerRepo.findOne({
      where: { customerNumber: dto.customerNumber },
    });
    if (!customer) throw new CustomerNotFoundException(dto.customerNumber);

    let user = await this.superAppUserRepo.findOne({
      where: { externalUserId: dto.externalUserId },
    });

    if (user) {
      user.status = SuperAppStatus.ACTIVE;
      user.customerId = customer.id;
      user.partnerCode = dto.partnerCode || user.partnerCode;
      await this.superAppUserRepo.save(user);
    } else {
      user = this.superAppUserRepo.create({
        externalUserId: dto.externalUserId,
        customerId: customer.id,
        partnerCode: dto.partnerCode || 'SUPERAPP_DEFAULT',
        status: SuperAppStatus.ACTIVE,
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

  async unsubscribe(dto: SuperAppUnsubscribeDto) {
    const user = await this.superAppUserRepo.findOne({
      where: { externalUserId: dto.externalUserId },
    });
    if (!user) {
      throw new NotFoundException(`SuperApp user '${dto.externalUserId}' not found`);
    }

    user.status = SuperAppStatus.DEACTIVATED;
    await this.superAppUserRepo.save(user);

    return {
      externalUserId: user.externalUserId,
      status: user.status,
      message: 'SuperApp subscription deactivated',
    };
  }

  async getStatus(dto: SuperAppStatusDto) {
    return this.getByExternalUserId(dto.externalUserId);
  }

  async getByExternalUserId(externalUserId: string) {
    const user = await this.superAppUserRepo.findOne({
      where: { externalUserId },
      relations: ['customer'],
    });

    if (!user) {
      throw new NotFoundException(`SuperApp user '${externalUserId}' not found`);
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

  async changeStatus(dto: SuperAppStatusChangeDto) {
    const user = await this.superAppUserRepo.findOne({
      where: { externalUserId: dto.externalUserId },
      relations: ['customer'],
    });

    if (!user) {
      throw new NotFoundException(`SuperApp user '${dto.externalUserId}' not found`);
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
}
