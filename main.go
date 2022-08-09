package main

import (
	"fmt"
	"github.com/agl7r/finance/console"
	"golang.org/x/exp/slices"
	"os"
)

func main() {
	command := findCommand()
	err := executeCommand(command)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func getCommandName() string {
	if len(os.Args) > 1 {
		commandName := os.Args[1]
		return commandName
	}
	return ""
}

func findCommand() console.Command {
	var commands []console.Command
	commands = append(commands, new(console.HelpCommand))
	commands = append(commands, new(console.AddPaymentCommand))
	commands = append(commands, new(console.ShowPayments))

	c := getCommandName()

	for _, command := range commands {
		if slices.Contains(command.Names(), c) {
			return command
		}
	}

	return new(console.HelpCommand)
}

func executeCommand(command console.Command) error {
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	return command.Execute(args)
}
