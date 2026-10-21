import {
  Controller,
  Post,
  Get,
  Body,
  Param,
  HttpCode,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse, ApiBearerAuth, ApiHeader } from '@nestjs/swagger';
import { CardsService } from './cards.service';
import { RequestCardDto, ReplaceCardDto, ListCardsDto } from './dto/card.dto';
import { CombinedAuthGuard } from '../../common/guards/combined-auth.guard';

@ApiTags('Cards')
@ApiBearerAuth()
@ApiHeader({ name: 'x-api-key', required: false, description: 'Optional API key authentication' })
@UseGuards(CombinedAuthGuard)
@Controller('api/v1/cards')
export class CardsController {
  constructor(private readonly cardsService: CardsService) {}

  @Post('request')
  @HttpCode(HttpStatus.CREATED)
  @ApiOperation({ summary: 'Submit debit or Mastercard card issuance request' })
  async requestCard(@Body() dto: RequestCardDto) {
    return this.cardsService.requestCard(dto);
  }

  @Post('replace')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Block compromised card and request immediate replacement' })
  async replaceCard(@Body() dto: ReplaceCardDto) {
    return this.cardsService.replaceCard(dto);
  }

  @Post('list')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'List cards by account or customer number' })
  async listCards(@Body() dto: ListCardsDto) {
    return this.cardsService.listCards(dto);
  }

  @Get('account/:accountNumber')
  @ApiOperation({ summary: 'GET active cards for account (REST convention)' })
  async getByAccount(@Param('accountNumber') accountNumber: string) {
    return this.cardsService.listCards({ accountNumber });
  }
}
