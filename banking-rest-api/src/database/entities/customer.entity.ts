import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  OneToMany,
  Index,
} from 'typeorm';
import { Account } from './account.entity';
import { CustomerLimit } from './customer-limit.entity';

export enum CustomerType {
  INDIVIDUAL = 'INDIVIDUAL',
  CORPORATE = 'CORPORATE',
}

export enum CustomerStatus {
  ACTIVE = 'ACTIVE',
  INACTIVE = 'INACTIVE',
  SUSPENDED = 'SUSPENDED',
  DECEASED = 'DECEASED',
}

export enum AMLStatus {
  CLEAR = 'CLEAR',
  SUSPECT = 'SUSPECT',
  BLOCKED = 'BLOCKED',
}

@Entity('customers')
export class Customer {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'customer_number', type: 'varchar', length: 20, unique: true })
  customerNumber!: string;

  @Column({
    name: 'customer_type',
    type: 'enum',
    enum: CustomerType,
    default: CustomerType.INDIVIDUAL,
  })
  customerType!: CustomerType;

  @Column({ name: 'first_name', type: 'varchar', length: 100, nullable: true })
  firstName!: string | null;

  @Column({ name: 'middle_name', type: 'varchar', length: 100, nullable: true })
  middleName!: string | null;

  @Column({ name: 'last_name', type: 'varchar', length: 100, nullable: true })
  lastName!: string | null;

  @Column({ name: 'legal_name', type: 'varchar', length: 255, nullable: true })
  legalName!: string | null;

  @Column({ name: 'date_of_birth', type: 'date', nullable: true })
  dateOfBirth!: string | null;

  @Column({ name: 'gender', type: 'varchar', length: 10, nullable: true })
  gender!: string | null;

  @Index()
  @Column({ name: 'phone_number', type: 'varchar', length: 20, nullable: true })
  phoneNumber!: string | null;

  @Column({ name: 'email', type: 'varchar', length: 100, nullable: true })
  email!: string | null;

  @Column({ name: 'address', type: 'text', nullable: true })
  address!: string | null;

  @Index()
  @Column({ name: 'national_id', type: 'varchar', length: 50, nullable: true })
  nationalId!: string | null; // Fayda / National ID

  @Column({ name: 'tin', type: 'varchar', length: 30, nullable: true })
  tin!: string | null;

  @Column({ name: 'kyc_level', type: 'varchar', length: 20, default: 'LEVEL_2' })
  kycLevel!: string;

  @Column({
    name: 'aml_status',
    type: 'enum',
    enum: AMLStatus,
    default: AMLStatus.CLEAR,
  })
  amlStatus!: AMLStatus;

  @Column({
    name: 'status',
    type: 'enum',
    enum: CustomerStatus,
    default: CustomerStatus.ACTIVE,
  })
  status!: CustomerStatus;

  @Column({ name: 'is_authorized', type: 'boolean', default: true })
  isAuthorized!: boolean;

  @Column({ name: 'created_by', type: 'varchar', length: 100, default: 'SYSTEM' })
  createdBy!: string;

  @Column({ name: 'authorized_by', type: 'varchar', length: 100, nullable: true })
  authorizedBy!: string | null;

  @Column({ name: 'authorized_at', type: 'timestamp', nullable: true })
  authorizedAt!: Date | null;

  @OneToMany(() => Account, (account) => account.customer)
  accounts!: Account[];

  @OneToMany(() => CustomerLimit, (limit) => limit.customer)
  limits!: CustomerLimit[];

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;

  get fullName(): string {
    if (this.legalName) return this.legalName;
    return [this.firstName, this.middleName, this.lastName].filter(Boolean).join(' ');
  }
}
