package progress_logger

import (
	"fmt"
	"io"
)

type ProgressLogger struct {
	writer io.Writer
	reader io.Reader
}

func (l *ProgressLogger) GetStepLogger(jobName string, total int) *StepLogger {
	return NewStepLogger(l.writer, l.reader, jobName, total)
}

func (l *ProgressLogger) Info(message string) error {
	msg := l.formatLog(Info, message)
	_, err := l.writer.Write([]byte(msg))
	return err
}

func (l *ProgressLogger) Warn(message string) error {
	msg := l.formatLog(Warn, message)
	_, err := l.writer.Write([]byte(msg))
	return err
}

func (l *ProgressLogger) Error(message string) error {
	msg := l.formatLog(Error, message)
	_, err := l.writer.Write([]byte(msg))
	return err
}

const (
	Info  = iota
	Warn  = iota
	Error = iota
)

func (l *ProgressLogger) formatLog(messageType int, message string) string {
	switch messageType {
	case Info:
		return fmt.Sprintf("Info: %s", message)
	case Warn:
		return fmt.Sprintf("Warn: %s", message)
	case Error:
		return fmt.Sprintf("Error: %s", message)
	default:
		panic("Unknown message type")
	}
}
