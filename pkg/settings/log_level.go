package settings

type LogLevelType = string

const (
	Error   LogLevelType = "error"
	Warn    LogLevelType = "warn"
	Info    LogLevelType = "info"
	Debug   LogLevelType = "debug"
	Verbose LogLevelType = "verbose"
)
