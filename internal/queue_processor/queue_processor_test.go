package queue_processor

import (
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/pkg/config"
	"github.com/TheAsda/skalka/pkg/settings"
	"testing"
)

type RunnerMock struct {
	calls int
}

func NewRunnerMock() *RunnerMock {
	return &RunnerMock{calls: 0}
}

func (r *RunnerMock) Run(command string, env config.Env, workdir string) error {
	r.calls++
	return nil
}

func TestQueueProcessor_FillQueue(t *testing.T) {
	gc := config.GlobalConfig{
		Settings: settings.Settings{
			LogLevel: "info",
		},
		Env:     []string{},
		Workdir: "/path/",
	}

	t.Run("fill queue", func(t *testing.T) {
		runner := NewRunnerMock()
		job := config.Job{
			FlushOnError: false,
			Env:          []string{},
			Steps: []config.Step{
				{
					Name: "Execute command 1",
					Env:  []string{},
					Run:  "command1",
				},
				{
					Name: "Execute command 2",
					Env:  []string{},
					Run:  "command2",
				},
			},
		}
		qp := NewQueueProcessor(*progress_logger.NewProgressLogger(progress_logger.NewIOMock(), gc.Settings), gc, runner)
		err := qp.FillQueue(job)
		if err != nil {
			t.Fatalf("Error while filling queue processor: %s", err.Error())
		}
		if qp.queue.IsEmpty() {
			t.Fatalf("Queue is empty")
		}
		if len(qp.queue.queue) != len(job.Steps) {
			t.Fatalf("Queue does not contain all steps")
		}
		if runner.calls != 0 {
			t.Fatalf("Runner was executed")
		}
	})
}

func TestQueueProcessor_Run(t *testing.T) {
	gc := config.GlobalConfig{
		Settings: settings.Settings{
			LogLevel: "info",
		},
		Env:     []string{},
		Workdir: "/path/",
	}

	t.Run("run", func(t *testing.T) {
		runner := NewRunnerMock()
		job := config.Job{
			FlushOnError: false,
			Env:          []string{},
			Steps: []config.Step{
				{
					Name: "Execute command 1",
					Env:  []string{},
					Run:  "command1",
				},
				{
					Name: "Execute command 2",
					Env:  []string{},
					Run:  "command2",
				},
			},
		}
		qp := NewQueueProcessor(*progress_logger.NewProgressLogger(progress_logger.NewIOMock(), gc.Settings), gc, runner)
		err := qp.FillQueue(job)
		if err != nil {
			t.Fatalf("Cannot fill queue: %s", err.Error())
		}
		qp.Run()
		if runner.calls != len(job.Steps) {
			t.Fatalf("Queue processor did not execute all steps")
		}
	})
}
