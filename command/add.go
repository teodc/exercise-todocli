package command

import (
	"fmt"

	"github.com/urfave/cli"
)

// Add adds a task to your ToDo list
func Add(context *cli.Context) error {
	if len(context.Args()) == 0 {
		return cli.NewExitError("Missing task description argument", 2)
	}

	description := context.Args().First()

	fmt.Println("Adding:", description)

	return nil
}
