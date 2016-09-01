package gochill

//JSONMarshalEngine used as json marshal activity
type JSONMarshalEngine func(v interface{}) ([]byte, error)

//OSHostEngine used as wrapper os golang native lib
type OSHostEngine func() (string, error)
