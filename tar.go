package main

import (
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

func tarLayer(c *cli.Context) {
	ts, graph, _ := initTagStore(c)
	id := c.Args().First()
	image, err := lookupImage(ts, id)
	if image == nil || err != nil {
		if err != nil {
			fmt.Printf("Failed to find image layer %s: %v\n", id, err)
		} else {
			fmt.Printf("Failed to find image layer %s\n", id)
		}
		os.Exit(1)
	}
	tar, err := graph.TarLayer(image)
	if tar == nil || err != nil {
		if err != nil {
			fmt.Printf("Failed to tar image layer %s: %v\n", id, err)
		} else {
			fmt.Printf("Failed to tar image layer %s\n", id)
		}
		os.Exit(1)
	}
	outfile := os.Stdout
	if c.String("output") != "" {
		outfile, err = os.Create(c.String("output"))
		if err != nil {
			fmt.Printf("Error opening %s for output: %s\n", c.String("output"), err)
		}
	}
	if _, err := io.Copy(outfile, tar); err != nil {
		if err != nil {
			fmt.Printf("Error writing image to output: %s\n", err)
		} else {
			fmt.Printf("Error writing image to output\n")
		}
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:  "tar",
		Usage: "produce a layer tarball",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "an output file",
			},
		},
		Action: tarLayer,
	})
}
