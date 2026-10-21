import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { ServiceEntity } from '../../database/entities/service.entity';
import { CustomerLimit } from '../../database/entities/customer-limit.entity';
import { Customer } from '../../database/entities/customer.entity';
import { ServiceDetailDto, CustomerLimitsQueryDto, AmendCustomerLimitDto } from './dto/service.dto';
import { CustomerNotFoundException } from '../../common/exceptions/domain.exceptions';

@Injectable()
export class ServicesService {
  constructor(
    @InjectRepository(ServiceEntity)
    private readonly serviceRepository: Repository<ServiceEntity>,
    @InjectRepository(CustomerLimit)
    private readonly customerLimitRepository: Repository<CustomerLimit>,
    @InjectRepository(Customer)
    private readonly customerRepository: Repository<Customer>,
  ) {}

  async getDetail(dto: ServiceDetailDto) {
    return this.getDetailByCode(dto.serviceCode);
  }

  async getDetailByCode(serviceCode: string) {
    const service = await this.serviceRepository.findOne({
      where: { serviceCode },
    });
    if (!service) {
      throw new NotFoundException(`Service '${serviceCode}' not found`);
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

  async getCustomerLimits(dto: CustomerLimitsQueryDto) {
    return this.getCustomerLimitsByNumber(dto.customerNumber, dto.serviceCode);
  }

  async getCustomerLimitsByNumber(customerNumber: string, serviceCode?: string) {
    const customer = await this.customerRepository.findOne({
      where: { customerNumber },
    });
    if (!customer) throw new CustomerNotFoundException(customerNumber);

    const where: any = { customerId: customer.id };
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

  async amendCustomerLimit(dto: AmendCustomerLimitDto) {
    const customer = await this.customerRepository.findOne({
      where: { customerNumber: dto.customerNumber },
    });
    if (!customer) throw new CustomerNotFoundException(dto.customerNumber);

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
    } else {
      if (dto.dailyLimit) limit.dailyLimit = dto.dailyLimit;
      if (dto.perTransactionLimit) limit.perTransactionLimit = dto.perTransactionLimit;
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
}
