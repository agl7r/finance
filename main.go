package main

import (
	"fmt"
	"github.com/agl7r/finance/apartment"
	"github.com/bojanz/currency"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	switch getCommand() {
	case "add-payment":
		handleAddPaymentCommand(os.Args[2:])
	case "payments", "show-payments", "list-payments":
		handleShowPayments(os.Args[2:])
	}
}

func getCommand() string {
	if len(os.Args) > 1 {
		command := os.Args[1]
		return command
	}
	return ""
}

func handleAddPaymentCommand(args []string) {
	monthId := args[0]

	if len(args) < 1 {
		fmt.Print("Укажите месяц в формате 2022-04\n")
		return
	}
	date, err := time.Parse("2006-01", monthId)
	if err != nil {
		fmt.Printf("Укажите месяц в формате 2022-04\n")
		return
	}
	month := &apartment.Month{Id: date.Format("2006-01")}

	if len(args) < 2 {
		fmt.Print("Укажите тип в формате идентификатора\n")
		return
	}
	typeId, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Print("Укажите тип в формате идентификатора\n")
		return
	}

	types := apartment.GetTypes()
	t := types.GetById(typeId)
	if t == nil {
		fmt.Printf("Тип не найден\n")
		return
	}

	if len(args) < 3 {
		fmt.Print("Укажите сумму\n")
		return
	}
	amount, err := currency.NewAmount(args[2], "RUB")

	payment := apartment.NewCommunalPayment(month, t, amount)

	repository := apartment.NewPaymentRepository()
	err = repository.Save(payment)
	if err != nil {
		fmt.Printf("Не удалось сохранить платёж\n")
	}

	fmt.Printf("Месяц: %s\n", month.Id)
	fmt.Printf("Тип: %s\n", t.Title)
	fmt.Printf("Сумма: %s\n", amount)

	fmt.Printf("Платёж добавлен\n")
}

func handleShowPayments(args []string) {
	repository := apartment.NewPaymentRepository()
	payments, _ := repository.FindAll()

	showMode := "table"
	if len(args) > 0 {
		showMode = args[0]
	}

	if showMode == "tree" {
		PrintTree(payments)
	} else {
		PrintTable(payments)
	}
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

	for y, subtree := range tree {
		fmt.Printf("%s", y)

		types := *apartment.GetTypes()
		for _, _type := range types {
			fmt.Printf("%15s", _type.Title)
		}
		fmt.Printf("%15s\n", "Итого")

		var keys []string
		for key, _ := range subtree {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, m := range keys {
			fmt.Printf("  %s", m)

			monthPayments := tree[y][m]

			for _, _type := range types {
				fmt.Printf("%15s", monthPayments.FindByType(_type).Amount.Number())
			}

			fmt.Printf("%15s", monthPayments.GetTotal().Number())
			fmt.Printf("\n")
		}
	}
}
