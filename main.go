package main

import (
	"fmt"
	"github.com/agl7r/finance/apartment"
	"github.com/bojanz/currency"
)

func main() {
	types := apartment.GetTypes()

	for _, t := range *types {
		fmt.Printf("%d %s\n", t.Id, t.Title)
	}

	internetType := types.GetById(5)
	fmt.Printf("\nInternet: %d %s\n", internetType.Id, internetType.Title)

	fmt.Printf("\n")
	month := apartment.CurrentMonth()
	fmt.Printf("Current month: %s\n", month.Id)
	nextMonth := month.NextMonth()
	fmt.Printf("Next month: %s\n", nextMonth.Id)
	nextNextMonth := nextMonth.NextMonth()
	fmt.Printf("Next next month: %s\n", nextNextMonth.Id)

	fmt.Printf("\n")

	repository := apartment.NewPaymentRepository()

	payments, _ := repository.FindAll()

	for _, payment := range payments {
		fmt.Printf("%s %d %s\n", payment.Month.Id, payment.Type.Id, payment.Amount)
	}
	fmt.Printf("Total: %s\n", payments.GetTotal())

	a, _ := currency.NewAmount("300", "RUB")
	month.Id = "2022-01"
	p := apartment.NewCommunalPayment(month, internetType, a)
	_ = repository.Save(p)
}
