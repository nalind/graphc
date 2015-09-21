package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func imageHeads(c *cli.Context) {
	graph := initGraph(c)
	heads := graph.Heads()
	if heads == nil {
		fmt.Printf("Failed to read heads list\n")
		os.Exit(1)
	}
	for id, _ := range heads {
		fmt.Printf("%s\n", id)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "heads",
		Usage:  "list images with no children",
		Action: imageHeads,
	})
}
