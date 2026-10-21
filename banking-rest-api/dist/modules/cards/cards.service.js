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
exports.CardsService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const card_request_entity_1 = require("../../database/entities/card-request.entity");
const card_entity_1 = require("../../database/entities/card.entity");
const account_entity_1 = require("../../database/entities/account.entity");
const branch_entity_1 = require("../../database/entities/branch.entity");
const customer_entity_1 = require("../../database/entities/customer.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
let CardsService = class CardsService {
    cardRequestRepo;
    cardRepo;
    accountRepo;
    branchRepo;
    customerRepo;
    constructor(cardRequestRepo, cardRepo, accountRepo, branchRepo, customerRepo) {
        this.cardRequestRepo = cardRequestRepo;
        this.cardRepo = cardRepo;
        this.accountRepo = accountRepo;
        this.branchRepo = branchRepo;
        this.customerRepo = customerRepo;
    }
    async requestCard(dto) {
        const account = await this.accountRepo.findOne({
            where: { accountNumber: dto.accountNumber },
            relations: ['customer'],
        });
        if (!account)
            throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
        const branch = await this.branchRepo.findOne({
            where: { branchCode: dto.targetBranchCode },
        });
        if (!branch)
            throw new common_1.BadRequestException(`Branch code '${dto.targetBranchCode}' not found`);
        const count = await this.cardRequestRepo.count();
        const requestReference = `CR-${new Date().toISOString().slice(0, 10).replace(/-/g, '')}-${(1000 + count + 1).toString()}`;
        const cardReq = this.cardRequestRepo.create({
            requestReference,
            accountId: account.id,
            customerId: account.customerId,
            cardType: dto.cardType || card_request_entity_1.CardType.DEBIT,
            cardBrand: dto.cardBrand || card_request_entity_1.CardBrand.MASTERCARD,
            requestType: card_request_entity_1.CardRequestType.NEW,
            status: card_request_entity_1.CardRequestStatus.PENDING,
            targetBranchId: branch.id,
        });
        const saved = await this.cardRequestRepo.save(cardReq);
        return {
            requestReference: saved.requestReference,
            accountNumber: account.accountNumber,
            customerName: account.customer.fullName,
            cardType: saved.cardType,
            cardBrand: saved.cardBrand,
            collectionBranch: branch.branchName,
            status: saved.status,
            createdAt: saved.createdAt,
            message: 'Card order queued for personalization and dispatch',
        };
    }
    async replaceCard(dto) {
        const card = await this.cardRepo.findOne({
            where: { cardNumberMasked: dto.cardNumberMasked },
            relations: ['account', 'customer'],
        });
        if (!card) {
            throw new common_1.NotFoundException(`Card '${dto.cardNumberMasked}' not found`);
        }
        const branch = await this.branchRepo.findOne({
            where: { branchCode: dto.targetBranchCode },
        });
        if (!branch)
            throw new common_1.BadRequestException(`Branch '${dto.targetBranchCode}' not found`);
        card.status = card_entity_1.CardStatus.BLOCKED;
        await this.cardRepo.save(card);
        const count = await this.cardRequestRepo.count();
        const requestReference = `REP-${new Date().toISOString().slice(0, 10).replace(/-/g, '')}-${(1000 + count + 1).toString()}`;
        const req = this.cardRequestRepo.create({
            requestReference,
            accountId: card.accountId,
            customerId: card.customerId,
            cardType: card.cardType,
            cardBrand: card.cardBrand,
            requestType: card_request_entity_1.CardRequestType.REPLACEMENT,
            status: card_request_entity_1.CardRequestStatus.PENDING,
            targetBranchId: branch.id,
        });
        const saved = await this.cardRequestRepo.save(req);
        return {
            requestReference: saved.requestReference,
            previousCardMasked: card.cardNumberMasked,
            previousCardStatus: card.status,
            replacementStatus: saved.status,
            collectionBranch: branch.branchName,
            createdAt: saved.createdAt,
            message: 'Previous card blocked; replacement card requested',
        };
    }
    async listCards(dto) {
        const where = {};
        if (dto.accountNumber) {
            const account = await this.accountRepo.findOne({ where: { accountNumber: dto.accountNumber } });
            if (!account)
                throw new domain_exceptions_1.AccountNotFoundException(dto.accountNumber);
            where.accountId = account.id;
        }
        if (dto.customerNumber) {
            const customer = await this.customerRepo.findOne({ where: { customerNumber: dto.customerNumber } });
            if (!customer)
                throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
            where.customerId = customer.id;
        }
        const cards = await this.cardRepo.find({
            where,
            relations: ['account'],
            order: { createdAt: 'DESC' },
        });
        return {
            count: cards.length,
            cards: cards.map((c) => ({
                id: c.id,
                cardNumberMasked: c.cardNumberMasked,
                cardType: c.cardType,
                cardBrand: c.cardBrand,
                accountNumber: c.account?.accountNumber,
                status: c.status,
                expiryDate: c.expiryDate,
                dailyLimit: c.dailyLimit,
                createdAt: c.createdAt,
            })),
        };
    }
};
exports.CardsService = CardsService;
exports.CardsService = CardsService = __decorate([
    (0, common_1.Injectable)(),
    __param(0, (0, typeorm_1.InjectRepository)(card_request_entity_1.CardRequest)),
    __param(1, (0, typeorm_1.InjectRepository)(card_entity_1.Card)),
    __param(2, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(3, (0, typeorm_1.InjectRepository)(branch_entity_1.Branch)),
    __param(4, (0, typeorm_1.InjectRepository)(customer_entity_1.Customer)),
    __metadata("design:paramtypes", [typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], CardsService);
//# sourceMappingURL=cards.service.js.map