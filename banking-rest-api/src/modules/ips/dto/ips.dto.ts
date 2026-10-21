import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsNumberString } from 'class-validator';

export class IpsAccountLookupDto {
  @ApiProperty({ example: '1000123456789', description: 'Account number to lookup via IPS' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;

  @ApiPropertyOptional({ example: 'CBETETAA', description: 'Bank identifier / BIC' })
  @IsOptional()
  @IsString()
  bankCode?: string;
}

export class IpsTransferDto {
  @ApiProperty({ example: '1000123456789', description: 'Local source account number' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: 'AWINETAA', description: 'Recipient bank BIC/code' })
  @IsNotEmpty()
  @IsString()
  destinationBankCode!: string;

  @ApiProperty({ example: '0132049281900', description: 'Recipient external account number' })
  @IsNotEmpty()
  @IsString()
  destinationAccountNumber!: string;

  @ApiProperty({ example: 'Dawit Yohannes', description: 'Beneficiary name' })
  @IsNotEmpty()
  @IsString()
  beneficiaryName!: string;

  @ApiProperty({ example: '2500.0000', description: 'Transfer amount' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'IPS P2P Interbank Transfer' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IPS-TX-20261020-0091', description: 'ISO 20022 EndToEndId' })
  @IsOptional()
  @IsString()
  endToEndId?: string;

  @ApiPropertyOptional({ example: 'IDEMP-IPS-98124' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}

export class IpsQrPaymentDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: 'ETH.IPS.MCH.8841029', description: 'Payload decoded from QR standard' })
  @IsNotEmpty()
  @IsString()
  qrPayload!: string;

  @ApiProperty({ example: '1000987654321', description: 'Merchant receiving account' })
  @IsNotEmpty()
  @IsString()
  merchantAccount!: string;

  @ApiProperty({ example: 'Fresh Supermarket Bole', description: 'Merchant legal trade name' })
  @IsNotEmpty()
  @IsString()
  merchantName!: string;

  @ApiProperty({ example: '450.0000' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'Retail Grocery Checkout' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-QR-8821' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}

export class IpsStatusDto {
  @ApiProperty({ example: 'IPS2699182341' })
  @IsNotEmpty()
  @IsString()
  transactionReference!: string;
}
