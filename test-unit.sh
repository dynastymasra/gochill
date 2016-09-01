#!/bin/bash

go test -v -cover -coverprofile=cover.out
$GOPATH/bin/gocov convert cover.out | $GOPATH/bin/gocov-xml > coverage.xml
