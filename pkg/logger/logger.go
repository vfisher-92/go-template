package logger

type Logger interface {
	Info(message string)
	Warn()
	Error(err error)
	Errorf(format string, a ...interface{})
}
