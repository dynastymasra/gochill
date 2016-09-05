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

func init() {
	Logger = log.New(os.Stdout, "", 0)
}

//NewCustomOutput used to change how Logger send log message
func NewCustomOutput(output io.Writer) {
	Logger.SetOutput(output)
}

//Alert used to create log message with alert level
func Alert(msg string, options ...Option) {
	message := _buildMessage(LevelAlert, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))
	Logger.Println(string(message))
}

//Info used to create log message with info level
func Info(msg string, options ...Option) {
	message := _buildMessage(LevelInfo, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//Critical used to create log message with critical level
func Critical(msg string, options ...Option) {
	message := _buildMessage(LevelCritical, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))
	Logger.Println(string(message))
}

//Error used to create log message with error level
func Error(msg string, options ...Option) {
	message := _buildMessage(LevelError, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stderr, CustomOutput))
	Logger.Println(string(message))
}

//Warn used to create log message with warning level
func Warn(msg string, options ...Option) {
	message := _buildMessage(LevelWarning, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//Notice used to create log message with warning level
func Notice(msg string, options ...Option) {
	message := _buildMessage(LevelNotice, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//Debug used to create log message with warning level
func Debug(msg string, options ...Option) {
	message := _buildMessage(LevelDebug, msg, options)
	NewCustomOutput(io.MultiWriter(os.Stdout, CustomOutput))
	Logger.Println(string(message))
}

//_buildMessage is a private function used to build message structure
func _buildMessage(level int, msg string, options []Option) []byte {
	m := NewMessage(level)
	m.ShortMessage = msg

	mergedOptions := MergeOptionsToMap(options)
	maps := CombineMaps(m.ToMap(), mergedOptions)

	return MapToJSON(maps)
}
