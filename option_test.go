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

	if o.Key == "" {
		t.Fail()
		t.Log("Expected result is not nil")
		t.Log("Given result : ", o.Key)
	}
}

func TestMergeOptions(t *testing.T) {
	o1 := NewOption()
	o1.Set("test1", "testing")

	o2 := NewOption()
	o2.Set("test2", "testing2")

	maps := MergeOptionsToMap([]Option{*o1, *o2})
	if len(maps) < 1 {
		t.Fail()
		t.Log("Expected results is maps at least 1")
		t.Log("Given result : ", len(maps))
	}

	if maps["_test1"] == nil {
		t.Fail()
		t.Log("Expected result is not nil")
		t.Log("Given result : ", maps["_test1"])
	}

	if maps["_test2"] == nil {
		t.Fail()
		t.Log("Expected result is not nil")
		t.Log("Given result : ", maps["_test2"])
	}
}

func TestO(t *testing.T) {
	o := O("key", "value")

	if o.Key == "" {
		t.Fail()
		t.Log("Expected result is not an empty string")
		t.Log("Given result : ", o.Key)
	}
}
