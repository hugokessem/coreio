# Banking Business Rules & Financial Invariants

## 1. Account Lifecycle & Restriction Rules

1. **Account Numbering:**
   - Account numbers must be represented as character strings (never integers or floats) to preserve leading zeros.
   - Standard CBE accounts are 13 digits, starting with prefix `1000` (individual) or `2000` (corporate).
2. **Account States:**
   - `ACTIVE`: Fully functional for debits, credits, inquiries, and standing orders.
   - `DORMANT`: Inactive for > 6 months. Inbound credits allowed; outbound debits blocked until re-activated in-branch.
   - `FROZEN`: Placed under judicial or compliance block. All debit and credit movements rejected.
   - `RESTRICTED`: Posting restrictions configured via `is_posting_restricted`, `is_debit_blocked`, or `is_credit_blocked`.
   - `CLOSED`: Account permanently terminated. Zero balance invariant strictly enforced before closure.
3. **Account Deletion Guard:**
   - Customers or accounts with existing transaction history or non-zero ledger balances cannot be deleted.

---

## 2. Balances & Hold (ACLOCKEDEVENTS) Mechanics

1. **Balance Separation:**
   - Every account maintains three distinct balances:
     - `ledger_balance`: Cumulative sum of all authoritative posted credits minus debits.
     - `locked_balance`: Sum of all currently active holds (`status = 'ACTIVE'`).
     - `available_balance`: Funds freely available for disbursement:
       $$\text{available\_balance} = \text{ledger\_balance} - \text{locked\_balance}$$
2. **Hold Creation:**
   - An account hold can only be placed if:
     $$\text{available\_balance} \ge \text{hold\_amount}$$
   - When placed, `locked_balance` increases by `hold_amount`, and `available_balance` decreases by `hold_amount`. `ledger_balance` remains unchanged.
3. **Hold Release:**
   - Unfreezes funds: `locked_balance` decreases by `hold_amount`, and `available_balance` increases by `hold_amount`.
4. **Hold Settlement (`transferFromHold`):**
   - Moves held funds to a destination account:
     - `locked_balance` decreases by `hold_amount`.
     - `ledger_balance` decreases by `hold_amount` (since `available_balance` was already deducted at hold creation).
     - Destination account receives credit to both `ledger_balance` and `available_balance`.

---

## 3. Financial Transaction Consistency & Invariants

1. **Atomicity & Double-Entry Postings:**
   - Every financial mutation occurs within a single PostgreSQL `SERIALIZABLE` or read-committed transaction protected by pessimistic row locks.
   - Every transfer generates balanced entries:
     $$\sum \text{Debit Entries} = \sum \text{Credit Entries}$$
2. **Deadlock Prevention Rule:**
   - When locking multiple account balances, IDs must be sorted deterministically:
     $$\text{Lock Order} = \text{sort}([A, B])$$
3. **Idempotency Guarantee:**
   - Every financial mutation accepts an optional `idempotencyKey`.
   - If a request is received with an identical idempotency key that has already committed, the API returns the cached transaction response immediately without modifying balances or generating new ledger entries.
   - If two identical idempotency keys race concurrently, the database unique constraint serializes execution; the runner-up catches the error and returns the winner's committed transaction receipt.
4. **Non-Negative Balance Invariant:**
   - Standard savings and demand accounts cannot be overdrawn. Transactions attempting to reduce `available_balance` below zero are rejected with HTTP 422 (`INSUFFICIENT_FUNDS`).

---

## 4. AML Screening & Risk Governance

1. **Customer AML Screening:**
   - Before opening accounts or executing high-value payments, customer `aml_status` is screened.
   - Customers flagged with `aml_status = 'SUSPECT'` or `'BLOCKED'` are barred from initiating outward remittances.
2. **Daily Channel Limits:**
   - Internal transfers: 5,000,000 ETB daily.
   - Interbank IPS transfers: 2,000,000 ETB daily.
   - CBE Birr wallet transfers: 300,000 ETB daily.
   - Transactions exceeding limits without custom override are rejected with HTTP 422 (`DAILY_LIMIT_EXCEEDED`).
