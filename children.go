package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func imageChildren(c *cli.Context) {
	graph := initGraph(c)
	id := c.Args().First()
	m := graph.ByParent()
	if m == nil {
		fmt.Printf("Failed to map of images\n")
		os.Exit(1)
	}
	images, found := m[id]
	if !found {
		fmt.Printf("No child images for %s\n", id)
		os.Exit(1)
	}
	for _, image := range images {
		fmt.Printf("%s\n", image.ID)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "children",
		Usage:  "list an image's child images",
		Action: imageChildren,
	})
}
