import {
  Controller,
  Post,
  Get,
  Put,
  Body,
  Param,
  Query,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { ServicesService } from './services.service';
import {
  ServiceDetailDto,
  CustomerLimitsQueryDto,
  AmendCustomerLimitDto,
} from './dto/service.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';
import { RolesGuard } from '../../common/guards/roles.guard';
import { Roles } from '../../common/decorators/roles.decorator';
import { UserRole } from '../../common/enums/user-role.enum';

@ApiTags('Services & Limits')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard, RolesGuard)
@Controller('api/v1/services')
export class ServicesController {
  constructor(private readonly servicesService: ServicesService) {}

  @Get()
  @ApiOperation({ summary: 'List all supported banking services (REST convention)' })
  async listAll() {
    return this.servicesService.listAll();
  }

  @Get('detail/:serviceCode')
  @ApiOperation({ summary: 'Get service tariff rules, channel limits, and configuration (Op 36: ServiceDetail Canonical)' })
  async getDetailByCode(@Param('serviceCode') serviceCode: string) {
    return this.servicesService.getDetailByCode(serviceCode);
  }

  @Post('detail')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Get service tariff rules, channel limits, and configuration (Op 36: ServiceDetail Legacy)' })
  async getDetail(@Body() dto: ServiceDetailDto) {
    return this.servicesService.getDetail(dto);
  }

  @Get('limits')
  @ApiOperation({ summary: 'Query tariff limits across all services (Op 37: ServiceLimit)' })
  async getServiceLimits() {
    return this.servicesService.getServiceLimits();
  }

  @Post('limits')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Query custom per-customer transaction limits (Op 38: CustomerLimitFetchByCIF)' })
  async getCustomerLimits(@Body() dto: CustomerLimitsQueryDto) {
    return this.servicesService.getCustomerLimits(dto);
  }

  @Get('limits/customer/:customerNumber')
  @ApiOperation({ summary: 'Query custom per-customer transaction limits - Canonical REST (Op 38: CustomerLimitFetchByCIF)' })
  async getCustomerLimitsCanonical(
    @Param('customerNumber') customerNumber: string,
    @Query('serviceCode') serviceCode?: string,
  ) {
    return this.servicesService.getCustomerLimitsByNumber(customerNumber, serviceCode);
  }

  @Put('limits/amend')
  @HttpCode(HttpStatus.OK)
  @Roles(UserRole.ADMIN)
  @ApiOperation({ summary: 'Amend per-customer transaction limits - Admin Dual Control (Op 39: CustomerLimitAmendByCIF)' })
  async amendCustomerLimitPut(@Body() dto: AmendCustomerLimitDto) {
    return this.servicesService.amendCustomerLimit(dto);
  }

  @Post('limits/amend')
  @HttpCode(HttpStatus.OK)
  @Roles(UserRole.ADMIN)
  @ApiOperation({ summary: 'Amend per-customer transaction limits - Legacy Alias (Op 39: CustomerLimitAmendByCIF)' })
  async amendCustomerLimitPost(@Body() dto: AmendCustomerLimitDto) {
    return this.servicesService.amendCustomerLimit(dto);
  }
}
