package main

import (
	s "tdd/stocks"
	"testing"
)

func TestMultiplication(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	twentyEuros := tenEuros.Times(2)
	expectedResult := s.NewMoney(20, "EUR")

	assertEqual(expectedResult, twentyEuros, t)
}
func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	actualMoneyAfterDivision := originalMoney.Divide(4)
	expectedMoneyAfterDivision := s.NewMoney(1000.5, "KRW")
	assertEqual(expectedMoneyAfterDivision, actualMoneyAfterDivision, t)
}

func TestAddition(t *testing.T) {
	var portfolio s.Portfolio
	var portfolioInDollars s.Money

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")
	fifteenDollars := s.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars, _ = portfolio.Evaluate("USD")

	assertEqual(fifteenDollars, portfolioInDollars, t)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio
	var portfolioInDollars s.Money

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "EUR")
	fifteenDollars := s.NewMoney(16, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars, _ = portfolio.Evaluate("USD")

	assertEqual(fifteenDollars, portfolioInDollars, t)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
    var portfolio s.Portfolio

    oneDollar := s.NewMoney(1, "USD")
    elevenHundredWon := s.NewMoney(1100, "KRW")

    portfolio = portfolio.Add(oneDollar)
    portfolio = portfolio.Add(elevenHundredWon)

    expectedValue := s.NewMoney(2200, "KRW")
    actualValue, _ := portfolio.Evaluate("KRW")

    assertEqual(expectedValue, actualValue, t)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
    var portfolio s.Portfolio

    oneDollar := s.NewMoney(1, "USD")
    oneEuro := s.NewMoney(1, "EUR")
    oneWon := s.NewMoney(1, "KRW")

    portfolio = portfolio.Add(oneDollar)
    portfolio = portfolio.Add(oneEuro)
    portfolio = portfolio.Add(oneWon)

    expectedErrorMessage := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
    _, actualError := portfolio.Evaluate("Kalganid")

    if expectedErrorMessage != actualError.Error() {
        t.Errorf("Expected %s Got %s",
            expectedErrorMessage, actualError.Error())
    }
}

func assertEqual(expected s.Money, actual s.Money, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v",
			expected, actual)
	}
}
