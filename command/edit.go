package command

import (
	"database/sql"
	"fmt"

	"github.com/urfave/cli/v2"
)

// Edit updates the given task
func Edit(context *cli.Context, db *sql.DB) error {
	if context.Args().Len() == 0 {
		return cli.NewExitError("Missing task ID argument", 1)
	}

	id := context.Args().First()

	fmt.Println("Editing:", id)

	return nil
}
