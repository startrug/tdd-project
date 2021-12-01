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

class Portfolio {
  constructor() {
    this.moneys = [];
  }

  add(...moneys) {
    this.moneys = this.moneys.concat(moneys);
  }

  evaluate(currency) {
    let total = this.moneys.reduce((sum, money) => {
      return sum + money.amount;
    }, 0);
    return new Money(total, currency);
  }
}

let fiveDollars = new Money(5, "USD");
let tenDollars = fiveDollars.times(2);
let expectedUsd = new Money(10, "USD");

assert.deepStrictEqual(tenDollars, expectedUsd);

let tenEuros = new Money(10, "EUR");
let twentyEuros = tenEuros.times(2);
let expectedEur = new Money(20, "EUR");

assert.deepStrictEqual(twentyEuros, expectedEur);

let originalMoney = new Money(4002, "KRW");
let actualMoneyAfterDivision = originalMoney.divide(4);
let expectedMoneyAfterDivision = new Money(1000.5, "KRW");

assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);

let fiveBucks = new Money(5, "USD");
let tenBucks = new Money(10, "USD");
let fifteenBucks = new Money(15, "USD");
let portfolio = new Portfolio();
portfolio.add(fiveBucks, tenBucks);
assert.deepStrictEqual(portfolio.evaluate("USD"), fifteenBucks);
