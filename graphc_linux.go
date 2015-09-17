// +build linux

package main

import (
	_ "github.com/docker/docker/daemon/graphdriver/aufs"
	_ "github.com/docker/docker/daemon/graphdriver/btrfs"
	_ "github.com/docker/docker/daemon/graphdriver/devmapper"
	_ "github.com/docker/docker/daemon/graphdriver/overlay"
	_ "github.com/docker/docker/daemon/graphdriver/vfs"
	_ "github.com/docker/docker/daemon/graphdriver/zfs"
)

var (
	defaultHome = "/var/lib/docker"
)
