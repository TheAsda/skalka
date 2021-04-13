package queue_processor

import (
	"github.com/TheAsda/skalka/pkg/config"
	"io"
	"os/exec"
)

type Task struct {
	name    string
	command string
	stdout  io.Writer
	stderr  io.Writer
	env     config.Env
	dir     string
}

func NewTask(name string, command string, stdout io.Writer, stderr io.Writer, env config.Env, dir string) *Task {
	return &Task{name: name, command: command, stdout: stdout, stderr: stderr, env: env, dir: dir}
}

func (t *Task) Execute() error {
	c := exec.Command("cmd", "/C", t.command)
	c.Stdout = t.stdout
	c.Stdout = t.stderr
	c.Env = t.env
	c.Dir = t.dir
	err := c.Run()
	return err
}
