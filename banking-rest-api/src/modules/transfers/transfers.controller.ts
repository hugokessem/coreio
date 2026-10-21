import {
  Controller,
  Get,
  Post,
  Body,
  Param,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { TransfersService } from './transfers.service';
import { ReconciliationService } from './reconciliation.service';
import {
  TransferVerifyDto,
  ExecuteTransferDto,
  TransferFromHoldDto,
  TransactionStatusDto,
  ReversalDto,
} from './dto/transfer.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';
import { RolesGuard } from '../../common/guards/roles.guard';
import { Roles } from '../../common/decorators/roles.decorator';
import { UserRole } from '../../common/enums/user-role.enum';

@ApiTags('Transfers')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard, RolesGuard)
@Controller('api/v1/transfers')
export class TransfersController {
  constructor(
    private readonly transfersService: TransfersService,
    private readonly reconciliationService: ReconciliationService,
  ) {}

  @Post('verify')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Pre-transfer verification, fee calculation, and balance sufficiency check (Op 23: FundTransferVerify)' })
  @ApiResponse({ status: 200, description: 'Transfer verified' })
  async verify(@Body() dto: TransferVerifyDto) {
    return this.transfersService.verify(dto);
  }

  @Post()
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Execute funds transfer - Canonical REST route (Op 24: FundTransfer)' })
  @ApiResponse({ status: 200, description: 'Transfer successfully posted' })
  @ApiResponse({ status: 422, description: 'Insufficient funds' })
  async executeCanonical(@Body() dto: ExecuteTransferDto) {
    return this.transfersService.execute(dto);
  }

  @Post('execute')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Execute funds transfer with double-entry ledger, row locking, and idempotency (Op 24: FundTransfer alias)' })
  @ApiResponse({ status: 200, description: 'Transfer successfully posted' })
  @ApiResponse({ status: 422, description: 'Insufficient funds' })
  async execute(@Body() dto: ExecuteTransferDto) {
    return this.transfersService.execute(dto);
  }

  @Get(':transactionReference')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Query transfer details and verification by reference (Op 25: FundTransferCheck)' })
  async getByReference(@Param('transactionReference') transactionReference: string) {
    return this.transfersService.getByReference(transactionReference);
  }

  @Get('status/:transactionReference')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Query transaction status by transaction reference (Op 26: StatusCheck GET)' })
  async getStatusByParam(@Param('transactionReference') transactionReference: string) {
    return this.transfersService.getByReference(transactionReference);
  }

  @Post('status')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Query transaction status by body reference (Op 26: StatusCheck POST)' })
  async status(@Body() dto: TransactionStatusDto) {
    return this.transfersService.getStatus(dto);
  }

  @Post('from-hold')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Settle and execute transfer from an active pre-allocated hold (Op 44: LockedAmountFT)' })
  async transferFromHold(@Body() dto: TransferFromHoldDto) {
    return this.transfersService.transferFromHold(dto);
  }

  @Post('reverse')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Perform compensating reversal of a completed transfer (Op 27: RevertFundTransfer)' })
  async reverse(@Body() dto: ReversalDto) {
    return this.transfersService.reverse(dto);
  }

  @Post('revert')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Perform compensating reversal of a completed transfer - Alias (Op 27: RevertFundTransfer)' })
  async revert(@Body() dto: ReversalDto) {
    return this.transfersService.reverse(dto);
  }

  @Post('reconcile')
  @HttpCode(HttpStatus.OK)
  @Roles(UserRole.ADMIN)
  @ApiOperation({ summary: 'Run system-wide double-entry ledger vs account balance reconciliation audit' })
  async reconcile() {
    return this.reconciliationService.reconcileAllAccounts();
  }
}
