# Domain Limitations & Known Operational Boundaries

## 1. External Clearing Interfaces & Simulation Boundaries

In legacy CoreIO, some operations contacted external national or partner switches through T24. In this standalone NestJS implementation, external clearing endpoints are simulated as follows:

1. **National Bank RTGS Switch:**
   - The application validates account balances, reserves RTGS tariff fees, and writes double-entry ledger entries to the internal clearing account. Real SWIFT MT103 / ISO 20022 pacs.008 network transmission to the National Bank of Ethiopia is modeled as an immediate local clearing confirmation.
2. **EthSwitch IPS Clearing:**
   - ISO 20022 interbank message formats and validation rules are strictly enforced. Interbank clearing between participating banks is settled internally against domestic accounts.
3. **Card Personalization Hardware:**
   - Card issuance orders and replacement requests transition into `PRINTING` and `DISPATCHED` states, simulating the physical card bureau printing workflow.

---

## 2. Legacy T24 Features Omitted from Scope

- **End of Day (EOD) COB Processing:**
   - Legacy T24 runs nightly Close of Business (COB) batch jobs for accruals and statement generation. In this application, balances and statements are computed dynamically in real-time from ledger entries.
- **Cheque Clearing & MICR Encoding:**
   - Physical cheque imaging and clearing houses are out of scope as they were not implemented in Go CoreIO.
- **Complex Multi-Tier Loan Amortization Schedules:**
   - Loan disbursement and interest amortization schedules outside of CoreIO account type categories (`6701`, `6501`, `6001`) are omitted.
