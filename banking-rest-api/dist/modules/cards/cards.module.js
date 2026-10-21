"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.CardsModule = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const card_request_entity_1 = require("../../database/entities/card-request.entity");
const card_entity_1 = require("../../database/entities/card.entity");
const account_entity_1 = require("../../database/entities/account.entity");
const branch_entity_1 = require("../../database/entities/branch.entity");
const customer_entity_1 = require("../../database/entities/customer.entity");
const cards_service_1 = require("./cards.service");
const cards_controller_1 = require("./cards.controller");
let CardsModule = class CardsModule {
};
exports.CardsModule = CardsModule;
exports.CardsModule = CardsModule = __decorate([
    (0, common_1.Module)({
        imports: [typeorm_1.TypeOrmModule.forFeature([card_request_entity_1.CardRequest, card_entity_1.Card, account_entity_1.Account, branch_entity_1.Branch, customer_entity_1.Customer])],
        controllers: [cards_controller_1.CardsController],
        providers: [cards_service_1.CardsService],
        exports: [cards_service_1.CardsService],
    })
], CardsModule);
//# sourceMappingURL=cards.module.js.map