package command

import (
	"fmt"

	"github.com/urfave/cli"
)

// Delete remoevs the given task from your ToDo list
func Delete(context *cli.Context) error {
	if len(context.Args()) == 0 {
		return cli.NewExitError("Missing task ID argument", 2)
	}

	id := context.Args().First()

	fmt.Println("Deleting:", id)

	return nil
}
