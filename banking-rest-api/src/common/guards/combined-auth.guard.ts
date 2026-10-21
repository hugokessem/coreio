import { Injectable, CanActivate, ExecutionContext, UnauthorizedException } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';
import { Reflector } from '@nestjs/core';

export const IS_PUBLIC_KEY = 'isPublic';

@Injectable()
export class CombinedAuthGuard extends AuthGuard('jwt') implements CanActivate {
  constructor(private reflector: Reflector) {
    super();
  }

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const isPublic = this.reflector.getAllAndOverride<boolean>(IS_PUBLIC_KEY, [
      context.getHandler(),
      context.getClass(),
    ]);

    if (isPublic) {
      return true;
    }

    const request = context.switchToHttp().getRequest();
    const apiKeyHeader = request.headers['x-api-key'];
    const configuredApiKey = process.env.API_KEY || 'cbe-banking-api-key-default-2026';

    // Allow API key authentication with least-privilege API_CLIENT role
    if (apiKeyHeader && configuredApiKey && apiKeyHeader === configuredApiKey) {
      request.user = {
        userId: 'system-api-consumer',
        username: 'api_gateway',
        roles: ['API_CLIENT'],
      };
      return true;
    }

    // Otherwise fallback to JWT AuthGuard
    try {
      const result = await super.canActivate(context);
      return result as boolean;
    } catch (err) {
      throw new UnauthorizedException('Authentication required. Provide a valid Bearer token or x-api-key header.');
    }
  }
}
