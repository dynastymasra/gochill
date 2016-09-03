package gochill

//Option struct used to set optional field to the log
type Option struct {
	Key    string
	Value  interface{}
	output map[string]interface{}
}

//Set used to set new key value
func (o *Option) Set(key string, value interface{}) {
	o.Key = "_" + key
	o.Value = value

	o._makeMap()
}

//Out used to get output data that should contains
//mapping Key and Value
func (o *Option) Out() map[string]interface{} {
	return o.output
}

//_makeMap is private function used to make mapping key value
func (o *Option) _makeMap() {
	o.output = make(map[string]interface{})
	o.output[o.Key] = o.Value
}

//NewOption used to create new Option instance
func NewOption() *Option {
	o := Option{}
	return &o
}
