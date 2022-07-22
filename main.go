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
}
