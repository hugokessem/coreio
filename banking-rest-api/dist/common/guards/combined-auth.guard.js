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
exports.CombinedAuthGuard = exports.IS_PUBLIC_KEY = void 0;
const common_1 = require("@nestjs/common");
const passport_1 = require("@nestjs/passport");
const core_1 = require("@nestjs/core");
exports.IS_PUBLIC_KEY = 'isPublic';
let CombinedAuthGuard = class CombinedAuthGuard extends (0, passport_1.AuthGuard)('jwt') {
    reflector;
    constructor(reflector) {
        super();
        this.reflector = reflector;
    }
    async canActivate(context) {
        const isPublic = this.reflector.getAllAndOverride(exports.IS_PUBLIC_KEY, [
            context.getHandler(),
            context.getClass(),
        ]);
        if (isPublic) {
            return true;
        }
        const request = context.switchToHttp().getRequest();
        const apiKeyHeader = request.headers['x-api-key'];
        const configuredApiKey = process.env.API_KEY || 'cbe-banking-api-key-default-2026';
        if (apiKeyHeader && configuredApiKey && apiKeyHeader === configuredApiKey) {
            request.user = {
                userId: 'system-api-consumer',
                username: 'api_gateway',
                roles: ['API_CLIENT'],
            };
            return true;
        }
        try {
            const result = await super.canActivate(context);
            return result;
        }
        catch (err) {
            throw new common_1.UnauthorizedException('Authentication required. Provide a valid Bearer token or x-api-key header.');
        }
    }
};
exports.CombinedAuthGuard = CombinedAuthGuard;
exports.CombinedAuthGuard = CombinedAuthGuard = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [core_1.Reflector])
], CombinedAuthGuard);
//# sourceMappingURL=combined-auth.guard.js.map