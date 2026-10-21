import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Customer } from '../../database/entities/customer.entity';
import { Account } from '../../database/entities/account.entity';
import { Branch } from '../../database/entities/branch.entity';
import { Currency } from '../../database/entities/currency.entity';
import { AccountType } from '../../database/entities/account-type.entity';
import { AccountParty } from '../../database/entities/account-party.entity';
import { AccountBalance } from '../../database/entities/account-balance.entity';
import { Transaction } from '../../database/entities/transaction.entity';
import { LedgerEntry } from '../../database/entities/ledger-entry.entity';
import { CustomersService } from './customers.service';
import { CustomersController } from './customers.controller';

@Module({
  imports: [
    TypeOrmModule.forFeature([
      Customer,
      Account,
      Branch,
      Currency,
      AccountType,
      AccountParty,
      AccountBalance,
      Transaction,
      LedgerEntry,
    ]),
  ],
  controllers: [CustomersController],
  providers: [CustomersService],
  exports: [CustomersService],
})
export class CustomersModule {}
