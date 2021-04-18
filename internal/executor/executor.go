package executor

import (
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/internal/queue_processor"
	"github.com/TheAsda/skalka/internal/transaction_manager"
	"github.com/TheAsda/skalka/internal/variables_store"
	"github.com/TheAsda/skalka/pkg/config"
	"sync"
)

type Executor struct {
	containers map[config.JobName]JobContainer
	gc         config.GlobalConfig
	logger     progress_logger.ProgressLogger
	runner     queue_processor.Runner
	wg         sync.WaitGroup
}

func NewExecutor(gc config.GlobalConfig, logger progress_logger.ProgressLogger, runner queue_processor.Runner) *Executor {
	return &Executor{containers: map[config.JobName]JobContainer{}, gc: gc, logger: logger, runner: runner}
}

func (e *Executor) Fill(jobs config.Jobs) error {
	for jobName, job := range jobs {
		container := JobContainer{
			job:   job,
			tm:    transaction_manager.NewTransactionManager(e.gc.Workdir, job.FlushOnError),
			store: variables_store.NewVariablesStore(),
		}
		container.processor = queue_processor.NewQueueProcessor(e.logger, e.gc, e.runner, *container.tm, *container.store)
		e.containers[jobName] = container
		err := container.processor.FillQueue(job)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Executor) Start() {
	e.logger.Verbose("Start executor")
	for _, c := range e.containers {
		e.wg.Add(1)
		go func(processor *queue_processor.QueueProcessor) {
			defer e.wg.Done()
			processor.Run()
		}(c.processor)
	}
}

func (e *Executor) Wait() {
	e.logger.Verbose("Waiting for processors to finish")
	e.wg.Wait()
	e.logger.Verbose("Processors finished")
}
