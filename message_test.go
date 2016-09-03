package gochill

import (
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
