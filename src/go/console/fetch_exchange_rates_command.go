package console

import (
	"fmt"
	"github.com/agl7r/finance/usecases"
	"time"
)

type FetchExchangeRatesCommand struct{}

func (c *FetchExchangeRatesCommand) Names() []string {
	return []string{"fetch-exchange-rates", "rates"}
}

func (c *FetchExchangeRatesCommand) Execute(args []string) error {
	date := time.Now()

	uc := usecases.FetchExchangeRates{}
	rates, err := uc.Execute(date)
	if err != nil {
		return fmt.Errorf("не удалось получить список валют: %s", err)
	}

	fmt.Printf("Курсы валют на %s:\n", rates.Date)

	currencies := []string{"USD", "EUR", "HKD", "CNY"}
	for _, currency := range currencies {
		if rate := rates.ByCode(currency); rate != nil {
			fmt.Printf("%s (%s): %s\n", rate.Name, rate.CharCode, rate.Value)
		}
	}

	return nil
}
