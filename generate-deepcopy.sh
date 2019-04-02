#!/usr/bin/env bash
go run $(pwd)/vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go --input-dirs ./pkg/apis/devconsole/v1alpha1/ -O zz_generated.deepcopy --bounding-dirs github.com/redhat-developer/devconsole-api/pkg/apis "devconsole:v1alpha1"

