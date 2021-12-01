import unittest

class Money:
  def __init__(self, amount, currency):
    self.amount = amount
    self.currency = currency

  def __eq__(self, other):
    return self.amount == other.amount and self.currency == other.currency

  def times(self, multiplier):
    return Money(self.amount * multiplier, self.currency)

  def divide(self, divisor):
    return Money(self.amount / divisor, self.currency)

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

if __name__ == '__main__':
    unittest.main()