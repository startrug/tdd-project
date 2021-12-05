package stocks

import (
	"errors"
	"strings"
)

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	p = append(p, money)
	return p
}

func (p Portfolio) Evaluate(bank Bank, currency string) (*Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, m := range p {
		if convertedCurrency, err := bank.Convert(m, currency); err == nil {
			total = total + convertedCurrency.amount
		} else {
			failedConversions = append(failedConversions, err.Error())
		}
	}
	if len(failedConversions) == 0 {
		totalMoney := NewMoney(total, currency)
		return &totalMoney, nil
	}
	failures := createFailureMessages(failedConversions)
	return nil, errors.New("Missing exchange rate(s):" + strings.Join(failures, ", "))
}

func createFailureMessages(failedConversions []string) []string {
	failures := make([]string, 0)
	failures = append(failures, failedConversions...)
	return failures
}
