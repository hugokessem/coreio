"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.SuperAppModule = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const super_app_user_entity_1 = require("../../database/entities/super-app-user.entity");
const customer_entity_1 = require("../../database/entities/customer.entity");
const superapp_service_1 = require("./superapp.service");
const superapp_controller_1 = require("./superapp.controller");
let SuperAppModule = class SuperAppModule {
};
exports.SuperAppModule = SuperAppModule;
exports.SuperAppModule = SuperAppModule = __decorate([
    (0, common_1.Module)({
        imports: [typeorm_1.TypeOrmModule.forFeature([super_app_user_entity_1.SuperAppUser, customer_entity_1.Customer])],
        controllers: [superapp_controller_1.SuperAppController],
        providers: [superapp_service_1.SuperAppService],
        exports: [superapp_service_1.SuperAppService],
    })
], SuperAppModule);
//# sourceMappingURL=superapp.module.js.map