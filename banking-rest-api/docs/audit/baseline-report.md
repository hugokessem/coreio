# CBE CoreIO Banking REST API — Baseline Audit Report

**Date:** 2026-10-20  
**Auditor:** Principal Banking Systems Architect & Senior NestJS Engineer  
**Workspace:** `/Users/jeo/Development/eagle/cbe/coreio/banking-rest-api`

---

## 1. Executive Summary

A comprehensive architectural and code audit of the NestJS banking application was conducted against the 55 operations discovered in the legacy Golang CoreIO reference codebase.

The current implementation provides a working modular monolith with real PostgreSQL persistence, double-entry ledger bookkeeping, decimal-safe arithmetic via BigNumber.js, and pessimistic locking.

### Baseline Verification Status
- **TypeScript Compilation (`npm run build`):** **PASS** (Zero compiler errors in strict mode)
- **Unit & Integration Tests (`npm test`):** **PASS** (31 tests passing across 4 test suites)
- **Database Engine:** PostgreSQL 18 on Docker container `kora-postgres` (ports: 5432)
- **Active Databases:** `cbe_banking` (development), `cbe_banking_test` (automated tests)
- **Total Discovered CoreIO Operations:** 55
- **Operations Currently Implemented with Endpoints:** 42
- **Operations Missing or Requiring Gap Remediation:** 13

---

## 2. API Coverage Audit Matrix (55 CoreIO Operations)

| # | CoreIO Operation | Documented Endpoint | Implemented Endpoint | Baseline Status | Gap & Required Remediation |
|---|---|---|---|:---:|---|
| 1 | `AccountLookup` | `POST /api/v1/accounts/lookup` | `POST /api/v1/accounts/lookup`<br>`GET /api/v1/accounts/:acc` | **IMPLEMENTED_AND_VERIFIED** | None. Returns relational metadata and balances. |
| 2 | `AccountList` | `GET /api/v1/accounts` | `POST /api/v1/accounts/list` | **IMPLEMENTED_WITH_GAPS** | Add canonical `GET /api/v1/accounts?customerNumber=...` alias. |
| 3 | `AccountCreation` | `POST /api/v1/accounts` | `POST /api/v1/accounts/create` | **IMPLEMENTED_WITH_GAPS** | Add canonical `POST /api/v1/accounts` REST alias. |
| 4 | `AccountCreate` | `POST /api/v1/accounts/onboard` | *None* | **MISSING** | Implement composite onboarding workflow (customer create + KYC check + account opening). |
| 5 | `CustomerLookup` | `POST /api/v1/customers/lookup` | `POST /api/v1/customers/lookup` | **IMPLEMENTED_AND_VERIFIED** | Lookup by CIF, phone, account, or national ID. |
| 6 | `CustomerDetail` | `GET /api/v1/customers/:num` | `POST /api/v1/customers/detail`<br>`GET /api/v1/customers/:num` | **IMPLEMENTED_AND_VERIFIED** | Detailed customer profile with accounts and limits. |
| 7 | `CustomerFetch` | `POST /api/v1/customers/fetch` | `POST /api/v1/customers/fetch` | **IMPLEMENTED_AND_VERIFIED** | Paginated customer inquiry. |
| 8 | `CreateCustomer` | `POST /api/v1/customers` | `POST /api/v1/customers/create`<br>`POST /api/v1/customers` | **IMPLEMENTED_WITH_GAPS** | Add canonical `POST /api/v1/customers` REST route. |
| 9 | `CustomerDelete` | `DELETE /api/v1/customers/:num` | `POST /api/v1/customers/delete`<br>`DELETE /api/v1/customers/:num` | **IMPLEMENTED_AND_VERIFIED** | Guarded delete for accounts with zero history. |
| 10 | `ActionAutorization` | `POST /api/v1/customers/:num/authorize` | *None* | **MISSING** | Maker-checker approval workflow with dual control. |
| 11 | `VerifyAML` | `POST /api/v1/customers/:num/verify-aml` | `POST /api/v1/customers/aml-verify` | **IMPLEMENTED_WITH_GAPS** | Add route alias `POST /api/v1/customers/:num/verify-aml`. |
| 12 | `AMLLookup` | `POST /api/v1/accounts/:acc/aml-lookup` | *None* | **MISSING** | AML status inquiry by account number. |
| 13 | `PhoneLookup` | `POST /api/v1/customers/phone-lookup` | `POST /api/v1/customers/phone-lookup` | **IMPLEMENTED_AND_VERIFIED** | Mobile phone customer lookup. |
| 14 | `NameLookup` | `POST /api/v1/accounts/name-lookup` | `POST /api/v1/accounts/name-lookup` | **IMPLEMENTED_AND_VERIFIED** | Beneficiary name and eligibility inquiry. |
| 15 | `Fayda` | `POST /api/v1/identity/fayda/verify` | *None* | **MISSING** | Local identity verification check for Fayda National ID records. |
| 16 | `CheckIfUserExists` | `POST /api/v1/customers/check-existence` | *None* | **MISSING** | Pre-onboarding existence check by phone or NID. |
| 17 | `Eligibility` | `POST /api/v1/customers/eligibility` | `POST /api/v1/accounts/eligibility` | **IMPLEMENTED_WITH_GAPS** | Add customer-level eligibility route. |
| 18 | `FundTransferVerify` | `POST /api/v1/transfers/verify` | `POST /api/v1/transfers/verify` | **IMPLEMENTED_AND_VERIFIED** | Pre-transfer validation and fee calculation. |
| 19 | `FundTransfer` | `POST /api/v1/transfers` | `POST /api/v1/transfers/execute` | **IMPLEMENTED_WITH_GAPS** | Add canonical `POST /api/v1/transfers` alias. |
| 20 | `FundTransferCheck` | `GET /api/v1/transfers/:ftNumber` | *None* | **MISSING** | Query transfer receipt by `ftNumber` reference. |
| 21 | `StatusCheck` | `GET /api/v1/transfers/status/:txId` | `POST /api/v1/transfers/status` | **IMPLEMENTED_WITH_GAPS** | Add GET route alias. |
| 22 | `RevertFundTransfer` | `POST /api/v1/transfers/revert` | `POST /api/v1/transfers/reverse` | **IMPLEMENTED_WITH_GAPS** | Add `/revert` alias. |
| 23 | `CreateLockedAmount` | `POST /api/v1/accounts/holds` | `POST /api/v1/accounts/holds/create` | **IMPLEMENTED_WITH_GAPS** | Add canonical `POST /api/v1/accounts/holds` route. |
| 24 | `ListLockedAmount` | `GET /api/v1/accounts/:acc/holds` | `POST /api/v1/accounts/holds/list` | **IMPLEMENTED_WITH_GAPS** | Add `GET /api/v1/accounts/:acc/holds` route. |
| 25 | `ReleaseLockedAmount` | `POST /api/v1/accounts/holds/release` | `POST /api/v1/accounts/holds/release` | **IMPLEMENTED_AND_VERIFIED** | Unfreezes locked balance. |
| 26 | `LockedAmountFT` | `POST /api/v1/transfers/from-hold` | `POST /api/v1/transfers/from-hold` | **IMPLEMENTED_AND_VERIFIED** | Settles hold directly to destination account. |
| 27 | `CreateStandingOrder` | `POST /api/v1/standing-orders` | `POST /api/v1/standing-orders/create` | **IMPLEMENTED_WITH_GAPS** | Add canonical `POST /api/v1/standing-orders` route. |
| 28 | `ListStandingOrder` | `GET /api/v1/standing-orders` | `POST /api/v1/standing-orders/list` | **IMPLEMENTED_WITH_GAPS** | Add canonical GET with query parameters. |
| 29 | `ListStandingOrderHistory` | `GET /api/v1/standing-orders/history` | *None* | **MISSING** | Execution history for standing orders. |
| 30 | `UpdateStandingOrder` | `PUT /api/v1/standing-orders/:orderId` | *None* | **MISSING** | Update amount, frequency, or end date. |
| 31 | `CancleStandingOrder` | `DELETE /api/v1/standing-orders/:orderId` | `POST /api/v1/standing-orders/cancel` | **IMPLEMENTED_WITH_GAPS** | Add canonical DELETE route alias. |
| 32 | `MiniStatementByLimit` | `GET /api/v1/accounts/:acc/statement?limit=...` | `POST /api/v1/accounts/mini-statement` | **IMPLEMENTED_AND_VERIFIED** | Returns recent posted entries. |
| 33 | `MiniStatementByDateRange` | `GET /api/v1/accounts/:acc/statement?from=...` | `POST /api/v1/accounts/statement`<br>`GET ...` | **IMPLEMENTED_AND_VERIFIED** | Filtered statement with opening/closing balances. |
| 34 | `ExchangeRates` | `GET /api/v1/exchange-rates` | `POST /api/v1/rates/lookup`<br>`GET /api/v1/rates` | **IMPLEMENTED_WITH_GAPS** | Add `/api/v1/exchange-rates` alias. |
| 35 | `ServiceDetail` | `GET /api/v1/services` | `POST /api/v1/services/detail`<br>`GET /api/v1/services` | **IMPLEMENTED_AND_VERIFIED** | Service parameters and limits. |
| 36 | `ServiceLimit` | `GET /api/v1/services/limits` | *None* | **MISSING** | Query all channel/service limits. |
| 37 | `CustomerLimitFetchByCIF` | `GET /api/v1/customers/:num/limits` | `POST /api/v1/services/limits` | **IMPLEMENTED_WITH_GAPS** | Add `GET /api/v1/customers/:num/limits` route. |
| 38 | `CustomerLimitAmendByCIF` | `PUT /api/v1/customers/:num/limits` | *None* | **MISSING** | Amend per-customer service limit overrides. |
| 39 | `BillPayment` | `POST /api/v1/payments/bill` | `POST /api/v1/payments/bill` | **IMPLEMENTED_AND_VERIFIED** | Utility bill payment with ledger entries. |
| 40 | `BillPaymentVerify` | `POST /api/v1/payments/bill/verify` | *None* | **MISSING** | Pre-validation of bill reference and biller code. |
| 41 | `SplitPayment` | `POST /api/v1/payments/split` | `POST /api/v1/payments/split` | **IMPLEMENTED_AND_VERIFIED** | Atomic 1-to-many payment distribution. |
| 42 | `RTGS` | `POST /api/v1/payments/rtgs` | `POST /api/v1/payments/rtgs` | **IMPLEMENTED_AND_VERIFIED** | High-value settlement (explicitly simulated locally). |
| 43 | `CardRequest` | `POST /api/v1/cards/request` | `POST /api/v1/cards/request` | **IMPLEMENTED_AND_VERIFIED** | Debit/Mastercard issuance request. |
| 44 | `CardReplace` | `POST /api/v1/cards/replace` | `POST /api/v1/cards/replace` | **IMPLEMENTED_AND_VERIFIED** | Block compromised card and issue replacement. |
| 45 | `SuperAppSubscribe` | `POST /api/v1/superapp/subscribe` | `POST /api/v1/superapp/subscribe` | **IMPLEMENTED_AND_VERIFIED** | External user binding. |
| 46 | `SuperAppUnsubscribe` | `POST /api/v1/superapp/unsubscribe` | `POST /api/v1/superapp/unsubscribe` | **IMPLEMENTED_AND_VERIFIED** | External user unbinding. |
| 47 | `SuperAppStatusChange` | `PUT /api/v1/superapp/status` | *None* | **MISSING** | Change status of superapp user (ACTIVE/SUSPENDED/DEACTIVATED). |
| 48 | `IPS:AccountLookup` | `POST /api/v1/ips/account-lookup` | `POST /api/v1/ips/account-lookup` | **IMPLEMENTED_AND_VERIFIED** | Interbank ISO 20022 account inquiry. |
| 49 | `IPS:FundTransfer` | `POST /api/v1/ips/transfer` | `POST /api/v1/ips/transfer` | **IMPLEMENTED_AND_VERIFIED** | ISO 20022 interbank transfer. |
| 50 | `IPS:StatusCheck` | `POST /api/v1/ips/status-check` | `POST /api/v1/ips/status` | **IMPLEMENTED_WITH_GAPS** | Add `/status-check` alias. |
| 51 | `IPS:QRPayment` | `POST /api/v1/ips/qr-payment` | `POST /api/v1/ips/qr-payment` | **IMPLEMENTED_AND_VERIFIED** | Merchant QR payment. |
| 52 | `Wallet:CustomerAccountLookup`| `POST /api/v1/wallet/customer-lookup`| `POST /api/v1/wallet/customer-lookup`| **IMPLEMENTED_AND_VERIFIED** | Mobile wallet customer lookup. |
| 53 | `Wallet:CustomerFundTransfer` | `POST /api/v1/wallet/transfer` | `POST /api/v1/wallet/transfer-to-wallet`<br>`POST /api/v1/wallet/transfer-from-wallet` | **IMPLEMENTED_WITH_GAPS** | Add composite `POST /api/v1/wallet/transfer`. |
| 54 | `Wallet:AgentAccountLookup` | `POST /api/v1/wallet/agent-lookup` | *None* | **MISSING** | Wallet agent credential and float lookup. |
| 55 | `Wallet:AgentFundTransfer` | `POST /api/v1/wallet/agent-transfer` | *None* | **MISSING** | Agent cash-in/cash-out float transfer. |

---

## 3. High-Priority Correctness & Security Findings

1. **Security Vulnerability in `CombinedAuthGuard`:**
   - Any client sending `x-api-key` is automatically assigned `['API_CLIENT', 'TELLER', 'ADMIN']` roles.
   - This elevates API key consumers to full administrative and teller privileges.
   - Fix: Assign `['API_CLIENT']` role only; restrict administrative endpoints to verified admin users or specific dual-control scopes.
2. **Hardcoded Secret Fallback:**
   - Default secret and API key fallbacks in `.env` must be hardened for non-development environments.
3. **Idempotency Fingerprinting:**
   - Idempotency is currently tracked as a column on `transactions`. If an idempotent key is submitted with a *different* payload, it should reject with `IDEMPOTENCY_CONFLICT` (409) rather than returning an existing transaction that had different parameters.
   - Solution: Dedicated `IdempotencyRecord` entity storing request payload SHA-256 hash.
4. **Standing Order Background Execution:**
   - Standing orders are defined in the database, but need an autonomous execution worker using PostgreSQL row locking (`SELECT FOR UPDATE SKIP LOCKED`) to execute due orders with double-entry ledger postings.
5. **Ledger Reconciliation Mechanism:**
   - Need an authoritative reconciliation service to verify that every account's `ledger_balance` matches the sum of its posted `LedgerEntry` records.
