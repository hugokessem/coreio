"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Transaction = exports.TransactionChannel = exports.TransactionStatus = exports.TransactionType = void 0;
const typeorm_1 = require("typeorm");
const account_entity_1 = require("./account.entity");
const ledger_entry_entity_1 = require("./ledger-entry.entity");
var TransactionType;
(function (TransactionType) {
    TransactionType["INTERNAL_TRANSFER"] = "INTERNAL_TRANSFER";
    TransactionType["TRANSFER_FROM_HOLD"] = "TRANSFER_FROM_HOLD";
    TransactionType["IPS_TRANSFER"] = "IPS_TRANSFER";
    TransactionType["IPS_QR_PAYMENT"] = "IPS_QR_PAYMENT";
    TransactionType["WALLET_TRANSFER"] = "WALLET_TRANSFER";
    TransactionType["BILL_PAYMENT"] = "BILL_PAYMENT";
    TransactionType["SPLIT_PAYMENT"] = "SPLIT_PAYMENT";
    TransactionType["RTGS_TRANSFER"] = "RTGS_TRANSFER";
    TransactionType["CASH_DEPOSIT"] = "CASH_DEPOSIT";
    TransactionType["CASH_WITHDRAWAL"] = "CASH_WITHDRAWAL";
    TransactionType["FEE"] = "FEE";
    TransactionType["REVERSAL"] = "REVERSAL";
})(TransactionType || (exports.TransactionType = TransactionType = {}));
var TransactionStatus;
(function (TransactionStatus) {
    TransactionStatus["PENDING"] = "PENDING";
    TransactionStatus["COMPLETED"] = "COMPLETED";
    TransactionStatus["FAILED"] = "FAILED";
    TransactionStatus["REVERSED"] = "REVERSED";
})(TransactionStatus || (exports.TransactionStatus = TransactionStatus = {}));
var TransactionChannel;
(function (TransactionChannel) {
    TransactionChannel["CORE"] = "CORE";
    TransactionChannel["WEB"] = "WEB";
    TransactionChannel["MOBILE"] = "MOBILE";
    TransactionChannel["IPS"] = "IPS";
    TransactionChannel["CBE_BIRR"] = "CBE_BIRR";
    TransactionChannel["ATM"] = "ATM";
    TransactionChannel["POS"] = "POS";
    TransactionChannel["SYSTEM"] = "SYSTEM";
})(TransactionChannel || (exports.TransactionChannel = TransactionChannel = {}));
let Transaction = class Transaction {
    id;
    transactionReference;
    idempotencyKey;
    transactionType;
    channel;
    sourceAccountId;
    sourceAccount;
    destinationAccountId;
    destinationAccount;
    externalAccount;
    externalBankCode;
    amount;
    feeAmount;
    currencyCode;
    exchangeRate;
    narration;
    status;
    failureReason;
    reversalOfTransactionId;
    initiatedAt;
    completedAt;
    ledgerEntries;
    createdAt;
    updatedAt;
};
exports.Transaction = Transaction;
__decorate([
    (0, typeorm_1.PrimaryGeneratedColumn)('uuid'),
    __metadata("design:type", String)
], Transaction.prototype, "id", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'transaction_reference', type: 'varchar', length: 50, unique: true }),
    __metadata("design:type", String)
], Transaction.prototype, "transactionReference", void 0);
__decorate([
    (0, typeorm_1.Index)({ unique: true }),
    (0, typeorm_1.Column)({ name: 'idempotency_key', type: 'varchar', length: 100, unique: true, nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "idempotencyKey", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'transaction_type',
        type: 'enum',
        enum: TransactionType,
        default: TransactionType.INTERNAL_TRANSFER,
    }),
    __metadata("design:type", String)
], Transaction.prototype, "transactionType", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'channel',
        type: 'enum',
        enum: TransactionChannel,
        default: TransactionChannel.CORE,
    }),
    __metadata("design:type", String)
], Transaction.prototype, "channel", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'source_account_id', type: 'uuid', nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "sourceAccountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account, { nullable: true }),
    (0, typeorm_1.JoinColumn)({ name: 'source_account_id' }),
    __metadata("design:type", Object)
], Transaction.prototype, "sourceAccount", void 0);
__decorate([
    (0, typeorm_1.Index)(),
    (0, typeorm_1.Column)({ name: 'destination_account_id', type: 'uuid', nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "destinationAccountId", void 0);
__decorate([
    (0, typeorm_1.ManyToOne)(() => account_entity_1.Account, { nullable: true }),
    (0, typeorm_1.JoinColumn)({ name: 'destination_account_id' }),
    __metadata("design:type", Object)
], Transaction.prototype, "destinationAccount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'external_account', type: 'varchar', length: 50, nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "externalAccount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'external_bank_code', type: 'varchar', length: 30, nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "externalBankCode", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
    }),
    __metadata("design:type", String)
], Transaction.prototype, "amount", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'fee_amount',
        type: 'numeric',
        precision: 18,
        scale: 4,
        default: '0.0000',
    }),
    __metadata("design:type", String)
], Transaction.prototype, "feeAmount", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'currency_code', type: 'varchar', length: 3 }),
    __metadata("design:type", String)
], Transaction.prototype, "currencyCode", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'exchange_rate',
        type: 'numeric',
        precision: 18,
        scale: 6,
        default: '1.000000',
    }),
    __metadata("design:type", String)
], Transaction.prototype, "exchangeRate", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'narration', type: 'varchar', length: 255 }),
    __metadata("design:type", String)
], Transaction.prototype, "narration", void 0);
__decorate([
    (0, typeorm_1.Column)({
        name: 'status',
        type: 'enum',
        enum: TransactionStatus,
        default: TransactionStatus.PENDING,
    }),
    __metadata("design:type", String)
], Transaction.prototype, "status", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'failure_reason', type: 'varchar', length: 255, nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "failureReason", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'reversal_of_transaction_id', type: 'uuid', nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "reversalOfTransactionId", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'initiated_at', type: 'timestamp', default: () => 'CURRENT_TIMESTAMP' }),
    __metadata("design:type", Date)
], Transaction.prototype, "initiatedAt", void 0);
__decorate([
    (0, typeorm_1.Column)({ name: 'completed_at', type: 'timestamp', nullable: true }),
    __metadata("design:type", Object)
], Transaction.prototype, "completedAt", void 0);
__decorate([
    (0, typeorm_1.OneToMany)(() => ledger_entry_entity_1.LedgerEntry, (entry) => entry.transaction, { cascade: true }),
    __metadata("design:type", Array)
], Transaction.prototype, "ledgerEntries", void 0);
__decorate([
    (0, typeorm_1.CreateDateColumn)({ name: 'created_at' }),
    __metadata("design:type", Date)
], Transaction.prototype, "createdAt", void 0);
__decorate([
    (0, typeorm_1.UpdateDateColumn)({ name: 'updated_at' }),
    __metadata("design:type", Date)
], Transaction.prototype, "updatedAt", void 0);
exports.Transaction = Transaction = __decorate([
    (0, typeorm_1.Entity)('transactions')
], Transaction);
//# sourceMappingURL=transaction.entity.js.map