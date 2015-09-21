package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func put(c *cli.Context) {
	driver := initDriver(c)
	id := c.Args().First()
	if err := driver.Put(id); err != nil {
		fmt.Printf("Failed to Put %s: %s\n", id, err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:      "put",
		ShortName: "p",
		Usage:     "unmount an image from the filesystem",
		Action:    put,
	})
}
