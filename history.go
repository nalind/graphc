package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/image"
)

func imageHistory(c *cli.Context) {
	ts, graph, _ := initTagStore(c)
	id := c.Args().First()
	img, err := lookupImage(ts, id)
	if err != nil {
		fmt.Printf("Error locating image: %s\n", err)
		os.Exit(1)
	}
	images := []string{}
	err = graph.WalkHistory(img, func(img image.Image) error {
		images = append([]string{img.ID}, images...)
		return nil
	})
	if err != nil {
		fmt.Printf("Error building list of layers: %s\n", err)
		os.Exit(1)
	}
	ids := ts.ByID()
	for _, image := range images {
		img, err := lookupImage(ts, image)
		if err != nil {
			fmt.Printf("Error reading image %s: %s\n", image, err)
			os.Exit(1)
		}
		listLayer(img, &ids)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "history",
		Usage:  "list layers in an image",
		Action: imageHistory,
	})
}
