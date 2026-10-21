import { FinancialMath } from '../../src/common/utils/financial-math.util';

describe('Banking Account Business Rules Unit Tests', () => {
  it('should derive available balance as ledgerBalance minus active holds', () => {
    const ledgerBalance = '150000.0000';
    const activeHold = '5000.0000';

    const availableBalance = FinancialMath.subtract(ledgerBalance, activeHold);
    expect(availableBalance).toBe('145000.0000');
  });

  it('should reject debit when required amount exceeds available balance even if ledger balance is high', () => {
    const availableBalance = '145000.0000';
    const requestedAmount = '145000.0001';

    const isSufficient = FinancialMath.isGreaterThanOrEqualTo(availableBalance, requestedAmount);
    expect(isSufficient).toBe(false);
  });

  it('should correctly restore available balance when hold is released', () => {
    let availableBalance = '145000.0000';
    let lockedBalance = '5000.0000';
    const holdToRelease = '5000.0000';

    lockedBalance = FinancialMath.subtract(lockedBalance, holdToRelease);
    availableBalance = FinancialMath.add(availableBalance, holdToRelease);

    expect(lockedBalance).toBe('0.0000');
    expect(availableBalance).toBe('150000.0000');
  });

  it('should enforce balanced double-entry invariant: debit sum equals credit sum', () => {
    const transferAmount = '12500.0000';
    const feeAmount = '5.0000';

    // Source customer account debited total
    const totalDebit = FinancialMath.add(transferAmount, feeAmount); // 12505.0000

    // Destination customer credited principal
    const destinationCredit = transferAmount; // 12500.0000
    // Bank fee income account credited fee
    const feeCredit = feeAmount; // 5.0000

    const totalCredits = FinancialMath.add(destinationCredit, feeCredit);

    expect(FinancialMath.isEqualTo(totalDebit, totalCredits)).toBe(true);
  });
});
