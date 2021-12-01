package main

import (
	"testing"
)

func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{
		amount:   5,
		currency: "USD",
	}
	tenner := fiver.Times(2)
	expectedResult := Money{amount: 10, currency: "USD"}
	assertEqual(expectedResult, tenner, t)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	twentyEuros := tenEuros.Times(2)
	expectedResult := Money{amount: 20, currency: "EUR"}

	assertEqual(expectedResult, twentyEuros, t)
}
func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(expectedMoneyAfterDivision, actualMoneyAfterDivision, t)
}

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(fifteenDollars, portfolioInDollars, t)
}

func assertEqual(expected Money, actual Money, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v",
			expected, actual)
	}
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}

type Money struct {
	amount   float64
	currency string
}

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	p = append(p, money)
	return p
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total = total + m.amount
	}
	return Money{amount: total, currency: currency}
}
