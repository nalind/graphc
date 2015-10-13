package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/cliconfig"
	"github.com/docker/docker/graph"
	"github.com/docker/docker/pkg/parsers"
	"github.com/docker/docker/registry"
)

func pullImage(c *cli.Context) {
	ts, _, _ := initTagStore(c)
	id := c.Args().First()
	image, tag := parsers.ParseRepositoryTag(id)
	if image == "" {
		tag = ""
	} else {
		if tag == "" {
			tag = "latest"
		}
	}
	conf, err := cliconfig.Load("")
	if err != nil {
		fmt.Printf("Error loading configuration: %s\n", err)
		os.Exit(1)
	}
	index, err := registry.ParseIndexInfo(image)
	if err != nil {
		fmt.Printf("Error finding index for %s: %s\n", image, err)
		os.Exit(1)
	}
	ac := registry.ResolveAuthConfig(conf, index)
	pullconfig := &graph.ImagePullConfig{
		AuthConfig: &ac,
		OutStream:  os.Stderr,
	}
	err = ts.Pull(image, tag, pullconfig)
	if err != nil {
		fmt.Printf("Error pulling image: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "pull",
		Usage:  "pull an image",
		Action: pullImage,
	})
}
