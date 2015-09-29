package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/image"
)

func listLayers(c *cli.Context) {
	ts, graph, _ := initTagStore(c)
	id := c.Args().First()
	img, err := ts.LookupImage(id)
	if err != nil {
		fmt.Printf("Error locating image: %s\n", err)
		os.Exit(1)
	}
	images := []*image.Image{}
	err = graph.WalkHistory(img, func(img image.Image) error {
		images = append(images, &img)
		return nil
	})
	if err != nil {
		fmt.Printf("Error building list of layers: %s\n", err)
		os.Exit(1)
	}
	ids := ts.ByID()
	for n, _ := range images {
		img = images[len(images)-n-1]
		listLayer(img, &ids)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "layers",
		Usage:  "list layers in an image",
		Action: listLayers,
	})
}
