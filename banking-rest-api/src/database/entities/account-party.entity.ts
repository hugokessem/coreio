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

export enum PartyRole {
  PRIMARY_OWNER = 'PRIMARY_OWNER',
  JOINT_OWNER = 'JOINT_OWNER',
  AUTHORIZED_SIGNATORY = 'AUTHORIZED_SIGNATORY',
  GUARANTOR = 'GUARANTOR',
  BENEFICIAL_OWNER = 'BENEFICIAL_OWNER',
}

@Entity('account_parties')
@Index(['accountId', 'customerId', 'partyRole'], { unique: true })
export class AccountParty {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index()
  @Column({ name: 'account_id', type: 'uuid' })
  accountId!: string;

  @ManyToOne(() => Account, { onDelete: 'CASCADE' })
  @JoinColumn({ name: 'account_id' })
  account!: Account;

  @Index()
  @Column({ name: 'customer_id', type: 'uuid' })
  customerId!: string;

  @ManyToOne(() => Customer, { onDelete: 'RESTRICT' })
  @JoinColumn({ name: 'customer_id' })
  customer!: Customer;

  @Column({
    name: 'party_role',
    type: 'enum',
    enum: PartyRole,
    default: PartyRole.PRIMARY_OWNER,
  })
  partyRole!: PartyRole;

  @Column({ name: 'authorization_level', type: 'varchar', length: 50, default: 'SOLE_SIGNATORY' })
  authorizationLevel!: string; // SOLE_SIGNATORY, JOINT_SIGNATORY, VIEW_ONLY

  @Column({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' })
  status!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
