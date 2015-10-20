package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/daemon/events"
	"github.com/docker/docker/daemon/graphdriver"
	"github.com/docker/docker/graph"
	"github.com/docker/docker/pkg/parsers"
)

func initTarget(c *cli.Context) (*graph.TagStore, *graph.Graph, graphdriver.Driver) {
	driver := c.String("target-driver")
	if driver == "" {
		driver = c.GlobalString("driver")
	}
	home := c.String("target-home")
	if home == "" {
		home = c.GlobalString("home")
	}
	opts := c.StringSlice("target-storage-opt")
	if len(opts) == 0 {
		opts = c.GlobalStringSlice("storage-opt")
	}
	graphdriver.DefaultDriver = driver
	if graphdriver.DefaultDriver == "" {
		fmt.Printf("No graphdriver specified.\n")
		os.Exit(1)
	}
	homedir := home
	d, err := graphdriver.New(homedir, opts, nil, nil)
	if err != nil {
		fmt.Printf("Failed to instantiate graphdriver: %s\n", err)
		os.Exit(1)
	}
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "[DEBUG] Using target driver %s.\n[DEBUG] %g\n[DEBUG] Target home directory: %s\n", d.String(), d.Status(), homedir)
	}

	homedir = filepath.Join(home, "graph")
	g, err := graph.NewGraph(homedir, d, nil, nil)
	if err != nil {
		fmt.Printf("Failed to instantiate graph: %s\n", err)
		os.Exit(1)
	}
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "[DEBUG] %d images.\n", len(g.Map()))
	}

	tsfile := filepath.Join(home, "repositories-"+d.String())
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

func copyLayer(c *cli.Context) {
	ts, _, _ := initTagStore(c)
	ts2, _, _ := initTarget(c)
	id := c.Args().First()
	layer, err := lookupImage(ts, id)
	if err != nil {
		fmt.Printf("Failed to locate layer %s: %s\n", id, err)
		os.Exit(1)
	}
	r, w := io.Pipe()
	go func() {
		err = ts.ImageExport([]string{layer.ID}, w)
		if err != nil {
			fmt.Printf("Error exporting: %s\n", err)
			os.Exit(1)
		}
		w.Close()
		if c.GlobalBool("debug") {
			fmt.Fprintf(os.Stderr, "[DEBUG] Finished exporting %s.\n", layer.ID)
		}
	}()
	err = ts2.Load(r, os.Stdout)
	if err != nil {
		fmt.Printf("Error importing: %s\n", err)
		os.Exit(1)
	}
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "[DEBUG] Finished importing %s.\n", layer.ID)
	}
	ids := ts.ByID()[layer.ID]
	for _, tags := range ids {
		repo, tag := parsers.ParseRepositoryTag(tags)
		if repo != "" {
			if tag == "" {
				tag = "latest"
			}
			if c.GlobalBool("debug") {
				fmt.Fprintf(os.Stderr, "[DEBUG] Tagging %s as %s:%s.\n", layer.ID, repo, tag)
			}
			ts2.Tag(repo, tag, layer.ID, c.Bool("force-tag"))
		}
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "copy",
		Usage:  "copy image to new storage",
		Action: copyLayer,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "target-home",
				Value:  "/var/lib/docker/",
				Usage:  "home directory for graphdriver storage operations",
				EnvVar: "GRAPHDRIVER_HOME",
			},
			cli.StringFlag{
				Name:   "target-storage-driver, target-driver, t",
				Value:  "",
				Usage:  "storage driver to use",
				EnvVar: "GRAPHDRIVER_BACKEND",
			},
			cli.StringSliceFlag{
				Name:   "target-storage-opt",
				Value:  &cli.StringSlice{},
				Usage:  "set storage driver options",
				EnvVar: "GRAPHDRIVER_OPTIONS",
			},
			cli.BoolFlag{
				Name:  "force-tag, f",
				Usage: "apply tags, even if they already point elsewhere",
			},
		},
	})
}
