import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsNumberString, IsEnum } from 'class-validator';
import { StandingOrderFrequency, StandingOrderStatus } from '../../../database/entities/standing-order.entity';

export class CreateStandingOrderDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: '1000987654321' })
  @IsNotEmpty()
  @IsString()
  beneficiaryAccount!: string;

  @ApiProperty({ example: 'Almaz Tadesse Mengesha' })
  @IsNotEmpty()
  @IsString()
  beneficiaryName!: string;

  @ApiPropertyOptional({ example: 'CBETETAA' })
  @IsOptional()
  @IsString()
  beneficiaryBankCode?: string;

  @ApiProperty({ example: '2500.0000' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'ETB' })
  @IsOptional()
  @IsString()
  currencyCode?: string;

  @ApiProperty({ enum: StandingOrderFrequency, example: StandingOrderFrequency.MONTHLY })
  @IsEnum(StandingOrderFrequency)
  frequency!: StandingOrderFrequency;

  @ApiProperty({ example: '2026-11-01' })
  @IsNotEmpty()
  @IsString()
  startDate!: string;

  @ApiPropertyOptional({ example: '2027-11-01' })
  @IsOptional()
  @IsString()
  endDate?: string;
}

export class ListStandingOrdersDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;
}

export class CancelStandingOrderDto {
  @ApiProperty({ example: 'SO-1000123456789-01' })
  @IsNotEmpty()
  @IsString()
  orderReference!: string;
}

export class UpdateStandingOrderDto {
  @ApiPropertyOptional({ example: '3000.0000' })
  @IsOptional()
  @IsNumberString()
  amount?: string;

  @ApiPropertyOptional({ enum: StandingOrderFrequency, example: StandingOrderFrequency.WEEKLY })
  @IsOptional()
  @IsEnum(StandingOrderFrequency)
  frequency?: StandingOrderFrequency;

  @ApiPropertyOptional({ example: '2027-12-31' })
  @IsOptional()
  @IsString()
  endDate?: string;

  @ApiPropertyOptional({ enum: StandingOrderStatus, example: StandingOrderStatus.PAUSED })
  @IsOptional()
  @IsEnum(StandingOrderStatus)
  status?: StandingOrderStatus;
}

export class ListStandingOrderHistoryDto {
  @ApiProperty({ example: 'SO-1000123456789-01' })
  @IsNotEmpty()
  @IsString()
  orderReference!: string;
}
