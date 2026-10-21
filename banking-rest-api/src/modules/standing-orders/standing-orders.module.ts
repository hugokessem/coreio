import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { StandingOrder } from '../../database/entities/standing-order.entity';
import { StandingOrderExecution } from '../../database/entities/standing-order-execution.entity';
import { Account } from '../../database/entities/account.entity';
import { StandingOrdersService } from './standing-orders.service';
import { StandingOrdersController } from './standing-orders.controller';
import { TransfersModule } from '../transfers/transfers.module';

@Module({
  imports: [
    TypeOrmModule.forFeature([StandingOrder, StandingOrderExecution, Account]),
    TransfersModule,
  ],
  controllers: [StandingOrdersController],
  providers: [StandingOrdersService],
  exports: [StandingOrdersService],
})
export class StandingOrdersModule {}
