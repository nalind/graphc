package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func inspectImage(c *cli.Context) {
	ts, _, _ := initTagStore(c)
	id := c.Args().First()
	idata, err := lookup(ts, id)
	if err != nil {
		fmt.Printf("Failed to locate image %s: %s\n", id, err)
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
	if err := j.Encode(idata); err != nil {
		fmt.Printf("Error writing JSON to output: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "inspect",
		Usage:  "inspect an image",
		Action: inspectImage,
	})
}
