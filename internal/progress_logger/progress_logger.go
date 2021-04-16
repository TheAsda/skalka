package progress_logger

import (
	"fmt"
	"github.com/TheAsda/skalka/pkg/settings"
	"io"
)

type ProgressLogger struct {
	writer   io.Writer
	reader   io.Reader
	logLevel MessageType
}

func NewProgressLogger(writer io.Writer, reader io.Reader, sett settings.Settings) *ProgressLogger {
	logLevel := Warn
	switch sett.LogLevel {
	case settings.Error:
		logLevel = Error
		break
	case settings.Warn:
		logLevel = Warn
		break
	case settings.Info:
		logLevel = Info
		break
	case settings.Debug:
		logLevel = Debug
		break
	case settings.Verbose:
		logLevel = Verbose
		break
	}
	return &ProgressLogger{writer: writer, reader: reader, logLevel: logLevel}
}

func (l *ProgressLogger) LogJob(jobName string) error {
	msg := fmt.Sprintf("Start job: %s", jobName)
	err := l.write(msg)
	return err
}

func (l *ProgressLogger) LogStep(current int, total int, name string) error {
	msg := fmt.Sprintf("[%d/%d] %s", current, total, name)
	err := l.write(msg)
	return err
}

func (l *ProgressLogger) Error(message string) error {
	msg := l.formatLog(Error, message)
	err := l.write(msg)
	return err
}

func (l *ProgressLogger) Warn(message string) error {
	if l.logLevel < Warn {
		return nil
	}
	msg := l.formatLog(Warn, message)
	err := l.write(msg)
	return err
}

func (l *ProgressLogger) Info(message string) error {
	if l.logLevel < Info {
		return nil
	}
	msg := l.formatLog(Info, message)
	err := l.write(msg)
	return err
}

func (l *ProgressLogger) Debug(message string) error {
	if l.logLevel < Debug {
		return nil
	}
	msg := l.formatLog(Debug, message)
	err := l.write(msg)
	return err
}

func (l *ProgressLogger) Verbose(message string) error {
	if l.logLevel < Verbose {
		return nil
	}
	msg := l.formatLog(Verbose, message)
	err := l.write(msg)
	return err
}

func (l *ProgressLogger) GetStdout() io.Writer {
	return l.writer
}

func (l *ProgressLogger) GetStderr() io.Writer {
	return l.writer
}

func (l *ProgressLogger) write(message string) error {
	_, err := l.writer.Write([]byte(message))
	return err
}

func (l *ProgressLogger) formatLog(messageType MessageType, message string) string {
	var format string
	switch messageType {
	case Info:
		format = "Info: %s"
	case Warn:
		format = "Warn: %s"
	case Error:
		format = "Error: %s"
	case Debug:
		format = "Debug: %s"
	case Verbose:
		format = "Verbose: %s"
	default:
		panic("Unknown message type")
	}
	return fmt.Sprintf(format, message)
}
