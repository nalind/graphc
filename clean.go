package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func clean(c *cli.Context) {
	driver := initDriver(c)
	if err := driver.Cleanup(); err != nil {
		fmt.Printf("Failed to clean up: %s\n", err)
	}
}

func init() {
	commands = append(commands,
		cli.Command{
			Name:   "clean",
			Usage:  "clean up stateful artifacts",
			Action: clean,
		})
}
