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
        total = 0.0
        failures = []
        for m in self.moneys:
            try:
                total += self.__convert(m, currency)
            except KeyError as ke:
                failures.append(ke)

        if len(failures) == 0:
            return Money(total, currency)

        failure_message = ",".join(f.args[0] for f in failures)
        raise Exception("Missing exchange rate(s):[" + failure_message + "]")

  def __convert(self, a_money, a_currency):
      exchange_rates = {'EUR->USD': 1.1, 'USD->KRW': 1100}

      if a_money.currency == a_currency:
            return a_money.amount
      else:
            key = a_money.currency + '->' + a_currency
            return a_money.amount * exchange_rates[key]