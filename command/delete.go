package command

import (
	"database/sql"
	"fmt"

	"github.com/urfave/cli/v2"
)

// Delete remoevs the given task from your ToDo list
func Delete(context *cli.Context, db *sql.DB) error {
	if context.Args().Len() == 0 {
		return cli.NewExitError("Missing task ID argument", 1)
	}

	id := context.Args().First()

	fmt.Println("Deleting:", id)

	return nil
}
