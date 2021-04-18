package queue_processor

import (
	"github.com/TheAsda/skalka/pkg/config"
	"io"
	"os/exec"
)

type Runner interface {
	Run(command string, env config.Env, workdir string) error
}

type TaskRunner struct {
	commandRunner string
	arguments     []string
	stdout        io.Writer
	stderr        io.Writer
}

func NewTaskRunner(commandRunner string, arguments []string, stdout io.Writer, stderr io.Writer) *TaskRunner {
	return &TaskRunner{commandRunner: commandRunner, arguments: arguments, stdout: stdout, stderr: stderr}
}

func (r TaskRunner) Run(command string, env config.Env, workdir string) error {
	c := exec.Command("cmd", "/C", command)
	c.Stdout = r.stdout
	c.Stderr = r.stderr
	c.Env = env
	c.Dir = workdir
	err := c.Run()
	return err
}
