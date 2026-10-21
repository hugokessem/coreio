import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Account } from '../../database/entities/account.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { AccountHold } from '../../database/entities/account-hold.entity';
import { Customer } from '../../database/entities/customer.entity';
import { Branch } from '../../database/entities/branch.entity';
import { Currency } from '../../database/entities/currency.entity';
import { AccountType } from '../../database/entities/account-type.entity';
import { Transaction } from '../../database/entities/transaction.entity';
import { LedgerEntry } from '../../database/entities/ledger-entry.entity';
import { AccountsService } from './accounts.service';
import { AccountsController } from './accounts.controller';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Account,
      AccountBalance,
      AccountHold,
      Customer,
      Branch,
      Currency,
      AccountType,
      Transaction,
      LedgerEntry,
    ]),
  ],
  controllers: [AccountsController],
  providers: [AccountsService],
  exports: [AccountsService],
})
export class AccountsModule {}
