package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func remove(c *cli.Context) {
	graph, driver := initGraph(c)
	id := c.Args().First()
	if id == "" {
		fmt.Printf("No image specified.\n")
		os.Exit(1)
	}
	if !graph.Exists(id) {
		fmt.Printf("No image named %s exists.\n", id)
		os.Exit(1)
	}
	if err := driver.Remove(id); err != nil {
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
