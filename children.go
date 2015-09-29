package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func imageChildren(c *cli.Context) {
	ts, graph, _ := initTagStore(c)
	id := c.Args().First()
	image, _ := ts.LookupImage(id)
	if image != nil {
		id = image.ID
	}
	m := graph.ByParent()
	if m == nil {
		fmt.Printf("Failed to read map of images\n")
		os.Exit(1)
	}
	images, found := m[id]
	if !found {
		fmt.Printf("No child images for %s\n", id)
		os.Exit(1)
	}
	ids := ts.ByID()
	for _, image := range images {
		fmt.Printf("%s", image.ID[:12])
		if nicks, ok := ids[image.ID]; ok {
			for i, nick := range nicks {
				if i > 0 {
					fmt.Printf(",");
				} else {
					fmt.Printf("\t");
				}
				fmt.Printf("%s", nick);
			}
		}
		fmt.Printf("\n");
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "children",
		Usage:  "list an image's child images",
		Action: imageChildren,
	})
}
