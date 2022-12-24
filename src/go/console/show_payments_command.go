package console

import (
	"errors"
	"fmt"
	"github.com/agl7r/finance/apartment"
	"sort"
	"strconv"
	"log"
)

type ShowPayments struct {}

func (c *ShowPayments) Names() []string  {
	return []string{"payments","show-payments","payments-list"}
}

func (c *ShowPayments) Execute(args []string) error {
	repository := apartment.NewPaymentRepository()

	year := 0
	showMode := "table"

	var payments apartment.CommunalPayments

	for _, arg := range args {
		if arg == "table" || arg == "tree" {
			showMode = arg
		} else {
			var err error
			year, err = strconv.Atoi(arg)
			if err != nil {
				return errors.New("не удалось распознать год")
			}
			if year < 100 {
				year += 2000
			}
		}
	}

	var err error
	if year > 0 {
		payments, err = repository.FindByYear(year)
		if err != nil {
			log.Panic(err)
		}
	} else {
		payments, err = repository.FindAll()
		if err != nil {
			log.Panic(err)
		}
	}

	if len(payments) == 0 {
		fmt.Println("Оплаты не найдены.")
	}

	if showMode == "tree" {
		PrintTree(payments)
	} else {
		PrintTable(payments)
	}

	return nil
}

func BuildTree(payments apartment.CommunalPayments) map[string]map[string]apartment.CommunalPayments {
	tree := make(map[string]map[string]apartment.CommunalPayments)

	for _, payment := range payments {
		if tree[payment.Month.Y()] == nil {
			tree[payment.Month.Y()] = make(map[string]apartment.CommunalPayments)
		}
		tree[payment.Month.Y()][payment.Month.M()] = append(tree[payment.Month.Y()][payment.Month.M()], payment)
	}

	return tree
}

func PrintTree(payments apartment.CommunalPayments) {
	tree := BuildTree(payments)

	for y, subtree := range tree {
		fmt.Printf("%s\n", y)
		for m, monthPayments := range subtree {
			fmt.Printf("    %s\n", m)
			for _, payment := range monthPayments {
				fmt.Printf("      %s: %s\n", payment.Type, payment.Amount)
			}
			fmt.Printf("          Итого: %s\n\n", monthPayments.GetTotal())
		}
	}
}

func PrintTable(payments apartment.CommunalPayments) {
	tree := BuildTree(payments)

	var years []string
	for key, _ := range tree {
		years = append(years, key)
	}
	sort.Strings(years)

	for _, year := range years {
		fmt.Printf("%s", year)

		subtree := tree[year]

		types := *apartment.GetTypes()
		for _, _type := range types {
			fmt.Printf("%15s", _type.Title)
		}
		fmt.Printf("%15s\n", "Итого")

		var months []string
		for key, _ := range subtree {
			months = append(months, key)
		}
		sort.Strings(months)

		for _, month := range months {
			fmt.Printf("  %s", month)

			monthPayments := tree[year][month]

			for _, _type := range types {
				payment := monthPayments.FindByType(_type)
				if payment != nil {
					fmt.Printf("%15s", monthPayments.FindByType(_type).Amount.Number())
				} else {
					fmt.Printf("%15s", " ")
				}
			}

			fmt.Printf("%15s", monthPayments.GetTotal().Number())
			fmt.Printf("\n")
		}
	}
}
