import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';
import { IsNotEmpty, IsOptional, IsString, IsEnum, IsEmail, IsNumberString } from 'class-validator';
import { CustomerType } from '../../../database/entities/customer.entity';

export class CustomerLookupDto {
  @ApiPropertyOptional({ example: '10000001', description: 'Customer number' })
  @IsOptional()
  @IsString()
  customerNumber?: string;

  @ApiPropertyOptional({ example: '+251911223344', description: 'Customer registered phone number' })
  @IsOptional()
  @IsString()
  phoneNumber?: string;

  @ApiPropertyOptional({ example: '1000123456789', description: 'Account number associated with customer' })
  @IsOptional()
  @IsString()
  accountNumber?: string;

  @ApiPropertyOptional({ example: 'FAYDA-9876543210', description: 'National ID / Fayda number' })
  @IsOptional()
  @IsString()
  nationalId?: string;
}

export class CustomerDetailDto {
  @ApiProperty({ example: '10000001', description: 'Unique customer number' })
  @IsNotEmpty()
  @IsString()
  customerNumber!: string;
}

export class CreateCustomerDto {
  @ApiProperty({ enum: CustomerType, example: CustomerType.INDIVIDUAL })
  @IsEnum(CustomerType)
  customerType!: CustomerType;

  @ApiPropertyOptional({ example: 'Abebe' })
  @IsOptional()
  @IsString()
  firstName?: string;

  @ApiPropertyOptional({ example: 'Kebede' })
  @IsOptional()
  @IsString()
  middleName?: string;

  @ApiPropertyOptional({ example: 'Bikila' })
  @IsOptional()
  @IsString()
  lastName?: string;

  @ApiPropertyOptional({ example: 'Abebe Kebede Bikila' })
  @IsOptional()
  @IsString()
  legalName?: string;

  @ApiPropertyOptional({ example: 'Abebe Kebede Bikila' })
  @IsOptional()
  @IsString()
  fullName?: string;

  @ApiPropertyOptional({ example: '1985-04-12' })
  @IsOptional()
  @IsString()
  dateOfBirth?: string;

  @ApiPropertyOptional({ example: 'MALE' })
  @IsOptional()
  @IsString()
  gender?: string;

  @ApiProperty({ example: '+251911223344' })
  @IsNotEmpty()
  @IsString()
  phoneNumber!: string;

  @ApiPropertyOptional({ example: 'abebe.kebede@example.com' })
  @IsOptional()
  @IsEmail()
  email?: string;

  @ApiPropertyOptional({ example: 'Kirkos Sub-City, Addis Ababa' })
  @IsOptional()
  @IsString()
  address?: string;

  @ApiPropertyOptional({ example: 'Bole Medhanialem' })
  @IsOptional()
  @IsString()
  addressLine1?: string;

  @ApiPropertyOptional({ example: 'Addis Ababa' })
  @IsOptional()
  @IsString()
  city?: string;

  @ApiPropertyOptional({ example: 'FAYDA-9876543210' })
  @IsOptional()
  @IsString()
  nationalId?: string;

  @ApiPropertyOptional({ example: '1234567890' })
  @IsOptional()
  @IsString()
  tin?: string;
}

export class CustomerPhoneLookupDto {
  @ApiProperty({ example: '+251911223344', description: 'Phone number to lookup' })
  @IsNotEmpty()
  @IsString()
  phoneNumber!: string;
}

export class AmlVerifyDto {
  @ApiProperty({ example: '10000001' })
  @IsNotEmpty()
  @IsString()
  customerNumber!: string;
}

export class CheckExistenceDto {
  @ApiPropertyOptional({ example: '+251911223344' })
  @IsOptional()
  @IsString()
  phoneNumber?: string;

  @ApiPropertyOptional({ example: 'FAYDA-9876543210' })
  @IsOptional()
  @IsString()
  nationalId?: string;
}

export class CustomerEligibilityDto {
  @ApiPropertyOptional({ example: '10000001' })
  @IsOptional()
  @IsString()
  customerNumber?: string;

  @ApiPropertyOptional({ example: '+251911223344' })
  @IsOptional()
  @IsString()
  phoneNumber?: string;

  @ApiPropertyOptional({ example: 'FAYDA-9876543210' })
  @IsOptional()
  @IsString()
  nationalId?: string;

  @ApiPropertyOptional({ example: '1995-05-15' })
  @IsOptional()
  @IsString()
  dateOfBirth?: string;

  @ApiPropertyOptional({ example: 'MOBILE_BANKING' })
  @IsOptional()
  @IsString()
  channel?: string;
}

export class FaydaVerifyDto {
  @ApiPropertyOptional({ example: 'FAYDA-9876543210', description: 'Ethiopian Fayda National Identity string' })
  @IsOptional()
  @IsString()
  nationalId?: string;

  @ApiPropertyOptional({ example: 'FIN-ETH-99281928' })
  @IsOptional()
  @IsString()
  faydaId?: string;

  @ApiPropertyOptional({ example: 'Abebe Bikila' })
  @IsOptional()
  @IsString()
  fullName?: string;

  @ApiPropertyOptional({ example: '1985-04-12' })
  @IsOptional()
  @IsString()
  dateOfBirth?: string;

  @ApiPropertyOptional({ example: '+251911223344' })
  @IsOptional()
  @IsString()
  phoneNumber?: string;
}

export class CompositeOnboardDto {
  @ApiProperty({ enum: CustomerType, example: CustomerType.INDIVIDUAL })
  @IsEnum(CustomerType)
  customerType!: CustomerType;

  @ApiPropertyOptional({ example: 'Tewodros' })
  @IsOptional()
  @IsString()
  firstName?: string;

  @ApiPropertyOptional({ example: 'Kassahun' })
  @IsOptional()
  @IsString()
  middleName?: string;

  @ApiPropertyOptional({ example: 'Bekele' })
  @IsOptional()
  @IsString()
  lastName?: string;

  @ApiProperty({ example: '+251911889900' })
  @IsNotEmpty()
  @IsString()
  phoneNumber!: string;

  @ApiPropertyOptional({ example: 'FAYDA-7788990011' })
  @IsOptional()
  @IsString()
  nationalId?: string;

  @ApiProperty({ example: '6501', description: 'Account product category code (6501 Savings)' })
  @IsNotEmpty()
  @IsString()
  accountTypeCode!: string;

  @ApiProperty({ example: '001', description: 'Branch code (001 Addis Ababa Main)' })
  @IsNotEmpty()
  @IsString()
  branchCode!: string;

  @ApiProperty({ example: 'ETB' })
  @IsNotEmpty()
  @IsString()
  currencyCode!: string;

  @ApiPropertyOptional({ example: '500.0000', description: 'Initial deposit amount' })
  @IsOptional()
  @IsNumberString()
  initialDeposit?: string;
}
