package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func imageParent(c *cli.Context) {
	ts, graph, _ := initTagStore(c)
	id := c.Args().First()
	image, err := lookupImage(ts, id)
	if err != nil {
		fmt.Printf("Failed to read image %s: %v\n", id, err)
		os.Exit(1)
	}
	parent, err := graph.Get(image.Parent)
	if err != nil {
		fmt.Printf("Failed to read image %s: %v\n", image.Parent, err)
		os.Exit(1)
	}
	ids := ts.ByID()
	listLayer(parent, &ids)
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "parent",
		Usage:  "list an image's parent image",
		Action: imageParent,
	})
}
