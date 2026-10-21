export declare class FinancialMath {
    static add(a: string | number, b: string | number): string;
    static subtract(a: string | number, b: string | number): string;
    static multiply(a: string | number, b: string | number, decimalPlaces?: number): string;
    static divide(a: string | number, b: string | number, decimalPlaces?: number): string;
    static isGreaterThanOrEqualTo(a: string | number, b: string | number): boolean;
    static isGreaterThan(a: string | number, b: string | number): boolean;
    static isLessThan(a: string | number, b: string | number): boolean;
    static isLessThanOrEqualTo(a: string | number, b: string | number): boolean;
    static isEqualTo(a: string | number, b: string | number): boolean;
    static isPositive(a: string | number): boolean;
    static toCurrencyString(a: string | number, decimals?: number): string;
}
