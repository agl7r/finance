package main

import (
	"fmt"
	"github.com/agl7r/finance/console"
	"github.com/joho/godotenv"
	"os"
	"flag"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "")
	flag.Parse()

	_ = godotenv.Load(".env.local")
	dsn := os.Getenv("DSN")

	if verbose {
		fmt.Printf("dsn: %s\n", dsn)
	}

	args := flag.Args()

	command := console.FindCommand(args)

	commandArgs := []string{}
	if len(args) > 0 {
		commandArgs = args[1:]
	}
	err := console.ExecuteCommand(command, commandArgs)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}
}
