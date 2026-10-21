import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  ManyToOne,
  JoinColumn,
  Index,
} from 'typeorm';
import { Transaction } from './transaction.entity';
import { Account } from './account.entity';

export enum EntryType {
  DEBIT = 'DEBIT',
  CREDIT = 'CREDIT',
}

@Entity('ledger_entries')
export class LedgerEntry {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index()
  @Column({ name: 'transaction_id', type: 'uuid' })
  transactionId!: string;

  @ManyToOne(() => Transaction, (tx) => tx.ledgerEntries, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'transaction_id' })
  transaction!: Transaction;

  @Index()
  @Column({ name: 'account_id', type: 'uuid' })
  accountId!: string;

  @ManyToOne(() => Account, { onDelete: 'RESTRICT' })
  @JoinColumn({ name: 'account_id' })
  account!: Account;

  @Column({
    name: 'entry_type',
    type: 'enum',
    enum: EntryType,
  })
  entryType!: EntryType;

  @Column({
    name: 'amount',
    type: 'numeric',
    precision: 18,
    scale: 4,
  })
  amount!: string;

  @Column({ name: 'currency_code', type: 'varchar', length: 3 })
  currencyCode!: string;

  @Column({
    name: 'balance_after',
    type: 'numeric',
    precision: 18,
    scale: 4,
  })
  balanceAfter!: string;

  @Column({ name: 'narration', type: 'varchar', length: 255 })
  narration!: string;

  @Column({ name: 'value_date', type: 'date', default: () => 'CURRENT_DATE' })
  valueDate!: string;

  @CreateDateColumn({ name: 'posted_at' })
  postedAt!: Date;
}
