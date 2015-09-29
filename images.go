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
		fmt.Printf("%s", image.ID[:12])
		if nicks, ok := ids[image.ID]; ok {
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
