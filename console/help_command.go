package console

import (
	"fmt"
)

type HelpCommand struct {}

func (c *HelpCommand) Names() []string {
	return []string{"help", "man"}
}

func (c *HelpCommand) Execute(args []string) error {
	fmt.Println("help - эта справка")
	fmt.Println("add-payment - добавить коммунальный платёж")
	fmt.Println("payments - показать список платежей")
	return nil
}
