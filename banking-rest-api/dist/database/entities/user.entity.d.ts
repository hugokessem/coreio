export declare enum UserRole {
    ADMIN = "ADMIN",
    TELLER = "TELLER",
    API_CLIENT = "API_CLIENT",
    AUDITOR = "AUDITOR"
}
export declare class User {
    id: string;
    username: string;
    passwordHash: string;
    email: string;
    roles: string[];
    status: string;
    createdAt: Date;
    updatedAt: Date;
}
