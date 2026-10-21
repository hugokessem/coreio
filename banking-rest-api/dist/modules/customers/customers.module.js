"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.CustomersModule = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const customer_entity_1 = require("../../database/entities/customer.entity");
const account_entity_1 = require("../../database/entities/account.entity");
const branch_entity_1 = require("../../database/entities/branch.entity");
const currency_entity_1 = require("../../database/entities/currency.entity");
const account_type_entity_1 = require("../../database/entities/account-type.entity");
const account_party_entity_1 = require("../../database/entities/account-party.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const customers_service_1 = require("./customers.service");
const customers_controller_1 = require("./customers.controller");
let CustomersModule = class CustomersModule {
};
exports.CustomersModule = CustomersModule;
exports.CustomersModule = CustomersModule = __decorate([
    (0, common_1.Module)({
        imports: [
            typeorm_1.TypeOrmModule.forFeature([
                customer_entity_1.Customer,
                account_entity_1.Account,
                branch_entity_1.Branch,
                currency_entity_1.Currency,
                account_type_entity_1.AccountType,
                account_party_entity_1.AccountParty,
                account_balance_entity_1.AccountBalance,
                transaction_entity_1.Transaction,
                ledger_entry_entity_1.LedgerEntry,
            ]),
        ],
        controllers: [customers_controller_1.CustomersController],
        providers: [customers_service_1.CustomersService],
        exports: [customers_service_1.CustomersService],
    })
], CustomersModule);
//# sourceMappingURL=customers.module.js.map