import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsNumberString, IsArray, ValidateNested } from 'class-validator';
import { Type } from 'class-transformer';

export class BillPaymentDto {
  @ApiProperty({ example: '1000123456789', description: 'Customer debit account' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: 'BILL_ETHIO', description: 'Biller service code' })
  @IsNotEmpty()
  @IsString()
  billerCode!: string;

  @ApiProperty({ example: '0911223344', description: 'Customer utility/phone/bill account number' })
  @IsNotEmpty()
  @IsString()
  billReference!: string;

  @ApiProperty({ example: '350.0000', description: 'Payment amount' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'Ethio Telecom Monthly Postpaid Bill' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-BILL-0099' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}

export class BillPaymentVerifyDto {
  @ApiProperty({ example: '1000123456789', description: 'Customer debit account' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: 'BILL_ETHIO', description: 'Biller service code' })
  @IsNotEmpty()
  @IsString()
  billerCode!: string;

  @ApiProperty({ example: '0911223344', description: 'Customer utility/phone/bill account number' })
  @IsNotEmpty()
  @IsString()
  billReference!: string;

  @ApiProperty({ example: '350.0000', description: 'Payment amount' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;
}

export class SplitDestinationDto {
  @ApiProperty({ example: '1000987654321' })
  @IsNotEmpty()
  @IsString()
  destinationAccountNumber!: string;

  @ApiProperty({ example: '500.0000' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'Employee Bonus Portion' })
  @IsOptional()
  @IsString()
  narration?: string;
}

export class SplitPaymentDto {
  @ApiProperty({ example: '2000123456789', description: 'Source corporate/master account' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ type: [SplitDestinationDto], description: 'List of recipients and split amounts' })
  @IsArray()
  @ValidateNested({ each: true })
  @Type(() => SplitDestinationDto)
  splits!: SplitDestinationDto[];

  @ApiPropertyOptional({ example: 'Vendor Bulk Settlement' })
  @IsOptional()
  @IsString()
  batchNarration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-SPLIT-BATCH-001' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}

export class RtgsPaymentDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: 'AWINETAA', description: 'Destination Bank SWIFT/BIC code' })
  @IsNotEmpty()
  @IsString()
  beneficiaryBankSwift!: string;

  @ApiProperty({ example: '0132049281900', description: 'Beneficiary account number at destination bank' })
  @IsNotEmpty()
  @IsString()
  beneficiaryAccountNumber!: string;

  @ApiProperty({ example: 'Awash Bank Commercial Customer PLC' })
  @IsNotEmpty()
  @IsString()
  beneficiaryName!: string;

  @ApiProperty({ example: '250000.0000' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'Import LC Settlement' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-RTGS-001' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}
