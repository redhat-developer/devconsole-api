# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash

include ./make/verbose.mk
.DEFAULT_GOAL := help
include ./make/help.mk
include ./make/out.mk
include ./make/find-tools.mk
include ./make/go.mk
include ./make/git.mk
include ./make/format.mk
include ./make/lint.mk
include ./make/docker.mk

.PHONY: build
## Build
build: ./vendor $(shell find . -path ./vendor -prune -o -name '*.go' -print)
	$(Q)CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
	    go build github.com/redhat-developer/devconsole-api/pkg/apis/

.PHONY: generate
## Generate deepcopy after modifying API
generate:
	$(Q)go run $(shell pwd)/vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go --input-dirs ./pkg/apis/devconsole/v1alpha1/ -O zz_generated.deepcopy --bounding-dirs github.com/redhat-developer/devconsole-api/pkg/apis "devconsole:v1alpha1"

.PHONY: clean
## Clean
clean:
	$(Q)-rm -rf ${V_FLAG} ./vendor
	$(Q)go clean ${X_FLAG} ./...

./vendor: Gopkg.toml Gopkg.lock
	$(Q)dep ensure ${V_FLAG} -vendor-only
