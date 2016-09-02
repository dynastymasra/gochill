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

log.Println(gochill.Alert("alert message"))
log.Println(gochill.Info("alert message"))
```

Available levels :

- Alert
- Critical
- Error
- Warning
- Notice
- Info
- Debug

# Unit Test

How to run :

```
go test -cover
```

Via `Makefile` :

```
make test
```
