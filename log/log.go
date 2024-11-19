package log

import (
	"encoding/json"
	"fmt"

	"github.com/Jack-Gledhill/robojack/env"
	"github.com/Jack-Gledhill/robojack/utils"
)

// Log represents a single log entry, with a level, message, and timestamp
type Log struct {
	Level     int    `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// Print is a helper than changes behaviour based on the environment. In production, logs are output as JSON
// In development, logs are output in a human-readable format
func (l *Log) Print() {
	if env.Development() {
		l.PrettyPrint()
	} else {
		l.JSONPrint()
	}
}

// JSONPrint outputs the log as a stringified JSON object, this is useful for logging to aggregators that can parse JSON
func (l *Log) JSONPrint() {
	b, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}

// PrettyPrint shows the log in a human-readable format that is helpful in development
func (l *Log) PrettyPrint() {
	o, err := utils.TemplateString(PrettyPrintFormat, l)
	if err != nil {
		panic(err)
	}

	fmt.Println(o)
}

// LevelName returns the pretty print name for the log level. All possible return values will have the same length to make logs easier to read
func (l *Log) LevelName() string {
	n := LevelStrings[LevelInfo]
	s, ok := LevelStrings[l.Level]
	if ok {
		n = s
	}

	return utils.FixedWidthAppend(n, len(LevelStrings[LevelError]))
}

// NewLog creates a new Log object with the current time, and a message fully expanded by fmt.Sprintf
func NewLog(level int, msg string, vals ...interface{}) Log {
	return Log{
		Level:     level,
		Message:   fmt.Sprintf(msg, vals...),
		Timestamp: GetTimestamp(),
	}
}
