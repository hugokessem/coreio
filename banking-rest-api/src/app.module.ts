import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import databaseConfig from './config/database.config';

import { AuthModule } from './modules/auth/auth.module';
import { CustomersModule } from './modules/customers/customers.module';
import { AccountsModule } from './modules/accounts/accounts.module';
import { TransfersModule } from './modules/transfers/transfers.module';
import { StandingOrdersModule } from './modules/standing-orders/standing-orders.module';
import { PaymentsModule } from './modules/payments/payments.module';
import { CardsModule } from './modules/cards/cards.module';
import { RatesModule } from './modules/rates/rates.module';
import { ServicesModule } from './modules/services/services.module';
import { SuperAppModule } from './modules/superapp/superapp.module';
import { IpsModule } from './modules/ips/ips.module';
import { WalletModule } from './modules/wallet/wallet.module';
import { HealthModule } from './modules/health/health.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: [databaseConfig],
    }),
    TypeOrmModule.forRootAsync({
      inject: [ConfigService],
      useFactory: (configService: ConfigService) => ({
        ...configService.get('database'),
        autoLoadEntities: true,
      }),
    }),
    AuthModule,
    CustomersModule,
    AccountsModule,
    TransfersModule,
    StandingOrdersModule,
    PaymentsModule,
    CardsModule,
    RatesModule,
    ServicesModule,
    SuperAppModule,
    IpsModule,
    WalletModule,
    HealthModule,
  ],
})
export class AppModule {}
