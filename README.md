## graphc

### A utility for managing docker layers and layer storage

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
