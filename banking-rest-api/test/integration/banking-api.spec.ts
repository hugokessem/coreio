import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication, ValidationPipe } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from '../../src/app.module';
import { AllExceptionsFilter } from '../../src/common/filters/all-exceptions.filter';

describe('CBE CoreIO Banking REST API Integration Suite', () => {
  let app: INestApplication;
  let jwtToken: string;

  beforeAll(async () => {
    // Override DB to use cbe_banking_test
    process.env.DB_DATABASE = 'cbe_banking_test';

    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    app.useGlobalFilters(new AllExceptionsFilter());
    app.useGlobalPipes(
      new ValidationPipe({
        whitelist: true,
        forbidNonWhitelisted: true,
        transform: true,
      }),
    );
    await app.init();

    // Authenticate to get JWT token
    const loginRes = await request(app.getHttpServer())
      .post('/api/v1/auth/login')
      .send({ username: 'admin', password: 'Admin123!' });

    expect(loginRes.status).toBe(200);
    expect(loginRes.body.accessToken).toBeDefined();
    jwtToken = loginRes.body.accessToken;
  });

  afterAll(async () => {
    await app.close();
  });

  describe('1. Authentication & Security', () => {
    it('POST /api/v1/auth/login should reject invalid credentials', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/auth/login')
        .send({ username: 'admin', password: 'WrongPassword' });

      expect(res.status).toBe(401);
    });

    it('should reject unauthenticated requests to protected endpoints', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/lookup')
        .send({ accountNumber: '1000123456789' });

      expect(res.status).toBe(401);
    });

    it('should allow requests with valid x-api-key header', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/lookup')
        .set('x-api-key', 'cbe-banking-api-key-default-2026')
        .send({ accountNumber: '1000123456789' });

      expect(res.status).toBe(200);
      expect(res.body.accountNumber).toBe('1000123456789');
    });
  });

  describe('2. Customers Domain', () => {
    it('POST /api/v1/customers/lookup by customer number', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/customers/lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ customerNumber: '10000001' });

      expect(res.status).toBe(200);
      expect(res.body.customerNumber).toBe('10000001');
      expect(res.body.fullName).toContain('Abebe');
    });

    it('POST /api/v1/customers/phone-lookup', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/customers/phone-lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ phoneNumber: '+251911223344' });

      expect(res.status).toBe(200);
      expect(res.body.customerNumber).toBe('10000001');
    });

    it('POST /api/v1/customers/aml-verify on clear customer', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/customers/aml-verify')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ customerNumber: '10000001' });

      expect(res.status).toBe(200);
      expect(res.body.isCleared).toBe(true);
      expect(res.body.amlStatus).toBe('CLEAR');
    });
  });

  describe('3. Accounts Domain & Balance Inquiry', () => {
    it('POST /api/v1/accounts/lookup should return real persisted relational details', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ accountNumber: '1000123456789' });

      expect(res.status).toBe(200);
      expect(res.body.accountNumber).toBe('1000123456789');
      expect(res.body.currency).toBe('ETB');
      expect(res.body.branchCode).toBe('001');
      expect(res.body.ledgerBalance).toBeDefined();
      expect(res.body.availableBalance).toBeDefined();
    });

    it('POST /api/v1/accounts/name-lookup for beneficiary title verification', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/name-lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ accountNumber: '1000987654321' });

      expect(res.status).toBe(200);
      expect(res.body.accountName).toContain('Almaz Tadesse');
      expect(res.body.isAllowedToTransact).toBe(true);
    });

    it('POST /api/v1/accounts/eligibility should check available balance and restrictions', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/eligibility')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          accountNumber: '1000123456789',
          amount: '500.0000',
          transactionType: 'DEBIT',
        });

      expect(res.status).toBe(200);
      expect(res.body.eligible).toBe(true);
    });

    it('GET /api/v1/accounts/:accountNumber/statement returns ledger entries', async () => {
      const res = await request(app.getHttpServer())
        .get('/api/v1/accounts/1000123456789/statement')
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.accountNumber).toBe('1000123456789');
      expect(Array.isArray(res.body.transactions)).toBe(true);
      expect(res.body.transactions.length).toBeGreaterThan(0);
    });
  });

  describe('4. Holds Management (ACLOCKEDEVENTS)', () => {
    let holdRef: string;

    it('POST /api/v1/accounts/holds/create should place hold and reduce available balance', async () => {
      // First check initial balance
      const initialBal = await request(app.getHttpServer())
        .get('/api/v1/accounts/1000987654321/balance')
        .set('Authorization', `Bearer ${jwtToken}`);

      const prevAvail = parseFloat(initialBal.body.availableBalance);

      const holdRes = await request(app.getHttpServer())
        .post('/api/v1/accounts/holds/create')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          accountNumber: '1000987654321',
          amount: '1000.0000',
          reason: 'Test Administrative Hold',
          lockEvent: 'PLEDGE',
        });

      expect(holdRes.status).toBe(201);
      expect(holdRes.body.holdReference).toBeDefined();
      holdRef = holdRes.body.holdReference;
      expect(parseFloat(holdRes.body.newAvailableBalance)).toBe(prevAvail - 1000);
    });

    it('POST /api/v1/accounts/holds/release should release hold and restore available balance', async () => {
      const releaseRes = await request(app.getHttpServer())
        .post('/api/v1/accounts/holds/release')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          holdReference: holdRef,
          reason: 'Test Release Completed',
        });

      expect(releaseRes.status).toBe(200);
      expect(releaseRes.body.status).toBe('RELEASED');
    });
  });

  describe('5. Transfers & Financial Consistency', () => {
    it('POST /api/v1/transfers/verify dry run verification', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers/verify')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '100.0000',
        });

      expect(res.status).toBe(200);
      expect(res.body.canExecute).toBe(true);
      expect(res.body.destinationAccountName).toContain('Almaz');
    });

    it('POST /api/v1/transfers/execute posts balanced double-entry ledger entries', async () => {
      const idempKey = `IDEMP-TEST-${Date.now()}`;

      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers/execute')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '200.0000',
          narration: 'Integration Test Fund Transfer',
          idempotencyKey: idempKey,
        });

      expect(res.status).toBe(200);
      expect(res.body.status).toBe('COMPLETED');
      expect(res.body.transactionReference).toBeDefined();

      // Test Idempotency Replay
      const replayRes = await request(app.getHttpServer())
        .post('/api/v1/transfers/execute')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '200.0000',
          narration: 'Integration Test Fund Transfer',
          idempotencyKey: idempKey,
        });

      expect(replayRes.status).toBe(200);
      expect(replayRes.body.transactionReference).toBe(res.body.transactionReference);
      expect(replayRes.body.message).toContain('Idempotent replay');
    });

    it('should reject transfer when available balance is insufficient', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers/execute')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '999999999.0000',
          narration: 'Astronomical transfer should fail',
        });

      expect(res.status).toBe(422);
      expect(res.body.errorCode).toBe('INSUFFICIENT_FUNDS');
    });
  });

  describe('6. Payments (Bill, Split, RTGS)', () => {
    it('POST /api/v1/payments/bill should process utility bill payment', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/payments/bill')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          billerCode: 'BILL_ETHIO',
          billReference: '0911223344',
          amount: '50.0000',
          narration: 'Ethio Telecom Phone Bill',
        });

      expect(res.status).toBe(200);
      expect(res.body.status).toBe('COMPLETED');
      expect(res.body.billerCode).toBe('BILL_ETHIO');
    });

    it('POST /api/v1/payments/split should distribute funds to multiple destinations', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/payments/split')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '2000123456789',
          splits: [
            { destinationAccountNumber: '1000123456789', amount: '100.0000', narration: 'Vendor A' },
            { destinationAccountNumber: '1000987654321', amount: '150.0000', narration: 'Vendor B' },
          ],
        });

      expect(res.status).toBe(200);
      expect(res.body.splitCount).toBe(2);
      expect(res.body.totalDistributed).toBe('250.0000');
    });
  });

  describe('7. IPS & CBE Birr Wallet', () => {
    it('POST /api/v1/ips/account-lookup returns valid interbank account title', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/ips/account-lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ accountNumber: '1000123456789' });

      expect(res.status).toBe(200);
      expect(res.body.accountName).toContain('Abebe');
      expect(res.body.bankCode).toBe('CBETETAA');
    });

    it('POST /api/v1/wallet/customer-lookup verifies CBE Birr customer', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/wallet/customer-lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ phoneNumber: '+251911223344' });

      expect(res.status).toBe(200);
      expect(res.body.phoneNumber).toBe('+251911223344');
      expect(res.body.customerName).toContain('Abebe');
    });
  });

  describe('8. Canonical Routes & Compatibility Aliases', () => {
    it('POST /api/v1/transfers should execute transfer via canonical route', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '150.0000',
          narration: 'Canonical route test',
        });

      expect(res.status).toBe(200);
      expect(res.body.transactionReference).toBeDefined();
      expect(res.body.status).toBe('COMPLETED');
    });

    it('GET /api/v1/transfers/:ref queries transfer details by reference (Op 25)', async () => {
      // First create a transfer
      const txRes = await request(app.getHttpServer())
        .post('/api/v1/transfers')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '50.0000',
          narration: 'Reference check test',
        });

      const ref = txRes.body.transactionReference;

      const res = await request(app.getHttpServer())
        .get(`/api/v1/transfers/${ref}`)
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.transactionReference).toBe(ref);
      expect(res.body.status).toBe('COMPLETED');
    });

    it('GET /api/v1/transfers/status/:ref queries status (Op 26)', async () => {
      const txRes = await request(app.getHttpServer())
        .post('/api/v1/transfers')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '60.0000',
          narration: 'Status check test',
        });

      const ref = txRes.body.transactionReference;

      const res = await request(app.getHttpServer())
        .get(`/api/v1/transfers/status/${ref}`)
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.transactionReference).toBe(ref);
    });

    it('POST /api/v1/transfers/revert should reverse transfer (Op 27)', async () => {
      const txRes = await request(app.getHttpServer())
        .post('/api/v1/transfers')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '75.0000',
          narration: 'Transfer to reverse',
        });

      const ref = txRes.body.transactionReference;

      const revRes = await request(app.getHttpServer())
        .post('/api/v1/transfers/revert')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          transactionReference: ref,
          reason: 'Erroneous transfer reversal test',
        });

      expect(revRes.status).toBe(200);
      expect(revRes.body.originalReference).toBe(ref);
      expect(revRes.body.reversalReference).toContain('REV-');
    });

    it('POST /api/v1/accounts/mini-statement-by-limit produces statement (Op 28)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/mini-statement-by-limit')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ accountNumber: '1000123456789', limit: 5 });

      expect(res.status).toBe(200);
      expect(res.body.accountNumber).toBe('1000123456789');
      expect(Array.isArray(res.body.entries)).toBe(true);
    });

    it('POST /api/v1/accounts/mini-statement-by-date-range produces statement (Op 29)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/mini-statement-by-date-range')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          accountNumber: '1000123456789',
          startDate: '2026-01-01',
          endDate: '2026-12-31',
        });

      expect(res.status).toBe(200);
      expect(res.body.accountNumber).toBe('1000123456789');
      expect(res.body.openingBalance).toBeDefined();
      expect(res.body.closingBalance).toBeDefined();
    });
  });

  describe('9. Customer Identity, Existence, and Maker-Checker', () => {
    it('POST /api/v1/customers/check-existence verifies user presence (Op 16)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/customers/check-existence')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ phoneNumber: '+251911223344' });

      expect(res.status).toBe(200);
      expect(res.body.exists).toBe(true);
      expect(res.body.customerNumber).toBe('10000001');
    });

    it('POST /api/v1/customers/eligibility evaluates customer onboarding eligibility (Op 17)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/customers/eligibility')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          phoneNumber: '+251911999888',
          nationalId: 'NAT-NEW-ELIG-01',
          dateOfBirth: '1995-05-15',
        });

      expect(res.status).toBe(200);
      expect(res.body.isEligible).toBe(true);
    });

    it('POST /api/v1/accounts/:acc/aml-lookup checks AML risk profile (Op 12)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/accounts/1000123456789/aml-lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({});

      expect(res.status).toBe(200);
      expect(res.body.amlStatus).toBe('CLEAR');
      expect(res.body.passedScreening).toBe(true);
    });

    it('POST /api/v1/identity/fayda/verify validates national ID (Op 15)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/identity/fayda/verify')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          faydaId: 'FIN-ETH-99281928',
          fullName: 'Abebe Bikila Demisse',
          dateOfBirth: '1985-04-12',
          phoneNumber: '+251911223344',
        });

      expect(res.status).toBe(200);
      expect(res.body.isVerified).toBe(true);
      expect(res.body.simulation).toBe(true);
    });

    it('POST /api/v1/customers/:cif/authorize enforces dual control maker-checker (Op 10)', async () => {
      // First create a new customer
      const newCust = await request(app.getHttpServer())
        .post('/api/v1/customers/create')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          customerType: 'INDIVIDUAL',
          fullName: 'Dual Control Test User',
          phoneNumber: '+251977665544',
          email: 'dual@test.cbe.et',
          nationalId: 'NAT-DUAL-001',
          dateOfBirth: '1990-01-01',
          addressLine1: 'Bole Medhanialem',
          city: 'Addis Ababa',
        });

      expect(newCust.status).toBe(201);
      const cif = newCust.body.customerNumber;

      // Creator self-approval must fail with 400 Bad Request
      const selfAuth = await request(app.getHttpServer())
        .post(`/api/v1/customers/${cif}/authorize`)
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ authorizerUsername: 'admin' });

      expect(selfAuth.status).toBe(400);
      expect(selfAuth.body.message).toContain('maker-checker');

      // Approval by different authorized officer succeeds
      const diffAuth = await request(app.getHttpServer())
        .post(`/api/v1/customers/${cif}/authorize`)
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ authorizerUsername: 'senior_supervisor' });

      expect(diffAuth.status).toBe(200);
      expect(diffAuth.body.isAuthorized).toBe(true);
      expect(diffAuth.body.authorizedBy).toBe('senior_supervisor');
    });
  });

  describe('10. Standing Orders Full Lifecycle & Execution Engine', () => {
    let orderRef: string;

    it('POST /api/v1/standing-orders creates instruction (Op 30)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/standing-orders')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          beneficiaryAccount: '1000987654321',
          beneficiaryName: 'Almaz Tadesse',
          amount: '300.0000',
          frequency: 'MONTHLY',
          startDate: '2026-10-01', // past/due date for testing execution
        });

      expect(res.status).toBe(201);
      expect(res.body.orderReference).toBeDefined();
      orderRef = res.body.orderReference;
    });

    it('GET /api/v1/standing-orders?accountNumber=... lists orders (Op 31)', async () => {
      const res = await request(app.getHttpServer())
        .get('/api/v1/standing-orders?accountNumber=1000123456789')
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.orders.length).toBeGreaterThan(0);
    });

    it('PUT /api/v1/standing-orders/:ref updates standing order parameters (Op 33)', async () => {
      const res = await request(app.getHttpServer())
        .put(`/api/v1/standing-orders/${orderRef}`)
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ amount: '350.0000' });

      expect(res.status).toBe(200);
      expect(res.body.amount).toBe('350.0000');
    });

    it('POST /api/v1/standing-orders/execute-due runs durable PostgreSQL locked scheduler', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/standing-orders/execute-due')
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.processed).toBeGreaterThan(0);
      expect(res.body.succeeded).toBeGreaterThan(0);
    });

    it('GET /api/v1/standing-orders/:ref/history verifies persistent execution audit (Op 32)', async () => {
      const res = await request(app.getHttpServer())
        .get(`/api/v1/standing-orders/${orderRef}/history`)
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.totalExecutions).toBeGreaterThan(0);
      expect(res.body.executions[0].status).toBe('SUCCESS');
    });

    it('DELETE /api/v1/standing-orders/:ref cancels order (Op 34)', async () => {
      const res = await request(app.getHttpServer())
        .delete(`/api/v1/standing-orders/${orderRef}`)
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.status).toBe('CANCELLED');
    });
  });

  describe('11. Idempotency Invariants & 409 Conflict Handling', () => {
    const key = `IDEMP-INTEGRATION-${Date.now()}`;

    it('1st request executes transfer and caches response', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '120.0000',
          narration: 'Idempotency invariant test',
          idempotencyKey: key,
        });

      expect(res.status).toBe(200);
      expect(res.body.transactionReference).toBeDefined();
    });

    it('2nd identical request returns exact cached replay', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '120.0000',
          narration: 'Idempotency invariant test',
          idempotencyKey: key,
        });

      expect(res.status).toBe(200);
      expect(res.body.transactionReference).toBeDefined();
    });

    it('3rd request with SAME key but ALTERED payload throws 409 Conflict', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          destinationAccountNumber: '1000987654321',
          amount: '999.0000', // DIFFERENT AMOUNT!
          narration: 'Idempotency conflict test',
          idempotencyKey: key,
        });

      expect(res.status).toBe(409);
      expect(res.body.errorCode).toBe('IDEMPOTENCY_CONFLICT');
    });
  });

  describe('12. Least-Privilege API Key & Ledger Reconciliation', () => {
    it('API key client is rejected from admin reconciliation endpoint with 403 Forbidden', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers/reconcile')
        .set('x-api-key', 'cbe-banking-api-key-default-2026');

      expect(res.status).toBe(403);
    });

    it('JWT Admin performs system-wide ledger reconciliation and confirms balanced books', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/transfers/reconcile')
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.systemBalanced).toBe(true);
      expect(res.body.imbalancedAccounts).toBe(0);
      expect(res.body.accountsAudited).toBeGreaterThan(0);
    });
  });

  describe('13. Payments, Services, Rates, and SuperApp', () => {
    it('POST /api/v1/payments/bill/verify validates bill payment tariff (Op 41)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/payments/bill/verify')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          sourceAccountNumber: '1000123456789',
          billerCode: 'BILL_ETHIO',
          billReference: '0911223344',
          amount: '200.0000',
        });

      expect(res.status).toBe(200);
      expect(res.body.canExecute).toBe(true);
      expect(res.body.status).toBe('VERIFIED');
    });

    it('GET /api/v1/services/limits returns service tariff limits (Op 37)', async () => {
      const res = await request(app.getHttpServer())
        .get('/api/v1/services/limits')
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.totalServices).toBeGreaterThan(0);
    });

    it('PUT /api/v1/services/limits/amend updates customer limits (Op 39)', async () => {
      const res = await request(app.getHttpServer())
        .put('/api/v1/services/limits/amend')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          customerNumber: '10000001',
          serviceCode: 'FT_INT',
          dailyLimit: '800000.0000',
          perTransactionLimit: '200000.0000',
        });

      expect(res.status).toBe(200);
      expect(res.body.dailyLimit).toBe('800000.0000');
    });

    it('GET /api/v1/rates/exchange-rates returns FX rates (Op 35)', async () => {
      const res = await request(app.getHttpServer())
        .get('/api/v1/rates/exchange-rates')
        .set('Authorization', `Bearer ${jwtToken}`);

      expect(res.status).toBe(200);
      expect(res.body.rates.length).toBeGreaterThan(0);
    });

    it('PUT /api/v1/superapp/status changes subscription status (Op 49)', async () => {
      // First subscribe external partner user
      const subRes = await request(app.getHttpServer())
        .post('/api/v1/superapp/subscribe')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          externalUserId: 'EXT-USER-TELCO-9921',
          customerNumber: '10000001',
          partnerCode: 'TELEBIRR_MINIAPP',
        });
      expect(subRes.status).toBe(200);

      const res = await request(app.getHttpServer())
        .put('/api/v1/superapp/status')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({
          externalUserId: 'EXT-USER-TELCO-9921',
          status: 'SUSPENDED',
          reason: 'Customer security request',
        });

      expect(res.status).toBe(200);
      expect(res.body.status).toBe('SUSPENDED');
    });

    it('POST /api/v1/wallet/agent-lookup checks agent profile (Op 56)', async () => {
      const res = await request(app.getHttpServer())
        .post('/api/v1/wallet/agent-lookup')
        .set('Authorization', `Bearer ${jwtToken}`)
        .send({ agentIdentifier: '+251911223344' });

      expect(res.status).toBe(200);
      expect(res.body.agentType).toBe('SUPER_AGENT');
      expect(res.body.simulation).toBe(true);
    });
  });

  describe('14. Observability & Health Probes', () => {
    it('GET /health returns liveness status UP', async () => {
      const res = await request(app.getHttpServer()).get('/health');
      expect(res.status).toBe(200);
      expect(res.body.status).toBe('UP');
      expect(res.body.uptimeSeconds).toBeGreaterThanOrEqual(0);
    });

    it('GET /health/ready returns PostgreSQL readiness status READY', async () => {
      const res = await request(app.getHttpServer()).get('/health/ready');
      expect(res.status).toBe(200);
      expect(res.body.status).toBe('READY');
      expect(res.body.database.status).toBe('UP');
      expect(res.body.database.latencyMs).toBeGreaterThanOrEqual(0);
    });
  });
});
