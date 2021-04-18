package queue_processor

import (
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/internal/transaction_manager"
	"github.com/TheAsda/skalka/pkg/config"
)

type Queue struct {
	queue    []Task
	current  int
	logger   progress_logger.ProgressLogger
	tm       transaction_manager.TransactionManager
	IsFailed bool
	config   config.GlobalConfig
}

func NewQueue(logger progress_logger.ProgressLogger, manager transaction_manager.TransactionManager, globalConfig config.GlobalConfig) *Queue {
	return &Queue{queue: []Task{}, logger: logger, current: 0, tm: manager, IsFailed: false, config: globalConfig}
}

func (q *Queue) Add(step config.Step) error {
	env := append(q.config.Env, step.Env...)
	path, err := q.tm.GetPath()
	if err != nil {
		q.logger.Error(err.Error())
		return err
	}
	q.queue = append(q.queue, *NewTask(step.Name, step.Run, q.logger.GetStdout(), q.logger.GetStderr(), env, path))
	return nil
}

func (q *Queue) ExecuteNext() error {
	task := q.queue[q.current]
	q.current++
	q.logger.LogStep(q.current, len(q.queue), task.name)
	err := task.Execute()
	if err != nil {
		q.IsFailed = true
	}
	return err
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q Queue) Finish() error {
	q.logger.Verbose("Finishing queue")
	err := q.tm.Commit()
	return err
}

func (q Queue) Cancel() error {
	q.logger.Verbose("Canceling queue")
	err := q.tm.Rollback()
	return err
}
