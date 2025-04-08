package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var L zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

	L = zerolog.New(output).
		Level(zerolog.TraceLevel).
		With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 1).
		Logger()
}

func Trace(msg string, v ...interface{}) {
	L.Trace().Msgf(msg, v...)
}

func Debug(msg string, v ...interface{}) {
	L.Debug().Msgf(msg, v...)
}

func Info(msg string, v ...interface{}) {
	L.Info().Msgf(msg, v...)
}

func Warn(msg string, v ...interface{}) {
	L.Warn().Msgf(msg, v...)
}

func Error(msg string, v ...interface{}) {
	L.Error().Msgf(msg, v...)
}

func Fatal(msg string, v ...interface{}) {
	L.Fatal().Msgf(msg, v...)
}

func Panic(msg string, v ...interface{}) {
	L.Panic().Msgf(msg, v...)
}
