import functools
import operator
from money import Money


class Portfolio:
  def __init__(self):
        self.moneys = []
        self._eur_to_usd = 1.1

  def add(self, *moneys):
        self.moneys.extend(moneys)

  def evaluate(self, currency):
      total = functools.reduce(operator.add, map(lambda m: self.__convert(m, currency), self.moneys), 0)
      return Money(total, currency)

  def __convert(self, a_money, a_currency):
      if a_money.currency == a_currency:
            return a_money.amount
      else:
            return a_money.amount * self._eur_to_usd