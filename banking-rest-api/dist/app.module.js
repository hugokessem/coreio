"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppModule = void 0;
const common_1 = require("@nestjs/common");
const config_1 = require("@nestjs/config");
const typeorm_1 = require("@nestjs/typeorm");
const database_config_1 = require("./config/database.config");
const auth_module_1 = require("./modules/auth/auth.module");
const customers_module_1 = require("./modules/customers/customers.module");
const accounts_module_1 = require("./modules/accounts/accounts.module");
const transfers_module_1 = require("./modules/transfers/transfers.module");
const standing_orders_module_1 = require("./modules/standing-orders/standing-orders.module");
const payments_module_1 = require("./modules/payments/payments.module");
const cards_module_1 = require("./modules/cards/cards.module");
const rates_module_1 = require("./modules/rates/rates.module");
const services_module_1 = require("./modules/services/services.module");
const superapp_module_1 = require("./modules/superapp/superapp.module");
const ips_module_1 = require("./modules/ips/ips.module");
const wallet_module_1 = require("./modules/wallet/wallet.module");
const health_module_1 = require("./modules/health/health.module");
let AppModule = class AppModule {
};
exports.AppModule = AppModule;
exports.AppModule = AppModule = __decorate([
    (0, common_1.Module)({
        imports: [
            config_1.ConfigModule.forRoot({
                isGlobal: true,
                load: [database_config_1.default],
            }),
            typeorm_1.TypeOrmModule.forRootAsync({
                inject: [config_1.ConfigService],
                useFactory: (configService) => ({
                    ...configService.get('database'),
                    autoLoadEntities: true,
                }),
            }),
            auth_module_1.AuthModule,
            customers_module_1.CustomersModule,
            accounts_module_1.AccountsModule,
            transfers_module_1.TransfersModule,
            standing_orders_module_1.StandingOrdersModule,
            payments_module_1.PaymentsModule,
            cards_module_1.CardsModule,
            rates_module_1.RatesModule,
            services_module_1.ServicesModule,
            superapp_module_1.SuperAppModule,
            ips_module_1.IpsModule,
            wallet_module_1.WalletModule,
            health_module_1.HealthModule,
        ],
    })
], AppModule);
//# sourceMappingURL=app.module.js.map