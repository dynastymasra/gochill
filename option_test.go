package gochill

import "testing"

func TestNewOption(t *testing.T) {
	o := NewOption()
	o.Set("key1", "value1")

	if o.Key != "_key1" {
		t.Fail()
		t.Log("Expected result : ", "key1")
		t.Log("Given result : ", o.Key)
	}

	if o.Value != "value1" {
		t.Fail()
		t.Log("Expected result : ", "value1")
		t.Log("Given result : ", o.Value)
	}

}

func TestOptionMaps(t *testing.T) {
	o := NewOption()
	o.Set("key1", "value1")

	maps := o.Out()

	if maps["_key1"] == nil {
		t.Fail()
		t.Log("Expected result is not nil")
		t.Log("Given result : ", maps["key1"])
	}
}
