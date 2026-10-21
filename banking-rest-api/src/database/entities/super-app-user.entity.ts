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

export enum SuperAppStatus {
  ACTIVE = 'ACTIVE',
  SUSPENDED = 'SUSPENDED',
  DEACTIVATED = 'DEACTIVATED',
}

@Entity('super_app_users')
export class SuperAppUser {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'external_user_id', type: 'varchar', length: 100, unique: true })
  externalUserId!: string;

  @Index()
  @Column({ name: 'customer_id', type: 'uuid' })
  customerId!: string;

  @ManyToOne(() => Customer, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'customer_id' })
  customer!: Customer;

  @Column({ name: 'partner_code', type: 'varchar', length: 50, default: 'SUPERAPP_DEFAULT' })
  partnerCode!: string;

  @Column({
    name: 'status',
    type: 'enum',
    enum: SuperAppStatus,
    default: SuperAppStatus.ACTIVE,
  })
  status!: SuperAppStatus;

  @CreateDateColumn({ name: 'registered_at' })
  registeredAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
