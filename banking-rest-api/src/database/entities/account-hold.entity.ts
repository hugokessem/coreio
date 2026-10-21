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

export enum HoldStatus {
  ACTIVE = 'ACTIVE',
  RELEASED = 'RELEASED',
  SETTLED = 'SETTLED',
  EXPIRED = 'EXPIRED',
}

@Entity('account_holds')
export class AccountHold {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'hold_reference', type: 'varchar', length: 50, unique: true })
  holdReference!: string; // Maps to CoreIO lock ID / ACLOCKEDEVENTS ID

  @Index()
  @Column({ name: 'account_id', type: 'uuid' })
  accountId!: string;

  @ManyToOne(() => Account, (account) => account.holds, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'account_id' })
  account!: Account;

  @Column({
    name: 'amount',
    type: 'numeric',
    precision: 18,
    scale: 4,
  })
  amount!: string;

  @Column({
    name: 'remaining_amount',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  remainingAmount!: string;

  @Column({ name: 'currency_code', type: 'varchar', length: 3 })
  currencyCode!: string;

  @Column({ name: 'reason', type: 'varchar', length: 255 })
  reason!: string;

  @Column({ name: 'lock_event', type: 'varchar', length: 50, default: 'HOLD' })
  lockEvent!: string; // e.g. "COURT_ORDER", "PLEDGE", "PENDING_SETTLEMENT"

  @Column({
    name: 'status',
    type: 'enum',
    enum: HoldStatus,
    default: HoldStatus.ACTIVE,
  })
  status!: HoldStatus;

  @Column({ name: 'from_date', type: 'date', nullable: true })
  fromDate!: string | null;

  @Column({ name: 'to_date', type: 'date', nullable: true })
  toDate!: string | null;

  @Column({ name: 'released_at', type: 'timestamp', nullable: true })
  releasedAt!: Date | null;

  @Column({ name: 'released_reason', type: 'varchar', length: 255, nullable: true })
  releasedReason!: string | null;

  @Column({ name: 'created_by', type: 'varchar', length: 100, default: 'SYSTEM' })
  createdBy!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
