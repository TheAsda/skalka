package queue_processor

import (
	"fmt"
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/internal/transaction_manager"
	"github.com/TheAsda/skalka/pkg/config"
	"sync"
)

type QueueProcessor struct {
	queue  Queue
	wg     sync.WaitGroup
	logger progress_logger.ProgressLogger
	config config.GlobalConfig
}

func NewQueueProcessor(logger progress_logger.ProgressLogger, config config.GlobalConfig) *QueueProcessor {
	return &QueueProcessor{logger: logger, config: config, queue: nil}
}

func (p *QueueProcessor) FillQueue(job config.Job) error {
	p.logger.Verbose("Filling queue")
	tm := transaction_manager.NewTransactionManager(p.config.Workdir, job.FlushOnError)
	queue := NewQueue(p.logger, *tm, p.config)
	for _, step := range job.Steps {
		p.logger.Verbose(fmt.Sprintf("Add step '%s' to queue", step.Name))
		err := queue.Add(step)
		if err != nil {
			return err
		}
	}
	p.queue = *queue
	return nil
}

func (p *QueueProcessor) Run() {
	p.logger.Verbose(fmt.Sprintf("Starting queue"))
	for p.queue.IsEmpty() {
		err := p.queue.ExecuteNext()
		if err != nil {
			p.logger.Error(err.Error())
			break
		}
	}
	p.logger.Verbose("Queue finished")
	if p.queue.IsFailed {
		err := p.queue.Cancel()
		if err != nil {
			p.logger.Error(err.Error())
		}
	} else {
		err := p.queue.Finish()
		if err != nil {
			p.logger.Error(err.Error())
		}
	}
}
