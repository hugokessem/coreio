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
import { Customer } from './customer.entity';

@Entity('customer_limits')
@Index(['customerId', 'serviceCode'], { unique: true })
export class CustomerLimit {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index()
  @Column({ name: 'customer_id', type: 'uuid' })
  customerId!: string;

  @ManyToOne(() => Customer, (c) => c.limits, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'customer_id' })
  customer!: Customer;

  @Column({ name: 'service_code', type: 'varchar', length: 50 })
  serviceCode!: string;

  @Column({
    name: 'daily_limit',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '500000.0000',
  })
  dailyLimit!: string;

  @Column({
    name: 'per_transaction_limit',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '100000.0000',
  })
  perTransactionLimit!: string;

  @Column({
    name: 'daily_used',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  dailyUsed!: string;

  @Column({ name: 'last_reset_date', type: 'date', default: () => 'CURRENT_DATE' })
  lastResetDate!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
