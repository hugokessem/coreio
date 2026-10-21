import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  UpdateDateColumn,
  OneToOne,
  JoinColumn,
  Index,
  VersionColumn,
} from 'typeorm';
import { Account } from './account.entity';

@Entity('account_balances')
export class AccountBalance {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'account_id', type: 'uuid', unique: true })
  accountId!: string;

  @OneToOne(() => Account, (account) => account.balance, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'account_id' })
  account!: Account;

  @Column({
    name: 'ledger_balance',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  ledgerBalance!: string;

  @Column({
    name: 'available_balance',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  availableBalance!: string;

  @Column({
    name: 'locked_balance',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  lockedBalance!: string; // Active holds sum

  @Column({ name: 'currency_code', type: 'varchar', length: 3 })
  currencyCode!: string;

  @VersionColumn({ default: 1 })
  version!: number;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
