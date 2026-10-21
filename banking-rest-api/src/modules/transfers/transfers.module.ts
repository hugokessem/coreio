import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Account } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { AccountHold } from '../../database/entities/account-hold.entity';
import { Transaction } from '../../database/entities/transaction.entity';
import { LedgerEntry } from '../../database/entities/ledger-entry.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import { IdempotencyRecord } from '../../database/entities/idempotency-record.entity';
import { TransfersService } from './transfers.service';
import { TransfersController } from './transfers.controller';
import { IdempotencyService } from '../../common/services/idempotency.service';
import { ReconciliationService } from './reconciliation.service';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Account,
      AccountBalance,
      AccountHold,
      Transaction,
      LedgerEntry,
      ServiceEntity,
      IdempotencyRecord,
    ]),
  ],
  controllers: [TransfersController],
  providers: [TransfersService, IdempotencyService, ReconciliationService],
  exports: [TransfersService, ReconciliationService],
})
export class TransfersModule {}
