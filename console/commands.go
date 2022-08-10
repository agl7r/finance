package console

import (
	"fmt"
	"golang.org/x/exp/slices"
	"os"
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

func getCommandName() string {
	if len(os.Args) > 1 {
		commandName := os.Args[1]
		return commandName
	}
	return ""
}

func FindCommand() Command {
	var commands []Command
	commands = append(commands, new(HelpCommand))
	commands = append(commands, new(AddPaymentCommand))
	commands = append(commands, new(ShowPayments))

	commandName := getCommandName()
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

func ExecuteCommand(command Command) error {
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	return command.Execute(args)
}

