import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  Index,
} from 'typeorm';

export enum UserRole {
  ADMIN = 'ADMIN',
  TELLER = 'TELLER',
  API_CLIENT = 'API_CLIENT',
  AUDITOR = 'AUDITOR',
}

@Entity('users')
export class User {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'username', type: 'varchar', length: 50, unique: true })
  username!: string;

  @Column({ name: 'password_hash', type: 'varchar', length: 255 })
  passwordHash!: string;

  @Column({ name: 'email', type: 'varchar', length: 100 })
  email!: string;

  @Column({
    name: 'roles',
    type: 'simple-array',
    default: UserRole.API_CLIENT,
  })
  roles!: string[];

  @Column({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' })
  status!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
