package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/context"
	"github.com/docker/docker/pkg/parsers"
	"github.com/docker/docker/runconfig"
)

func importImage(c *cli.Context) {
	ts, _, _ := initTagStore(c)
	ctx := context.Context{}
	infile := os.Stdin
	src := c.String("input")
	if src == "" {
		src = "-"
	} else {
		infile = nil
	}
	repo, tag := parsers.ParseRepositoryTag(c.String("tag"))
	if repo == "" {
		tag = ""
	} else {
		if tag == "" {
			tag = "latest"
		}
	}
	msg := c.String("message")
	runconfig := &runconfig.Config{}
	err := ts.Import(ctx, src, repo, tag, msg, infile, nil, runconfig)
	if err != nil {
		fmt.Printf("Error writing JSON to output: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:  "import",
		Usage: "import an image",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "input, i",
				Value: "",
				Usage: "input file or URL",
			},
			cli.StringFlag{
				Name:  "tag, t",
				Value: "",
				Usage: "tag name",
			},
			cli.StringFlag{
				Name:  "message, m",
				Value: "",
				Usage: "commit message",
			},
		},
		Action: importImage,
	})
}
