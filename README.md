## graphc

### A utility for managing docker layers and layer storage

####*Building:*
 `go get -d -n -u github.com/docker/docker/daemon/graphdriver github.com/docker/docker/graph`
 `( cd ``echo $GOPATH | cut -f1 -d:``/src/github.com/docker/docker ; env GOPATH=``pwd``/vendor:"$GOPATH" sh hack/make.sh dynbinary )`
 `go get -u github.com/codegangsta/cli github.com/mistifyio/go-zfs github.com/opencontainers/runc/libcontainer/label`
 `go get -u github.com/tchap/go-patricia/patricia github.com/vbatts/tar-split/tar/asm github.com/vbatts/tar-split/tar/storage`
 `go build -tags daemon`

This will build the binary in the current directory.

####*Invocation:*
`graphc --help`

The program has subcommands which are analogues of the public `Driver` API in
Docker's `graphdriver`.

######Disclaimer:
> *This program is strictly a work-in-progress.  It directly accesses data
> which is normally only accessed by a docker daemon.  It can (and probably
> will) do everything up to and including eat your laundry.*
