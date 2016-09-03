package gochill

import (
	"io"
	"log"
	"os"
)

//Logger used to create new log instance
var (
	Logger       *log.Logger
	CustomOutput io.Writer
)

//NewCustomeLogger used to set custom log output destination
func NewCustomeLogger(std, custom io.Writer) *log.Logger {
	logger := log.New(io.MultiWriter(std, custom), "", 0)
	return logger
}

//Alert used to create log message with alert level
func Alert(msg string) {
	message := buildMessage(LevelAlert, msg)
	Logger = NewCustomeLogger(os.Stderr, CustomOutput)
	Logger.Println(string(message))
}

//Info used to create log message with info level
func Info(msg string) {
	message := buildMessage(LevelInfo, msg)
	Logger = NewCustomeLogger(os.Stdout, CustomOutput)
	Logger.Println(string(message))
}

//Critical used to create log message with critical level
func Critical(msg string) {
	message := buildMessage(LevelCritical, msg)
	Logger = NewCustomeLogger(os.Stderr, CustomOutput)
	Logger.Println(string(message))
}

//Error used to create log message with error level
func Error(msg string) {
	message := buildMessage(LevelError, msg)
	Logger = NewCustomeLogger(os.Stderr, CustomOutput)
	Logger.Println(string(message))
}

//Warn used to create log message with warning level
func Warn(msg string) {
	message := buildMessage(LevelWarning, msg)
	Logger = NewCustomeLogger(os.Stdout, CustomOutput)
	Logger.Println(string(message))
}

//Notice used to create log message with warning level
func Notice(msg string) {
	message := buildMessage(LevelNotice, msg)
	Logger = NewCustomeLogger(os.Stdout, CustomOutput)
	Logger.Println(string(message))
}

//Debug used to create log message with warning level
func Debug(msg string) {
	message := buildMessage(LevelDebug, msg)
	Logger = NewCustomeLogger(os.Stdout, CustomOutput)
	Logger.Println(string(message))
}

//buildMessage is a private function used to build message structure
func buildMessage(level int, msg string) []byte {
	m := NewMessage(level)
	m.ShortMessage = msg

	return m.ToJSON()
}
