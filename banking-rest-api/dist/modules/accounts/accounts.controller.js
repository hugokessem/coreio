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
exports.AccountsController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const accounts_service_1 = require("./accounts.service");
const account_dto_1 = require("./dto/account.dto");
const combined_auth_guard_1 = require("../../common/guards/combined-auth.guard");
let AccountsController = class AccountsController {
    accountsService;
    constructor(accountsService) {
        this.accountsService = accountsService;
    }
    async lookup(dto) {
        return this.accountsService.lookup(dto);
    }
    async nameLookup(dto) {
        return this.accountsService.nameLookup(dto);
    }
    async list(dto) {
        return this.accountsService.list(dto);
    }
    async listRest(customerNumber) {
        return this.accountsService.list({ customerNumber });
    }
    async create(dto) {
        return this.accountsService.create(dto);
    }
    async eligibility(dto) {
        return this.accountsService.checkEligibility(dto);
    }
    async createHold(dto) {
        return this.accountsService.createHold(dto);
    }
    async listHoldsPost(dto) {
        return this.accountsService.listHolds(dto.accountNumber);
    }
    async listHoldsGet(accountNumber) {
        return this.accountsService.listHolds(accountNumber);
    }
    async releaseHold(dto) {
        return this.accountsService.releaseHold(dto);
    }
    async statement(dto) {
        return this.accountsService.getStatement(dto);
    }
    async statementByDateRange(dto) {
        return this.accountsService.getStatement(dto);
    }
    async miniStatement(dto) {
        return this.accountsService.getMiniStatement(dto);
    }
    async miniStatementByLimit(dto) {
        return this.accountsService.getMiniStatement(dto);
    }
    async getByNumber(accountNumber) {
        return this.accountsService.lookup({ accountNumber });
    }
    async getBalance(accountNumber) {
        const acc = await this.accountsService.lookup({ accountNumber });
        return {
            accountNumber: acc.accountNumber,
            currency: acc.currency,
            ledgerBalance: acc.ledgerBalance,
            availableBalance: acc.availableBalance,
            lockedBalance: acc.lockedBalance,
        };
    }
    async getStatementRest(accountNumber, startDate, endDate, limit = '50') {
        return this.accountsService.getStatement({
            accountNumber,
            startDate,
            endDate,
            limit: parseInt(limit, 10),
        });
    }
};
exports.AccountsController = AccountsController;
__decorate([
    (0, common_1.Post)('lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Lookup account details, balances, status, branch, and customer' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Account found' }),
    (0, swagger_1.ApiResponse)({ status: 404, description: 'Account not found' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountLookupDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "lookup", null);
__decorate([
    (0, common_1.Post)('name-lookup'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Account title/name enquiry for transfer beneficiaries' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountNameLookupDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "nameLookup", null);
__decorate([
    (0, common_1.Post)('list'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'List all accounts belonging to a customer' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountListDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "list", null);
__decorate([
    (0, common_1.Get)(),
    (0, swagger_1.ApiOperation)({ summary: 'List accounts for a customer (REST convention)' }),
    __param(0, (0, common_1.Query)('customerNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "listRest", null);
__decorate([
    (0, common_1.Post)(['create', '']),
    (0, common_1.HttpCode)(common_1.HttpStatus.CREATED),
    (0, swagger_1.ApiOperation)({ summary: 'Open new bank account with initial deposit option' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.CreateAccountDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "create", null);
__decorate([
    (0, common_1.Post)('eligibility'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Check account eligibility for debit/credit operations' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountEligibilityDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "eligibility", null);
__decorate([
    (0, common_1.Post)(['holds/create', 'holds']),
    (0, common_1.HttpCode)(common_1.HttpStatus.CREATED),
    (0, swagger_1.ApiOperation)({ summary: 'Place administrative, court, or pledge hold on account' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.CreateHoldDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "createHold", null);
__decorate([
    (0, common_1.Post)('holds/list'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'List active holds for an account' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountLookupDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "listHoldsPost", null);
__decorate([
    (0, common_1.Get)(':accountNumber/holds'),
    (0, swagger_1.ApiOperation)({ summary: 'GET active holds for account (REST convention)' }),
    __param(0, (0, common_1.Param)('accountNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "listHoldsGet", null);
__decorate([
    (0, common_1.Post)('holds/release'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Release active hold, unfreezing the locked balance' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.ReleaseHoldDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "releaseHold", null);
__decorate([
    (0, common_1.Post)('statement'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Generate detailed account ledger statement (Op 29: MiniStatementByDateRange)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountStatementDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "statement", null);
__decorate([
    (0, common_1.Post)('mini-statement-by-date-range'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Generate detailed account ledger statement - CoreIO Alias (Op 29: MiniStatementByDateRange)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountStatementDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "statementByDateRange", null);
__decorate([
    (0, common_1.Post)('mini-statement'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Generate mini-statement (recent N transactions) (Op 28: MiniStatementByLimit)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountMiniStatementDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "miniStatement", null);
__decorate([
    (0, common_1.Post)('mini-statement-by-limit'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Generate mini-statement - CoreIO Alias (Op 28: MiniStatementByLimit)' }),
    __param(0, (0, common_1.Body)()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [account_dto_1.AccountMiniStatementDto]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "miniStatementByLimit", null);
__decorate([
    (0, common_1.Get)(':accountNumber'),
    (0, swagger_1.ApiOperation)({ summary: 'GET account details by account number (REST convention)' }),
    __param(0, (0, common_1.Param)('accountNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "getByNumber", null);
__decorate([
    (0, common_1.Get)(':accountNumber/balance'),
    (0, swagger_1.ApiOperation)({ summary: 'GET account balance by account number (REST convention)' }),
    __param(0, (0, common_1.Param)('accountNumber')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "getBalance", null);
__decorate([
    (0, common_1.Get)(':accountNumber/statement'),
    (0, swagger_1.ApiOperation)({ summary: 'GET statement by account number (REST convention)' }),
    __param(0, (0, common_1.Param)('accountNumber')),
    __param(1, (0, common_1.Query)('startDate')),
    __param(2, (0, common_1.Query)('endDate')),
    __param(3, (0, common_1.Query)('limit')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String, String, String, Object]),
    __metadata("design:returntype", Promise)
], AccountsController.prototype, "getStatementRest", null);
exports.AccountsController = AccountsController = __decorate([
    (0, swagger_1.ApiTags)('Accounts'),
    (0, swagger_1.ApiBearerAuth)(),
    (0, swagger_1.ApiHeader)({ name: 'x-api-key', required: false, description: 'Optional API key authentication' }),
    (0, common_1.UseGuards)(combined_auth_guard_1.CombinedAuthGuard),
    (0, common_1.Controller)('api/v1/accounts'),
    __metadata("design:paramtypes", [accounts_service_1.AccountsService])
], AccountsController);
//# sourceMappingURL=accounts.controller.js.map