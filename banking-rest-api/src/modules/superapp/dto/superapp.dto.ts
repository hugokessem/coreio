import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsEnum } from 'class-validator';
import { SuperAppStatus } from '../../../database/entities/super-app-user.entity';

export class SuperAppSubscribeDto {
  @ApiProperty({ example: 'EXT-USER-TELCO-9921', description: 'External ecosystem partner user ID' })
  @IsNotEmpty()
  @IsString()
  externalUserId!: string;

  @ApiProperty({ example: '10000001', description: 'Bank customer number to link' })
  @IsNotEmpty()
  @IsString()
  customerNumber!: string;

  @ApiPropertyOptional({ example: 'TELEBIRR_MINIAPP', description: 'Partner platform identifier' })
  @IsOptional()
  @IsString()
  partnerCode?: string;
}

export class SuperAppUnsubscribeDto {
  @ApiProperty({ example: 'EXT-USER-TELCO-9921' })
  @IsNotEmpty()
  @IsString()
  externalUserId!: string;
}

export class SuperAppStatusDto {
  @ApiProperty({ example: 'EXT-USER-TELCO-9921' })
  @IsNotEmpty()
  @IsString()
  externalUserId!: string;
}

export class SuperAppStatusChangeDto {
  @ApiProperty({ example: 'EXT-USER-TELCO-9921' })
  @IsNotEmpty()
  @IsString()
  externalUserId!: string;

  @ApiProperty({ enum: SuperAppStatus, example: SuperAppStatus.SUSPENDED })
  @IsEnum(SuperAppStatus)
  status!: SuperAppStatus;

  @ApiPropertyOptional({ example: 'Temporary suspension requested by customer' })
  @IsOptional()
  @IsString()
  reason?: string;
}
