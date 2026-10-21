export declare class AuditLog {
    id: string;
    action: string;
    entityName: string;
    entityId: string | null;
    userId: string | null;
    ipAddress: string | null;
    userAgent: string | null;
    details: Record<string, any> | null;
    timestamp: Date;
}
