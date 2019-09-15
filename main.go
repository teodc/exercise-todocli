package main

import (
	"log"
	"os"
	"sort"

	"github.com/teodc/todo/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "ToDo"
	app.Usage = "Keep track of what you have to do"
	app.Version = "0.0.0"
	app.Authors = []cli.Author{cli.Author{Name: "teodc"}}

	app.Action = func(context *cli.Context) error {
		return cli.ShowAppHelp(context)
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l", "ls"},
			Usage:   "List my tasks",
			Action:  command.List,
		},
		{
			Name:      "add",
			Aliases:   []string{"a"},
			Usage:     "Add a task",
			ArgsUsage: "[task description]",
			Action:    command.Add,
		},
		{
			Name:      "edit",
			Aliases:   []string{"e"},
			Usage:     "Edit the given task",
			ArgsUsage: "[task ID]",
			Action:    command.Edit,
		},
		{
			Name:      "complete",
			Aliases:   []string{"c"},
			Usage:     "Mark the given task as done",
			ArgsUsage: "[task ID]",
			Action:    command.Complete,
		},
		{
			Name:      "delete",
			Aliases:   []string{"d", "rm"},
			Usage:     "Delete the given task",
			ArgsUsage: "[task ID]",
			Action:    command.Delete,
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
