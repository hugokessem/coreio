import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  Index,
} from 'typeorm';

export enum AccountCategory {
  SAVINGS = 'SAVINGS',
  CURRENT = 'CURRENT',
  FIXED_DEPOSIT = 'FIXED_DEPOSIT',
  OVERDRAFT = 'OVERDRAFT',
  LOAN = 'LOAN',
}

@Entity('account_types')
export class AccountType {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'code', type: 'varchar', length: 20, unique: true })
  code!: string; // e.g. "6501" (Savings), "6001" (Current)

  @Column({ name: 'name', type: 'varchar', length: 100 })
  name!: string;

  @Column({
    name: 'category',
    type: 'enum',
    enum: AccountCategory,
    default: AccountCategory.SAVINGS,
  })
  category!: AccountCategory;

  @Column({
    name: 'minimum_balance',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '0.0000',
  })
  minimumBalance!: string;

  @Column({ name: 'is_interest_bearing', type: 'boolean', default: false })
  isInterestBearing!: boolean;

  @Column({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' })
  status!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
