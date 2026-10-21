import {
  Entity,
  PrimaryColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
} from 'typeorm';

@Entity('currencies')
export class Currency {
  @PrimaryColumn({ type: 'varchar', length: 3 })
  code!: string; // e.g. "ETB", "USD", "EUR", "GBP"

  @Column({ name: 'name', type: 'varchar', length: 50 })
  name!: string;

  @Column({ name: 'numeric_code', type: 'varchar', length: 3 })
  numericCode!: string; // e.g. "230" for ETB, "840" for USD

  @Column({ name: 'decimal_places', type: 'int', default: 2 })
  decimalPlaces!: number;

  @Column({ name: 'is_base_currency', type: 'boolean', default: false })
  isBaseCurrency!: boolean;

  @Column({ name: 'status', type: 'varchar', length: 20, default: 'ACTIVE' })
  status!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
