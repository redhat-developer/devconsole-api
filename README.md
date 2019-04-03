## Overview

The canonical location of the OpenShift DevConsole API definition. This repo holds the API type definitions and serialization code used by [DevConsole Operator](https://github.com/redhat-developer/devconsole-operator)

## Prerequisites

- [dep][dep_tool] version v0.5.0+.
- [go][go_tool] version v1.10+.

## Build API

```sh
$ dep ensure
$ go build github.com/redhat-developer/devconsole-api/pkg/apis/
```

Regenerate code after modifying API:

```sh
$ ./generate-deepcopy.sh
```

## Add a new API for the custom resource

New API for the custom resource can be generated in [DevConsole Operator](https://github.com/redhat-developer/devconsole-operator) and then moved to this repo:

```sh
$ cd $GOPATH/src/github.com/redhat-developer/devconsole-operator

# Add a new API for the custom resource <NewResourceKind>
$ operator-sdk add api --api-version=devconsole.openshift.io/v1alpha1 --kind=<NewResourceKind>

# Move generated API to devconsole-api
$ rsync -avh --progress ./pkg/apis/devconsole/ $GOPATH/src/github.com/redhat-developer/devconsole-api/pkg/apis/devconsole/
```
[dep_tool]:https://golang.github.io/dep/docs/installation.html
[go_tool]:https://golang.org/dl/
