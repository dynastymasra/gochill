#!/bin/bash

go test -v -cover -covermode count -coverprofile=/go/src/gochill/test-results/cover.out | $GOPATH/bin/go-junit-report > test-results/results.xml
$GOPATH/bin/gocov convert /go/src/gochill/test-results/cover.out | $GOPATH/bin/gocov-xml > test-results/coverage.xml
echo "Inside container : $(pwd)"
echo "Test results : ..."
ls -l test-results
