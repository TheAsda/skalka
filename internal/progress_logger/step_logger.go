package progress_logger

import (
	"fmt"
	"io"
)

type StepLogger struct {
	writer  io.Writer
	reader  io.Reader
	jobName string
	total   int
	current int
}

func NewStepLogger(writer io.Writer, reader io.Reader, jobName string, total int) *StepLogger {
	return &StepLogger{writer: writer, reader: reader, jobName: jobName, total: total, current: 1}
}

func (l *StepLogger) GetStdout() io.Writer {
	return l.writer
}

func (l *StepLogger) GetStderr() io.Writer {
	return l.writer
}

func (l *StepLogger) LogStep(name string) error {
	msg := l.formatStep(name)
	_, err := l.writer.Write([]byte(msg))
	l.nextStep()
	return err
}

func (l *StepLogger) formatStep(name string) string {
	return fmt.Sprintf("[%s] %d/%d: %s", l.jobName, l.current, l.total, name)
}

func (l *StepLogger) nextStep() {
	if l.current == l.total {
		panic("Step limit reached")
	}
	l.current += 1
}
