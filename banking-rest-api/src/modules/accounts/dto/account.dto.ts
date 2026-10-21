import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsNumberString, IsInt, Min, Max } from 'class-validator';

export class AccountLookupDto {
  @ApiProperty({ example: '1000123456789', description: 'Account number to lookup' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;
}

export class AccountNameLookupDto {
  @ApiProperty({ example: '1000123456789', description: 'Account number for beneficiary/name lookup' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;
}

export class CreateAccountDto {
  @ApiProperty({ example: '10000001', description: 'Customer number owning the account' })
  @IsNotEmpty()
  @IsString()
  customerNumber!: string;

  @ApiProperty({ example: '6501', description: 'Account category code (6501 Savings, 6001 Current)' })
  @IsNotEmpty()
  @IsString()
  accountTypeCode!: string;

  @ApiProperty({ example: '001', description: 'Branch code (001 Addis Ababa Main, 100 Bole)' })
  @IsNotEmpty()
  @IsString()
  branchCode!: string;

  @ApiProperty({ example: 'ETB', description: 'Currency code (ETB, USD, EUR, GBP)' })
  @IsNotEmpty()
  @IsString()
  currencyCode!: string;

  @ApiPropertyOptional({ example: 'Abebe Kebede Personal Account' })
  @IsOptional()
  @IsString()
  accountName?: string;

  @ApiPropertyOptional({ example: '1000.0000', description: 'Initial deposit amount' })
  @IsOptional()
  @IsNumberString()
  initialDeposit?: string;
}

export class AccountListDto {
  @ApiProperty({ example: '10000001', description: 'Customer number to list all accounts for' })
  @IsNotEmpty()
  @IsString()
  customerNumber!: string;
}

export class AccountEligibilityDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;

  @ApiProperty({ example: '500.0000', description: 'Amount intended for transaction' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'DEBIT', description: 'DEBIT or CREDIT' })
  @IsOptional()
  @IsString()
  transactionType?: string;
}

export class CreateHoldDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;

  @ApiProperty({ example: '2500.0000' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiProperty({ example: 'Court Order Freeze' })
  @IsNotEmpty()
  @IsString()
  reason!: string;

  @ApiPropertyOptional({ example: 'COURT_ORDER' })
  @IsOptional()
  @IsString()
  lockEvent?: string;

  @ApiPropertyOptional({ example: '2026-10-20' })
  @IsOptional()
  @IsString()
  fromDate?: string;

  @ApiPropertyOptional({ example: '2026-12-31' })
  @IsOptional()
  @IsString()
  toDate?: string;
}

export class ReleaseHoldDto {
  @ApiProperty({ example: 'HLD-001', description: 'Hold / Lock reference to release' })
  @IsNotEmpty()
  @IsString()
  holdReference!: string;

  @ApiPropertyOptional({ example: 'Court order lifted' })
  @IsOptional()
  @IsString()
  reason?: string;
}

export class AccountStatementDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;

  @ApiPropertyOptional({ example: '2026-01-01' })
  @IsOptional()
  @IsString()
  startDate?: string;

  @ApiPropertyOptional({ example: '2026-12-31' })
  @IsOptional()
  @IsString()
  endDate?: string;

  @ApiPropertyOptional({ example: 50, default: 50 })
  @IsOptional()
  @IsInt()
  @Min(1)
  @Max(200)
  limit?: number;
}

export class AccountMiniStatementDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;

  @ApiPropertyOptional({ example: 5, default: 5 })
  @IsOptional()
  @IsInt()
  @Min(1)
  @Max(50)
  count?: number;

  @ApiPropertyOptional({ example: 5, default: 5 })
  @IsOptional()
  @IsInt()
  @Min(1)
  @Max(50)
  limit?: number;
}
