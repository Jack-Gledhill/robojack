package log

func Warn(msg string, v ...interface{}) {
	l.Warn().Msgf(msg, v...)
}
