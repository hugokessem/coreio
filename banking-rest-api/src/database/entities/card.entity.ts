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
import { Customer } from './customer.entity';
import { CardType, CardBrand } from './card-request.entity';

export enum CardStatus {
  ACTIVE = 'ACTIVE',
  BLOCKED = 'BLOCKED',
  EXPIRED = 'EXPIRED',
  CANCELLED = 'CANCELLED',
}

@Entity('cards')
export class Card {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Column({ name: 'card_number_masked', type: 'varchar', length: 30 })
  cardNumberMasked!: string; // e.g. "541234******7890"

  @Index({ unique: true })
  @Column({ name: 'card_hash', type: 'varchar', length: 64, unique: true })
  cardHash!: string; // SHA-256 hash of full PAN for duplicate prevention & lookup

  @Column({
    name: 'card_type',
    type: 'enum',
    enum: CardType,
    default: CardType.DEBIT,
  })
  cardType!: CardType;

  @Column({
    name: 'card_brand',
    type: 'enum',
    enum: CardBrand,
    default: CardBrand.MASTERCARD,
  })
  cardBrand!: CardBrand;

  @Index()
  @Column({ name: 'account_id', type: 'uuid' })
  accountId!: string;

  @ManyToOne(() => Account)
  @JoinColumn({ name: 'account_id' })
  account!: Account;

  @Index()
  @Column({ name: 'customer_id', type: 'uuid' })
  customerId!: string;

  @ManyToOne(() => Customer)
  @JoinColumn({ name: 'customer_id' })
  customer!: Customer;

  @Column({
    name: 'status',
    type: 'enum',
    enum: CardStatus,
    default: CardStatus.ACTIVE,
  })
  status!: CardStatus;

  @Column({ name: 'expiry_date', type: 'varchar', length: 5 })
  expiryDate!: string; // MM/YY

  @Column({
    name: 'daily_limit',
    type: 'numeric',
    precision: 18,
    scale: 4,
    default: '50000.0000',
  })
  dailyLimit!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
