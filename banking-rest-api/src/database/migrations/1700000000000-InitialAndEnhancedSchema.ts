import { MigrationInterface, QueryRunner } from 'typeorm';

export class InitialAndEnhancedSchema1700000000000 implements MigrationInterface {
  name = 'InitialAndEnhancedSchema1700000000000';

  public async up(queryRunner: QueryRunner): Promise<void> {
    // 1. Account Parties (Normalized Joint/Authorized Signatories/Corporate Ownership)
    await queryRunner.query(`
      CREATE TABLE IF NOT EXISTS "account_parties" (
        "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
        "account_id" uuid NOT NULL,
        "customer_id" uuid NOT NULL,
        "party_role" varchar(30) NOT NULL DEFAULT 'PRIMARY_OWNER',
        "authorization_level" varchar(30) NOT NULL DEFAULT 'FULL',
        "effective_from" TIMESTAMP NOT NULL DEFAULT now(),
        "effective_to" TIMESTAMP,
        "status" varchar(20) NOT NULL DEFAULT 'ACTIVE',
        "created_at" TIMESTAMP NOT NULL DEFAULT now(),
        "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
        CONSTRAINT "PK_account_parties" PRIMARY KEY ("id"),
        CONSTRAINT "UQ_account_parties_account_customer" UNIQUE ("account_id", "customer_id"),
        CONSTRAINT "FK_account_parties_account" FOREIGN KEY ("account_id") REFERENCES "accounts"("id") ON DELETE CASCADE,
        CONSTRAINT "FK_account_parties_customer" FOREIGN KEY ("customer_id") REFERENCES "customers"("id") ON DELETE CASCADE
      );
    `);
    await queryRunner.query(`CREATE INDEX IF NOT EXISTS "IDX_account_parties_account" ON "account_parties" ("account_id");`);
    await queryRunner.query(`CREATE INDEX IF NOT EXISTS "IDX_account_parties_customer" ON "account_parties" ("customer_id");`);

    // 2. Idempotency Records (Durable request fingerprint and response caching)
    await queryRunner.query(`
      CREATE TABLE IF NOT EXISTS "idempotency_records" (
        "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
        "idempotency_key" varchar(128) NOT NULL,
        "operation" varchar(100) NOT NULL,
        "scope" varchar(100) NOT NULL DEFAULT 'DEFAULT',
        "request_hash" varchar(64) NOT NULL,
        "status" varchar(30) NOT NULL DEFAULT 'COMPLETED',
        "transaction_reference" varchar(100),
        "status_code" integer NOT NULL DEFAULT 200,
        "response_body" jsonb,
        "caller_identifier" varchar(100),
        "expires_at" TIMESTAMP,
        "created_at" TIMESTAMP NOT NULL DEFAULT now(),
        "updated_at" TIMESTAMP NOT NULL DEFAULT now(),
        CONSTRAINT "PK_idempotency_records" PRIMARY KEY ("id"),
        CONSTRAINT "UQ_idempotency_key_scope" UNIQUE ("idempotency_key", "scope")
      );
    `);
    await queryRunner.query(`CREATE INDEX IF NOT EXISTS "IDX_idempotency_records_key_scope" ON "idempotency_records" ("idempotency_key", "scope");`);

    // 3. Standing Order Executions (Durable execution history and audit)
    await queryRunner.query(`
      CREATE TABLE IF NOT EXISTS "standing_order_executions" (
        "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
        "standing_order_id" uuid NOT NULL,
        "execution_date" date NOT NULL,
        "amount" numeric(18,4) NOT NULL,
        "transaction_reference" varchar(50),
        "status" varchar(30) NOT NULL DEFAULT 'SUCCESS',
        "failure_reason" varchar(255),
        "created_at" TIMESTAMP NOT NULL DEFAULT now(),
        CONSTRAINT "PK_standing_order_executions" PRIMARY KEY ("id"),
        CONSTRAINT "FK_standing_order_executions_so" FOREIGN KEY ("standing_order_id") REFERENCES "standing_orders"("id") ON DELETE CASCADE
      );
    `);
    await queryRunner.query(`CREATE INDEX IF NOT EXISTS "IDX_standing_order_executions_so" ON "standing_order_executions" ("standing_order_id");`);

    // 4. Ensure customer dual-control maker-checker columns exist
    await queryRunner.query(`
      ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "is_authorized" boolean NOT NULL DEFAULT true;
      ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "created_by" varchar(100);
      ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "authorized_by" varchar(100);
      ALTER TABLE "customers" ADD COLUMN IF NOT EXISTS "authorized_at" TIMESTAMP;
    `);

    // 5. Ensure account_holds remaining_amount column exists
    await queryRunner.query(`
      ALTER TABLE "account_holds" ADD COLUMN IF NOT EXISTS "remaining_amount" numeric(18,4);
    `);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`DROP TABLE IF EXISTS "standing_order_executions";`);
    await queryRunner.query(`DROP TABLE IF EXISTS "idempotency_records";`);
    await queryRunner.query(`DROP TABLE IF EXISTS "account_parties";`);
    await queryRunner.query(`
      ALTER TABLE "customers" DROP COLUMN IF EXISTS "authorized_at";
      ALTER TABLE "customers" DROP COLUMN IF EXISTS "authorized_by";
      ALTER TABLE "customers" DROP COLUMN IF EXISTS "created_by";
      ALTER TABLE "customers" DROP COLUMN IF EXISTS "is_authorized";
    `);
    await queryRunner.query(`
      ALTER TABLE "account_holds" DROP COLUMN IF EXISTS "remaining_amount";
    `);
  }
}
