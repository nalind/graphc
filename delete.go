package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func deleteImage(c *cli.Context) {
	ts, g, _ := initTagStore(c)
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
	if err := g.Delete(image.ID); err != nil {
		fmt.Printf("Failed to delete %s: %s\n", id, err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands,
		cli.Command{
			Name:      "delete",
			ShortName: "d",
			Usage:     "delete an image",
			Action:    deleteImage,
		})
}
