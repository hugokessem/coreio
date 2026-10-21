"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.IpsModule = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const service_entity_1 = require("../../database/entities/service.entity");
const ips_service_1 = require("./ips.service");
const ips_controller_1 = require("./ips.controller");
let IpsModule = class IpsModule {
};
exports.IpsModule = IpsModule;
exports.IpsModule = IpsModule = __decorate([
    (0, common_1.Module)({
        imports: [
            typeorm_1.TypeOrmModule.forFeature([
                account_entity_1.Account,
                account_balance_entity_1.AccountBalance,
                transaction_entity_1.Transaction,
                ledger_entry_entity_1.LedgerEntry,
                service_entity_1.ServiceEntity,
            ]),
        ],
        controllers: [ips_controller_1.IpsController],
        providers: [ips_service_1.IpsService],
        exports: [ips_service_1.IpsService],
    })
], IpsModule);
//# sourceMappingURL=ips.module.js.map