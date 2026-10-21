import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsNumberString } from 'class-validator';

export class ServiceDetailDto {
  @ApiProperty({ example: 'FT_INT', description: 'Service code (FT_INT, FT_IPS, CBE_BIRR, BILL_ETHIO)' })
  @IsNotEmpty()
  @IsString()
  serviceCode!: string;
}

export class CustomerLimitsQueryDto {
  @ApiProperty({ example: '10000001' })
  @IsNotEmpty()
  @IsString()
  customerNumber!: string;

  @ApiPropertyOptional({ example: 'FT_INT' })
  @IsOptional()
  @IsString()
  serviceCode?: string;
}

export class AmendCustomerLimitDto {
  @ApiProperty({ example: '10000001' })
  @IsNotEmpty()
  @IsString()
  customerNumber!: string;

  @ApiProperty({ example: 'FT_INT' })
  @IsNotEmpty()
  @IsString()
  serviceCode!: string;

  @ApiPropertyOptional({ example: '750000.0000' })
  @IsOptional()
  @IsNumberString()
  dailyLimit?: string;

  @ApiPropertyOptional({ example: '150000.0000' })
  @IsOptional()
  @IsNumberString()
  perTransactionLimit?: string;
}
