import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsNumberString } from 'class-validator';

export class WalletCustomerLookupDto {
  @ApiProperty({ example: '+251911223344', description: 'CBE Birr registered mobile phone number' })
  @IsNotEmpty()
  @IsString()
  phoneNumber!: string;
}

export class AccountToWalletTransferDto {
  @ApiProperty({ example: '1000123456789', description: 'CBE Bank Account Number' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: '+251911223344', description: 'CBE Birr Wallet Phone Number' })
  @IsNotEmpty()
  @IsString()
  walletPhoneNumber!: string;

  @ApiProperty({ example: '500.0000', description: 'Amount in ETB' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'Top-up CBE Birr Wallet' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-WALLET-TOPUP-001' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}

export class WalletToAccountTransferDto {
  @ApiProperty({ example: '+251911223344', description: 'CBE Birr Wallet Phone Number' })
  @IsNotEmpty()
  @IsString()
  walletPhoneNumber!: string;

  @ApiProperty({ example: '1000123456789', description: 'CBE Bank Account Number' })
  @IsNotEmpty()
  @IsString()
  destinationAccountNumber!: string;

  @ApiProperty({ example: '500.0000', description: 'Amount in ETB' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'Transfer from CBE Birr to Savings' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-WALLET-CASHOUT-001' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}

export class WalletAgentLookupDto {
  @ApiProperty({ example: '+251911223344', description: 'Agent phone number or unique agent identifier' })
  @IsNotEmpty()
  @IsString()
  agentIdentifier!: string;
}

export class WalletAgentTransferDto {
  @ApiProperty({ example: '1000123456789', description: 'CBE Bank Account Number' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: '+251911223344', description: 'Agent mobile phone or agent wallet ID' })
  @IsNotEmpty()
  @IsString()
  agentIdentifier!: string;

  @ApiProperty({ example: '1000.0000', description: 'Amount in ETB' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'Float disbursement to agent' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-AGENT-TRF-001' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}
