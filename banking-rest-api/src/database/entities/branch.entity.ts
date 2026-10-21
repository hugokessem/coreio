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

@Entity('branches')
export class Branch {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'branch_code', type: 'varchar', length: 10, unique: true })
  branchCode!: string;

  @Column({ name: 'branch_name', type: 'varchar', length: 150 })
  branchName!: string;

  @Column({ name: 'address', type: 'text', nullable: true })
  address!: string | null;

  @Column({ name: 'city', type: 'varchar', length: 100, default: 'Addis Ababa' })
  city!: string;

  @Column({ name: 'swift_code', type: 'varchar', length: 20, nullable: true })
  swiftCode!: string | null;

  @Column({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' })
  status!: string;

  @OneToMany(() => Account, (account) => account.branch)
  accounts!: Account[];

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
