# Testing & Verification Strategy

## 1. Automated Test Suite Summary

The backend includes a comprehensive, automated test suite built with **Jest** and **Supertest**, verifying business rules, decimal math, API contracts, PostgreSQL integration, and concurrent transaction safety.

| Test Suite | File Location | Tests | Status | Scope |
|---|---|:---:|:---:|---|
| **Financial Math Unit Tests** | `test/unit/financial-math.spec.ts` | 6 | **PASS** | BigNumber decimal precision, division by zero guard, inequality checks |
| **Banking Rules Unit Tests** | `test/unit/accounts-rules.spec.ts` | 4 | **PASS** | Available balance calculation, hold releases, double-entry equality |
| **API Integration Suite** | `test/integration/banking-api.spec.ts` | 19 | **PASS** | Live PostgreSQL integration against `cbe_banking_test` (Auth, Customers, Accounts, Holds, Transfers, Payments, IPS, Wallet) |
| **Concurrency & Deadlock Suite** | `test/integration/concurrency.spec.ts` | 2 | **PASS** | Simultaneous concurrent transfers and duplicate idempotency racing |
| **Total** | | **31** | **ALL PASS** | **100% Passing** |

---

## 2. Concurrency Test Verification

The concurrency test (`test/integration/concurrency.spec.ts`) verifies two critical banking scenarios:

### Scenario A: High-Concurrency Debit Load
- **Execution:** Spawns 5 simultaneous asynchronous debit transfers against a single source account.
- **Verification:**
  - Row-level pessimistic locks (`SELECT ... FOR UPDATE`) ensure each transfer checks the balance and deducts funds sequentially.
  - Final account balance strictly equals initial balance minus total sum of transfers.
  - Zero double-spending or negative balances occur.

### Scenario B: Simultaneous Idempotent Collision
- **Execution:** Spawns 3 simultaneous requests with the exact same `idempotencyKey` submitted in parallel before any single request has committed.
- **Verification:**
  - One request acquires the write lock and inserts the `Transaction` row.
  - Concurrent runners hit PostgreSQL unique key constraint `23505`.
  - Caught gracefully: the system waits for the winner to commit and returns the identical transaction reference receipt without executing duplicate debits.

---

## 3. Running the Tests Locally

Ensure the test database is seeded:
```bash
DB_DATABASE=cbe_banking_test npm run seed
```

Run all unit and integration tests:
```bash
npm test -- --runInBand
```
