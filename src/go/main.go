package main

import (
	"flag"
	"fmt"
	"github.com/agl7r/finance/config"
	"github.com/agl7r/finance/console"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "")
	flag.Parse()

	_ = godotenv.Load(".env.local")
	config.Config.Dsn = os.Getenv("DSN")
	config.Config.Port, _ = strconv.Atoi(os.Getenv("PORT"))

	if verbose {
		fmt.Printf("dsn: %s\n", config.Config.Dsn)
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
