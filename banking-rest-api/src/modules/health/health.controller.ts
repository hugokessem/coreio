import { Controller, Get, HttpCode, HttpStatus } from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse } from '@nestjs/swagger';
import { DataSource } from 'typeorm';

@ApiTags('Observability & Health')
@Controller()
export class HealthController {
  constructor(private readonly dataSource: DataSource) {}

  @Get('health')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Liveness probe - checks if the HTTP process is responsive' })
  @ApiResponse({ status: 200, description: 'Service is alive' })
  getLiveness() {
    return {
      status: 'UP',
      timestamp: new Date().toISOString(),
      uptimeSeconds: Math.floor(process.uptime()),
    };
  }

  @Get('health/ready')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'Readiness probe - verifies PostgreSQL connection and operational dependencies' })
  @ApiResponse({ status: 200, description: 'Service is ready to accept traffic' })
  @ApiResponse({ status: 503, description: 'PostgreSQL or dependent service is unavailable' })
  async getReadiness() {
    let dbStatus = 'DOWN';
    let dbLatencyMs = -1;

    try {
      const start = Date.now();
      await this.dataSource.query('SELECT 1');
      dbLatencyMs = Date.now() - start;
      dbStatus = 'UP';
    } catch (err: any) {
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

  @Get('api/v1/health')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ summary: 'API Versioned health check endpoint' })
  getApiHealth() {
    return this.getReadiness();
  }
}
