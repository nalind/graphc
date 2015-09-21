package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func create(c *cli.Context) {
	driver := initDriver(c)
	id := c.Args().First()
	if err := driver.Create(id, c.String("parent")); err != nil {
		fmt.Printf("Failed to create %s: %s\n", id, err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:      "create",
		ShortName: "c",
		Usage:     "create a new storage for id",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parent, p",
				Value: "",
				Usage: "an id of which the new image will initially be a copy",
			},
		},
		Action: create,
	})
}
