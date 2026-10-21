import {
  Controller,
  Post,
  Get,
  Delete,
  Body,
  Param,
  Query,
  HttpCode,
  HttpStatus,
  UseGuards,
  Req,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { CustomersService } from './customers.service';
import {
  CustomerLookupDto,
  CustomerDetailDto,
  CreateCustomerDto,
  CustomerPhoneLookupDto,
  AmlVerifyDto,
  CheckExistenceDto,
  CustomerEligibilityDto,
  FaydaVerifyDto,
  CompositeOnboardDto,
} from './dto/customer.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';
import { RolesGuard } from '../../common/guards/roles.guard';
import { Roles } from '../../common/decorators/roles.decorator';
import { UserRole } from '../../database/entities/user.entity';

@ApiTags('Customers & Identity')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard, RolesGuard)
@Controller('api/v1')
export class CustomersController {
  constructor(private readonly customersService: CustomersService) {}

  @Post('customers/lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Lookup customer by customer number, phone, NID, or account number' })
  @ApiResponse({ status: 200, description: 'Customer details retrieved' })
  async lookup(@Body() dto: CustomerLookupDto) {
    return this.customersService.lookup(dto);
  }

  @Post('customers/detail')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Get full customer detail by customer number' })
  async detail(@Body() dto: CustomerDetailDto) {
    return this.customersService.getDetail(dto);
  }

  @Post('customers/phone-lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Lookup customer by registered phone number' })
  async phoneLookup(@Body() dto: CustomerPhoneLookupDto) {
    return this.customersService.phoneLookup(dto);
  }

  @Post(['customers/create', 'customers'])
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({ summary: 'Register/create new customer record' })
  async create(@Body() dto: CreateCustomerDto, @Req() req: any) {
    const creator = req.user?.username || 'SYSTEM';
    return this.customersService.create(dto, creator);
  }

  @Post('customers/fetch')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Fetch paginated list of customers' })
  async fetch(@Body() body: { page?: number; limit?: number }) {
    return this.customersService.fetchAll(body.page || 1, body.limit || 50);
  }

  @Roles(UserRole.ADMIN)
  @Post('customers/delete')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Delete customer record (Requires ADMIN role)' })
  async deletePost(@Body() body: { customerNumber: string }) {
    return this.customersService.deleteCustomer(body.customerNumber);
  }

  @Post(['customers/aml-verify', 'customers/:customerNumber/verify-aml'])
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Verify customer AML compliance' })
  async amlVerify(@Body() dto: AmlVerifyDto, @Param('customerNumber') paramNumber?: string) {
    const customerNumber = paramNumber || dto.customerNumber;
    return this.customersService.amlVerify({ customerNumber });
  }

  @Roles(UserRole.TELLER, UserRole.ADMIN)
  @Post('customers/:customerNumber/authorize')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Maker-checker approval of a customer record (Enforces dual control)' })
  async authorizeCustomer(
    @Param('customerNumber') customerNumber: string,
    @Body() body: { authorizerUsername?: string },
    @Req() req: any,
  ) {
    const authorizer = body?.authorizerUsername || req.user?.username || 'ADMIN';
    return this.customersService.authorizeCustomer(customerNumber, authorizer);
  }

  @Post('customers/check-existence')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Pre-onboarding existence check: validates if phone or national ID is registered' })
  async checkExistence(@Body() dto: CheckExistenceDto) {
    return this.customersService.checkExistence(dto);
  }

  @Post('customers/eligibility')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Check customer eligibility for digital channels and services' })
  async checkEligibility(@Body() dto: CustomerEligibilityDto) {
    return this.customersService.checkCustomerEligibility(dto);
  }

  @Post('identity/fayda/verify')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Verify national citizen identity against local Fayda identity persistence store' })
  async verifyFayda(@Body() dto: FaydaVerifyDto) {
    return this.customersService.verifyFayda(dto);
  }

  @Post('accounts/onboard')
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({ summary: 'Composite onboarding workflow: Customer registration, AML check, and account opening' })
  async compositeOnboard(@Body() dto: CompositeOnboardDto, @Req() req: any) {
    const creator = req.user?.username || 'SYSTEM';
    return this.customersService.compositeOnboard(dto, creator);
  }

  @Post('accounts/:accountNumber/aml-lookup')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Check AML status associated with an account holder' })
  async amlLookupByAccount(@Param('accountNumber') accountNumber: string) {
    return this.customersService.amlLookupByAccount(accountNumber);
  }

  // RESTful standard conventions
  @Get('customers')
  @ApiOperation({ summary: 'GET all customers (REST convention)' })
  async getAll(@Query('page') page = '1', @Query('limit') limit = '50') {
    return this.customersService.fetchAll(parseInt(page, 10), parseInt(limit, 10));
  }

  @Get('customers/:customerNumber')
  @ApiOperation({ summary: 'GET customer by customer number (REST convention)' })
  async getOne(@Param('customerNumber') customerNumber: string) {
    return this.customersService.getDetail({ customerNumber });
  }

  @Roles(UserRole.ADMIN)
  @Delete('customers/:customerNumber')
  @ApiOperation({ summary: 'DELETE customer by customer number (Requires ADMIN role)' })
  async deleteOne(@Param('customerNumber') customerNumber: string) {
    return this.customersService.deleteCustomer(customerNumber);
  }
}
