import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  ManyToOne,
  JoinColumn,
  Index,
} from 'typeorm';
import { StandingOrder } from './standing-order.entity';

export enum ExecutionStatus {
  SUCCESS = 'SUCCESS',
  FAILED = 'FAILED',
  SKIPPED = 'SKIPPED',
}

@Entity('standing_order_executions')
export class StandingOrderExecution {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index()
  @Column({ name: 'standing_order_id', type: 'uuid' })
  standingOrderId!: string;

  @ManyToOne(() => StandingOrder, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'standing_order_id' })
  standingOrder!: StandingOrder;

  @Column({ name: 'execution_date', type: 'date' })
  executionDate!: string;

  @Column({ name: 'amount', type: 'numeric', precision: 18, scale: 4 })
  amount!: string;

  @Column({ name: 'transaction_reference', type: 'varchar', length: 50, nullable: true })
  transactionReference!: string | null;

  @Column({
    name: 'status',
    type: 'enum',
    enum: ExecutionStatus,
    default: ExecutionStatus.SUCCESS,
  })
  status!: ExecutionStatus;

  @Column({ name: 'failure_reason', type: 'varchar', length: 255, nullable: true })
  failureReason!: string | null;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;
}
