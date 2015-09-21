package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func jsonImage(c *cli.Context) {
	graph := initGraph(c)
	id := c.Args().First()
	b, err := graph.RawJSON(id)
	if err != nil {
		fmt.Printf("Failed to obtain a JSON representation of image %s: %s\n", id, err)
		os.Exit(1)
	}
	outfile := os.Stdout
	if c.String("output") != "" {
		outfile, err = os.Create(c.String("output"))
		if err != nil {
			fmt.Printf("Error opening %s for output: %s\n", c.String("output"), err)
		}
	}
	if n, err := outfile.Write(b); n != len(b) || err != nil {
		if err != nil {
			fmt.Printf("Error writing JSON to output: %s\n", err)
		} else {
			fmt.Printf("Error writing JSON to output\n")
		}
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:  "json-image",
		Usage: "export an image as JSON",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "an output file",
			},
		},
		Action: jsonImage,
	})
}
