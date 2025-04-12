package log

func Panic(msg string, v ...interface{}) {
	l.Panic().Msgf(msg, v...)
}
