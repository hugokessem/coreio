import {
  Controller,
  Post,
  Body,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { WalletService } from './wallet.service';
import {
  WalletCustomerLookupDto,
  AccountToWalletTransferDto,
  WalletToAccountTransferDto,
  WalletAgentLookupDto,
  WalletAgentTransferDto,
} from './dto/wallet.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';

@ApiTags('CBE Birr Wallet')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard)
@Controller('api/v1/wallet')
export class WalletController {
  constructor(private readonly walletService: WalletService) {}

  @Post('customer-lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Lookup customer profile and linked accounts by CBE Birr phone (Op 54: Wallet:CustomerAccountLookup)' })
  async customerLookup(@Body() dto: WalletCustomerLookupDto) {
    return this.walletService.customerLookup(dto);
  }

  @Post('transfer')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Transfer funds from bank account to CBE Birr mobile wallet - Canonical (Op 55: Wallet:CustomerFundTransfer)' })
  async transfer(@Body() dto: AccountToWalletTransferDto) {
    return this.walletService.transferToWallet(dto);
  }

  @Post('transfer-to-wallet')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Transfer funds from bank account to CBE Birr mobile wallet - Legacy Alias (Op 55: Wallet:CustomerFundTransfer)' })
  async transferToWallet(@Body() dto: AccountToWalletTransferDto) {
    return this.walletService.transferToWallet(dto);
  }

  @Post('transfer-from-wallet')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Transfer funds from CBE Birr mobile wallet to bank account' })
  async transferFromWallet(@Body() dto: WalletToAccountTransferDto) {
    return this.walletService.transferFromWallet(dto);
  }

  @Post('agent-lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Lookup CBE Birr registered agent details and float status (Op 56: Wallet:AgentAccountLookup)' })
  async agentLookup(@Body() dto: WalletAgentLookupDto) {
    return this.walletService.agentLookup(dto);
  }

  @Post('agent-transfer')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Disburse liquidity float to CBE Birr Agent wallet (Op 57: Wallet:AgentFundTransfer)' })
  async agentTransfer(@Body() dto: WalletAgentTransferDto) {
    return this.walletService.agentTransfer(dto);
  }
}
