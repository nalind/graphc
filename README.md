## graphc

### A utility for managing docker layers and layer storage

####*Building:*
`go get github.com/codegangsta/cli`
`go get github.com/docker/docker/daemon/graphdriver`
`go get github.com/mistifyio/go-zfs`
`go get github.com/opencontainers/runc/libcontainer/label`
`go build -tags daemon`

This will build the binary in the current directory.

####*Invocation:*
`graphc --help`

The program has subcommands which are analogues of the public `Driver` API in
Docker's `graphdriver`.

######Disclaimer:
> *This program is strictly a work-in-progress and can (and probably will) do
> everything up to and including eat your laundry.*
