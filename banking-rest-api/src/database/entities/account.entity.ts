import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  ManyToOne,
  OneToOne,
  OneToMany,
  JoinColumn,
  Index,
} from 'typeorm';
import { Customer } from './customer.entity';
import { Branch } from './branch.entity';
import { Currency } from './currency.entity';
import { AccountType } from './account-type.entity';
import { AccountBalance } from './account-balance.entity';
import { AccountHold } from './account-hold.entity';

export enum AccountStatus {
  ACTIVE = 'ACTIVE',
  DORMANT = 'DORMANT',
  FROZEN = 'FROZEN',
  RESTRICTED = 'RESTRICTED',
  CLOSED = 'CLOSED',
}

@Entity('accounts')
export class Account {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'account_number', type: 'varchar', length: 30, unique: true })
  accountNumber!: string;

  @Index()
  @Column({ name: 'customer_id', type: 'uuid' })
  customerId!: string;

  @ManyToOne(() => Customer, (customer) => customer.accounts)
  @JoinColumn({ name: 'customer_id' })
  customer!: Customer;

  @Index()
  @Column({ name: 'account_type_id', type: 'uuid' })
  accountTypeId!: string;

  @ManyToOne(() => AccountType)
  @JoinColumn({ name: 'account_type_id' })
  accountType!: AccountType;

  @Index()
  @Column({ name: 'branch_id', type: 'uuid' })
  branchId!: string;

  @ManyToOne(() => Branch, (branch) => branch.accounts)
  @JoinColumn({ name: 'branch_id' })
  branch!: Branch;

  @Column({ name: 'currency_code', type: 'varchar', length: 3 })
  currencyCode!: string;

  @ManyToOne(() => Currency)
  @JoinColumn({ name: 'currency_code' })
  currency!: Currency;

  @Column({ name: 'account_name', type: 'varchar', length: 255 })
  accountName!: string;

  @Column({
    name: 'status',
    type: 'enum',
    enum: AccountStatus,
    default: AccountStatus.ACTIVE,
  })
  status!: AccountStatus;

  @Column({ name: 'is_posting_restricted', type: 'boolean', default: false })
  isPostingRestricted!: boolean;

  @Column({ name: 'posting_restriction_reason', type: 'varchar', length: 255, nullable: true })
  postingRestrictionReason!: string | null;

  @Column({ name: 'is_debit_blocked', type: 'boolean', default: false })
  isDebitBlocked!: boolean;

  @Column({ name: 'is_credit_blocked', type: 'boolean', default: false })
  isCreditBlocked!: boolean;

  @Column({ name: 'opened_at', type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' })
  openedAt!: Date;

  @Column({ name: 'closed_at', type: 'timestamp', nullable: true })
  closedAt!: Date | null;

  @Column({ name: 'last_activity_at', type: 'timestamp', nullable: true })
  lastActivityAt!: Date | null;

  @OneToOne(() => AccountBalance, (balance) => balance.account, { cascade: true })
  balance!: AccountBalance;

  @OneToMany(() => AccountHold, (hold) => hold.account)
  holds!: AccountHold[];

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
