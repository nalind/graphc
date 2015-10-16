package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func exportImage(c *cli.Context) {
	ts, _, _ := initTagStore(c)
	outfile := os.Stdout
	if c.String("output") != "" {
		newfile, err := os.Open(c.String("output"))
		if err != nil {
			fmt.Printf("Error opening output file: %s\n", err)
			os.Exit(1)
		}
		outfile = newfile
	}
	images := []string{}
	for _, arg := range c.Args() {
		id := lookupID(ts, arg)
		if id != "" {
			images = append(images, id)
		}
	}
	err := ts.ImageExport(images, outfile)
	if err != nil {
		fmt.Printf("Error exporting: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:  "export",
		Usage: "export an image",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "an output file",
			},
		},
		Action: exportImage,
	})
}
