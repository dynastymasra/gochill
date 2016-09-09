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

//NilWriter used as custom output writer when CustomOutput nil
type NilWriter struct{}

func (nw NilWriter) Write(p []byte) (int, error) { return 0, nil }

func init() {
	Logger = log.New(os.Stdout, "", 0)
	if CustomOutput == nil {
		CustomOutput = NilWriter{}
	}
}

//NewCustomOutput used to change how Logger send log message
func NewCustomOutput(output io.Writer) {
	Logger.SetOutput(output)
}

//Alert used to create log message with alert level
func Alert(msg MessageLog, options ...Option) {
	message := buildMessage(LevelAlert, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))
	Logger.Println(string(message))
}

//Info used to create log message with info level
func Info(msg MessageLog, options ...Option) {
	message := buildMessage(LevelInfo, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//Critical used to create log message with critical level
func Critical(msg MessageLog, options ...Option) {
	message := buildMessage(LevelCritical, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))
	Logger.Println(string(message))
}

//Error used to create log message with error level
func Error(msg MessageLog, options ...Option) {
	message := buildMessage(LevelError, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))
	Logger.Println(string(message))
}

//Warn used to create log message with warning level
func Warn(msg MessageLog, options ...Option) {
	message := buildMessage(LevelWarning, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//Notice used to create log message with warning level
func Notice(msg MessageLog, options ...Option) {
	message := buildMessage(LevelNotice, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//Debug used to create log message with warning level
func Debug(msg MessageLog, options ...Option) {
	message := buildMessage(LevelDebug, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//buildMessage is a private function used to build message structure
func buildMessage(level int, msg MessageLog, options []Option) []byte {
	m := NewMessage(level)
	m.ShortMessage = msg.Short
	m.FullMessage = msg.Full

	mergedOptions := MergeOptionsToMap(options)
	maps := CombineMaps(m.ToMap(), mergedOptions)

	return MapToJSON(maps)
}
