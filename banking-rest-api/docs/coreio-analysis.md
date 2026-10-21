# CoreIO Architecture & Reverse Engineering Analysis

## 1. Executive Summary

**CoreIO** is a Go-based integration library originally developed to bridge CBE (Commercial Bank of Ethiopia) SuperApp backend services with legacy enterprise banking systems:
- **Temenos T24 Core Banking System** (via SOAP 1.1 / XML Web Services over HTTP)
- **Huawei Mobile Money / CBE Birr Wallet** (via SOAP / SyncCPS interface)
- **National Interbank Payment System (IPS / EthSwitch)** (via ISO 20022 XML messages `acmt.023`, `pacs.008`, `pacs.028`)
- **Ethiopian National ID System (Fayda)** (for biometric & demographic national ID verification)
- **MasterCard Support System & IBM Filenet/CMIS** (document upload/download)

In the reference Go implementation, CoreIO was purely a translation shim: it received Go struct parameters, constructed SOAP XML envelopes containing Temenos/Huawei/ISO 20022 payloads, executed HTTP POST requests against upstream gateway endpoints, and parsed the XML response back into Go structs.

## 2. Architectural Deficiencies of CoreIO

1. **Tight Coupling to T24 XML Envelopes:** The application logic is intermingled with raw XML string templates and brittle XML unmarshalling structs.
2. **Lack of Relational State:** CoreIO has zero internal persistence. It treats upstream core banking as a black box. If T24 is slow, unavailable, or down, operations fail immediately.
3. **No Financial Transaction Consistency:** Because CoreIO is stateless, operations like `FundTransfer` perform an external remote call without local ACID guarantees, double-entry ledgers, or deterministic rollbacks.
4. **Security Vulnerabilities:**
   - Credentials (`Username`, `Password`) were embedded directly in SOAP headers in cleartext or passed inside HTTP bodies.
   - PII was reflected in XML log outputs.
   - No row-level authorization was enforced locally.
5. **No Native Idempotency:** Outside of passing an optional client reference string to T24, duplicate requests submitted concurrently risked double-debiting.

## 3. Re-engineering Strategy for NestJS + PostgreSQL

The new NestJS application replaces the XML/SOAP/T24 transport layer with an **authoritative, independent relational banking engine**:

| Concern | Legacy CoreIO (Go) | New Banking REST API (NestJS + PostgreSQL) |
|---|---|---|
| **Architecture** | Stateless XML integration library | Domain-Driven Modular Monolith |
| **Persistence** | None (T24 Core Banking upstream) | Normalized PostgreSQL Relational Engine |
| **Communication** | SOAP 1.1 / XML over HTTP | REST / JSON APIs with OpenAPI/Swagger |
| **Financial Ledger** | Delegated to T24 | Authoritative Double-Entry General Ledger |
| **Hold Management** | XML requests (`ACLOCKEDEVENTS`) | Native `account_holds` table with real-time balance calculations |
| **Concurrency Control** | External T24 locking | PostgreSQL `SELECT ... FOR UPDATE` row-level pessimistic locks |
| **Idempotency** | Passed through to T24 | First-class `idempotency_keys` table with atomic deduplication |
| **Security** | Hardcoded credentials | JWT Bearer Authentication, Role-Based Access Control, Auditing |
| **Data Types** | Strings in XML | Strictly validated TypeScript DTOs, PostgreSQL `NUMERIC(18,4)` |

## 4. Discovered Subsystems & Classification

### 4.1 Fully Understood Core Banking Operations
- **Account Lookup & Inquiry:** Look up account metadata, branch, currency, balance, account type, customer ownership.
- **Account Creation:** Provision new checking/savings account under customer and branch.
- **Customer Lookup & Details:** Inquire customer demographic, contact, and KYC info.
- **Customer Creation & Onboarding:** Create customer profile with Fayda NID, AML verification, contact details.
- **Fund Transfer (Internal & Interbank):** Concurrency-safe atomic transfer between accounts with debit/credit ledger postings, fees, and disaster reserve fund deductions.
- **Fund Transfer Verification (Simulation):** Dry-run pre-validation of balance, limits, fees, and account statuses.
- **Account Holds / Locked Amounts:** Create hold, list active holds, release hold, execute transfer from held amount.
- **Standing Orders:** Schedule recurring transfers, list active orders, history, and cancel order.
- **Mini Statements & Transaction History:** Query chronological ledger statements filtered by date range or transaction limit.
- **Exchange Rates:** Query official buying, selling, and mid rates for major trading pairs (USD, EUR, GBP against ETB).
- **Service Limits:** Inquire and enforce daily amount and frequency limits per channel (APP, USSD).
- **Fayda Verification & User Existence Check:** Verify national identification number against citizen registry.

### 4.2 External Gateways Adapted to Core Banking Domain
- **IPS (EthSwitch Interoperable Payment System):** Re-engineered as inter-bank routing with virtual clearing accounts and BICs.
- **CBE Birr Wallet:** Re-engineered as phone-linked digital wallet accounts with internal wallet clearing ledger accounts.
- **SuperApp Subscription:** Modeled as user-to-customer channel subscription lifecycle.

### 4.3 Unsupported / External Ancillary
- CMIS document storage (handled as standard file/attachment metadata).
- External survey analytics HTTP triggers (refactored as non-blocking event hooks).
