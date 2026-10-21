import { Repository, DataSource, EntityManager } from 'typeorm';
import { IdempotencyRecord } from '../../database/entities/idempotency-record.entity';
export declare class IdempotencyService {
    private readonly idempotencyRepo;
    private readonly dataSource;
    constructor(idempotencyRepo: Repository<IdempotencyRecord>, dataSource: DataSource);
    hashPayload(payload: any): string;
    checkOrRegister(key: string | undefined, operation: string, payload: any, callerId?: string): Promise<{
        isReplay: boolean;
        cachedResponse?: any;
        recordId?: string;
    }>;
    recordSuccess(key: string | undefined, operation: string, payload: any, transactionReference: string, response: any, httpStatus?: number, callerId?: string, entityManager?: EntityManager): Promise<void>;
}
