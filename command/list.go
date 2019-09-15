package command

import (
	"fmt"

	"github.com/urfave/cli"
)

// List shows the tasks in your ToDo list
func List(context *cli.Context) error {
	fmt.Println("Listing tasks...")

	return nil
}
