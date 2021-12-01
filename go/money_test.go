package main

import (
	"testing"
)

func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{
		amount: 5,
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
	amount float64;
	currency string
}