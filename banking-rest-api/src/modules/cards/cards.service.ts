import { Injectable, BadRequestException, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import {
  CardRequest,
  CardRequestType,
  CardRequestStatus,
  CardType,
  CardBrand,
} from '../../database/entities/card-request.entity';
import { Card, CardStatus } from '../../database/entities/card.entity';
import { Account } from '../../database/entities/account.entity';
import { Branch } from '../../database/entities/branch.entity';
import { Customer } from '../../database/entities/customer.entity';
import { RequestCardDto, ReplaceCardDto, ListCardsDto } from './dto/card.dto';
import { AccountNotFoundException, CustomerNotFoundException } from '../../common/exceptions/domain.exceptions';

@Injectable()
export class CardsService {
  constructor(
    @InjectRepository(CardRequest)
    private readonly cardRequestRepo: Repository<CardRequest>,
    @InjectRepository(Card)
    private readonly cardRepo: Repository<Card>,
    @InjectRepository(Account)
    private readonly accountRepo: Repository<Account>,
    @InjectRepository(Branch)
    private readonly branchRepo: Repository<Branch>,
    @InjectRepository(Customer)
    private readonly customerRepo: Repository<Customer>,
  ) {}

  async requestCard(dto: RequestCardDto) {
    const account = await this.accountRepo.findOne({
      where: { accountNumber: dto.accountNumber },
      relations: ['customer'],
    });
    if (!account) throw new AccountNotFoundException(dto.accountNumber);

    const branch = await this.branchRepo.findOne({
      where: { branchCode: dto.targetBranchCode },
    });
    if (!branch) throw new BadRequestException(`Branch code '${dto.targetBranchCode}' not found`);

    const count = await this.cardRequestRepo.count();
    const requestReference = `CR-${new Date().toISOString().slice(0, 10).replace(/-/g, '')}-${(1000 + count + 1).toString()}`;

    const cardReq = this.cardRequestRepo.create({
      requestReference,
      accountId: account.id,
      customerId: account.customerId,
      cardType: dto.cardType || CardType.DEBIT,
      cardBrand: dto.cardBrand || CardBrand.MASTERCARD,
      requestType: CardRequestType.NEW,
      status: CardRequestStatus.PENDING,
      targetBranchId: branch.id,
    });

    const saved = await this.cardRequestRepo.save(cardReq);

    return {
      requestReference: saved.requestReference,
      accountNumber: account.accountNumber,
      customerName: account.customer.fullName,
      cardType: saved.cardType,
      cardBrand: saved.cardBrand,
      collectionBranch: branch.branchName,
      status: saved.status,
      createdAt: saved.createdAt,
      message: 'Card order queued for personalization and dispatch',
    };
  }

  async replaceCard(dto: ReplaceCardDto) {
    const card = await this.cardRepo.findOne({
      where: { cardNumberMasked: dto.cardNumberMasked },
      relations: ['account', 'customer'],
    });

    if (!card) {
      throw new NotFoundException(`Card '${dto.cardNumberMasked}' not found`);
    }

    const branch = await this.branchRepo.findOne({
      where: { branchCode: dto.targetBranchCode },
    });
    if (!branch) throw new BadRequestException(`Branch '${dto.targetBranchCode}' not found`);

    // Block current card
    card.status = CardStatus.BLOCKED;
    await this.cardRepo.save(card);

    const count = await this.cardRequestRepo.count();
    const requestReference = `REP-${new Date().toISOString().slice(0, 10).replace(/-/g, '')}-${(1000 + count + 1).toString()}`;

    const req = this.cardRequestRepo.create({
      requestReference,
      accountId: card.accountId,
      customerId: card.customerId,
      cardType: card.cardType,
      cardBrand: card.cardBrand,
      requestType: CardRequestType.REPLACEMENT,
      status: CardRequestStatus.PENDING,
      targetBranchId: branch.id,
    });
    const saved = await this.cardRequestRepo.save(req);

    return {
      requestReference: saved.requestReference,
      previousCardMasked: card.cardNumberMasked,
      previousCardStatus: card.status,
      replacementStatus: saved.status,
      collectionBranch: branch.branchName,
      createdAt: saved.createdAt,
      message: 'Previous card blocked; replacement card requested',
    };
  }

  async listCards(dto: ListCardsDto) {
    const where: any = {};

    if (dto.accountNumber) {
      const account = await this.accountRepo.findOne({ where: { accountNumber: dto.accountNumber } });
      if (!account) throw new AccountNotFoundException(dto.accountNumber);
      where.accountId = account.id;
    }

    if (dto.customerNumber) {
      const customer = await this.customerRepo.findOne({ where: { customerNumber: dto.customerNumber } });
      if (!customer) throw new CustomerNotFoundException(dto.customerNumber);
      where.customerId = customer.id;
    }

    const cards = await this.cardRepo.find({
      where,
      relations: ['account'],
      order: { createdAt: 'DESC' },
    });

    return {
      count: cards.length,
      cards: cards.map((c) => ({
        id: c.id,
        cardNumberMasked: c.cardNumberMasked,
        cardType: c.cardType,
        cardBrand: c.cardBrand,
        accountNumber: c.account?.accountNumber,
        status: c.status,
        expiryDate: c.expiryDate,
        dailyLimit: c.dailyLimit,
        createdAt: c.createdAt,
      })),
    };
  }
}
