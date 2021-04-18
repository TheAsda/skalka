package queue_processor

import (
	"github.com/TheAsda/skalka/pkg/config"
)

type Task struct {
	name    string
	command string
	env     config.Env
	dir     string
	runner  Runner
	// TODO: add variables store
}

func NewTask(name string, command string, env config.Env, dir string, runner Runner) *Task {
	return &Task{name: name, command: command, env: env, dir: dir, runner: runner}
}

func (t *Task) Execute() error {
	err := t.runner.Run(t.command, t.env, t.dir)
	return err
}
