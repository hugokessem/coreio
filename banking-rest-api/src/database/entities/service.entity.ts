import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  Index,
} from 'typeorm';

export enum FeeType {
  FLAT = 'FLAT',
  PERCENTAGE = 'PERCENTAGE',
}

@Entity('services')
export class ServiceEntity {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'service_code', type: 'varchar', length: 50, unique: true })
  serviceCode!: string; // e.g. "FT_INT", "FT_IPS", "BILL_ETHIO"

  @Column({ name: 'service_name', type: 'varchar', length: 100 })
  serviceName!: string;

  @Column({ name: 'category', type: 'varchar', length: 50 })
  category!: string;

  @Column({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' })
  status!: string;

  @Column({
    name: 'min_amount',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '1.0000',
  })
  minAmount!: string;

  @Column({
    name: 'max_amount',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '1000000.0000',
  })
  maxAmount!: string;

  @Column({
    name: 'daily_limit',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '5000000.0000',
  })
  dailyLimit!: string;

  @Column({
    name: 'fee_type',
    type: 'enum',
    enum: FeeType,
    default: FeeType.FLAT,
  })
  feeType!: FeeType;

  @Column({
    name: 'fee_value',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  feeValue!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
