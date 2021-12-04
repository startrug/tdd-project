const assert = require("assert");
const Money = require("./money");
const Portfolio = require("./portfolio");

class MoneyTest {
  testMultiplication() {
    let fiveDollars = new Money(5, "USD");
    let tenDollars = fiveDollars.times(2);
    let expectedUsd = new Money(10, "USD");

    assert.deepStrictEqual(tenDollars, expectedUsd);
  }

  testDivision() {
    let originalMoney = new Money(4002, "KRW");
    let actualMoneyAfterDivision = originalMoney.divide(4);
    let expectedMoneyAfterDivision = new Money(1000.5, "KRW");

    assert.deepStrictEqual(
      actualMoneyAfterDivision,
      expectedMoneyAfterDivision
    );
  }

  testAddition() {
    let fiveBucks = new Money(5, "USD");
    let tenBucks = new Money(10, "USD");
    let fifteenBucks = new Money(15, "USD");
    let portfolio = new Portfolio();
    portfolio.add(fiveBucks, tenBucks);

    assert.deepStrictEqual(portfolio.evaluate("USD"), fifteenBucks);
  }

  testAdditionOfDollarsAndEuros() {
    let fiveDollars = new Money(5, "USD");
    let tenEuros = new Money(10, "EUR");
    let portfolio = new Portfolio();
    portfolio.add(fiveDollars, tenEuros);
    let expectedValue = new Money(16, "USD");
    assert.deepStrictEqual(portfolio.evaluate("USD"), expectedValue);
  }

  testAdditionOfDollarsAndWons() {
    let oneDollar = new Money(1, "USD");
    let elevenHundredWon = new Money(1100, "KRW");
    let portfolio = new Portfolio();
    portfolio.add(oneDollar, elevenHundredWon);
    let expectedValue = new Money(2200, "KRW");
    assert.deepStrictEqual(portfolio.evaluate("KRW"), expectedValue);
  }

  testAdditionWithMultipleMissingExchangeRates() {
    let oneDollar = new Money(1, "USD");
    let oneEuro = new Money(1, "EUR");
    let oneWon = new Money(1, "KRW");
    let portfolio = new Portfolio();
    portfolio.add(oneDollar, oneEuro, oneWon);
    let expectedError = new Error(
      "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid]"
    );
    assert.throws(function () {
      portfolio.evaluate("Kalganid");
    }, expectedError);
  }

  getAllTestMethods() {
    let moneyPrototype = MoneyTest.prototype;
    let allProps = Object.getOwnPropertyNames(moneyPrototype);
    let testMethods = allProps.filter((p) => {
      return typeof moneyPrototype[p] === "function" && p.startsWith("test");
    });
    return testMethods;
  }

  runAllTests() {
    let testMethods = this.getAllTestMethods();
    testMethods.forEach((m) => {
      console.log("Running: %s()", m);
      let method = Reflect.get(this, m);
      try {
        Reflect.apply(method, this, []);
        console.log("Pass");
      } catch (e) {
        if (e instanceof assert.AssertionError) {
          console.log(e);
        } else {
          throw e;
        }
      }
    });
  }
}

new MoneyTest().runAllTests();
