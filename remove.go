package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func remove(c *cli.Context) {
	driver := initDriver(c)
	id := c.Args().First()
	if err := driver.Remove(id); err != nil {
		fmt.Printf("Failed to remove %s: %s\n", id, err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands,
		cli.Command{
			Name:      "remove",
			ShortName: "r",
			Usage:     "remove storage for id",
			Action:    remove,
		})
}
