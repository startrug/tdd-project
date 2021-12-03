import unittest
from money import Money
from portfolio import Portfolio


class TestMoney(unittest.TestCase):
  def testMultiplication(self):
    fiver = Money(5, "USD")
    tenner = fiver.times(2)
    expected_usd = Money(10, "USD")
    self.assertEqual(expected_usd, tenner)

  def testMultiplicationInEuros(self):
    ten_euros = Money(10, "EUR")
    twenty_euros = ten_euros.times(2)
    expected_eur = Money(20, "EUR")
    self.assertEqual(expected_eur, twenty_euros)

  def testDivision(self):
    original_money = Money(4002, "KRW")
    actual_money_after_division = original_money.divide(4)
    expected_money_after_division = Money(1000.5, "KRW")
    self.assertEqual(expected_money_after_division,
                      actual_money_after_division)

  def testAddition(self):
    five_dollars = Money(5, "USD")
    ten_dollars = Money(10, "USD")
    fifteen_dollars = Money(15, "USD")
    portfolio = Portfolio()
    portfolio.add(five_dollars, ten_dollars)
    self.assertEqual(fifteen_dollars, portfolio.evaluate("USD"))

  def testAdditionOfDollarsAndEuros(self):
        fiveDollars = Money(5, "USD")
        tenEuros = Money(10, "EUR")
        portfolio = Portfolio()
        portfolio.add(fiveDollars, tenEuros)
        expectedValue = Money(16, "USD")
        actualValue = portfolio.evaluate("USD")
        self.assertEqual(expectedValue, actualValue, "%s != %s" % (expectedValue, actualValue))

if __name__ == '__main__':
    unittest.main()