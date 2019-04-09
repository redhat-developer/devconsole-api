## Overview

The canonical location of the OpenShift DevConsole API definition. This repo holds the API type definitions and serialization code used by [DevConsole Operator](https://github.com/redhat-developer/devconsole-operator)

## Prerequisites

- [dep][dep_tool] version v0.5.0+.
- [git][git_tool]
- [go][go_tool] version v1.10+.
- [docker][docker_tool] version 17.03+.

## Build API

```sh
$ make build
```

Regenerate deepcopy after modifying API:

```sh
$ make generate
```

## Add a new API for the custom resource

New API for the custom resource can be generated in [DevConsole Operator](https://github.com/redhat-developer/devconsole-operator) and then moved to this repo:

```sh
$ cd $GOPATH/src/github.com/redhat-developer/devconsole-operator

# Add a new API for the custom resource <NewResourceKind>
$ operator-sdk add api --api-version=devconsole.openshift.io/v1alpha1 --kind=<NewResourceKind>

# Move generated API to devconsole-api
$ rsync -avh --progress ./pkg/apis/devconsole/ $GOPATH/src/github.com/redhat-developer/devconsole-api/pkg/apis/devconsole/

# Remove generated API from devconsole-operator
$ rm -rf ./pkg/apis

$ cd $GOPATH/src/github.com/redhat-developer/devconsole-api

# Edit your API
# ...

# Re-generate deepcopy
$ make generate
```
[dep_tool]:https://golang.github.io/dep/docs/installation.html
[go_tool]:https://golang.org/dl/
[git_tool]:https://git-scm.com/downloads
[docker_tool]:https://docs.docker.com/install/
