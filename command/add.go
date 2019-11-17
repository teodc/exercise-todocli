package command

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

// Add adds a task to your ToDo list
func Add(context *cli.Context, db *sql.DB) error {
	if context.Args().Len() == 0 {
		return cli.NewExitError("Missing task description argument", 1)
	}

	description := context.Args().First()

	transaction, err := db.Begin()
	if err != nil {
		return cli.NewExitError("Begin transaction error", 1)
	}

	statement, err := db.Prepare("INSERT INTO tasks(content, created) values(?,?)")
	if err != nil {
		return cli.NewExitError("Prepare statement error", 1)
	}

	_, err = statement.Exec(description, time.Now().Unix())
	if err != nil {
		return cli.NewExitError("Execute statement error", 1)
	}

	transaction.Commit()

	fmt.Println("Task added!")

	return nil
}
