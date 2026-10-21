import {
  Controller,
  Post,
  Get,
  Body,
  Param,
  Query,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { AccountsService } from './accounts.service';
import {
  AccountLookupDto,
  AccountNameLookupDto,
  CreateAccountDto,
  AccountListDto,
  AccountEligibilityDto,
  CreateHoldDto,
  ReleaseHoldDto,
  AccountStatementDto,
  AccountMiniStatementDto,
} from './dto/account.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';

@ApiTags('Accounts')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard)
@Controller('api/v1/accounts')
export class AccountsController {
  constructor(private readonly accountsService: AccountsService) {}

  @Post('lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Lookup account details, balances, status, branch, and customer' })
  @ApiResponse({ status: 200, description: 'Account found' })
  @ApiResponse({ status: 404, description: 'Account not found' })
  async lookup(@Body() dto: AccountLookupDto) {
    return this.accountsService.lookup(dto);
  }

  @Post('name-lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Account title/name enquiry for transfer beneficiaries' })
  async nameLookup(@Body() dto: AccountNameLookupDto) {
    return this.accountsService.nameLookup(dto);
  }

  @Post('list')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'List all accounts belonging to a customer' })
  async list(@Body() dto: AccountListDto) {
    return this.accountsService.list(dto);
  }

  @Get()
  @ApiOperation({ summary: 'List accounts for a customer (REST convention)' })
  async listRest(@Query('customerNumber') customerNumber: string) {
    return this.accountsService.list({ customerNumber });
  }

  @Post(['create', ''])
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({ summary: 'Open new bank account with initial deposit option' })
  async create(@Body() dto: CreateAccountDto) {
    return this.accountsService.create(dto);
  }

  @Post('eligibility')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Check account eligibility for debit/credit operations' })
  async eligibility(@Body() dto: AccountEligibilityDto) {
    return this.accountsService.checkEligibility(dto);
  }

  @Post(['holds/create', 'holds'])
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({ summary: 'Place administrative, court, or pledge hold on account' })
  async createHold(@Body() dto: CreateHoldDto) {
    return this.accountsService.createHold(dto);
  }

  @Post('holds/list')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'List active holds for an account' })
  async listHoldsPost(@Body() dto: AccountLookupDto) {
    return this.accountsService.listHolds(dto.accountNumber);
  }

  @Get(':accountNumber/holds')
  @ApiOperation({ summary: 'GET active holds for account (REST convention)' })
  async listHoldsGet(@Param('accountNumber') accountNumber: string) {
    return this.accountsService.listHolds(accountNumber);
  }

  @Post('holds/release')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Release active hold, unfreezing the locked balance' })
  async releaseHold(@Body() dto: ReleaseHoldDto) {
    return this.accountsService.releaseHold(dto);
  }

  @Post('statement')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Generate detailed account ledger statement (Op 29: MiniStatementByDateRange)' })
  async statement(@Body() dto: AccountStatementDto) {
    return this.accountsService.getStatement(dto);
  }

  @Post('mini-statement-by-date-range')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Generate detailed account ledger statement - CoreIO Alias (Op 29: MiniStatementByDateRange)' })
  async statementByDateRange(@Body() dto: AccountStatementDto) {
    return this.accountsService.getStatement(dto);
  }

  @Post('mini-statement')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Generate mini-statement (recent N transactions) (Op 28: MiniStatementByLimit)' })
  async miniStatement(@Body() dto: AccountMiniStatementDto) {
    return this.accountsService.getMiniStatement(dto);
  }

  @Post('mini-statement-by-limit')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Generate mini-statement - CoreIO Alias (Op 28: MiniStatementByLimit)' })
  async miniStatementByLimit(@Body() dto: AccountMiniStatementDto) {
    return this.accountsService.getMiniStatement(dto);
  }

  // RESTful routes
  @Get(':accountNumber')
  @ApiOperation({ summary: 'GET account details by account number (REST convention)' })
  async getByNumber(@Param('accountNumber') accountNumber: string) {
    return this.accountsService.lookup({ accountNumber });
  }

  @Get(':accountNumber/balance')
  @ApiOperation({ summary: 'GET account balance by account number (REST convention)' })
  async getBalance(@Param('accountNumber') accountNumber: string) {
    const acc = await this.accountsService.lookup({ accountNumber });
    return {
      accountNumber: acc.accountNumber,
      currency: acc.currency,
      ledgerBalance: acc.ledgerBalance,
      availableBalance: acc.availableBalance,
      lockedBalance: acc.lockedBalance,
    };
  }

  @Get(':accountNumber/statement')
  @ApiOperation({ summary: 'GET statement by account number (REST convention)' })
  async getStatementRest(
    @Param('accountNumber') accountNumber: string,
    @Query('startDate') startDate?: string,
    @Query('endDate') endDate?: string,
    @Query('limit') limit = '50',
  ) {
    return this.accountsService.getStatement({
      accountNumber,
      startDate,
      endDate,
      limit: parseInt(limit, 10),
    });
  }
}
