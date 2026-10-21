"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FinancialMath = void 0;
const bignumber_js_1 = require("bignumber.js");
bignumber_js_1.default.config({
    DECIMAL_PLACES: 4,
    ROUNDING_MODE: bignumber_js_1.default.ROUND_HALF_UP,
    EXPONENTIAL_AT: [-10, 20],
});
class FinancialMath {
    static add(a, b) {
        return new bignumber_js_1.default(a).plus(new bignumber_js_1.default(b)).toFixed(4);
    }
    static subtract(a, b) {
        return new bignumber_js_1.default(a).minus(new bignumber_js_1.default(b)).toFixed(4);
    }
    static multiply(a, b, decimalPlaces = 4) {
        return new bignumber_js_1.default(a).times(new bignumber_js_1.default(b)).toFixed(decimalPlaces);
    }
    static divide(a, b, decimalPlaces = 4) {
        const divisor = new bignumber_js_1.default(b);
        if (divisor.isZero()) {
            throw new Error('Division by zero in financial calculation');
        }
        return new bignumber_js_1.default(a).dividedBy(divisor).toFixed(decimalPlaces);
    }
    static isGreaterThanOrEqualTo(a, b) {
        return new bignumber_js_1.default(a).isGreaterThanOrEqualTo(new bignumber_js_1.default(b));
    }
    static isGreaterThan(a, b) {
        return new bignumber_js_1.default(a).isGreaterThan(new bignumber_js_1.default(b));
    }
    static isLessThan(a, b) {
        return new bignumber_js_1.default(a).isLessThan(new bignumber_js_1.default(b));
    }
    static isLessThanOrEqualTo(a, b) {
        return new bignumber_js_1.default(a).isLessThanOrEqualTo(new bignumber_js_1.default(b));
    }
    static isEqualTo(a, b) {
        return new bignumber_js_1.default(a).isEqualTo(new bignumber_js_1.default(b));
    }
    static isPositive(a) {
        return new bignumber_js_1.default(a).isGreaterThan(0);
    }
    static toCurrencyString(a, decimals = 2) {
        return new bignumber_js_1.default(a).toFixed(decimals);
    }
}
exports.FinancialMath = FinancialMath;
//# sourceMappingURL=financial-math.util.js.map