package logger

type LogLevel uint

const (
	Info LogLevel = iota + 1
	Fail
	Warn
)
