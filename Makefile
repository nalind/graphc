GOPATH = ${PWD}/vendor:${PWD}/vendor/src/github.com/docker/docker/vendor:${PWD}:$(shell go env GOPATH)
TAGS = selinux
all:
	GOPATH="$(GOPATH)" go build -tags "daemon $(TAGS)"
