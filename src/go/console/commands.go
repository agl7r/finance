package console

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type Command interface {
	Names() []string
	Execute(args []string) error
}

type EmptyCommand struct {}

func (c *EmptyCommand) Names() []string {
	return []string{}
}

func (c *EmptyCommand) Execute(args []string) error {
	fmt.Println("Введите команду (help для получения списка команд)")
	return nil
}

type NotFoundCommand struct {}

func (c *NotFoundCommand) Names() []string {
	return []string{}
}

func (c *NotFoundCommand) Execute(args []string) error {
	fmt.Println("Команда не найдена (введите help для получения списка команд)")
	return nil
}

func getCommandName(args []string) string {
	if len(args) > 0 {
		commandName := args[0]
		return commandName
	}
	return ""
}

func FindCommand(args []string) Command {
	var commands []Command
	commands = append(commands, new(HelpCommand))
	commands = append(commands, new(StatusCommand))
	commands = append(commands, new(AddPaymentCommand))
	commands = append(commands, new(ShowPayments))
	commands = append(commands, new(FetchExchangeRatesCommand))
	commands = append(commands, new(ServerCommand))

	commandName := getCommandName(args)
	if commandName == "" {
		return new(EmptyCommand)
	}

	for _, command := range commands {
		if slices.Contains(command.Names(), commandName) {
			return command
		}
	}

	return new(NotFoundCommand)
}

func ExecuteCommand(command Command, args []string) error {
	return command.Execute(args)
}

