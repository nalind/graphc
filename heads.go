package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func imageHeads(c *cli.Context) {
	ts, graph, _ := initTagStore(c)
	heads := graph.Heads()
	if heads == nil {
		fmt.Printf("Failed to read heads list\n")
		os.Exit(1)
	}
	ids := ts.ByID()
	for _, img := range heads {
		listLayer(img, &ids)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "heads",
		Usage:  "list images with no children",
		Action: imageHeads,
	})
}
