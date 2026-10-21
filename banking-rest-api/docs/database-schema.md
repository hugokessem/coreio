# PostgreSQL Relational Database Schema & Domain Model

## 1. Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    CUSTOMERS ||--o{ ACCOUNTS : owns
    CUSTOMERS ||--o{ CUSTOMER_LIMITS : has
    CUSTOMERS ||--o{ CARDS : holds
    CUSTOMERS ||--o{ SUPER_APP_USERS : links
    BRANCHES ||--o{ ACCOUNTS : domiciles
    CURRENCIES ||--o{ ACCOUNTS : denominate
    CURRENCIES ||--o{ EXCHANGE_RATES : rates
    ACCOUNT_TYPES ||--o{ ACCOUNTS : classifies
    ACCOUNTS ||--|| ACCOUNT_BALANCES : tracks
    ACCOUNTS ||--o{ ACCOUNT_HOLDS : locks
    ACCOUNTS ||--o{ TRANSACTIONS : participates
    ACCOUNTS ||--o{ LEDGER_ENTRIES : posted_to
    ACCOUNTS ||--o{ STANDING_ORDERS : instructs
    ACCOUNTS ||--o{ CARDS : linked_to
    TRANSACTIONS ||--o{ LEDGER_ENTRIES : records

    CUSTOMERS {
        uuid id PK
        varchar customer_number UK
        varchar legal_name
        varchar phone_number
        varchar national_id
        varchar kyc_level
        enum aml_status
        enum status
    }

    ACCOUNTS {
        uuid id PK
        varchar account_number UK
        uuid customer_id FK
        uuid branch_id FK
        uuid account_type_id FK
        varchar currency_code FK
        varchar account_name
        enum status
        boolean is_posting_restricted
        boolean is_debit_blocked
        boolean is_credit_blocked
    }

    ACCOUNT_BALANCES {
        uuid id PK
        uuid account_id FK,UK
        numeric ledger_balance
        numeric available_balance
        numeric locked_balance
        int version
    }

    ACCOUNT_HOLDS {
        uuid id PK
        varchar hold_reference UK
        uuid account_id FK
        numeric amount
        varchar reason
        varchar lock_event
        enum status
    }

    TRANSACTIONS {
        uuid id PK
        varchar transaction_reference UK
        varchar idempotency_key UK
        enum transaction_type
        enum channel
        uuid source_account_id FK
        uuid destination_account_id FK
        numeric amount
        numeric fee_amount
        enum status
    }

    LEDGER_ENTRIES {
        uuid id PK
        uuid transaction_id FK
        uuid account_id FK
        enum entry_type
        numeric amount
        numeric balance_after
        varchar narration
        date value_date
        timestamp posted_at
    }
```

---

## 2. Table Specifications & Constraints

### 2.1 `customers`
- **Primary Key:** `id` (UUID)
- **Unique Constraints:** `customer_number`
- **Indexes:** `phone_number`, `national_id` (Fayda NID)
- **Key Columns:**
  - `customer_number`: varchar(20), bank-wide customer identifier
  - `aml_status`: enum (`CLEAR`, `SUSPECT`, `BLOCKED`)
  - `status`: enum (`ACTIVE`, `INACTIVE`, `SUSPENDED`, `DECEASED`)

### 2.2 `accounts`
- **Primary Key:** `id` (UUID)
- **Unique Constraints:** `account_number` (13-character string, e.g. `1000123456789`)
- **Foreign Keys:**
  - `customer_id` -> `customers.id` (RESTRICT)
  - `branch_id` -> `branches.id` (RESTRICT)
  - `account_type_id` -> `account_types.id` (RESTRICT)
  - `currency_code` -> `currencies.code` (RESTRICT)
- **Restrictions Flags:**
  - `is_posting_restricted`, `is_debit_blocked`, `is_credit_blocked`

### 2.3 `account_balances`
- **Primary Key:** `id` (UUID)
- **Foreign Keys:** `account_id` -> `accounts.id` (CASCADE, UNIQUE)
- **Columns:**
  - `ledger_balance`: `numeric(18, 4)`
  - `available_balance`: `numeric(18, 4)`
  - `locked_balance`: `numeric(18, 4)`
  - `version`: integer optimistic locking counter
- **Invariant:**
  $$\text{available\_balance} = \text{ledger\_balance} - \text{locked\_balance}$$

### 2.4 `account_holds` (`ACLOCKEDEVENTS` Reference)
- **Primary Key:** `id` (UUID)
- **Unique Constraints:** `hold_reference` (e.g. `HLD-001`)
- **Foreign Keys:** `account_id` -> `accounts.id` (CASCADE)
- **Columns:**
  - `amount`: `numeric(18, 4)`
  - `lock_event`: `COURT_ORDER`, `PLEDGE`, `PENDING_SETTLEMENT`
  - `status`: enum (`ACTIVE`, `RELEASED`, `SETTLED`, `EXPIRED`)

### 2.5 `transactions`
- **Primary Key:** `id` (UUID)
- **Unique Constraints:** `transaction_reference` (e.g. `FT2629399881`), `idempotency_key`
- **Columns:**
  - `transaction_type`: `INTERNAL_TRANSFER`, `IPS_TRANSFER`, `WALLET_TRANSFER`, `BILL_PAYMENT`, `SPLIT_PAYMENT`, `RTGS_TRANSFER`, `CASH_DEPOSIT`, `FEE`, `REVERSAL`
  - `channel`: `CORE`, `WEB`, `MOBILE`, `IPS`, `CBE_BIRR`, `ATM`, `POS`, `SYSTEM`
  - `amount`, `fee_amount`: `numeric(18, 4)`
  - `status`: enum (`PENDING`, `COMPLETED`, `FAILED`, `REVERSED`)

### 2.6 `ledger_entries` (Double-Entry Bookkeeping)
- **Primary Key:** `id` (UUID)
- **Foreign Keys:**
  - `transaction_id` -> `transactions.id` (CASCADE)
  - `account_id` -> `accounts.id` (RESTRICT)
- **Columns:**
  - `entry_type`: enum (`DEBIT`, `CREDIT`)
  - `amount`: `numeric(18, 4)`
  - `balance_after`: `numeric(18, 4)`
  - `narration`: varchar(255)
  - `value_date`: date
  - `posted_at`: timestamp
