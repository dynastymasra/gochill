# Gochill
Golang library for building logging data structure.

# Abstract

Logging data structure that supported is following [GELF](http://docs.graylog.org/en/2.0/pages/gelf.html#).
All data structure should be in json mode.  Required fields :

- version
- host
- service
- short_message
- timestamp
- level

We need a library that support this logging data structure but also should be easy
to use for other programmers.

# Rationale

This library should be only for building logging data structure, it means for now,
this library not do anything that print the message to `stdout`.

# Environment Variable

This library will listen to environment variable with registered key `SERVICE_NAME`.
This variable used to fill `service` json field.

If environment variable not exists, then `service` will be an empty string.

# Hostname

This library will automatically load `hostname` from os using golang's native library

```go
os.Hostname()
```

But if there is some error, then `hostname` will be an empty string.

# Usage

```go
import "gochill"

gochill.Alert("alert message")
```

Available levels :

- Alert
- Critical
- Error
- Warning
- Notice
- Info
- Debug

Gochill designed with `io.MultiWriter` it means you can add your own custom
`io.Writer` to proceed your log.

```go
package main

import "gochill"

type CustomWriter struct{}
func (cw CustomWriter) Write(message []byte) (int, error) {
  fmt.Println("hello world")
}

gochill.CustomOutput = CustomWriter{}
gochill.Info("log info")
```

These code should produce an output will send a structured json code and also
  print "hello world" to stdout.

# Unit Test

How to run :

```
go test -cover
```

Via `Makefile` :

```
make test
```
