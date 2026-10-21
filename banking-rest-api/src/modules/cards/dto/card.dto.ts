import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsEnum } from 'class-validator';
import { CardType, CardBrand } from '../../../database/entities/card-request.entity';

export class RequestCardDto {
  @ApiProperty({ example: '1000123456789' })
  @IsNotEmpty()
  @IsString()
  accountNumber!: string;

  @ApiPropertyOptional({ enum: CardType, default: CardType.DEBIT })
  @IsOptional()
  @IsEnum(CardType)
  cardType?: CardType;

  @ApiPropertyOptional({ enum: CardBrand, default: CardBrand.MASTERCARD })
  @IsOptional()
  @IsEnum(CardBrand)
  cardBrand?: CardBrand;

  @ApiProperty({ example: '001', description: 'Branch code for collection' })
  @IsNotEmpty()
  @IsString()
  targetBranchCode!: string;
}

export class ReplaceCardDto {
  @ApiProperty({ example: '541234******7890', description: 'Masked card number or card hash' })
  @IsNotEmpty()
  @IsString()
  cardNumberMasked!: string;

  @ApiProperty({ example: 'Card damaged / magnetic stripe unreadable' })
  @IsNotEmpty()
  @IsString()
  reason!: string;

  @ApiProperty({ example: '100', description: 'Collection branch code' })
  @IsNotEmpty()
  @IsString()
  targetBranchCode!: string;
}

export class ListCardsDto {
  @ApiPropertyOptional({ example: '1000123456789' })
  @IsOptional()
  @IsString()
  accountNumber?: string;

  @ApiPropertyOptional({ example: '10000001' })
  @IsOptional()
  @IsString()
  customerNumber?: string;
}
