package gochill

//Alert used to create log message with alert level
func Alert(msg string) []byte {
	return buildMessage(LevelAlert, msg)
}

//Info used to create log message with info level
func Info(msg string) []byte {
	return buildMessage(LevelInfo, msg)
}

//Critical used to create log message with critical level
func Critical(msg string) []byte {
	return buildMessage(LevelCritical, msg)
}

//Error used to create log message with error level
func Error(msg string) []byte {
	return buildMessage(LevelError, msg)
}

//Warn used to create log message with warning level
func Warn(msg string) []byte {
	return buildMessage(LevelWarning, msg)
}

//Notice used to create log message with warning level
func Notice(msg string) []byte {
	return buildMessage(LevelNotice, msg)
}

//Debug used to create log message with warning level
func Debug(msg string) []byte {
	return buildMessage(LevelDebug, msg)
}

//buildMessage is a private function used to build message structure
func buildMessage(level int, msg string) []byte {
	m := NewMessage(level)
	m.ShortMessage = msg

	return m.ToJSON()
}
