import { DataSource } from 'typeorm';
export declare class HealthController {
    private readonly dataSource;
    constructor(dataSource: DataSource);
    getLiveness(): {
        status: string;
        timestamp: string;
        uptimeSeconds: number;
    };
    getReadiness(): Promise<{
        status: string;
        timestamp: string;
        database: {
            status: string;
            latencyMs: number;
        };
        memory: {
            heapUsedMb: number;
            rssMb: number;
        };
    }>;
    getApiHealth(): Promise<{
        status: string;
        timestamp: string;
        database: {
            status: string;
            latencyMs: number;
        };
        memory: {
            heapUsedMb: number;
            rssMb: number;
        };
    }>;
}
