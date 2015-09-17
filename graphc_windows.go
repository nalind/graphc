// +build windows

package main

import (
	"os"

	_ "github.com/docker/docker/daemon/graphdriver/windows"
)

var (
	defaultHome = os.Getenv("programdata") + string(os.PathSeparator) + "docker"
)
