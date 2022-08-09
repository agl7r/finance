package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"os"
)

func main() {
	var commands []Command
	commands = append(commands, new(HelpCommand))
	commands = append(commands, new(AddPaymentCommand))
	commands = append(commands, new(ShowPayments))

	c := getCommand()

	for _, command := range commands {
		if slices.Contains(command.Names(), c) {
			err := command.Execute(os.Args[2:])
			if err != nil {
				fmt.Printf("%s", err)
			}
		}
	}
}

func getCommand() string {
	if len(os.Args) > 1 {
		command := os.Args[1]
		return command
	}
	return ""
}

type Command interface {
	Names() []string
	Execute(args []string) error
}
