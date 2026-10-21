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
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.CardsController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const cards_service_1 = require("./cards.service");
const card_dto_1 = require("./dto/card.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
let CardsController = class CardsController {
    cardsService;
    constructor(cardsService) {
        this.cardsService = cardsService;
    }
    async requestCard(dto) {
        return this.cardsService.requestCard(dto);
    }
    async replaceCard(dto) {
        return this.cardsService.replaceCard(dto);
    }
    async listCards(dto) {
        return this.cardsService.listCards(dto);
    }
    async getByAccount(accountNumber) {
        return this.cardsService.listCards({ accountNumber });
    }
};
exports.CardsController = CardsController;
__decorate([
    (0, common_1.Post)('request'),
    (0, common_1.HttpCode)(common_1.HttpStatus.CREATED),
    (0, swagger_1.ApiOperation)({ summary: 'Submit debit or Mastercard card issuance request' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [card_dto_1.RequestCardDto]),
    __metadata("design:returntype", Promise)
], CardsController.prototype, "requestCard", null);
__decorate([
    (0, common_1.Post)('replace'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Block compromised card and request immediate replacement' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [card_dto_1.ReplaceCardDto]),
    __metadata("design:returntype", Promise)
], CardsController.prototype, "replaceCard", null);
__decorate([
    (0, common_1.Post)('list'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'List cards by account or customer number' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [card_dto_1.ListCardsDto]),
    __metadata("design:returntype", Promise)
], CardsController.prototype, "listCards", null);
__decorate([
    (0, common_1.Get)('account/:accountNumber'),
    (0, swagger_1.ApiOperation)({ summary: 'GET active cards for account (REST convention)' }),
    __param(0, (0, common_1.Param)('accountNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], CardsController.prototype, "getByAccount", null);
exports.CardsController = CardsController = __decorate([
    (0, swagger_1.ApiTags)('Cards'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard),
    (0, common_1.Controller)('api/v1/cards'),
    __metadata("design:paramtypes", [cards_service_1.CardsService])
], CardsController);
//# sourceMappingURL=cards.controller.js.map