package log

func Trace(msg string, v ...interface{}) {
	l.Trace().Msgf(msg, v...)
}
