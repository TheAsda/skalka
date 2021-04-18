package progress_logger

import (
	"fmt"
	"github.com/TheAsda/skalka/pkg/settings"
	"io"
)

type ProgressLogger struct {
	writer   io.Writer
	logLevel MessageType
}

func NewProgressLogger(writer io.Writer, sett settings.Settings) *ProgressLogger {
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
	return &ProgressLogger{writer: writer, logLevel: logLevel}
}

func (l *ProgressLogger) LogJob(jobName string) {
	msg := fmt.Sprintf("Start job: %s", jobName)
	l.write(msg)
}

func (l *ProgressLogger) LogStep(current int, total int, name string) {
	msg := fmt.Sprintf("[%d/%d] %s", current, total, name)
	l.write(msg)
}

func (l *ProgressLogger) Error(message string) {
	msg := l.formatLog(Error, message)
	l.write(msg)
}

func (l *ProgressLogger) Warn(message string) {
	if l.logLevel < Warn {
		return
	}
	msg := l.formatLog(Warn, message)
	l.write(msg)
}

func (l *ProgressLogger) Info(message string) {
	if l.logLevel < Info {
		return
	}
	msg := l.formatLog(Info, message)
	l.write(msg)
}

func (l *ProgressLogger) Debug(message string) {
	if l.logLevel < Debug {
		return
	}
	msg := l.formatLog(Debug, message)
	l.write(msg)
}

func (l *ProgressLogger) Verbose(message string) {
	if l.logLevel < Verbose {
		return
	}
	msg := l.formatLog(Verbose, message)
	l.write(msg)
}

func (l *ProgressLogger) GetStdout() io.Writer {
	return l.writer
}

func (l *ProgressLogger) GetStderr() io.Writer {
	return l.writer
}

func (l *ProgressLogger) write(message string) {
	_, err := l.writer.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func (l *ProgressLogger) formatLog(messageType MessageType, message string) string {
	var format string
	switch messageType {
	case Info:
		format = "[INFO]: %s"
	case Warn:
		format = "[WARN]: %s"
	case Error:
		format = "[ERROR]: %s"
	case Debug:
		format = "[DEBUG]: %s"
	case Verbose:
		format = "[VERBOSE]: %s"
	default:
		panic("Unknown message type")
	}
	return fmt.Sprintf(format, message)
}
