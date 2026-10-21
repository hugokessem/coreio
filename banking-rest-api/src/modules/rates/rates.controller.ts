import {
  Controller,
  Post,
  Get,
  Body,
  Query,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { RatesService } from './rates.service';
import { RateLookupDto } from './dto/rate.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';

@ApiTags('Exchange Rates')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard)
@Controller('api/v1/rates')
export class RatesController {
  constructor(private readonly ratesService: RatesService) {}

  @Post('lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Lookup FX exchange rates by currency pair or date (Op 35: ExchangeRates)' })
  async lookup(@Body() dto: RateLookupDto) {
    return this.ratesService.lookup(dto);
  }

  @Get()
  @ApiOperation({ summary: 'GET all active FX exchange rates - REST Canonical (Op 35: ExchangeRates)' })
  async getAllRates(
    @Query('fromCurrency') fromCurrency?: string,
    @Query('toCurrency') toCurrency?: string,
  ) {
    return this.ratesService.lookup({ fromCurrency, toCurrency });
  }

  @Get('exchange-rates')
  @ApiOperation({ summary: 'GET FX exchange rates - CoreIO Route Alias (Op 35: ExchangeRates)' })
  async getExchangeRates(
    @Query('fromCurrency') fromCurrency?: string,
    @Query('toCurrency') toCurrency?: string,
  ) {
    return this.ratesService.lookup({ fromCurrency, toCurrency });
  }

  @Post('exchange-rates')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'POST FX exchange rates lookup - CoreIO Route Alias (Op 35: ExchangeRates)' })
  async postExchangeRates(@Body() dto: RateLookupDto) {
    return this.ratesService.lookup(dto);
  }
}
