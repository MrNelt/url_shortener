package logger

type ILogger interface {
	Info(msg string)
	Debug(msg string)
	Trace(msg string)
	Error(msg string)
	Fatal(msg string)
}
