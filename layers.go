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
		img = images[len(images) - n - 1]
		fmt.Printf("%s", img.ID[:12])
		if nicks, ok := ids[img.ID]; ok {
			for i, nick := range nicks {
				if i > 0 {
					fmt.Printf(",");
				} else {
					fmt.Printf("\t");
				}
				fmt.Printf("%s", nick);
			}
		}
		fmt.Printf("\n");
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:  "layers",
		Usage: "list layers in an image",
		Action: listLayers,
	})
}
