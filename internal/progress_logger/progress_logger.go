package progress_logger

import (
	"fmt"
	"io"
)

type ProgressLogger struct {
	writer  io.Writer
	reader  io.Reader
	jobName string
	total   int
	current int
}

func NewProgressLogger(writer io.Writer, reader io.Reader, jobName string, total int) *ProgressLogger {
	return &ProgressLogger{writer: writer, reader: reader, jobName: jobName, total: total, current: 1}
}

func (l *ProgressLogger) LogInfo(message string) {

}

func (l *ProgressLogger) LogStep(name string) error {
	msg := l.formatStep(name)
	_, err := l.writer.Write([]byte(msg))
	l.nextStep()
	return err
}

func (l *ProgressLogger) formatStep(name string) string {
	return fmt.Sprintf("[%s] %d/%d: %s", l.jobName, l.current, l.total, name)
}

func (l *ProgressLogger) nextStep() {
	if l.current == l.total {
		panic("Step limit reached")
	}
	l.current += 1
}
