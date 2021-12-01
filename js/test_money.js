const assert = require("assert");

class Money {
  constructor(amount, currency) {
    this.amount = amount;
    this.currency = currency;
  }

  times(multiplier) {
    return new Money(this.amount * multiplier, this.currency);
  }

  divide(divisor) {
    return new Money(this.amount / divisor, this.currency);
  }
}

let fiver = new Money(5, "USD");
let tenner = fiver.times(2);
let expectedUsd = new Money(10, "USD");

assert.deepStrictEqual(tenner, expectedUsd);

let tenEuros = new Money(10, "EUR");
let twentyEuros = tenEuros.times(2);
let expectedEur = new Money(20, "EUR");

assert.deepStrictEqual(twentyEuros, expectedEur);

let originalMoney = new Money(4002, "KRW");
let actualMoneyAfterDivision = originalMoney.divide(4);
let expectedMoneyAfterDivision = new Money(1000.5, "KRW");

assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);
