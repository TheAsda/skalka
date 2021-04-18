package queue_processor

import (
	"github.com/TheAsda/skalka/pkg/config"
	"io"
)

type Task struct {
	name    string
	command string
	stdout  io.Writer
	stderr  io.Writer
	env     config.Env
	dir     string
	runner  Runner
	// TODO: add variables store
}

func NewTask(name string, command string, stdout io.Writer, stderr io.Writer, env config.Env, dir string, runner Runner) *Task {
	return &Task{name: name, command: command, stdout: stdout, stderr: stderr, env: env, dir: dir, runner: runner}
}

func (t *Task) Execute() error {
	err := t.runner.Run(t.command, t.env, t.dir)
	return err
}
