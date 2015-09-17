// +build freebsd

package main

import (
	_ "github.com/docker/docker/daemon/graphdriver/zfs"
)

var (
	defaultHome = "/var/lib/docker"
)
