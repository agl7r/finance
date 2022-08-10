package main

import (
	"fmt"
	"github.com/agl7r/finance/console"
)

func main() {
	command := console.FindCommand()
	err := console.ExecuteCommand(command)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

