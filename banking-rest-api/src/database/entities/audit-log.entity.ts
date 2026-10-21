import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  Index,
} from 'typeorm';

@Entity('audit_logs')
export class AuditLog {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Index()
  @Column({ name: 'action', type: 'varchar', length: 100 })
  action!: string;

  @Column({ name: 'entity_name', type: 'varchar', length: 100 })
  entityName!: string;

  @Index()
  @Column({ name: 'entity_id', type: 'varchar', length: 100, nullable: true })
  entityId!: string | null;

  @Index()
  @Column({ name: 'user_id', type: 'varchar', length: 100, nullable: true })
  userId!: string | null;

  @Column({ name: 'ip_address', type: 'varchar', length: 45, nullable: true })
  ipAddress!: string | null;

  @Column({ name: 'user_agent', type: 'varchar', length: 255, nullable: true })
  userAgent!: string | null;

  @Column({ name: 'details', type: 'jsonb', nullable: true })
  details!: Record<string, any> | null;

  @CreateDateColumn({ name: 'timestamp' })
  timestamp!: Date;
}
