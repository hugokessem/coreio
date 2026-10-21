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
exports.CustomersService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const typeorm_2 = require("typeorm");
const customer_entity_1 = require("../../database/entities/customer.entity");
const account_entity_1 = require("../../database/entities/account.entity");
const account_balance_entity_1 = require("../../database/entities/account-balance.entity");
const branch_entity_1 = require("../../database/entities/branch.entity");
const currency_entity_1 = require("../../database/entities/currency.entity");
const account_type_entity_1 = require("../../database/entities/account-type.entity");
const account_party_entity_1 = require("../../database/entities/account-party.entity");
const transaction_entity_1 = require("../../database/entities/transaction.entity");
const ledger_entry_entity_1 = require("../../database/entities/ledger-entry.entity");
const domain_exceptions_1 = require("../../common/exceptions/domain.exceptions");
const financial_math_util_1 = require("../../common/utils/financial-math.util");
let CustomersService = class CustomersService {
    dataSource;
    customerRepository;
    accountRepository;
    branchRepository;
    currencyRepository;
    accountTypeRepository;
    constructor(dataSource, customerRepository, accountRepository, branchRepository, currencyRepository, accountTypeRepository) {
        this.dataSource = dataSource;
        this.customerRepository = customerRepository;
        this.accountRepository = accountRepository;
        this.branchRepository = branchRepository;
        this.currencyRepository = currencyRepository;
        this.accountTypeRepository = accountTypeRepository;
    }
    async lookup(dto) {
        let customer = null;
        if (dto.customerNumber) {
            customer = await this.customerRepository.findOne({
                where: { customerNumber: dto.customerNumber },
                relations: ['accounts'],
            });
        }
        else if (dto.phoneNumber) {
            customer = await this.customerRepository.findOne({
                where: { phoneNumber: dto.phoneNumber },
                relations: ['accounts'],
            });
        }
        else if (dto.nationalId) {
            customer = await this.customerRepository.findOne({
                where: { nationalId: dto.nationalId },
                relations: ['accounts'],
            });
        }
        else if (dto.accountNumber) {
            const account = await this.accountRepository.findOne({
                where: { accountNumber: dto.accountNumber },
                relations: ['customer', 'customer.accounts'],
            });
            if (account) {
                customer = account.customer;
            }
        }
        else {
            throw new common_1.BadRequestException('At least one lookup parameter must be provided');
        }
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber || dto.phoneNumber || dto.accountNumber || dto.nationalId || 'unknown');
        }
        return this.mapToResponse(customer);
    }
    async getDetail(dto) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber: dto.customerNumber },
            relations: ['accounts', 'accounts.branch', 'accounts.currency', 'accounts.accountType', 'limits'],
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
        }
        return this.mapToResponse(customer, true);
    }
    async phoneLookup(dto) {
        const customer = await this.customerRepository.findOne({
            where: { phoneNumber: dto.phoneNumber },
            relations: ['accounts'],
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(dto.phoneNumber);
        }
        return this.mapToResponse(customer);
    }
    async create(dto, creatorUsername = 'SYSTEM') {
        const count = await this.customerRepository.count();
        const customerNumber = (10000001 + count).toString();
        const fullName = dto.fullName ||
            [dto.firstName, dto.middleName, dto.lastName].filter(Boolean).join(' ') ||
            dto.legalName ||
            'Customer';
        const address = dto.address || [dto.addressLine1, dto.city].filter(Boolean).join(', ') || null;
        const customer = this.customerRepository.create({
            ...dto,
            fullName,
            address,
            customerNumber,
            status: customer_entity_1.CustomerStatus.ACTIVE,
            amlStatus: customer_entity_1.AMLStatus.CLEAR,
            kycLevel: 'LEVEL_2',
            isAuthorized: false,
            createdBy: creatorUsername,
            authorizedBy: null,
            authorizedAt: null,
        });
        const saved = await this.customerRepository.save(customer);
        return this.mapToResponse(saved);
    }
    async authorizeCustomer(customerNumber, authorizerUsername) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber },
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(customerNumber);
        }
        if (customer.createdBy && customer.createdBy === authorizerUsername) {
            throw new common_1.BadRequestException('Dual control violation: Maker cannot authorize their own registered customer record (maker-checker policy)');
        }
        customer.isAuthorized = true;
        customer.authorizedBy = authorizerUsername;
        customer.authorizedAt = new Date();
        await this.customerRepository.save(customer);
        return {
            customerNumber: customer.customerNumber,
            isAuthorized: true,
            authorizedBy: customer.authorizedBy,
            authorizedAt: customer.authorizedAt,
            message: 'Customer successfully authorized under maker-checker dual control policy',
        };
    }
    async checkExistence(dto) {
        let exists = false;
        let matchType = null;
        let customerNumber = null;
        if (dto.phoneNumber) {
            const byPhone = await this.customerRepository.findOne({ where: { phoneNumber: dto.phoneNumber } });
            if (byPhone) {
                exists = true;
                matchType = 'PHONE_NUMBER';
                customerNumber = byPhone.customerNumber;
            }
        }
        if (!exists && dto.nationalId) {
            const byNid = await this.customerRepository.findOne({ where: { nationalId: dto.nationalId } });
            if (byNid) {
                exists = true;
                matchType = 'NATIONAL_ID';
                customerNumber = byNid.customerNumber;
            }
        }
        return {
            exists,
            matchType,
            customerNumber,
            phoneNumber: dto.phoneNumber || null,
            nationalId: dto.nationalId || null,
            canProceedWithOnboarding: !exists,
        };
    }
    async checkCustomerEligibility(dto) {
        if (dto.customerNumber) {
            const customer = await this.customerRepository.findOne({
                where: { customerNumber: dto.customerNumber },
                relations: ['accounts'],
            });
            if (!customer)
                throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
            const reasons = [];
            if (customer.status !== customer_entity_1.CustomerStatus.ACTIVE) {
                reasons.push(`Customer status is ${customer.status}`);
            }
            if (customer.amlStatus !== customer_entity_1.AMLStatus.CLEAR) {
                reasons.push(`AML screening flag: ${customer.amlStatus}`);
            }
            const eligible = reasons.length === 0;
            return {
                customerNumber: customer.customerNumber,
                customerName: customer.fullName,
                eligible,
                isEligible: eligible,
                amlStatus: customer.amlStatus,
                kycLevel: customer.kycLevel,
                channel: dto.channel || 'MOBILE_BANKING',
                reasons,
            };
        }
        return {
            customerNumber: null,
            customerName: null,
            eligible: true,
            isEligible: true,
            amlStatus: 'CLEAR',
            channel: dto.channel || 'MOBILE_BANKING',
            reasons: [],
        };
    }
    async verifyFayda(dto) {
        const nid = dto.nationalId || dto.faydaId;
        const customer = nid
            ? await this.customerRepository.findOne({ where: { nationalId: nid } })
            : null;
        return {
            nationalId: nid || null,
            customerId: customer?.customerNumber || null,
            customerNumber: customer?.customerNumber || null,
            customerName: customer?.fullName || dto.fullName || 'Verified National ID Holder',
            dateOfBirth: customer?.dateOfBirth || dto.dateOfBirth || null,
            isVerified: true,
            verified: true,
            customerFlag: customer ? 'FOUND' : 'EXTERNAL_REGISTRY_MATCH',
            simulation: true,
            source: 'LOCAL_IDENTITY_PERSISTENCE_STORE',
            message: 'Fayda National Identity validated against local identity persistence store (sandbox simulation)',
        };
    }
    async compositeOnboard(dto, creatorUsername = 'SYSTEM') {
        return await this.dataSource.transaction(async (manager) => {
            const branch = await manager.findOne(branch_entity_1.Branch, { where: { branchCode: dto.branchCode } });
            if (!branch)
                throw new common_1.BadRequestException(`Invalid branch code '${dto.branchCode}'`);
            const currency = await manager.findOne(currency_entity_1.Currency, { where: { code: dto.currencyCode } });
            if (!currency)
                throw new common_1.BadRequestException(`Invalid currency '${dto.currencyCode}'`);
            const accType = await manager.findOne(account_type_entity_1.AccountType, { where: { code: dto.accountTypeCode } });
            if (!accType)
                throw new common_1.BadRequestException(`Invalid account type '${dto.accountTypeCode}'`);
            const custCount = await manager.count(customer_entity_1.Customer);
            const customerNumber = (10000001 + custCount).toString();
            const customer = manager.create(customer_entity_1.Customer, {
                customerNumber,
                customerType: dto.customerType,
                firstName: dto.firstName || null,
                middleName: dto.middleName || null,
                lastName: dto.lastName || null,
                legalName: [dto.firstName, dto.middleName, dto.lastName].filter(Boolean).join(' '),
                phoneNumber: dto.phoneNumber,
                nationalId: dto.nationalId || null,
                status: customer_entity_1.CustomerStatus.ACTIVE,
                amlStatus: customer_entity_1.AMLStatus.CLEAR,
                kycLevel: 'LEVEL_2',
                isAuthorized: true,
                createdBy: creatorUsername,
                authorizedBy: creatorUsername,
                authorizedAt: new Date(),
            });
            const savedCustomer = await manager.save(customer_entity_1.Customer, customer);
            const accCount = await manager.count(account_entity_1.Account);
            const accountNumber = `1000${(100000000 + accCount + 1).toString()}`;
            const account = manager.create(account_entity_1.Account, {
                accountNumber,
                customerId: savedCustomer.id,
                branchId: branch.id,
                currencyCode: currency.code,
                accountTypeId: accType.id,
                accountName: savedCustomer.legalName || 'Account Holder',
                status: account_entity_1.AccountStatus.ACTIVE,
            });
            const savedAccount = await manager.save(account_entity_1.Account, account);
            const party = manager.create(account_party_entity_1.AccountParty, {
                accountId: savedAccount.id,
                customerId: savedCustomer.id,
                partyRole: account_party_entity_1.PartyRole.PRIMARY_OWNER,
                authorizationLevel: 'SOLE_SIGNATORY',
                status: 'ACTIVE',
            });
            await manager.save(account_party_entity_1.AccountParty, party);
            const depositStr = dto.initialDeposit ? financial_math_util_1.FinancialMath.add(dto.initialDeposit, '0') : '0.0000';
            const balance = manager.create(account_balance_entity_1.AccountBalance, {
                accountId: savedAccount.id,
                ledgerBalance: depositStr,
                availableBalance: depositStr,
                lockedBalance: '0.0000',
                currencyCode: currency.code,
            });
            await manager.save(account_balance_entity_1.AccountBalance, balance);
            if (financial_math_util_1.FinancialMath.isPositive(depositStr)) {
                const tx = manager.create(transaction_entity_1.Transaction, {
                    transactionReference: `INIT-${accountNumber}`,
                    transactionType: transaction_entity_1.TransactionType.CASH_DEPOSIT,
                    channel: transaction_entity_1.TransactionChannel.SYSTEM,
                    destinationAccountId: savedAccount.id,
                    amount: depositStr,
                    currencyCode: currency.code,
                    narration: 'Composite Onboarding Opening Deposit',
                    status: transaction_entity_1.TransactionStatus.COMPLETED,
                    completedAt: new Date(),
                });
                await manager.save(transaction_entity_1.Transaction, tx);
                await manager.save(ledger_entry_entity_1.LedgerEntry, manager.create(ledger_entry_entity_1.LedgerEntry, {
                    transactionId: tx.id,
                    accountId: savedAccount.id,
                    entryType: ledger_entry_entity_1.EntryType.CREDIT,
                    amount: depositStr,
                    currencyCode: currency.code,
                    balanceAfter: depositStr,
                    narration: 'Opening Deposit',
                }));
            }
            return {
                customerNumber: savedCustomer.customerNumber,
                customerName: savedCustomer.legalName,
                accountNumber: savedAccount.accountNumber,
                currency: currency.code,
                branchCode: branch.branchCode,
                accountTypeCode: accType.code,
                initialBalance: depositStr,
                status: 'ONBOARDED_SUCCESSFULLY',
                timestamp: new Date().toISOString(),
            };
        });
    }
    async amlLookupByAccount(accountNumber) {
        const account = await this.accountRepository.findOne({
            where: { accountNumber },
            relations: ['customer'],
        });
        if (!account)
            throw new domain_exceptions_1.AccountNotFoundException(accountNumber);
        const customer = account.customer;
        const isCleared = customer.amlStatus === customer_entity_1.AMLStatus.CLEAR;
        return {
            accountNumber: account.accountNumber,
            customerNumber: customer.customerNumber,
            customerName: customer.fullName,
            amlStatus: customer.amlStatus,
            isCleared,
            passedScreening: isCleared,
            screenedAt: new Date().toISOString(),
            evaluation: isCleared ? 'PASS' : 'REVIEW_REQUIRED',
        };
    }
    async fetchAll(page = 1, limit = 50) {
        const [customers, total] = await this.customerRepository.findAndCount({
            skip: (page - 1) * limit,
            take: limit,
            order: { createdAt: 'DESC' },
            relations: ['accounts'],
        });
        return {
            data: customers.map((c) => this.mapToResponse(c)),
            meta: {
                page,
                limit,
                total,
                totalPages: Math.ceil(total / limit),
            },
        };
    }
    async deleteCustomer(customerNumber) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber },
            relations: ['accounts'],
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(customerNumber);
        }
        if (customer.accounts && customer.accounts.length > 0) {
            throw new common_1.BadRequestException('Cannot delete customer with existing active or historical accounts');
        }
        await this.customerRepository.remove(customer);
        return {
            success: true,
            message: `Customer ${customerNumber} successfully deleted`,
        };
    }
    async amlVerify(dto) {
        const customer = await this.customerRepository.findOne({
            where: { customerNumber: dto.customerNumber },
        });
        if (!customer) {
            throw new domain_exceptions_1.CustomerNotFoundException(dto.customerNumber);
        }
        const isCleared = customer.amlStatus === customer_entity_1.AMLStatus.CLEAR;
        return {
            customerNumber: customer.customerNumber,
            customerName: customer.fullName,
            amlStatus: customer.amlStatus,
            isCleared,
            screenedAt: new Date().toISOString(),
            flags: isCleared ? [] : ['Suspicious transaction profile or sanctioned entity flag'],
        };
    }
    mapToResponse(c, includeFullRelations = false) {
        return {
            customerNumber: c.customerNumber,
            customerType: c.customerType,
            fullName: c.fullName,
            firstName: c.firstName,
            middleName: c.middleName,
            lastName: c.lastName,
            legalName: c.legalName,
            phoneNumber: c.phoneNumber,
            email: c.email,
            address: c.address,
            dateOfBirth: c.dateOfBirth,
            gender: c.gender,
            nationalId: c.nationalId,
            tin: c.tin,
            kycLevel: c.kycLevel,
            amlStatus: c.amlStatus,
            status: c.status,
            isAuthorized: c.isAuthorized,
            authorizedBy: c.authorizedBy,
            authorizedAt: c.authorizedAt,
            accounts: c.accounts
                ? c.accounts.map((a) => ({
                    accountNumber: a.accountNumber,
                    accountName: a.accountName,
                    currency: a.currencyCode,
                    status: a.status,
                    branchCode: a.branch?.branchCode || null,
                }))
                : [],
            createdAt: c.createdAt,
            updatedAt: c.updatedAt,
        };
    }
};
exports.CustomersService = CustomersService;
exports.CustomersService = CustomersService = __decorate([
    (0, common_1.Injectable)(),
    __param(1, (0, typeorm_1.InjectRepository)(customer_entity_1.Customer)),
    __param(2, (0, typeorm_1.InjectRepository)(account_entity_1.Account)),
    __param(3, (0, typeorm_1.InjectRepository)(branch_entity_1.Branch)),
    __param(4, (0, typeorm_1.InjectRepository)(currency_entity_1.Currency)),
    __param(5, (0, typeorm_1.InjectRepository)(account_type_entity_1.AccountType)),
    __metadata("design:paramtypes", [typeorm_2.DataSource,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository,
        typeorm_2.Repository])
], CustomersService);
//# sourceMappingURL=customers.service.js.map