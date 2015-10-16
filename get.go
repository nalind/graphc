package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func get(c *cli.Context) {
	ts, _, driver := initTagStore(c)
	id := c.Args().First()
	image, err := lookupImage(ts, id)
	if err != nil {
		fmt.Printf("Failed to locate image %s: %s\n", id, err)
		os.Exit(1)
	}
	loc, err := driver.Get(image.ID, c.GlobalString("context"))
	if err != nil {
		fmt.Printf("Failed to Get %s: %s\n", id, err)
		os.Exit(1)
	}
	fmt.Printf("%s is available at %s\n", id, loc)
}

func init() {
	commands = append(commands, cli.Command{
		Name:      "get",
		ShortName: "g",
		Usage:     "mount an image to the filesystem",
		Action:    get,
	})
}
