import BigNumber from 'bignumber.js';

// Configure BigNumber for banking precision
BigNumber.config({
  DECIMAL_PLACES: 4,
  ROUNDING_MODE: BigNumber.ROUND_HALF_UP,
  EXPONENTIAL_AT: [-10, 20],
});

export class FinancialMath {
  /**
   * Safe decimal addition
   */
  static add(a: string | number, b: string | number): string {
    return new BigNumber(a).plus(new BigNumber(b)).toFixed(4);
  }

  /**
   * Safe decimal subtraction
   */
  static subtract(a: string | number, b: string | number): string {
    return new BigNumber(a).minus(new BigNumber(b)).toFixed(4);
  }

  /**
   * Safe decimal multiplication
   */
  static multiply(a: string | number, b: string | number, decimalPlaces = 4): string {
    return new BigNumber(a).times(new BigNumber(b)).toFixed(decimalPlaces);
  }

  /**
   * Safe decimal division
   */
  static divide(a: string | number, b: string | number, decimalPlaces = 4): string {
    const divisor = new BigNumber(b);
    if (divisor.isZero()) {
      throw new Error('Division by zero in financial calculation');
    }
    return new BigNumber(a).dividedBy(divisor).toFixed(decimalPlaces);
  }

  /**
   * Check if a >= b
   */
  static isGreaterThanOrEqualTo(a: string | number, b: string | number): boolean {
    return new BigNumber(a).isGreaterThanOrEqualTo(new BigNumber(b));
  }

  /**
   * Check if a > b
   */
  static isGreaterThan(a: string | number, b: string | number): boolean {
    return new BigNumber(a).isGreaterThan(new BigNumber(b));
  }

  /**
   * Check if a < b
   */
  static isLessThan(a: string | number, b: string | number): boolean {
    return new BigNumber(a).isLessThan(new BigNumber(b));
  }

  /**
   * Check if a <= b
   */
  static isLessThanOrEqualTo(a: string | number, b: string | number): boolean {
    return new BigNumber(a).isLessThanOrEqualTo(new BigNumber(b));
  }

  /**
   * Check if a === b
   */
  static isEqualTo(a: string | number, b: string | number): boolean {
    return new BigNumber(a).isEqualTo(new BigNumber(b));
  }

  /**
   * Check if value is positive (> 0)
   */
  static isPositive(a: string | number): boolean {
    return new BigNumber(a).isGreaterThan(0);
  }

  /**
   * Format for display (2 decimals for fiat currency)
   */
  static toCurrencyString(a: string | number, decimals = 2): string {
    return new BigNumber(a).toFixed(decimals);
  }
}
