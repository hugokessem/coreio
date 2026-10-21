import { DataSource } from 'typeorm';
import * as bcrypt from 'bcrypt';
import {
  Customer,
  CustomerType,
  CustomerStatus,
  AMLStatus,
  Branch,
  Currency,
  ExchangeRate,
  AccountType,
  AccountCategory,
  Account,
  AccountStatus,
  AccountBalance,
  AccountHold,
  HoldStatus,
  Transaction,
  TransactionType,
  TransactionStatus,
  TransactionChannel,
  LedgerEntry,
  EntryType,
  ServiceEntity,
  FeeType,
  CustomerLimit,
  SuperAppUser,
  SuperAppStatus,
  CardRequest,
  CardRequestType,
  CardRequestStatus,
  CardType,
  CardBrand,
  Card,
  CardStatus,
  StandingOrder,
  StandingOrderFrequency,
  StandingOrderStatus,
  User,
  UserRole,
  AccountParty,
  PartyRole,
  IdempotencyRecord,
  StandingOrderExecution,
} from '../entities';
import { AppDataSource } from '../data-source';

export async function runSeeds(dataSource: DataSource) {
  console.log('--- Starting Database Seeding ---');

  // 1. Seed Users
  const userRepo = dataSource.getRepository(User);
  const adminPassword = await bcrypt.hash('Admin123!', 10);
  const apiClientPassword = await bcrypt.hash('ApiSecret123!', 10);

  const existingAdmin = await userRepo.findOne({ where: { username: 'admin' } });
  if (!existingAdmin) {
    const admin = userRepo.create({
      username: 'admin',
      passwordHash: adminPassword,
      email: 'admin@cbe.com.et',
      roles: [UserRole.ADMIN, UserRole.TELLER, UserRole.API_CLIENT],
      status: 'ACTIVE',
    });
    await userRepo.save(admin);
    console.log('Seeded User: admin');
  }

  const existingClient = await userRepo.findOne({ where: { username: 'api_client' } });
  if (!existingClient) {
    const client = userRepo.create({
      username: 'api_client',
      passwordHash: apiClientPassword,
      email: 'apiconsumer@partner.et',
      roles: [UserRole.API_CLIENT],
      status: 'ACTIVE',
    });
    await userRepo.save(client);
    console.log('Seeded User: api_client');
  }

  // 2. Seed Currencies
  const currencyRepo = dataSource.getRepository(Currency);
  const currenciesData = [
    { code: 'ETB', name: 'Ethiopian Birr', numericCode: '230', decimalPlaces: 2, isBaseCurrency: true },
    { code: 'USD', name: 'US Dollar', numericCode: '840', decimalPlaces: 2, isBaseCurrency: false },
    { code: 'EUR', name: 'Euro', numericCode: '978', decimalPlaces: 2, isBaseCurrency: false },
    { code: 'GBP', name: 'British Pound', numericCode: '826', decimalPlaces: 2, isBaseCurrency: false },
  ];

  for (const c of currenciesData) {
    let curr = await currencyRepo.findOne({ where: { code: c.code } });
    if (!curr) {
      curr = currencyRepo.create(c);
      await currencyRepo.save(curr);
      console.log(`Seeded Currency: ${c.code}`);
    }
  }

  // 3. Seed Exchange Rates
  const rateRepo = dataSource.getRepository(ExchangeRate);
  const today = new Date().toISOString().split('T')[0];
  const ratesData = [
    { fromCurrency: 'USD', toCurrency: 'ETB', buyRate: '135.500000', sellRate: '138.210000', midRate: '136.855000', effectiveDate: today },
    { fromCurrency: 'EUR', toCurrency: 'ETB', buyRate: '147.200000', sellRate: '150.140000', midRate: '148.670000', effectiveDate: today },
    { fromCurrency: 'GBP', toCurrency: 'ETB', buyRate: '172.100000', sellRate: '175.540000', midRate: '173.820000', effectiveDate: today },
  ];

  for (const r of ratesData) {
    const existingRate = await rateRepo.findOne({
      where: { fromCurrency: r.fromCurrency, toCurrency: r.toCurrency, effectiveDate: r.effectiveDate },
    });
    if (!existingRate) {
      await rateRepo.save(rateRepo.create(r));
      console.log(`Seeded FX Rate: ${r.fromCurrency}/${r.toCurrency}`);
    }
  }

  // 4. Seed Branches
  const branchRepo = dataSource.getRepository(Branch);
  const branchesData = [
    { branchCode: '001', branchName: 'Addis Ababa Main Branch', address: 'Churchill Road', city: 'Addis Ababa', swiftCode: 'CBETETAA' },
    { branchCode: '100', branchName: 'Bole Branch', address: 'Africa Avenue, Bole', city: 'Addis Ababa', swiftCode: 'CBETETAA' },
    { branchCode: '101', branchName: 'Mexico Branch', address: 'Mexico Square', city: 'Addis Ababa', swiftCode: 'CBETETAA' },
    { branchCode: '201', branchName: 'Hawassa Main Branch', address: 'Piazza', city: 'Hawassa', swiftCode: 'CBETETAA' },
  ];

  const branchMap: Record<string, Branch> = {};
  for (const b of branchesData) {
    let branch = await branchRepo.findOne({ where: { branchCode: b.branchCode } });
    if (!branch) {
      branch = await branchRepo.save(branchRepo.create(b));
      console.log(`Seeded Branch: ${b.branchCode} - ${b.branchName}`);
    }
    branchMap[b.branchCode] = branch;
  }

  // 5. Seed Account Types
  const accTypeRepo = dataSource.getRepository(AccountType);
  const accTypesData = [
    { code: '6501', name: 'Ordinary Savings Account', category: AccountCategory.SAVINGS, minimumBalance: '50.0000', isInterestBearing: true },
    { code: '6001', name: 'Demand Deposit (Current) Account', category: AccountCategory.CURRENT, minimumBalance: '500.0000', isInterestBearing: false },
    { code: '6601', name: 'Interest Free Banking (Wadiah) Savings', category: AccountCategory.SAVINGS, minimumBalance: '50.0000', isInterestBearing: false },
    { code: '6701', name: 'Fixed Time Deposit', category: AccountCategory.FIXED_DEPOSIT, minimumBalance: '5000.0000', isInterestBearing: true },
  ];

  const accTypeMap: Record<string, AccountType> = {};
  for (const at of accTypesData) {
    let accType = await accTypeRepo.findOne({ where: { code: at.code } });
    if (!accType) {
      accType = await accTypeRepo.save(accTypeRepo.create(at));
      console.log(`Seeded Account Type: ${at.code} - ${at.name}`);
    }
    accTypeMap[at.code] = accType;
  }

  // 6. Seed Services & Service Limits
  const serviceRepo = dataSource.getRepository(ServiceEntity);
  const servicesData = [
    { serviceCode: 'FT_INT', serviceName: 'Internal Fund Transfer', category: 'FUNDS_TRANSFER', minAmount: '10.0000', maxAmount: '1000000.0000', dailyLimit: '5000000.0000', feeType: FeeType.FLAT, feeValue: '0.0000' },
    { serviceCode: 'FT_IPS', serviceName: 'Interbank IPS Transfer', category: 'FUNDS_TRANSFER', minAmount: '10.0000', maxAmount: '500000.0000', dailyLimit: '2000000.0000', feeType: FeeType.FLAT, feeValue: '5.0000' },
    { serviceCode: 'CBE_BIRR', serviceName: 'CBE Birr Wallet Transfer', category: 'FUNDS_TRANSFER', minAmount: '5.0000', maxAmount: '100000.0000', dailyLimit: '300000.0000', feeType: FeeType.FLAT, feeValue: '2.0000' },
    { serviceCode: 'BILL_ETHIO', serviceName: 'Ethio Telecom Bill Payment', category: 'BILL_PAY', minAmount: '5.0000', maxAmount: '200000.0000', dailyLimit: '1000000.0000', feeType: FeeType.FLAT, feeValue: '0.0000' },
  ];

  for (const s of servicesData) {
    let srv = await serviceRepo.findOne({ where: { serviceCode: s.serviceCode } });
    if (!srv) {
      await serviceRepo.save(serviceRepo.create(s));
      console.log(`Seeded Service: ${s.serviceCode}`);
    }
  }

  // 7. Seed Customers
  const customerRepo = dataSource.getRepository(Customer);
  const customersData = [
    {
      customerNumber: '10000001',
      customerType: CustomerType.INDIVIDUAL,
      firstName: 'Abebe',
      middleName: 'Kebede',
      lastName: 'Bikila',
      legalName: 'Abebe Kebede Bikila',
      dateOfBirth: '1985-04-12',
      gender: 'MALE',
      phoneNumber: '+251911223344',
      email: 'abebe.kebede@example.com',
      address: 'Kirkos Sub-City, Addis Ababa',
      nationalId: 'FAYDA-9876543210',
      tin: '1234567890',
      kycLevel: 'LEVEL_3',
      amlStatus: AMLStatus.CLEAR,
      status: CustomerStatus.ACTIVE,
    },
    {
      customerNumber: '10000002',
      customerType: CustomerType.INDIVIDUAL,
      firstName: 'Almaz',
      middleName: 'Tadesse',
      lastName: 'Mengesha',
      legalName: 'Almaz Tadesse Mengesha',
      dateOfBirth: '1990-08-25',
      gender: 'FEMALE',
      phoneNumber: '+251922334455',
      email: 'almaz.t@example.com',
      address: 'Bole Sub-City, Addis Ababa',
      nationalId: 'FAYDA-1234567890',
      tin: '2233445566',
      kycLevel: 'LEVEL_3',
      amlStatus: AMLStatus.CLEAR,
      status: CustomerStatus.ACTIVE,
    },
    {
      customerNumber: '10000003',
      customerType: CustomerType.INDIVIDUAL,
      firstName: 'Chala',
      middleName: 'Bekele',
      lastName: 'Tolessa',
      legalName: 'Chala Bekele Tolessa',
      dateOfBirth: '1978-11-03',
      gender: 'MALE',
      phoneNumber: '+251933445566',
      email: 'chala.b@example.com',
      address: 'Lideta Sub-City, Addis Ababa',
      nationalId: 'FAYDA-5544332211',
      tin: '3344556677',
      kycLevel: 'LEVEL_2',
      amlStatus: AMLStatus.CLEAR,
      status: CustomerStatus.ACTIVE,
    },
    {
      customerNumber: '10000004',
      customerType: CustomerType.INDIVIDUAL,
      firstName: 'Solomon',
      middleName: 'Girma',
      lastName: 'Hailu',
      legalName: 'Solomon Girma Hailu',
      dateOfBirth: '1975-01-15',
      gender: 'MALE',
      phoneNumber: '+251944556677',
      email: 'solomon.girma@example.com',
      address: 'Yeka Sub-City, Addis Ababa',
      nationalId: 'FAYDA-9988776655',
      tin: '4455667788',
      kycLevel: 'LEVEL_2',
      amlStatus: AMLStatus.CLEAR,
      status: CustomerStatus.ACTIVE,
    },
    {
      customerNumber: '10000005',
      customerType: CustomerType.CORPORATE,
      firstName: null,
      middleName: null,
      lastName: null,
      legalName: 'Ethiopian Trading Enterprise SC',
      dateOfBirth: null,
      gender: null,
      phoneNumber: '+251115510000',
      email: 'finance@ete.com.et',
      address: 'Arada Sub-City, Addis Ababa',
      nationalId: null,
      tin: '0012345678',
      kycLevel: 'LEVEL_3',
      amlStatus: AMLStatus.CLEAR,
      status: CustomerStatus.ACTIVE,
    },
    {
      customerNumber: '10000006',
      customerType: CustomerType.CORPORATE,
      firstName: null,
      middleName: null,
      lastName: null,
      legalName: 'Suspicious Entity Holdings Ltd',
      dateOfBirth: null,
      gender: null,
      phoneNumber: '+251999999999',
      email: 'alert@suspicious.com',
      address: 'Unknown',
      nationalId: null,
      tin: '9999999999',
      kycLevel: 'LEVEL_1',
      amlStatus: AMLStatus.SUSPECT,
      status: CustomerStatus.SUSPENDED,
    },
  ];

  const customerMap: Record<string, Customer> = {};
  for (const c of customersData) {
    let cust = await customerRepo.findOne({ where: { customerNumber: c.customerNumber } });
    if (!cust) {
      cust = await customerRepo.save(customerRepo.create(c));
      console.log(`Seeded Customer: ${c.customerNumber} - ${c.legalName}`);
    }
    customerMap[c.customerNumber] = cust;
  }

  // 8. Seed Accounts and Balances
  const accountRepo = dataSource.getRepository(Account);
  const balanceRepo = dataSource.getRepository(AccountBalance);
  const holdRepo = dataSource.getRepository(AccountHold);
  const txRepo = dataSource.getRepository(Transaction);
  const ledgerRepo = dataSource.getRepository(LedgerEntry);

  const accountsData = [
    {
      accountNumber: '1000123456789',
      customerNumber: '10000001',
      accountTypeCode: '6501',
      branchCode: '001',
      currencyCode: 'ETB',
      accountName: 'Abebe Kebede Bikila',
      status: AccountStatus.ACTIVE,
      initialLedger: '150000.0000',
      holdAmount: '5000.0000',
      holdReason: 'Court Order Freeze Case #9921',
      holdReference: 'HLD-001',
    },
    {
      accountNumber: '1000987654321',
      customerNumber: '10000002',
      accountTypeCode: '6001',
      branchCode: '100',
      currencyCode: 'ETB',
      accountName: 'Almaz Tadesse Mengesha',
      status: AccountStatus.ACTIVE,
      initialLedger: '75000.0000',
      holdAmount: '0.0000',
      holdReason: null,
      holdReference: null,
    },
    {
      accountNumber: '1000555566667',
      customerNumber: '10000003',
      accountTypeCode: '6501',
      branchCode: '101',
      currencyCode: 'ETB',
      accountName: 'Chala Bekele Tolessa',
      status: AccountStatus.ACTIVE,
      initialLedger: '25000.0000',
      holdAmount: '0.0000',
      holdReason: null,
      holdReference: null,
    },
    {
      accountNumber: '1000111122223',
      customerNumber: '10000004',
      accountTypeCode: '6501',
      branchCode: '001',
      currencyCode: 'ETB',
      accountName: 'Solomon Girma Hailu',
      status: AccountStatus.DORMANT,
      initialLedger: '1200.0000',
      holdAmount: '0.0000',
      holdReason: null,
      holdReference: null,
    },
    {
      accountNumber: '1000999988887',
      customerNumber: '10000001',
      accountTypeCode: '6501',
      branchCode: '001',
      currencyCode: 'USD',
      accountName: 'Abebe Kebede Bikila - USD',
      status: AccountStatus.ACTIVE,
      initialLedger: '5000.0000',
      holdAmount: '0.0000',
      holdReason: null,
      holdReference: null,
    },
    {
      accountNumber: '2000123456789',
      customerNumber: '10000005',
      accountTypeCode: '6001',
      branchCode: '001',
      currencyCode: 'ETB',
      accountName: 'Ethiopian Trading Enterprise SC',
      status: AccountStatus.ACTIVE,
      initialLedger: '1500000.0000',
      holdAmount: '0.0000',
      holdReason: null,
      holdReference: null,
    },
  ];

  for (const ad of accountsData) {
    let acc = await accountRepo.findOne({ where: { accountNumber: ad.accountNumber } });
    if (!acc) {
      const cust = customerMap[ad.customerNumber];
      const br = branchMap[ad.branchCode];
      const at = accTypeMap[ad.accountTypeCode];

      acc = accountRepo.create({
        accountNumber: ad.accountNumber,
        customerId: cust.id,
        accountTypeId: at.id,
        branchId: br.id,
        currencyCode: ad.currencyCode,
        accountName: ad.accountName,
        status: ad.status,
      });
      acc = await accountRepo.save(acc);
      console.log(`Seeded Account: ${ad.accountNumber} (${ad.accountName})`);

      // Compute balance
      const initialLedgerNum = parseFloat(ad.initialLedger);
      const holdNum = parseFloat(ad.holdAmount);
      const availNum = initialLedgerNum - holdNum;

      const bal = balanceRepo.create({
        accountId: acc.id,
        ledgerBalance: ad.initialLedger,
        availableBalance: availNum.toFixed(4),
        lockedBalance: ad.holdAmount,
        currencyCode: ad.currencyCode,
        version: 1,
      });
      await balanceRepo.save(bal);

      // Create opening transaction & ledger entry
      const tx = txRepo.create({
        transactionReference: `INIT-${ad.accountNumber}`,
        transactionType: TransactionType.CASH_DEPOSIT,
        channel: TransactionChannel.SYSTEM,
        sourceAccountId: null,
        destinationAccountId: acc.id,
        amount: ad.initialLedger,
        currencyCode: ad.currencyCode,
        narration: 'Account Opening Initial Deposit',
        status: TransactionStatus.COMPLETED,
        completedAt: new Date(),
      });
      await txRepo.save(tx);

      const entry = ledgerRepo.create({
        transactionId: tx.id,
        accountId: acc.id,
        entryType: EntryType.CREDIT,
        amount: ad.initialLedger,
        currencyCode: ad.currencyCode,
        balanceAfter: ad.initialLedger,
        narration: 'Opening Deposit',
      });
      await ledgerRepo.save(entry);

      // Create AccountParty (Primary Owner)
      const partyRepo = dataSource.getRepository(AccountParty);
      const party = partyRepo.create({
        accountId: acc.id,
        customerId: cust.id,
        partyRole: PartyRole.PRIMARY_OWNER,
        authorizationLevel: 'SOLE_SIGNATORY',
        status: 'ACTIVE',
      });
      await partyRepo.save(party);

      // Create hold if specified
      if (holdNum > 0 && ad.holdReference) {
        const hold = holdRepo.create({
          holdReference: ad.holdReference,
          accountId: acc.id,
          amount: ad.holdAmount,
          remainingAmount: ad.holdAmount,
          currencyCode: ad.currencyCode,
          reason: ad.holdReason || 'Administrative Hold',
          lockEvent: 'COURT_ORDER',
          status: HoldStatus.ACTIVE,
          fromDate: today,
          createdBy: 'JUDICIAL_ORDER',
        });
        await holdRepo.save(hold);
        console.log(`Seeded Active Hold ${ad.holdReference} for ${ad.accountNumber}`);
      }
    }
  }

  // 9. Seed Standing Orders
  const soRepo = dataSource.getRepository(StandingOrder);
  const acc1 = await accountRepo.findOne({ where: { accountNumber: '1000123456789' } });
  const acc2 = await accountRepo.findOne({ where: { accountNumber: '1000987654321' } });

  if (acc1 && acc2) {
    const existingSO = await soRepo.findOne({ where: { orderReference: 'SO-1000123456789-01' } });
    if (!existingSO) {
      const so = soRepo.create({
        orderReference: 'SO-1000123456789-01',
        sourceAccountId: acc1.id,
        destinationAccountId: acc2.id,
        beneficiaryName: 'Almaz Tadesse Mengesha',
        beneficiaryAccount: '1000987654321',
        amount: '2500.0000',
        currencyCode: 'ETB',
        frequency: StandingOrderFrequency.MONTHLY,
        startDate: today,
        nextExecutionDate: today,
        status: StandingOrderStatus.ACTIVE,
      });
      await soRepo.save(so);
      console.log('Seeded Standing Order: SO-1000123456789-01');
    }
  }

  // 10. Seed Cards
  const cardRepo = dataSource.getRepository(Card);
  if (acc1) {
    const cust1 = customerMap['10000001'];
    const existingCard = await cardRepo.findOne({ where: { cardHash: 'seed_card_hash_001' } });
    if (!existingCard) {
      const card = cardRepo.create({
        cardNumberMasked: '541234******7890',
        cardHash: 'seed_card_hash_001',
        cardType: CardType.DEBIT,
        cardBrand: CardBrand.MASTERCARD,
        accountId: acc1.id,
        customerId: cust1.id,
        status: CardStatus.ACTIVE,
        expiryDate: '12/28',
        dailyLimit: '50000.0000',
      });
      await cardRepo.save(card);
      console.log('Seeded Card: 541234******7890');
    }
  }

  console.log('--- Database Seeding Completed Successfully ---');
}

// Allow direct execution via ts-node
if (require.main === module) {
  AppDataSource.initialize()
    .then(async (ds) => {
      // Synchronize database schema for development seeding
      await ds.synchronize();
      await runSeeds(ds);
      await ds.destroy();
      process.exit(0);
    })
    .catch((err) => {
      console.error('Seeding Error:', err);
      process.exit(1);
    });
}
