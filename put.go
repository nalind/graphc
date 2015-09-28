package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func put(c *cli.Context) {
	ts, _, driver := initTagStore(c)
	id := c.Args().First()
	image, err := ts.LookupImage(id)
	if err != nil {
		fmt.Printf("Failed to locate image %s: %s\n", id, err)
		os.Exit(1)
	}
	if err := driver.Put(image.ID); err != nil {
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
