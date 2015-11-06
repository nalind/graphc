package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/daemon"
	"github.com/docker/docker/daemon/events"
	"github.com/docker/docker/daemon/graphdriver"
	"github.com/docker/docker/graph"
	"github.com/docker/docker/image"
	_ "github.com/docker/docker/pkg/chrootarchive"
	"github.com/docker/docker/pkg/mflag"
	"github.com/docker/docker/pkg/parsers"
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

func initTagStoreAndConfig(c *cli.Context) (*graph.TagStore, *graph.TagStoreConfig, *graph.Graph, graphdriver.Driver) {
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
	return t, &config, g, d
}

func initTagStore(c *cli.Context) (*graph.TagStore, *graph.Graph, graphdriver.Driver) {
	t, _, g, d := initTagStoreAndConfig(c)
	return t, g, d
}

func initDaemon(c *cli.Context) (*daemon.Daemon, *graph.TagStore, *graph.Graph, graphdriver.Driver) {
	t, tc, g, d := initTagStoreAndConfig(c)
	config := &daemon.Config{}
	config.DisableBridge = true
	config.GraphDriver = c.GlobalString("driver")
	config.GraphOptions = c.GlobalStringSlice("storage-opt")
	home := c.GlobalString("home")
	config.Root = home
	config.TrustKeyPath = filepath.Join(c.GlobalString("configdir"), "key.json")
	flags := mflag.NewFlagSet("graphc", mflag.ExitOnError)
	config.InstallFlags(flags, func(string) string { return "" })
	daemon, err := daemon.NewDaemon(config, tc.Registry)
	if err != nil {
		fmt.Printf("Failed to instantiate daemon: %s\n", err)
		os.Exit(1)
	}
	return daemon, t, g, d
}

func lookupID(s *graph.TagStore, name string) string {
	s.Lock()
	defer s.Unlock()
	repo, tag := parsers.ParseRepositoryTag(name)
	if r, exists := s.Repositories[repo]; exists {
		if tag == "" {
			tag = "latest"
		}
		if id, exists := r[tag]; exists {
			return id
		}
	}
	if r, exists := s.Repositories[registry.IndexName+"/"+repo]; exists {
		if tag == "" {
			tag = "latest"
		}
		if id, exists := r[tag]; exists {
			return id
		}
	}
	names := strings.Split(name, "/")
	if len(names) > 1 {
		if r, exists := s.Repositories[strings.Join(names[1:], "/")]; exists {
			if tag == "" {
				tag = "latest"
			}
			if id, exists := r[tag]; exists {
				return id
			}
		}
	}
	return ""
}

func lookup(s *graph.TagStore, name string) (*types.ImageInspect, error) {
	img, err := s.Lookup(name)
	if img != nil {
		return img, err
	}
	id := lookupID(s, name)
	if id != "" {
		return s.Lookup(id)
	}
	return nil, err
}

func lookupImage(s *graph.TagStore, name string) (*image.Image, error) {
	img, err := s.LookupImage(name)
	if img != nil {
		return img, err
	}
	id := lookupID(s, name)
	if id != "" {
		return s.LookupImage(id)
	}
	return nil, err
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
			Value:  "/var/lib/docker",
			Usage:  "home directory for graphdriver storage operations",
			EnvVar: "GRAPHDRIVER_HOME",
		},
		cli.StringFlag{
			Name:  "configdir",
			Value: "/etc/docker",
			Usage: "directory for docker configuration",
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
