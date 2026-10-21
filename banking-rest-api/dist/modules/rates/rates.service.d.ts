import { Repository } from 'typeorm';
import { ExchangeRate } from '../../database/entities/exchange-rate.entity';
import { RateLookupDto } from './dto/rate.dto';
export declare class RatesService {
    private readonly rateRepository;
    constructor(rateRepository: Repository<ExchangeRate>);
    lookup(dto: RateLookupDto): Promise<{
        count: number;
        rates: {
            fromCurrency: string;
            toCurrency: string;
            buyRate: string;
            sellRate: string;
            midRate: string;
            effectiveDate: string;
        }[];
    }>;
    getAllRates(): Promise<{
        count: number;
        rates: {
            fromCurrency: string;
            toCurrency: string;
            buyRate: string;
            sellRate: string;
            midRate: string;
            effectiveDate: string;
        }[];
    }>;
}
