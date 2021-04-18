package progress_logger

type MessageType = int

const (
	Error   MessageType = 0
	Warn    MessageType = 1
	Info    MessageType = 2
	Debug   MessageType = 3
	Verbose MessageType = 4
)
