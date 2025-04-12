package log

import (
	"os"
	"time"

	"github.com/Jack-Gledhill/robojack/config"

	"github.com/rs/zerolog"
)

var Level zerolog.Level

var l zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

	Level = zerolog.InfoLevel
	if config.IsDevelopment() {
		Level = zerolog.DebugLevel
	}

	l = zerolog.New(output).
		Level(Level).
		With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 1).
		Logger()
}
