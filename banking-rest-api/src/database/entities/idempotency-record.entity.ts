import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  Index,
} from 'typeorm';

export enum IdempotencyStatus {
  PENDING = 'PENDING',
  COMPLETED = 'COMPLETED',
  FAILED = 'FAILED',
}

@Entity('idempotency_records')
export class IdempotencyRecord {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index({ unique: true })
  @Column({ name: 'idempotency_key', type: 'varchar', length: 100, unique: true })
  idempotencyKey!: string;

  @Column({ name: 'request_hash', type: 'varchar', length: 64 })
  requestHash!: string; // SHA-256 of canonical request payload

  @Column({ name: 'operation', type: 'varchar', length: 100 })
  operation!: string; // e.g. "FUND_TRANSFER", "BILL_PAYMENT"

  @Column({ name: 'caller_id', type: 'varchar', length: 100, nullable: true })
  callerId!: string | null;

  @Column({
    name: 'status',
    type: 'enum',
    enum: IdempotencyStatus,
    default: IdempotencyStatus.PENDING,
  })
  status!: IdempotencyStatus;

  @Column({ name: 'transaction_reference', type: 'varchar', length: 50, nullable: true })
  transactionReference!: string | null;

  @Column({ name: 'response_payload', type: 'jsonb', nullable: true })
  responsePayload!: Record<string, any> | null;

  @Column({ name: 'http_status_code', type: 'int', nullable: true })
  httpStatusCode!: number | null;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
