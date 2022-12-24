package console

import (
	"fmt"
	"github.com/agl7r/finance/config"
)

type StatusCommand struct {}

func (c *StatusCommand) Names() []string {
	return []string{"status"}
}

func (c *StatusCommand) Execute(args []string) error {
	fmt.Println(fmt.Sprintf("dsn: %s", config.Config.Dsn))
	fmt.Println(fmt.Sprintf("port: %d", config.Config.Port))
	return nil
}
