import { RatesService } from './rates.service';
import { RateLookupDto } from './dto/rate.dto';
export declare class RatesController {
    private readonly ratesService;
    constructor(ratesService: RatesService);
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
    getAllRates(fromCurrency?: string, toCurrency?: string): Promise<{
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
    getExchangeRates(fromCurrency?: string, toCurrency?: string): Promise<{
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
    postExchangeRates(dto: RateLookupDto): Promise<{
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
