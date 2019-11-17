package main

import (
	"database/sql"
	"log"
	"os"
	"sort"

	"github.com/teodc/todo/command"
	"github.com/teodc/todo/storage"
	"github.com/urfave/cli/v2"
)

// Store struct
type (
	Store struct {
		db *sql.DB
	}
)

func main() {
	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := cli.NewApp()
	app.Name = "ToDo"
	app.Usage = "Keep track of what you have to do"
	app.Version = "0.0.0"
	app.Authors = []*cli.Author{
		&cli.Author{Name: "teodc"},
	}

	app.Action = func(context *cli.Context) error {
		return cli.ShowAppHelp(context)
	}

	app.Commands = []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l", "ls"},
			Usage:   "List my tasks",
			Action: func(context *cli.Context) error {
				return command.List(context, db)
			},
		},
		{
			Name:      "add",
			Aliases:   []string{"a"},
			Usage:     "Add a task",
			ArgsUsage: "[task description]",
			Action: func(context *cli.Context) error {
				return command.Add(context, db)
			},
		},
		{
			Name:      "edit",
			Aliases:   []string{"e"},
			Usage:     "Edit the given task",
			ArgsUsage: "[task ID]",
			Action: func(context *cli.Context) error {
				return command.Edit(context, db)
			},
		},
		{
			Name:      "complete",
			Aliases:   []string{"c"},
			Usage:     "Mark the given task as done",
			ArgsUsage: "[task ID]",
			Action: func(context *cli.Context) error {
				return command.Complete(context, db)
			},
		},
		{
			Name:      "delete",
			Aliases:   []string{"d", "rm"},
			Usage:     "Delete the given task",
			ArgsUsage: "[task ID]",
			Action: func(context *cli.Context) error {
				return command.Delete(context, db)
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))

	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
