import { FinancialMath } from '../../src/common/utils/financial-math.util';

describe('FinancialMath Unit Tests', () => {
  it('should correctly add decimal amounts without floating point errors', () => {
    // 0.1 + 0.2 in standard JS IEEE-754 is 0.30000000000000004
    const result = FinancialMath.add('0.1000', '0.2000');
    expect(result).toBe('0.3000');
  });

  it('should correctly subtract amounts', () => {
    const result = FinancialMath.subtract('100.5000', '25.2500');
    expect(result).toBe('75.2500');
  });

  it('should accurately multiply with custom precision', () => {
    const result = FinancialMath.multiply('100.0000', '1.1500', 4);
    expect(result).toBe('115.0000');
  });

  it('should throw on division by zero', () => {
    expect(() => FinancialMath.divide('100', '0')).toThrow('Division by zero');
  });

  it('should correctly evaluate inequalities', () => {
    expect(FinancialMath.isGreaterThanOrEqualTo('50.0000', '50.0000')).toBe(true);
    expect(FinancialMath.isGreaterThanOrEqualTo('50.0001', '50.0000')).toBe(true);
    expect(FinancialMath.isLessThan('49.9999', '50.0000')).toBe(true);
    expect(FinancialMath.isEqualTo('100.0000', '100')).toBe(true);
  });

  it('should check positivity accurately', () => {
    expect(FinancialMath.isPositive('0.0001')).toBe(true);
    expect(FinancialMath.isPositive('0.0000')).toBe(false);
    expect(FinancialMath.isPositive('-5.0000')).toBe(false);
  });
});
