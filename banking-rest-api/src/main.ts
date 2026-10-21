import { NestFactory } from '@nestjs/core';
import { ValidationPipe, Logger } from '@nestjs/common';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import { AppModule } from './app.module';
import { AllExceptionsFilter } from './common/filters/all-exceptions.filter';

async function bootstrap() {
  const logger = new Logger('BankingBootstrap');
  const app = await NestFactory.create(AppModule);

  // Global exception filter
  app.useGlobalFilters(new AllExceptionsFilter());

  // Strict validation pipe
  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
      forbidNonWhitelisted: true,
      transform: true,
      transformOptions: {
        enableImplicitConversion: true,
      },
    }),
  );

  // Enable CORS
  app.enableCors({
    origin: '*',
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS',
    allowedHeaders: 'Content-Type, Accept, Authorization, x-api-key, x-correlation-id',
  });

  // OpenAPI / Swagger Documentation
  const config = new DocumentBuilder()
    .setTitle('CBE CoreIO Enterprise Banking REST API')
    .setDescription(
      'Standalone, independent, high-performance NestJS banking REST API re-engineered from legacy CoreIO T24 integration. ' +
      'Implements real PostgreSQL relational persistence, double-entry immutable ledger accounting, pessimistic balance concurrency locking, and ISO 20022 / EthSwitch interbank payments.',
    )
    .setVersion('1.0.0')
    .addBearerAuth()
    .addApiKey({ type: 'apiKey', name: 'x-api-key', in: 'header' }, 'x-api-key')
    .build();

  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('docs/swagger', app, document);

  const port = process.env.PORT || 3000;
  await app.listen(port);

  logger.log(`========================================================================`);
  logger.log(`  CBE CoreIO Banking REST API successfully started on port ${port}`);
  logger.log(`  Swagger API Documentation: http://localhost:${port}/docs/swagger`);
  logger.log(`  Standalone Mode: NO T24 or Go CoreIO runtime dependencies required`);
  logger.log(`========================================================================`);
}

bootstrap();
