package console

import (
	"errors"
	"fmt"
	"github.com/agl7r/finance/apartment"
	"github.com/bojanz/currency"
	"strconv"
	"strings"
	"time"
)

type AddPaymentCommand struct {}

func (c *AddPaymentCommand) Names() []string {
	return []string{"add-payment"}
}

func (c *AddPaymentCommand) Execute(args []string) error {
	if len(args) < 1 {
		fmt.Print("Укажите месяц в формате 2022-04: ")
		var dateInput string
		_, err := fmt.Scanf("%s", &dateInput)
		if err != nil {
			return errors.New("не удалось получить значение месяца")
		}
		args = append(args, dateInput)
	}
	monthId := args[0]
	date, err := time.Parse("2006-01", monthId)
	if err != nil {
		return errors.New("неверный формат месяца (пример: 2022-04)")
	}
	month := &apartment.Month{Id: date.Format("2006-01")}

	if len(args) < 2 {
		typesOutput := ""
		for i, _type := range *apartment.GetTypes() {
			if i > 0 {
				typesOutput = typesOutput + ", "
			}
			typesOutput = typesOutput + fmt.Sprintf("%d: %s", _type.Id, _type.Title)
		}
		fmt.Printf("Укажите тип в формате идентификатора (%s): ", typesOutput)

		var typeInput string
		_, err := fmt.Scanf("%s", &typeInput)
		if err != nil {
			return errors.New("Не удалось получить значение")
		}
		args = append(args, typeInput)
	}
	typeId, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Укажите тип в формате идентификатора\n")
	}

	types := apartment.GetTypes()
	t := types.GetById(typeId)
	if t == nil {
		return fmt.Errorf("Тип не найден\n")
	}

	if len(args) < 3 {
		fmt.Print("Укажите сумму: ")
		var amountInput string
		_, err := fmt.Scanf("%s", &amountInput)
		if err != nil {
			return errors.New("Не удалось получить значение")
		}
		args = append(args, amountInput)
	}
	amount, err := currency.NewAmount(strings.Replace(args[2], ",", ".", 1), "RUB")

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

	return nil
}
