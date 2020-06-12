SHELL := /bin/bash
version := $(shell git rev-list --count HEAD)
commit := $(shell git describe --always --long --dirty)
built_at := $(shell date +%FT%T%z)
built_by := ${USER}

flags := -gcflags="all=-N -l -c 2"
ldflags := -X main.version=v${version} -X main.commit=${commit}
ldflags += -X main.builtAt=${built_at} -X main.builtBy=${built_by}

dist := ./dist/stress
env := GO111MODULE=on
DIR := ${CURDIR}

all:
	source $(shell go env GOPATH)/src/github.com/SebastianJ/elrond-sdk/scripts/bls_build_flags.sh && $(env) go build -o $(dist) -ldflags="$(ldflags)" cmd/main.go

static:
	make -C $(shell go env GOPATH)/src/github.com/herumi/mcl
	make -C $(shell go env GOPATH)/src/github.com/herumi/bls BLS_SWAP_G=1
	source $(shell go env GOPATH)/src/github.com/SebastianJ/elrond-sdk/scripts/bls_build_flags.sh && $(env) go build -o $(dist) -ldflags="$(ldflags) -w -extldflags \"-static\"" cmd/main.go

debug:
	source $(shell go env GOPATH)/src/github.com/SebastianJ/elrond-sdk/scripts/bls_build_flags.sh && $(env) go build $(flags) -o $(dist) -ldflags="$(ldflags)" cmd/main.go

.PHONY:clean upload-linux

clean:
	@rm -f $(dist)
	@rm -rf ./dist
