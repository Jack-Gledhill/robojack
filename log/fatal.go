package log

func Fatal(msg string, v ...interface{}) {
	l.Fatal().Msgf(msg, v...)
}
