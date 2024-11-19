package log

import "os"

const (
	LevelDebug   = iota
	LevelInfo    = iota
	LevelWarning = iota
	LevelError   = iota
	LevelFatal   = iota
	// PrettyPrintFormat is a text/template string that will format the output when in development
	PrettyPrintFormat = "[{{.Timestamp}}] [{{.LevelName}}] {{.Message}}"
)

// LevelStrings maps each log level to a name that will be used when pretty printing
var LevelStrings = map[int]string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO",
	LevelWarning: "WARN",
	LevelError:   "ERROR",
	LevelFatal:   "FATAL",
}

// Debug logs a message at the debug level
func Debug(m string, v ...interface{}) {
	l := NewLog(LevelDebug, m, v...)
	go l.Print()
}

// Info logs a message at the info level
func Info(m string, v ...interface{}) {
	l := NewLog(LevelInfo, m, v...)
	go l.Print()
}

// Warning logs a message at the warning level
func Warning(m string, v ...interface{}) {
	l := NewLog(LevelWarning, m, v...)
	go l.Print()
}

// Error logs a message at the error level
func Error(m string, v ...interface{}) {
	l := NewLog(LevelError, m, v...)
	go l.Print()
}

// Fatal logs a message at the fatal level and causes the program to exit immediately
func Fatal(m string, v ...interface{}) {
	l := NewLog(LevelFatal, m, v...)
	l.Print()
	os.Exit(1)
}
