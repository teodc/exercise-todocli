package command

import (
	"fmt"

	"github.com/urfave/cli"
)

// Edit updates the given task
func Edit(context *cli.Context) error {
	if len(context.Args()) == 0 {
		return cli.NewExitError("Missing task ID argument", 2)
	}

	id := context.Args().First()

	fmt.Println("Editing:", id)

	return nil
}
