package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

func searchImages(c *cli.Context) {
	s := initRegistry()
	term := c.Args().First()
	results, err := s.Search(term, nil, nil)
	if err != nil {
		fmt.Printf("Error searching registry: %s\n", err)
		os.Exit(1)
	}
	maxcount := 0
	maxname := 0
	for _, result := range results.Results {
		s := fmt.Sprintf("%d", result.StarCount)
		if len(s) > maxcount {
			maxcount = len(s)
		}
		if len(result.Name) > maxname {
			maxname = len(result.Name)
		}
	}
	if maxcount > 0 {
		maxcount++
	}
	if maxname > 0 {
		maxname++
	}
	for _, result := range results.Results {
		official := ' '
		if result.IsOfficial {
			official = '✩'
		}
		automated := ' '
		if result.IsAutomated {
			automated = '⚙'
		}
		trusted := ' '
		if result.IsTrusted {
			trusted = '✅'
		}
		stars := fmt.Sprintf("%d", result.StarCount)
		stars = stars + strings.Repeat(" ", maxcount - len(stars))
		name := strings.Replace(result.Name, "\n", " ", -1)
		name = strings.Trim(name, "\r\n\t ")
		name = name + strings.Repeat(" ", maxname - len(name))
		description := strings.Replace(result.Description, "\n", " ", -1)
		description = strings.Trim(description, "\r\n\t ")
		fmt.Printf("%c %c %c %s%s %s\n", official, trusted, automated, stars, name, description)
	}
}

func init() {
	commands = append(commands, cli.Command{
		Name:   "search",
		Usage:  "search for images",
		Action: searchImages,
	})
}
