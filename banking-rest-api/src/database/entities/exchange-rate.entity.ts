import {
  Entity,
  PrimaryGeneratedColumn,
  Column,
  CreateDateColumn,
  UpdateDateColumn,
  Index,
} from 'typeorm';

@Entity('exchange_rates')
@Index(['fromCurrency', 'toCurrency', 'effectiveDate'], { unique: true })
export class ExchangeRate {
  @PrimaryGeneratedColumn('uuid')
  id!: string;

  @Column({ name: 'from_currency', type: 'varchar', length: 3 })
  fromCurrency!: string;

  @Column({ name: 'to_currency', type: 'varchar', length: 3 })
  toCurrency!: string;

  @Column({ name: 'buy_rate', type: 'numeric', precision: 18, scale: 6 })
  buyRate!: string;

  @Column({ name: 'sell_rate', type: 'numeric', precision: 18, scale: 6 })
  sellRate!: string;

  @Column({ name: 'mid_rate', type: 'numeric', precision: 18, scale: 6 })
  midRate!: string;

  @Column({ name: 'effective_date', type: 'date' })
  effectiveDate!: string;

  @CreateDateColumn({ name: 'created_at' })
  createdAt!: Date;

  @UpdateDateColumn({ name: 'updated_at' })
  updatedAt!: Date;
}
