import {
  Controller,
  Post,
  Body,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { IpsService } from './ips.service';
import {
  IpsAccountLookupDto,
  IpsTransferDto,
  IpsQrPaymentDto,
  IpsStatusDto,
} from './dto/ips.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';

@ApiTags('IPS (EthSwitch Interbank)')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard)
@Controller('api/v1/ips')
export class IpsController {
  constructor(private readonly ipsService: IpsService) {}

  @Post('account-lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'IPS Interbank ISO 20022 account & title lookup (Op 50: IPS:AccountLookup)' })
  async accountLookup(@Body() dto: IpsAccountLookupDto) {
    return this.ipsService.accountLookup(dto);
  }

  @Post('transfer')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Execute outgoing/incoming ISO 20022 IPS interbank fund transfer (Op 51: IPS:FundTransfer)' })
  async transfer(@Body() dto: IpsTransferDto) {
    return this.ipsService.executeTransfer(dto);
  }

  @Post('qr-payment')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Process merchant QR payment via IPS switch (Op 53: IPS:QRPayment)' })
  async qrPayment(@Body() dto: IpsQrPaymentDto) {
    return this.ipsService.qrPayment(dto);
  }

  @Post('status')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Inquire clearing status of an IPS transaction (Op 52: IPS:StatusCheck)' })
  async status(@Body() dto: IpsStatusDto) {
    return this.ipsService.getStatus(dto);
  }

  @Post('status-check')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Inquire clearing status of an IPS transaction - CoreIO Route Alias (Op 52: IPS:StatusCheck)' })
  async statusCheck(@Body() dto: IpsStatusDto) {
    return this.ipsService.getStatus(dto);
  }
}
