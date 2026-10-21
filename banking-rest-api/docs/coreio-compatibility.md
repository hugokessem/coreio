# CoreIO Compatibility & Migration Analysis

## 1. Compatibility Overview

The **NestJS Banking REST API** provides behavioral equivalence for the core banking capabilities previously mediated through the Golang `CoreIO` XML gateway.

While legacy CoreIO acted as a protocol converter (translating JSON HTTP into T24 OFS XML requests and parsing XML responses), the new application replaces the backend entirely with a native, relational PostgreSQL data store.

### Key Architectural Differences

| Dimension | Legacy Go CoreIO | New NestJS Banking Backend |
|---|---|---|
| **Data Source** | Remote Temenos T24 via XML/TCP | Local PostgreSQL 18 Relational DB |
| **Transport** | Custom XML over HTTP/SOAP | Pure REST / JSON over HTTPS |
| **Response Reliability** | Dependent on external T24 gateway latency | Sub-millisecond indexed relational queries |
| **Ledger Model** | Relies on opaque upstream core posting | Authoritative, immutable double-entry ledger |
| **Concurrency Control** | External (T24 lock tables) | Native `SELECT ... FOR UPDATE` row locks |
| **Hold Handling** | Opaque `ACLOCKEDEVENTS` XML records | Relational `account_holds` with state machines |
| **Idempotency** | Partial or upstream-managed | Guaranteed at database level with unique keys |
| **Security & Auth** | Often unprotected internal proxies | JWT Bearer tokens + RBAC + API Keys |

---

## 2. Intentional Modernization & Security Hardening

1. **Strict Authentication & Least Privilege:**
   - In Go CoreIO, multiple routes lacked bearer token validation. The new NestJS backend secures all endpoints behind `CombinedAuthGuard`, requiring either a valid JWT or a configured partner `x-api-key`.
2. **Elimination of IEEE-754 Floating-Point Arithmetic:**
   - Legacy Go code parsed float64 values from XML strings in several handlers. The new backend uses `NUMERIC(18,4)` and `BigNumber.js`, ensuring zero financial calculation loss.
3. **Pessimistic Concurrency Serialization:**
   - Legacy CoreIO relied on T24 returning an error code when an account was locked. The new service proactively enforces lock order sorting to eliminate database deadlocks under high concurrency.
4. **Structured JSON Error Contracts:**
   - Legacy CoreIO frequently returned raw XML fault strings or unstructured error objects. The NestJS backend provides standardized error responses including specific domain `errorCode` values (`INSUFFICIENT_FUNDS`, `ACCOUNT_RESTRICTED`, `IDEMPOTENCY_CONFLICT`).
