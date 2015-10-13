package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func imageParent(c *cli.Context) {
	ts, graph, _ := initTagStore(c)
	id := c.Args().First()
	image, err := ts.LookupImage(id)
	if err != nil {
		fmt.Printf("Failed to read image %s: %v\n", id, err)
		os.Exit(1)
	}
	parent, err := graph.Get(image.Parent)
	if err != nil {
		fmt.Printf("Failed to read image %s: %v\n", image.Parent, err)
		os.Exit(1)
	}
	fmt.Printf("%s", parent.ID[:12])
	ids := ts.ByID()
	if nicks, ok := ids[parent.ID]; ok {
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

func init() {
	commands = append(commands, cli.Command{
		Name:   "parent",
		Usage:  "list an image's parent image",
		Action: imageParent,
	})
}
