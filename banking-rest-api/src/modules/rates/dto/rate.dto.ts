import { ApiPropertyOptional } from '@nestjs/swagger';
import { IsOptional, IsString } from 'class-validator';

export class RateLookupDto {
  @ApiPropertyOptional({ example: 'USD' })
  @IsOptional()
  @IsString()
  fromCurrency?: string;

  @ApiPropertyOptional({ example: 'ETB' })
  @IsOptional()
  @IsString()
  toCurrency?: string;

  @ApiPropertyOptional({ example: '2026-10-20' })
  @IsOptional()
  @IsString()
  date?: string;
}
