package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func listImages(c *cli.Context) {
	ts, _, _ := initTagStore(c)
	ids := ts.ByID()
	images, err := ts.Images(c.GlobalString("filter"), c.Args().First(), c.GlobalBool("all"))
	if err != nil {
		fmt.Printf("Error locating images: %s\n", err)
		os.Exit(1)
	}
	for _, image := range images {
		img, err := lookupImage(ts, image.ID)
		if err != nil {
			fmt.Printf("Error locating image %s: %s\n", image.ID, err)
			os.Exit(1)
		}
		listLayer(img, &ids)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:  "images",
		Usage: "list images",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "all, a",
				Usage: "list all images",
			},
			cli.StringFlag{
				Name:  "filter",
				Value: "",
				Usage: "JSON filter",
			},
		},
		Action: listImages,
	})
}
