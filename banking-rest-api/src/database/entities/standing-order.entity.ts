import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  ManyToOne,
  JoinColumn,
  Index,
} from 'typeorm';
import { Account } from './account.entity';

export enum StandingOrderFrequency {
  DAILY = 'DAILY',
  WEEKLY = 'WEEKLY',
  MONTHLY = 'MONTHLY',
  QUARTERLY = 'QUARTERLY',
  ANNUALLY = 'ANNUALLY',
}

export enum StandingOrderStatus {
  ACTIVE = 'ACTIVE',
  PAUSED = 'PAUSED',
  CANCELLED = 'CANCELLED',
  COMPLETED = 'COMPLETED',
}

@Entity('standing_orders')
export class StandingOrder {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'order_reference', type: 'varchar', length: 50, unique: true })
  orderReference!: string;

  @Index()
  @Column({ name: 'source_account_id', type: 'uuid' })
  sourceAccountId!: string;

  @ManyToOne(() => Account, { onDelete: 'RESTRICT' })
  @JoinColumn({ name: 'source_account_id' })
  sourceAccount!: Account;

  @Column({ name: 'destination_account_id', type: 'uuid', nullable: true })
  destinationAccountId!: string | null;

  @ManyToOne(() => Account, { nullable: true, onDelete: 'SET NULL' })
  @JoinColumn({ name: 'destination_account_id' })
  destinationAccount!: Account | null;

  @Column({ name: 'beneficiary_name', type: 'varchar', length: 150 })
  beneficiaryName!: string;

  @Column({ name: 'beneficiary_account', type: 'varchar', length: 50 })
  beneficiaryAccount!: string;

  @Column({ name: 'beneficiary_bank_code', type: 'varchar', length: 30, nullable: true })
  beneficiaryBankCode!: string | null;

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
    name: 'frequency',
    type: 'enum',
    enum: StandingOrderFrequency,
    default: StandingOrderFrequency.MONTHLY,
  })
  frequency!: StandingOrderFrequency;

  @Column({ name: 'start_date', type: 'date' })
  startDate!: string;

  @Column({ name: 'end_date', type: 'date', nullable: true })
  endDate!: string | null;

  @Column({ name: 'next_execution_date', type: 'date' })
  nextExecutionDate!: string;

  @Column({
    name: 'status',
    type: 'enum',
    enum: StandingOrderStatus,
    default: StandingOrderStatus.ACTIVE,
  })
  status!: StandingOrderStatus;

  @Column({ name: 'failure_count', type: 'int', default: 0 })
  failureCount!: number;

  @Column({ name: 'last_executed_at', type: 'timestamp', nullable: true })
  lastExecutedAt!: Date | null;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
