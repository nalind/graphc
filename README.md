## graphc

### A utility for managing docker layers and layer storage

####*Description:*
The `graphc` utility aims to provide its user with the ability to import
images, either directly or from a registry.  It aims to also provide the
ability to export, inspect, and remove images from the local image store.  (It
can do most of this now.)

It should also provide options for creating containers, starting containers
with the help of `runc`, and deleting containers.  (Currently, it can't do any
of this.)

When the layer storage backend exports a mountable filesystem, `graphc` should
be able to mount a container's composed filesystem at a specified location, and
then unmount it later.

And it should do all of this in a way that doesn't corrupt the image store, and
which doesn't create surprises for a running docker daemon.  (Nope.)

####*Building:*
The makefile should be set up so that running `make` will be sufficient to
produce a graphc binary in the current directory.

####*Invocation:*
`graphc --help`

The program has subcommands which are analogues of the public `Driver` API in
Docker's `daemon/graphdriver` subdirectory and the `Graph` and `TagStore` APIs
in Docker's `graph` subdirectory.

######Disclaimer:
> *This program is strictly a work-in-progress.  It directly accesses data
> which is normally only accessed by a docker daemon.  It can (and probably
> will) do everything up to and including eat your laundry.*
