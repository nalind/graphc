package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func loadImage(c *cli.Context) {
	ts, _, _ := initTagStore(c)
	infile := os.Stdin
	if c.String("input") != "" {
		f, err := os.Open(c.String("input"))
		if err != nil {
			fmt.Printf("Error opening file to load: %s\n", err)
			os.Exit(1)
		}
		infile = f
	}
	err := ts.Load(infile, os.Stdout)
	if err != nil {
		fmt.Printf("Error loading image: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:  "load",
		Usage: "load an image",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "input, i",
				Value: "",
				Usage: "an input file",
			},
		},
		Action: loadImage,
	})
}
