package gochill

//OSHostEngine used as wrapper os golang native lib
type OSHostEngine func() (string, error)
