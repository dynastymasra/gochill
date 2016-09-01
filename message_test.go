package gochill

import (
	"encoding/json"
	"errors"
	"os"
	"testing"
)

func TestNewMessage(t *testing.T) {
	os.Setenv(EnvServiceKeyName, "testing_gochill_service")

	m := NewMessage(LevelAlert)
	if m.Level != LevelAlert {
		t.Fail()
		t.Log("Unexpected Level : ", m.Level)
	}

	if m.Host == "" {
		t.Fail()
		t.Log("Host should not be empty.")
	}

	if m.Service != os.Getenv(EnvServiceKeyName) {
		t.Fail()
		t.Log("Unexpected Service Name : ", m.Service)
	}
}

func TestMessageToJSON(t *testing.T) {
	m := NewMessage(LevelAlert)
	m.ShortMessage = "halo"

	toJSON := m.ToJSON()
	toString := string(toJSON)

	if toString == "" {
		t.Fail()
		t.Log("Unexpected string json : ", toString)
	}

	newm := Message{}
	err := json.Unmarshal([]byte(toString), &newm)
	if err != nil {
		t.Error(err)
	}

	if m.ShortMessage != newm.ShortMessage {
		t.Fail()
		t.Log("Requested short message : ", m.ShortMessage)
		t.Log("Given short message : ", newm.ShortMessage)
	}

	if m.Level != newm.Level {
		t.Fail()
		t.Log("Requested level : ", m.Level)
		t.Log("Given Level : ", newm.Level)
	}
}

func TestMessageJSONError(t *testing.T) {

	fakeJSON := func(v interface{}) ([]byte, error) {
		return nil, errors.New("error")
	}

	m := NewMessage(LevelAlert)
	m.ReplaceJSONEngine(fakeJSON)
	jsonByte := m.ToJSON()

	if jsonByte != nil {
		t.Fail()
		t.Log("Expected result is nil")
		t.Log("Given result is : ", jsonByte)
		t.Log("To string : ", string(jsonByte))
	}
}

func TestMessageHostnameError(t *testing.T) {
	fakeHostname := func() (string, error) {
		return "", errors.New("error")
	}

	m := NewMessage(LevelAlert)
	m.ReplaceOSEngine(fakeHostname)

	if m.Host != "" {
		t.Fail()
		t.Log("Expected hostname : nil")
		t.Log("Given hostname value : ", m.Host)
	}
}
