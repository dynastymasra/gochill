#!/bin/bash

go test -v -cover -covermode count -coverprofile=cover.out | $GOPATH/bin/go-junit-report > results.xml 
$GOPATH/bin/gocov convert cover.out | $GOPATH/bin/gocov-xml > coverage.xml
# $GOPATH/bin/go2xunit -fail -input cover.out -output results.xml
