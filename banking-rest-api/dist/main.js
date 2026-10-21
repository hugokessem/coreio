"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const core_1 = require("@nestjs/core");
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const app_module_1 = require("./app.module");
const all_exceptions_filter_1 = require("./common/filters/all-exceptions.filter");
async function bootstrap() {
    const logger = new common_1.Logger('BankingBootstrap');
    const app = await core_1.NestFactory.create(app_module_1.AppModule);
    app.useGlobalFilters(new all_exceptions_filter_1.AllExceptionsFilter());
    app.useGlobalPipes(new common_1.ValidationPipe({
        whitelist: true,
        forbidNonWhitelisted: true,
        transform: true,
        transformOptions: {
            enableImplicitConversion: true,
        },
    }));
    app.enableCors({
        origin: '*',
        methods: 'GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS',
        allowedHeaders: 'Content-Type, Accept, Authorization, x-api-key, x-correlation-id',
    });
    const config = new swagger_1.DocumentBuilder()
        .setTitle('CBE CoreIO Enterprise Banking REST API')
        .setDescription('Standalone, independent, high-performance NestJS banking REST API re-engineered from legacy CoreIO T24 integration. ' +
        'Implements real PostgreSQL relational persistence, double-entry immutable ledger accounting, pessimistic balance concurrency locking, and ISO 20022 / EthSwitch interbank payments.')
        .setVersion('1.0.0')
        .addBearerAuth()
        .addApiKey({ type: 'apiKey', name: 'x-api-key', in: 'header' }, 'x-api-key')
        .build();
    const document = swagger_1.SwaggerModule.createDocument(app, config);
    swagger_1.SwaggerModule.setup('docs/swagger', app, document);
    const port = process.env.PORT || 3000;
    await app.listen(port);
    logger.log(`========================================================================`);
    logger.log(`  CBE CoreIO Banking REST API successfully started on port ${port}`);
    logger.log(`  Swagger API Documentation: http://localhost:${port}/docs/swagger`);
    logger.log(`  Standalone Mode: NO T24 or Go CoreIO runtime dependencies required`);
    logger.log(`========================================================================`);
}
bootstrap();
//# sourceMappingURL=main.js.map