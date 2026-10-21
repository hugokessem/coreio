import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { CardRequest } from '../../database/entities/card-request.entity';
import { Card } from '../../database/entities/card.entity';
import { Account } from '../../database/entities/account.entity';
import { Branch } from '../../database/entities/branch.entity';
import { Customer } from '../../database/entities/customer.entity';
import { CardsService } from './cards.service';
import { CardsController } from './cards.controller';

@Module({
  imports: [TypeOrmModule.forFeature([CardRequest, Card, Account, Branch, Customer])],
  controllers: [CardsController],
  providers: [CardsService],
  exports: [CardsService],
})
export class CardsModule {}
