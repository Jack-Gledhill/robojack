package log

func Error(msg string, v ...interface{}) {
	l.Error().Msgf(msg, v...)
}
