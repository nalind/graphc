package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func inspectContainer(c *cli.Context) {
	d, _, _, _ := initDaemon(c)
	id := c.Args().First()
	cdata, err := d.ContainerInspect(id, true)
	if err != nil {
		fmt.Printf("Failed to locate container %s: %s\n", id, err)
		os.Exit(1)
	}
	outfile := os.Stdout
	if c.String("output") != "" {
		outfile, err = os.Create(c.String("output"))
		if err != nil {
			fmt.Printf("Error opening %s for output: %s\n", c.String("output"), err)
		}
	}
	j := json.NewEncoder(outfile)
	if err := j.Encode(cdata); err != nil {
		fmt.Printf("Error writing JSON to output: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "cinspect",
		Usage:  "inspect a container",
		Action: inspectContainer,
	})
}
