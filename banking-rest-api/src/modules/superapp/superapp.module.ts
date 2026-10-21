import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { SuperAppUser } from '../../database/entities/super-app-user.entity';
import { Customer } from '../../database/entities/customer.entity';
import { SuperAppService } from './superapp.service';
import { SuperAppController } from './superapp.controller';

@Module({
  imports: [TypeOrmModule.forFeature([SuperAppUser, Customer])],
  controllers: [SuperAppController],
  providers: [SuperAppService],
  exports: [SuperAppService],
})
export class SuperAppModule {}
