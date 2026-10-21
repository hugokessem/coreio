import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Account } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { Transaction } from '../../database/entities/transaction.entity';
import { LedgerEntry } from '../../database/entities/ledger-entry.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import { PaymentsService } from './payments.service';
import { PaymentsController } from './payments.controller';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Account,
      AccountBalance,
      Transaction,
      LedgerEntry,
      ServiceEntity,
    ]),
  ],
  controllers: [PaymentsController],
  providers: [PaymentsService],
  exports: [PaymentsService],
})
export class PaymentsModule {}
