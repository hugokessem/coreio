import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsNumberString, IsEnum } from 'class-validator';
import { TransactionChannel } from '../../../database/entities/transaction.entity';

export class TransferVerifyDto {
  @ApiProperty({ example: '1000123456789', description: 'Debiting account number' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: '1000987654321', description: 'Crediting account number' })
  @IsNotEmpty()
  @IsString()
  destinationAccountNumber!: string;

  @ApiProperty({ example: '1500.0000', description: 'Transfer amount' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'ETB', description: 'Transfer currency' })
  @IsOptional()
  @IsString()
  currencyCode?: string;

  @ApiPropertyOptional({ enum: TransactionChannel, example: TransactionChannel.MOBILE })
  @IsOptional()
  @IsEnum(TransactionChannel)
  channel?: TransactionChannel;
}

export class ExecuteTransferDto {
  @ApiProperty({ example: '1000123456789', description: 'Debiting account number' })
  @IsNotEmpty()
  @IsString()
  sourceAccountNumber!: string;

  @ApiProperty({ example: '1000987654321', description: 'Crediting account number' })
  @IsNotEmpty()
  @IsString()
  destinationAccountNumber!: string;

  @ApiProperty({ example: '1500.0000', description: 'Transfer amount' })
  @IsNotEmpty()
  @IsNumberString()
  amount!: string;

  @ApiPropertyOptional({ example: 'ETB', description: 'Currency code' })
  @IsOptional()
  @IsString()
  currencyCode?: string;

  @ApiProperty({ example: 'Invoice #8472 Payment', description: 'Payment narration/memo' })
  @IsNotEmpty()
  @IsString()
  narration!: string;

  @ApiPropertyOptional({ example: 'IDEMP-UUID-98234-ABCD', description: 'Idempotency key for preventing duplicates' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;

  @ApiPropertyOptional({ enum: TransactionChannel, example: TransactionChannel.MOBILE })
  @IsOptional()
  @IsEnum(TransactionChannel)
  channel?: TransactionChannel;
}

export class TransferFromHoldDto {
  @ApiProperty({ example: 'HLD-001', description: 'Active Hold reference to settle' })
  @IsNotEmpty()
  @IsString()
  holdReference!: string;

  @ApiProperty({ example: '1000987654321', description: 'Crediting destination account' })
  @IsNotEmpty()
  @IsString()
  destinationAccountNumber!: string;

  @ApiPropertyOptional({ example: 'Settlement of judicial hold' })
  @IsOptional()
  @IsString()
  narration?: string;

  @ApiPropertyOptional({ example: 'IDEMP-HOLD-SETTLE-001' })
  @IsOptional()
  @IsString()
  idempotencyKey?: string;
}

export class TransactionStatusDto {
  @ApiProperty({ example: 'FT2629399881', description: 'Transaction reference number' })
  @IsNotEmpty()
  @IsString()
  transactionReference!: string;
}

export class ReversalDto {
  @ApiProperty({ example: 'FT2629399881', description: 'Transaction reference to reverse' })
  @IsNotEmpty()
  @IsString()
  transactionReference!: string;

  @ApiProperty({ example: 'Erroneous debit reversal requested by customer' })
  @IsNotEmpty()
  @IsString()
  reason!: string;
}
