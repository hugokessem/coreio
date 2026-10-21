import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  ManyToOne,
  OneToMany,
  JoinColumn,
  Index,
} from 'typeorm';
import { Account } from './account.entity';
import { LedgerEntry } from './ledger-entry.entity';

export enum TransactionType {
  INTERNAL_TRANSFER = 'INTERNAL_TRANSFER',
  TRANSFER_FROM_HOLD = 'TRANSFER_FROM_HOLD',
  IPS_TRANSFER = 'IPS_TRANSFER',
  IPS_QR_PAYMENT = 'IPS_QR_PAYMENT',
  WALLET_TRANSFER = 'WALLET_TRANSFER',
  BILL_PAYMENT = 'BILL_PAYMENT',
  SPLIT_PAYMENT = 'SPLIT_PAYMENT',
  RTGS_TRANSFER = 'RTGS_TRANSFER',
  CASH_DEPOSIT = 'CASH_DEPOSIT',
  CASH_WITHDRAWAL = 'CASH_WITHDRAWAL',
  FEE = 'FEE',
  REVERSAL = 'REVERSAL',
}

export enum TransactionStatus {
  PENDING = 'PENDING',
  COMPLETED = 'COMPLETED',
  FAILED = 'FAILED',
  REVERSED = 'REVERSED',
}

export enum TransactionChannel {
  CORE = 'CORE',
  WEB = 'WEB',
  MOBILE = 'MOBILE',
  IPS = 'IPS',
  CBE_BIRR = 'CBE_BIRR',
  ATM = 'ATM',
  POS = 'POS',
  SYSTEM = 'SYSTEM',
}

@Entity('transactions')
export class Transaction {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'transaction_reference', type: 'varchar', length: 50, unique: true })
  transactionReference!: string; // e.g. FT2629399881

  @Index({ unique: true })
  @Column({ name: 'idempotency_key', type: 'varchar', length: 100, unique: true, nullable: true })
  idempotencyKey!: string | null;

  @Column({
    name: 'transaction_type',
    type: 'enum',
    enum: TransactionType,
    default: TransactionType.INTERNAL_TRANSFER,
  })
  transactionType!: TransactionType;

  @Column({
    name: 'channel',
    type: 'enum',
    enum: TransactionChannel,
    default: TransactionChannel.CORE,
  })
  channel!: TransactionChannel;

  @Index()
  @Column({ name: 'source_account_id', type: 'uuid', nullable: true })
  sourceAccountId!: string | null;

  @ManyToOne(() => Account, { nullable: true })
  @JoinColumn({ name: 'source_account_id' })
  sourceAccount!: Account | null;

  @Index()
  @Column({ name: 'destination_account_id', type: 'uuid', nullable: true })
  destinationAccountId!: string | null;

  @ManyToOne(() => Account, { nullable: true })
  @JoinColumn({ name: 'destination_account_id' })
  destinationAccount!: Account | null;

  @Column({ name: 'external_account', type: 'varchar', length: 50, nullable: true })
  externalAccount!: string | null;

  @Column({ name: 'external_bank_code', type: 'varchar', length: 30, nullable: true })
  externalBankCode!: string | null;

  @Column({
    name: 'amount',
    type: 'numeric',
    precision: 18,
    scale: 4,
  })
  amount!: string;

  @Column({
    name: 'fee_amount',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  feeAmount!: string;

  @Column({ name: 'currency_code', type: 'varchar', length: 3 })
  currencyCode!: string;

  @Column({
    name: 'exchange_rate',
    type: 'numeric',
    precision: 18,
    scale: 6,
    default: '1.000000',
  })
  exchangeRate!: string;

  @Column({ name: 'narration', type: 'varchar', length: 255 })
  narration!: string;

  @Column({
    name: 'status',
    type: 'enum',
    enum: TransactionStatus,
    default: TransactionStatus.PENDING,
  })
  status!: TransactionStatus;

  @Column({ name: 'failure_reason', type: 'varchar', length: 255, nullable: true })
  failureReason!: string | null;

  @Column({ name: 'reversal_of_transaction_id', type: 'uuid', nullable: true })
  reversalOfTransactionId!: string | null;

  @Column({ name: 'initiated_at', type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' })
  initiatedAt!: Date;

  @Column({ name: 'completed_at', type: 'timestamp', nullable: true })
  completedAt!: Date | null;

  @OneToMany(() => LedgerEntry, (entry) => entry.transaction, { cascade: true })
  ledgerEntries!: LedgerEntry[];

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
