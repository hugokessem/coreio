import {
  Controller,
  Post,
  Get,
  Put,
  Body,
  Param,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { SuperAppService } from './superapp.service';
import {
  SuperAppSubscribeDto,
  SuperAppUnsubscribeDto,
  SuperAppStatusDto,
  SuperAppStatusChangeDto,
} from './dto/superapp.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';

@ApiTags('SuperApp')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard)
@Controller('api/v1/superapp')
export class SuperAppController {
  constructor(private readonly superAppService: SuperAppService) {}

  @Post('subscribe')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Link/subscribe external partner user to bank customer profile (Op 47: SuperAppSubscribe)' })
  async subscribe(@Body() dto: SuperAppSubscribeDto) {
    return this.superAppService.subscribe(dto);
  }

  @Post('unsubscribe')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Deactivate external partner subscription link (Op 48: SuperAppUnsubscribe)' })
  async unsubscribe(@Body() dto: SuperAppUnsubscribeDto) {
    return this.superAppService.unsubscribe(dto);
  }

  @Put('status')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Update status of external partner subscription - Canonical REST (Op 49: SuperAppStatusChange)' })
  async changeStatusPut(@Body() dto: SuperAppStatusChangeDto) {
    return this.superAppService.changeStatus(dto);
  }

  @Post('status/change')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Update status of external partner subscription - Alias (Op 49: SuperAppStatusChange)' })
  async changeStatusPost(@Body() dto: SuperAppStatusChangeDto) {
    return this.superAppService.changeStatus(dto);
  }

  @Post('status')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Inquire status of external partner user link - Body Query' })
  async status(@Body() dto: SuperAppStatusDto) {
    return this.superAppService.getStatus(dto);
  }

  @Get('status/:externalUserId')
  @ApiOperation({ summary: 'Inquire status of external partner user link - Route Param' })
  async getStatusParam(@Param('externalUserId') externalUserId: string) {
    return this.superAppService.getByExternalUserId(externalUserId);
  }
}
