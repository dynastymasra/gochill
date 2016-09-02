package gochill

import (
	"encoding/json"
	"testing"
)

func TestAlert(t *testing.T) {
	log := Alert("ini log")

	logM := Message{}
	err := json.Unmarshal(log, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelAlert {
		t.Fail()
		t.Log("Expected Level : ", LevelAlert)
		t.Log("Given Level : ", logM.Level)
	}
}

func TestInfo(t *testing.T) {
	log := Info("ini log")

	logM := Message{}
	err := json.Unmarshal(log, &logM)

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
	log := Critical("ini log")

	logM := Message{}
	err := json.Unmarshal(log, &logM)

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
	log := Error("ini log")

	logM := Message{}
	err := json.Unmarshal(log, &logM)

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
	log := Warn("ini log")

	logM := Message{}
	err := json.Unmarshal(log, &logM)

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
	log := Notice("ini log")

	logM := Message{}
	err := json.Unmarshal(log, &logM)

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
	log := Debug("ini log")

	logM := Message{}
	err := json.Unmarshal(log, &logM)

	if err != nil {
		t.Error(err)
	}

	if logM.Level != LevelDebug {
		t.Fail()
		t.Log("Expected Level : ", LevelDebug)
		t.Log("Given Level : ", logM.Level)
	}
}
