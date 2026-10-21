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
import { Branch } from './branch.entity';

export enum CardType {
  DEBIT = 'DEBIT',
  PREPAID = 'PREPAID',
  CREDIT = 'CREDIT',
}

export enum CardBrand {
  MASTERCARD = 'MASTERCARD',
  VISA = 'VISA',
  LOCAL = 'LOCAL',
}

export enum CardRequestType {
  NEW = 'NEW',
  REPLACEMENT = 'REPLACEMENT',
  RENEWAL = 'RENEWAL',
}

export enum CardRequestStatus {
  PENDING = 'PENDING',
  APPROVED = 'APPROVED',
  PRINTING = 'PRINTING',
  DISPATCHED = 'DISPATCHED',
  DELIVERED = 'DELIVERED',
  REJECTED = 'REJECTED',
}

@Entity('card_requests')
export class CardRequest {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'request_reference', type: 'varchar', length: 50, unique: true })
  requestReference!: string;

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

  @Column({
    name: 'request_type',
    type: 'enum',
    enum: CardRequestType,
    default: CardRequestType.NEW,
  })
  requestType!: CardRequestType;

  @Column({
    name: 'status',
    type: 'enum',
    enum: CardRequestStatus,
    default: CardRequestStatus.PENDING,
  })
  status!: CardRequestStatus;

  @Column({ name: 'target_branch_id', type: 'uuid' })
  targetBranchId!: string;

  @ManyToOne(() => Branch)
  @JoinColumn({ name: 'target_branch_id' })
  targetBranch!: Branch;

  @Column({ name: 'rejection_reason', type: 'varchar', length: 255, nullable: true })
  rejectionReason!: string | null;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
