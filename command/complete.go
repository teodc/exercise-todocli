package command

import (
	"fmt"

	"github.com/urfave/cli"
)

// Complete marks the given task as completed
func Complete(context *cli.Context) error {
	if len(context.Args()) == 0 {
		return cli.NewExitError("Missing task ID argument", 2)
	}

	id := context.Args().First()

	fmt.Println("Completing:", id)

	return nil
}
