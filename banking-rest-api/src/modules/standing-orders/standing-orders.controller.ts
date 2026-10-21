import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Body,
  Param,
  Query,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { StandingOrdersService } from './standing-orders.service';
import {
  CreateStandingOrderDto,
  ListStandingOrdersDto,
  CancelStandingOrderDto,
  UpdateStandingOrderDto,
  ListStandingOrderHistoryDto,
} from './dto/standing-order.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';
import { RolesGuard } from '../../common/guards/roles.guard';
import { Roles } from '../../common/decorators/roles.decorator';
import { UserRole } from '../../common/enums/user-role.enum';

@ApiTags('Standing Orders')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard, RolesGuard)
@Controller('api/v1/standing-orders')
export class StandingOrdersController {
  constructor(private readonly standingOrdersService: StandingOrdersService) {}

  @Post()
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({ summary: 'Setup recurring standing order instruction - Canonical (Op 30: CreateStandingOrder)' })
  async createCanonical(@Body() dto: CreateStandingOrderDto) {
    return this.standingOrdersService.create(dto);
  }

  @Post('create')
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({ summary: 'Setup recurring standing order instruction - Legacy Alias (Op 30: CreateStandingOrder)' })
  async create(@Body() dto: CreateStandingOrderDto) {
    return this.standingOrdersService.create(dto);
  }

  @Get()
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'List all standing orders for an account - Canonical (Op 31: ListStandingOrder)' })
  async listQuery(@Query('accountNumber') accountNumber: string) {
    return this.standingOrdersService.listByAccount(accountNumber);
  }

  @Post('list')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'List all standing orders for an account - Legacy Alias (Op 31: ListStandingOrder)' })
  async list(@Body() dto: ListStandingOrdersDto) {
    return this.standingOrdersService.list(dto);
  }

  @Get(':orderReference/history')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Query execution history of a standing order (Op 32: ListStandingOrderHistory)' })
  async getHistory(@Param('orderReference') orderReference: string) {
    return this.standingOrdersService.getHistory(orderReference);
  }

  @Post('history')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Query execution history of a standing order - Body Alias (Op 32: ListStandingOrderHistory)' })
  async getHistoryBody(@Body() dto: ListStandingOrderHistoryDto) {
    return this.standingOrdersService.getHistory(dto.orderReference);
  }

  @Put(':orderReference')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Update parameters of an active standing order (Op 33: UpdateStandingOrder)' })
  async update(
    @Param('orderReference') orderReference: string,
    @Body() dto: UpdateStandingOrderDto,
  ) {
    return this.standingOrdersService.update(orderReference, dto);
  }

  @Delete(':orderReference')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Cancel an active standing order - Canonical (Op 34: CancleStandingOrder)' })
  async deleteCanonical(@Param('orderReference') orderReference: string) {
    return this.standingOrdersService.cancel(orderReference);
  }

  @Post('cancel')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Cancel an active standing order instruction - Legacy Alias (Op 34: CancleStandingOrder)' })
  async cancel(@Body() dto: CancelStandingOrderDto) {
    return this.standingOrdersService.cancel(dto);
  }

  @Post('execute-due')
  @HttpCode(HttpStatus.OK)
  @Roles(UserRole.ADMIN)
  @ApiOperation({ summary: 'Execute all due standing orders with FOR UPDATE SKIP LOCKED concurrency' })
  async executeDue() {
    return this.standingOrdersService.executeDueOrders();
  }
}
