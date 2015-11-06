package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func systemInfo(c *cli.Context) {
	d, _, _, _ := initDaemon(c)
	info, err := d.SystemInfo()
	if err != nil {
		fmt.Printf("Failed to read system info: %s\n", err)
		os.Exit(1)
	}
	outfile := os.Stdout
	if c.String("output") != "" {
		outfile, err = os.Create(c.String("output"))
		if err != nil {
			fmt.Printf("Error opening %s for output: %s\n", c.String("output"), err)
		}
	}
	j := json.NewEncoder(outfile)
	if err := j.Encode(info); err != nil {
		fmt.Printf("Error writing JSON to output: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "sysinfo",
		Usage:  "display system info",
		Action: systemInfo,
	})
}
