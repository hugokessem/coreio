import {
  Controller,
  Post,
  Body,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { PaymentsService } from './payments.service';
import {
  BillPaymentDto,
  BillPaymentVerifyDto,
  SplitPaymentDto,
  RtgsPaymentDto,
} from './dto/payment.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';

@ApiTags('Payments')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard)
@Controller('api/v1/payments')
export class PaymentsController {
  constructor(private readonly paymentsService: PaymentsService) {}

  @Post('bill/verify')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Pre-verify bill payment parameters, fee tariff, and customer balance (Op 41: BillPaymentVerify)' })
  @ApiResponse({ status: 200, description: 'Bill payment verified' })
  async verifyBill(@Body() dto: BillPaymentVerifyDto) {
    return this.paymentsService.verifyBill(dto);
  }

  @Post('bill')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Process bill payment to utility/merchant (Op 40: BillPayment)' })
  @ApiResponse({ status: 200, description: 'Bill payment posted' })
  async payBill(@Body() dto: BillPaymentDto) {
    return this.paymentsService.payBill(dto);
  }

  @Post('split')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Atomic split payment distribution from 1 source to multiple destinations (Op 42: SplitPayment)' })
  @ApiResponse({ status: 200, description: 'All split transfer legs posted atomically' })
  async splitPayment(@Body() dto: SplitPaymentDto) {
    return this.paymentsService.splitPayment(dto);
  }

  @Post('rtgs')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Interbank RTGS payment instruction (Op 43: RTGS - Simulated Sandbox Settlement)' })
  @ApiResponse({ status: 200, description: 'RTGS debit posted and clearing simulated' })
  async rtgsPayment(@Body() dto: RtgsPaymentDto) {
    return this.paymentsService.rtgsPayment(dto);
  }
}
