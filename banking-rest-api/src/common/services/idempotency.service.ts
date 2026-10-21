import { Injectable, ConflictException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository, DataSource, EntityManager } from 'typeorm';
import * as crypto from 'crypto';
import {
  IdempotencyRecord,
  IdempotencyStatus,
} from '../../database/entities/idempotency-record.entity';
import { IdempotencyConflictException } from '../exceptions/domain.exceptions';

@Injectable()
export class IdempotencyService {
  constructor(
    @InjectRepository(IdempotencyRecord)
    private readonly idempotencyRepo: Repository<IdempotencyRecord>,
    private readonly dataSource: DataSource,
  ) {}

  /**
   * Generates a deterministic SHA-256 hash of the canonical request payload
   */
  hashPayload(payload: any): string {
    const canonical = JSON.stringify(payload, Object.keys(payload || {}).sort());
    return crypto.createHash('sha256').update(canonical).digest('hex');
  }

  /**
   * Checks or registers an idempotency key.
   * If a previous identical request exists and completed, returns the cached result.
   * If a previous request with the same key exists but differing payload, throws 409 Conflict.
   */
  async checkOrRegister(
    key: string | undefined,
    operation: string,
    payload: any,
    callerId?: string,
  ): Promise<{ isReplay: boolean; cachedResponse?: any; recordId?: string }> {
    if (!key) {
      return { isReplay: false };
    }

    const requestHash = this.hashPayload(payload);
    const existing = await this.idempotencyRepo.findOne({
      where: { idempotencyKey: key },
    });

    if (existing) {
      if (existing.requestHash !== requestHash) {
        throw new IdempotencyConflictException(
          `Idempotency key '${key}' was previously used with a different request payload. Duplicate execution rejected.`,
        );
      }

      if (existing.status === IdempotencyStatus.COMPLETED && existing.responsePayload) {
        return {
          isReplay: true,
          cachedResponse: {
            ...existing.responsePayload,
            message: `Idempotent replay: previously completed ${operation} returned`,
          },
        };
      }
    }

    return { isReplay: false };
  }

  /**
   * Saves or updates an idempotency record with the completed transaction response
   */
  async recordSuccess(
    key: string | undefined,
    operation: string,
    payload: any,
    transactionReference: string,
    response: any,
    httpStatus = 200,
    callerId?: string,
    entityManager?: EntityManager,
  ) {
    if (!key) return;

    const manager = entityManager || this.dataSource.manager;
    const requestHash = this.hashPayload(payload);

    let record = await manager.findOne(IdempotencyRecord, {
      where: { idempotencyKey: key },
    });

    if (record) {
      record.status = IdempotencyStatus.COMPLETED;
      record.transactionReference = transactionReference;
      record.responsePayload = response;
      record.httpStatusCode = httpStatus;
      await manager.save(IdempotencyRecord, record);
    } else {
      record = manager.create(IdempotencyRecord, {
        idempotencyKey: key,
        requestHash,
        operation,
        callerId: callerId || null,
        status: IdempotencyStatus.COMPLETED,
        transactionReference,
        responsePayload: response,
        httpStatusCode: httpStatus,
      });
      await manager.save(IdempotencyRecord, record);
    }
  }
}
