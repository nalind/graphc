package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func remove(c *cli.Context) {
	ts, _, driver := initTagStore(c)
	id := c.Args().First()
	if id == "" {
		fmt.Printf("No image specified.\n")
		os.Exit(1)
	}
	image, err := lookupImage(ts, id)
	if err != nil {
		fmt.Printf("Error locating image %s: %s.\n", id, err)
		os.Exit(1)
	}
	if err := driver.Remove(image.ID); err != nil {
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
