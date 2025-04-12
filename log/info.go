package log

func Info(msg string, v ...interface{}) {
	l.Info().Msgf(msg, v...)
}
