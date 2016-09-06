package gochill

import (
	"errors"
	"os"
	"testing"
)

func TestMsgEmptyFullMsg(t *testing.T) {
	short, full := Msg("test")

	if short != "test" {
		t.Fail()
		t.Log("Expected result short variable is equal with test")
		t.Log("Given result : ", short)
	}

	if full != "" {
		t.Fail()
		t.Log("Expected result full variable is empty string")
		t.Log("Given result : ", full)
	}
}

func TestMsgFullMsgNotEmpty(t *testing.T) {
	short, full := Msg("test short", "test full")

	if short != "test short" {
		t.Fail()
		t.Log("Expected result short variable is equal with test sort")
		t.Log("Given result : ", short)
	}

	if full != "test full" {
		t.Fail()
		t.Log("Expected result full variable is equal with test full")
		t.Log("Given result : ", full)
	}

}

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
