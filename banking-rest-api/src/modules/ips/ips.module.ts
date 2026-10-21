import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Account } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { Transaction } from '../../database/entities/transaction.entity';
import { LedgerEntry } from '../../database/entities/ledger-entry.entity';
import { ServiceEntity } from '../../database/entities/service.entity';
import { IpsService } from './ips.service';
import { IpsController } from './ips.controller';

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
  controllers: [IpsController],
  providers: [IpsService],
  exports: [IpsService],
})
export class IpsModule {}
