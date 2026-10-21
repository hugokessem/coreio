import { ApiProperty } from '@nestjs/swagger';
import { IsNotEmpty, IsString } from 'class-validator';

export class LoginDto {
  @ApiProperty({ example: 'admin', description: 'Username for banking portal or API consumer' })
  @IsNotEmpty()
  @IsString()
  username!: string;

  @ApiProperty({ example: 'Admin123!', description: 'User password' })
  @IsNotEmpty()
  @IsString()
  password!: string;
}
