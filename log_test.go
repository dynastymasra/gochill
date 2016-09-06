package gochill

import (
	"encoding/json"
	"testing"

	simplejson "github.com/bitly/go-simplejson"
)

var message []byte

type CustomLogger struct{}

func (cl CustomLogger) Write(p []byte) (int, error) {
	message = p
	return 0, nil
}

func TestAlert(t *testing.T) {

	CustomOutput = CustomLogger{}
	Alert(Msg("ini log"))

	logM := Message{}
	err := json.Unmarshal(message, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.ShortMessage != "ini log" {
		t.Fail()
		t.Log("Expected result : ", "ini log")
		t.Log("Given result : ", logM.ShortMessage)
	}
}

func TestAlertWithOptional(t *testing.T) {

	CustomOutput = CustomLogger{}
	Alert(Msg("ini log"), O("key1", "value1"), O("key2", 10))

	j, _ := simplejson.NewJson(message)
	maps := j.MustMap()

	if maps["_key1"] == nil {
		t.Fail()
		t.Log("Expected result is _key1 exists")
		t.Log("Given result : ", maps["_key1"])
	}

	if maps["_key2"] == nil {
		t.Fail()
		t.Log("Expected result is _key2 exists")
		t.Log("Given result : ", maps["_key2"])
	}

	if maps["version"] == nil {
		t.Fail()
		t.Log("Expected result is version exists")
		t.Log("Given result : ", maps["version"])
	}

	if maps["short_message"] == nil {
		t.Fail()
		t.Log("Expected result is short_message exists")
		t.Log("Given result : ", maps["short_message"])
	}

}

//TestAlertWithoutCustomWriter used to make sure that
//logging still working even without custom io writer implemented
func TestAlertWithoutCustomWriter(t *testing.T) {
	Alert(Msg("ini log"))
	t.Log("logged")
}

func TestInfo(t *testing.T) {

	CustomOutput = CustomLogger{}
	Info(Msg("ini log"))

	logM := Message{}
	err := json.Unmarshal(message, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelInfo {
		t.Fail()
		t.Log("Expected Level : ", LevelInfo)
		t.Log("Given Level : ", logM.Level)
	}
}

func TestCritical(t *testing.T) {
	CustomOutput = CustomLogger{}
	Critical(Msg("ini log"))

	logM := Message{}
	err := json.Unmarshal(message, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelCritical {
		t.Fail()
		t.Log("Expected Level : ", LevelCritical)
		t.Log("Given Level : ", logM.Level)
	}
}

func TestError(t *testing.T) {
	CustomOutput = CustomLogger{}
	Error(Msg("ini log"))

	logM := Message{}
	err := json.Unmarshal(message, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelError {
		t.Fail()
		t.Log("Expected Level : ", LevelError)
		t.Log("Given Level : ", logM.Level)
	}
}

func TestWarn(t *testing.T) {
	CustomOutput = CustomLogger{}
	Warn(Msg("ini log"))

	logM := Message{}
	err := json.Unmarshal(message, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelWarning {
		t.Fail()
		t.Log("Expected Level : ", LevelWarning)
		t.Log("Given Level : ", logM.Level)
	}
}

func TestNotice(t *testing.T) {
	CustomOutput = CustomLogger{}
	Notice(Msg("ini log"))

	logM := Message{}
	err := json.Unmarshal(message, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelNotice {
		t.Fail()
		t.Log("Expected Level : ", LevelNotice)
		t.Log("Given Level : ", logM.Level)
	}
}

func TestDebug(t *testing.T) {
	CustomOutput = CustomLogger{}
	Debug(Msg("ini log"))

	logM := Message{}
	err := json.Unmarshal(message, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelDebug {
		t.Fail()
		t.Log("Expected Level : ", LevelDebug)
		t.Log("Given Level : ", logM.Level)
	}
}
