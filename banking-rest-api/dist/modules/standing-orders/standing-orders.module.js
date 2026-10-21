"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.StandingOrdersModule = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const standing_order_entity_1 = require("../../database/entities/standing-order.entity");
const standing_order_execution_entity_1 = require("../../database/entities/standing-order-execution.entity");
const account_entity_1 = require("../../database/entities/account.entity");
const standing_orders_service_1 = require("./standing-orders.service");
const standing_orders_controller_1 = require("./standing-orders.controller");
const transfers_module_1 = require("../transfers/transfers.module");
let StandingOrdersModule = class StandingOrdersModule {
};
exports.StandingOrdersModule = StandingOrdersModule;
exports.StandingOrdersModule = StandingOrdersModule = __decorate([
    (0, common_1.Module)({
        imports: [
            typeorm_1.TypeOrmModule.forFeature([standing_order_entity_1.StandingOrder, standing_order_execution_entity_1.StandingOrderExecution, account_entity_1.Account]),
            transfers_module_1.TransfersModule,
        ],
        controllers: [standing_orders_controller_1.StandingOrdersController],
        providers: [standing_orders_service_1.StandingOrdersService],
        exports: [standing_orders_service_1.StandingOrdersService],
    })
], StandingOrdersModule);
//# sourceMappingURL=standing-orders.module.js.map