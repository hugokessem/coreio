import { AuthService } from './auth.service';
import { LoginDto } from './dto/login.dto';
export declare class AuthController {
    private readonly authService;
    constructor(authService: AuthService);
    login(loginDto: LoginDto): Promise<{
        accessToken: string;
        tokenType: string;
        expiresIn: string;
        user: {
            id: string;
            username: string;
            email: string;
            roles: string[];
        };
    }>;
}
