"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.HealthController = void 0;
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const typeorm_1 = require("typeorm");
let HealthController = class HealthController {
    dataSource;
    constructor(dataSource) {
        this.dataSource = dataSource;
    }
    getLiveness() {
        return {
            status: 'UP',
            timestamp: new Date().toISOString(),
            uptimeSeconds: Math.floor(process.uptime()),
        };
    }
    async getReadiness() {
        let dbStatus = 'DOWN';
        let dbLatencyMs = -1;
        try {
            const start = Date.now();
            await this.dataSource.query('SELECT 1');
            dbLatencyMs = Date.now() - start;
            dbStatus = 'UP';
        }
        catch (err) {
            dbStatus = `DOWN: ${err.message}`;
        }
        const isReady = dbStatus === 'UP';
        return {
            status: isReady ? 'READY' : 'NOT_READY',
            timestamp: new Date().toISOString(),
            database: {
                status: dbStatus,
                latencyMs: dbLatencyMs,
            },
            memory: {
                heapUsedMb: Math.round(process.memoryUsage().heapUsed / 1024 / 1024),
                rssMb: Math.round(process.memoryUsage().rss / 1024 / 1024),
            },
        };
    }
    getApiHealth() {
        return this.getReadiness();
    }
};
exports.HealthController = HealthController;
__decorate([
    (0, common_1.Get)('health'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Liveness probe - checks if the HTTP process is responsive' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Service is alive' }),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", void 0)
], HealthController.prototype, "getLiveness", null);
__decorate([
    (0, common_1.Get)('health/ready'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'Readiness probe - verifies PostgreSQL connection and operational dependencies' }),
    (0, swagger_1.ApiResponse)({ status: 200, description: 'Service is ready to accept traffic' }),
    (0, swagger_1.ApiResponse)({ status: 503, description: 'PostgreSQL or dependent service is unavailable' }),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], HealthController.prototype, "getReadiness", null);
__decorate([
    (0, common_1.Get)('api/v1/health'),
    (0, common_1.HttpCode)(common_1.HttpStatus.OK),
    (0, swagger_1.ApiOperation)({ summary: 'API Versioned health check endpoint' }),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", void 0)
], HealthController.prototype, "getApiHealth", null);
exports.HealthController = HealthController = __decorate([
    (0, swagger_1.ApiTags)('Observability & Health'),
    (0, common_1.Controller)(),
    __metadata("design:paramtypes", [typeorm_1.DataSource])
], HealthController);
//# sourceMappingURL=health.controller.js.map