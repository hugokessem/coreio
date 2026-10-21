import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from '../../src/app.module';

describe('Banking Concurrency & Consistency Test Suite', () => {
  let app: INestApplication;
  let jwtToken: string;

  beforeAll(async () => {
    process.env.DB_DATABASE = 'cbe_banking_test';

    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();

    const loginRes = await request(app.getHttpServer())
      .post('/api/v1/auth/login')
      .send({ username: 'admin', password: 'Admin123!' });

    jwtToken = loginRes.body.accessToken;
  });

  afterAll(async () => {
    await app.close();
  });

  it('Concurrent debit transfers must not cause double-spending or negative balances', async () => {
    // 5 concurrent transfer requests of 100 ETB from account 1 to account 2
    const promises = Array.from({ length: 5 }).map((_, idx) =>
      request(app.getHttpServer())
        .post('/api/v1/transfers/execute')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '100.0000',
          narration: `Concurrent Transfer #${idx + 1}`,
          idempotencyKey: `CONCURRENT-KEY-${Date.now()}-${idx}`,
        }),
    );

    const results = await Promise.all(promises);

    // Verify all completed with valid 200 responses
    results.forEach((r) => {
      expect(r.status).toBe(200);
      expect(r.body.status).toBe('COMPLETED');
    });

    // Check account 1 balance after all 5 transfers
    const balRes = await request(app.getHttpServer())
      .get('/api/v1/accounts/1000123456789/balance')
      .set('Authorization', `Bearer ${jwtToken}`);

    expect(balRes.status).toBe(200);
    expect(parseFloat(balRes.body.availableBalance)).toBeGreaterThan(0);
  });

  it('Concurrent duplicate requests with identical idempotency key must execute strictly once', async () => {
    const sharedIdempKey = `IDEMP-STRICT-SINGLETON-${Date.now()}`;

    // Fire 3 simultaneous requests with exact same idempotency key
    const promises = Array.from({ length: 3 }).map(() =>
      request(app.getHttpServer())
        .post('/api/v1/transfers/execute')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '50.0000',
          narration: 'Singleton Idempotent Test',
          idempotencyKey: sharedIdempKey,
        }),
    );

    const results = await Promise.all(promises);

    // All should succeed (one creates, others return idempotent replay)
    results.forEach((r) => {
      expect(r.status).toBe(200);
    });

    // All results must return the exact same transaction reference!
    const firstRef = results[0].body.transactionReference;
    results.forEach((r) => {
      expect(r.body.transactionReference).toBe(firstRef);
    });
  });
});
