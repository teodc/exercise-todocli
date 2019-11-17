package command

import (
	"database/sql"
	"fmt"

	"github.com/urfave/cli/v2"
)

// List shows the uncompleted tasks in your ToDo list
func List(context *cli.Context, db *sql.DB) error {
	rows, err := db.Query("SELECT id, content, created_at from tasks")
	if err != nil {
		return cli.NewExitError("Query error", 1)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string
		var createdAt int
		err = rows.Scan(&id, &content, &createdAt)
		if err != nil {
			return cli.NewExitError("Scan row error", 1)
		}
		fmt.Println(id, content, createdAt)
	}

	err = rows.Err()
	if err != nil {
		return cli.NewExitError("Error", 1)
	}

	return nil
}
