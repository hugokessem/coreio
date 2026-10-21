"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.TransfersModule = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const account_hold_entity_1 = require("../../database/entities/account-hold.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const service_entity_1 = require("../../database/entities/service.entity");
const idempotency_record_entity_1 = require("../../database/entities/idempotency-record.entity");
const transfers_service_1 = require("./transfers.service");
const transfers_controller_1 = require("./transfers.controller");
const idempotency_service_1 = require("../../common/services/idempotency.service");
const reconciliation_service_1 = require("./reconciliation.service");
let TransfersModule = class TransfersModule {
};
exports.TransfersModule = TransfersModule;
exports.TransfersModule = TransfersModule = __decorate([
    (0, common_1.Module)({
        imports: [
            typeorm_1.TypeOrmModule.forFeature([
                account_entity_1.Account,
                account_balance_entity_1.AccountBalance,
                account_hold_entity_1.AccountHold,
                transaction_entity_1.Transaction,
                ledger_entry_entity_1.LedgerEntry,
                service_entity_1.ServiceEntity,
                idempotency_record_entity_1.IdempotencyRecord,
            ]),
        ],
        controllers: [transfers_controller_1.TransfersController],
        providers: [transfers_service_1.TransfersService, idempotency_service_1.IdempotencyService, reconciliation_service_1.ReconciliationService],
        exports: [transfers_service_1.TransfersService, reconciliation_service_1.ReconciliationService],
    })
], TransfersModule);
//# sourceMappingURL=transfers.module.js.map