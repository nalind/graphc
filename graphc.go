package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/daemon/graphdriver"
	"github.com/docker/docker/graph"
)

func initDriver(c *cli.Context) graphdriver.Driver {
	graphdriver.DefaultDriver = c.GlobalString("driver")
	homedir := c.GlobalString("home")
	drv, err := graphdriver.New(homedir, c.GlobalStringSlice("storage-opt"))
	if err != nil {
		fmt.Printf("Failed to instantiate graphdriver: %s\n", err)
		os.Exit(1)
	}
	if c.GlobalBool("debug") {
		fmt.Printf("[DEBUG] Using driver %s.\n[DEBUG] %g\n[DEBUG] Home directory: %s\n", drv.String(), drv.Status(), homedir)
	}
	return drv
}

func initGraph(c *cli.Context) *graph.Graph {
	drv := initDriver(c)
	homedir := c.GlobalString("home") + "/graph/"
	g, err := graph.NewGraph(homedir, drv)
	if err != nil {
		fmt.Printf("Failed to instantiate graph: %s\n", err)
		os.Exit(1)
	}
	return g
}

var commands []cli.Command

func main() {
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
			Name:   "driver, s",
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
}
