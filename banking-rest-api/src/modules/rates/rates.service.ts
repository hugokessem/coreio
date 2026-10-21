import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { ExchangeRate } from '../../database/entities/exchange-rate.entity';
import { RateLookupDto } from './dto/rate.dto';

@Injectable()
export class RatesService {
  constructor(
    @InjectRepository(ExchangeRate)
    private readonly rateRepository: Repository<ExchangeRate>,
  ) {}

  async lookup(dto: RateLookupDto) {
    const query = this.rateRepository.createQueryBuilder('rate');

    if (dto.fromCurrency) {
      query.andWhere('rate.fromCurrency = :from', { from: dto.fromCurrency });
    }
    if (dto.toCurrency) {
      query.andWhere('rate.toCurrency = :to', { to: dto.toCurrency });
    }
    if (dto.date) {
      query.andWhere('rate.effectiveDate = :date', { date: dto.date });
    }

    query.orderBy('rate.effectiveDate', 'DESC');
    const rates = await query.getMany();

    return {
      count: rates.length,
      rates: rates.map((r) => ({
        fromCurrency: r.fromCurrency,
        toCurrency: r.toCurrency,
        buyRate: r.buyRate,
        sellRate: r.sellRate,
        midRate: r.midRate,
        effectiveDate: r.effectiveDate,
      })),
    };
  }

  async getAllRates() {
    return this.lookup({});
  }
}
