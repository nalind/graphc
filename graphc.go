package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/daemon/events"
	"github.com/docker/docker/daemon/graphdriver"
	"github.com/docker/docker/graph"
	"github.com/docker/docker/image"
	_ "github.com/docker/docker/pkg/chrootarchive"
	"github.com/docker/docker/pkg/reexec"
	"github.com/docker/docker/registry"
)

func listLayer(img *image.Image, tags *map[string][]string) {
	fmt.Printf("%s", img.ID[:12])
	if tags != nil {
		if taglist, ok := (*tags)[img.ID]; ok {
			for i, tag := range taglist {
				if i > 0 {
					fmt.Printf(",")
				} else {
					fmt.Printf("\t")
				}
				fmt.Printf("%s", tag)
			}
		}
	}
	fmt.Printf("\n")
}

func initDriver(c *cli.Context) graphdriver.Driver {
	graphdriver.DefaultDriver = c.GlobalString("driver")
	if graphdriver.DefaultDriver == "" {
		fmt.Printf("No graphdriver specified.\n")
		os.Exit(1)
	}
	homedir := c.GlobalString("home")
	drv, err := graphdriver.New(homedir, c.GlobalStringSlice("storage-opt"), nil, nil)
	if err != nil {
		fmt.Printf("Failed to instantiate graphdriver: %s\n", err)
		os.Exit(1)
	}
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "[DEBUG] Using driver %s.\n[DEBUG] %g\n[DEBUG] Home directory: %s\n", drv.String(), drv.Status(), homedir)
	}
	return drv
}

func initGraph(c *cli.Context) (*graph.Graph, graphdriver.Driver) {
	drv := initDriver(c)
	homedir := filepath.Join(c.GlobalString("home"), "graph")
	g, err := graph.NewGraph(homedir, drv, nil, nil)
	if err != nil {
		fmt.Printf("Failed to instantiate graph: %s\n", err)
		os.Exit(1)
	}
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "[DEBUG] %d images.\n", len(g.Map()))
	}
	return g, drv
}

func initRegistry() *registry.Service {
	return registry.NewService(nil)
}

func initTagStore(c *cli.Context) (*graph.TagStore, *graph.Graph, graphdriver.Driver) {
	g, d := initGraph(c)
	tsfile := filepath.Join(c.GlobalString("home"), "repositories-"+d.String())
	e := events.New()
	r := initRegistry()
	config := graph.TagStoreConfig{
		Events:   e,
		Graph:    g,
		Registry: r,
	}
	t, err := graph.NewTagStore(tsfile, &config)
	if err != nil {
		fmt.Printf("Failed to instantiate tag store: %s\n", err)
		os.Exit(1)
	}
	return t, g, d
}

var commands []cli.Command

func main() {
	if reexec.Init() {
		return
	}
	graphc := cli.NewApp()
	graphc.Name = "graphc"
	graphc.Usage = "manage graphc storage"
	graphc.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "home",
			Value:  "/var/lib/docker/",
			Usage:  "home directory for graphdriver storage operations",
			EnvVar: "GRAPHDRIVER_HOME",
		},
		cli.StringFlag{
			Name:   "storage-driver, driver, s",
			Value:  "",
			Usage:  "storage driver to use",
			EnvVar: "GRAPHDRIVER_BACKEND",
		},
		cli.StringSliceFlag{
			Name:   "storage-opt",
			Value:  &cli.StringSlice{},
			Usage:  "set storage driver options",
			EnvVar: "GRAPHDRIVER_OPTIONS",
		},
		cli.StringFlag{
			Name:  "context, c",
			Value: "",
			Usage: "optional mountlabel (SELinux context)",
		},
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "print debugging information",
		},
	}
	graphc.EnableBashCompletion = true
	graphc.Commands = commands

	graphc.Run(os.Args)

	os.Exit(0)
}
