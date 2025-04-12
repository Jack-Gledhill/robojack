package log

func Debug(msg string, v ...interface{}) {
	l.Debug().Msgf(msg, v...)
}
