package main

import (
	"fmt"
	"github.com/agl7r/finance/apartment"
)

func main() {
	types := apartment.GetTypes()

	for _, t := range types {
		fmt.Printf("%d %s\n", t.Id, t.Title)
	}

	fmt.Printf("\n")
	month := apartment.CurrentMonth()
	fmt.Printf("Current month: %s\n", month.Id)
	nextMonth := month.NextMonth()
	fmt.Printf("Next month: %s\n", nextMonth.Id)
	nextNextMonth := nextMonth.NextMonth()
	fmt.Printf("Next next month: %s\n", nextNextMonth.Id)
}
