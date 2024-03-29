package core

type LogLevel int8

const (
	LogLevelDebug LogLevel = iota - 1
	LogLevelInfo
)

type LogFormat string

const (
	LogFormatJSON    LogFormat = "json"
	LogFormatConsole           = "console"
)

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Panic(msg string, args ...interface{})

	Flush()
}
