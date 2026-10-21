# CBE Banking REST API Contracts & Endpoint Specification

## 1. Global Conventions & Standards

- **Base URL:** `http://localhost:3000/api/v1`
- **Authentication:**
  - `Authorization: Bearer <jwt-token>`
  - Or header: `x-api-key: cbe-banking-api-key-default-2026`
- **Content-Type:** `application/json`
- **Error Response Standard:**
```json
{
  "statusCode": 422,
  "errorCode": "INSUFFICIENT_FUNDS",
  "message": "Insufficient available funds. Available: 50.0000, Required: 500.0000",
  "details": null,
  "timestamp": "2026-10-20T12:00:00.000Z",
  "path": "/api/v1/transfers/execute",
  "correlationId": null
}
```

---

## 2. Authentication API

### POST `/api/v1/auth/login`
**Request:**
```json
{
  "username": "admin",
  "password": "Admin123!"
}
```
**Response (200 OK):**
```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIsIn...",
  "tokenType": "Bearer",
  "expiresIn": "86400s",
  "user": {
    "id": "e4f5a6b7-c8d9-...",
    "username": "admin",
    "email": "admin@cbe.com.et",
    "roles": ["ADMIN", "TELLER", "API_CLIENT"]
  }
}
```

---

## 3. Customers API

### POST `/api/v1/customers/lookup`
**Request:**
```json
{
  "customerNumber": "10000001"
}
```
**Response (200 OK):**
```json
{
  "customerNumber": "10000001",
  "customerType": "INDIVIDUAL",
  "fullName": "Abebe Kebede Bikila",
  "phoneNumber": "+251911223344",
  "nationalId": "FAYDA-9876543210",
  "kycLevel": "LEVEL_3",
  "amlStatus": "CLEAR",
  "status": "ACTIVE",
  "accounts": [
    {
      "accountNumber": "1000123456789",
      "accountName": "Abebe Kebede Bikila",
      "currency": "ETB",
      "status": "ACTIVE"
    }
  ]
}
```

---

## 4. Accounts & Holds API

### POST `/api/v1/accounts/lookup`
**Request:**
```json
{
  "accountNumber": "1000123456789"
}
```
**Response (200 OK):**
```json
{
  "accountNumber": "1000123456789",
  "accountName": "Abebe Kebede Bikila",
  "customerNumber": "10000001",
  "accountTypeCode": "6501",
  "accountTypeName": "Ordinary Savings Account",
  "branchCode": "001",
  "branchName": "Addis Ababa Main Branch",
  "currency": "ETB",
  "status": "ACTIVE",
  "ledgerBalance": "150000.0000",
  "availableBalance": "145000.0000",
  "lockedBalance": "5000.0000",
  "openedAt": "2026-10-20T11:49:27.000Z"
}
```

### POST `/api/v1/accounts/holds/create`
**Request:**
```json
{
  "accountNumber": "1000123456789",
  "amount": "2500.0000",
  "reason": "Administrative Guarantee Hold",
  "lockEvent": "PLEDGE"
}
```
**Response (201 Created):**
```json
{
  "holdReference": "HLD-20261020-1002",
  "accountNumber": "1000123456789",
  "amount": "2500.0000",
  "currency": "ETB",
  "reason": "Administrative Guarantee Hold",
  "lockEvent": "PLEDGE",
  "status": "ACTIVE",
  "newAvailableBalance": "142500.0000",
  "newLockedBalance": "7500.0000",
  "createdAt": "2026-10-20T12:01:00.000Z"
}
```

---

## 5. Transfers API

### POST `/api/v1/transfers/execute`
**Request:**
```json
{
  "sourceAccountNumber": "1000123456789",
  "destinationAccountNumber": "1000987654321",
  "amount": "500.0000",
  "narration": "Monthly Rental Payment",
  "idempotencyKey": "IDEMP-RENT-20261020-001"
}
```
**Response (200 OK):**
```json
{
  "transactionReference": "FT2638491823",
  "status": "COMPLETED",
  "transactionType": "INTERNAL_TRANSFER",
  "channel": "MOBILE",
  "sourceAccountNumber": "1000123456789",
  "destinationAccountNumber": "1000987654321",
  "amount": "500.0000",
  "fee": "0.0000",
  "currency": "ETB",
  "narration": "Monthly Rental Payment",
  "idempotencyKey": "IDEMP-RENT-20261020-001",
  "initiatedAt": "2026-10-20T12:02:00.000Z",
  "completedAt": "2026-10-20T12:02:00.000Z",
  "message": "Transfer completed successfully"
}
```
