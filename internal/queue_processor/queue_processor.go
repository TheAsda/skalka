package queue_processor

import (
	"fmt"
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/internal/transaction_manager"
	"github.com/TheAsda/skalka/pkg/config"
	"sync"
)

type QueueProcessor struct {
	queues map[string]Queue
	wg     sync.WaitGroup
	logger progress_logger.ProgressLogger
	config config.GlobalConfig
}

func NewQueueProcessor(logger progress_logger.ProgressLogger, config config.GlobalConfig) *QueueProcessor {
	return &QueueProcessor{logger: logger, config: config, queues: map[string]Queue{}}
}

func (p *QueueProcessor) FillQueues(jobs config.Jobs) error {
	p.logger.Verbose("Filling queues")
	for name, job := range jobs {
		tm := transaction_manager.NewTransactionManager(p.config.Workdir, job.FlushOnError)
		queue := NewQueue(p.logger, *tm, p.config)
		for _, step := range job.Steps {
			err := queue.Add(step)
			if err != nil {
				return err
			}
		}
		p.queues[name] = *queue
	}
	return nil
}

func (p *QueueProcessor) Start() {
	p.logger.Verbose(fmt.Sprintf("Starting %d queues", len(p.queues)))
	for job, queue := range p.queues {
		p.wg.Add(1)
		go func(queue Queue, jobName string) {
			defer p.wg.Done()
			p.logger.LogJob(jobName)
			p.RunQueue(queue)
		}(queue, job)
	}
}

func (p *QueueProcessor) Wait() map[string]error {
	p.logger.Verbose("Waiting for jobs to complete")
	p.wg.Wait()
	p.logger.Verbose("Jobs completed")
	errors := map[string]error{}
	for jobName, queue := range p.queues {
		if queue.IsFailed {
			p.logger.Debug(fmt.Sprintf("Queue '%s' failed", jobName))
			err := queue.Cancel()
			if err != nil {
				p.logger.Error(err.Error())
				errors[jobName] = err
			}
		} else {
			p.logger.Debug(fmt.Sprintf("Queue '%s' did not fail", jobName))
			err := queue.Finish()
			if err != nil {
				p.logger.Error(err.Error())
				errors[jobName] = err
			}
		}
	}
	return errors
}

func (p *QueueProcessor) RunQueue(queue Queue) {
	for queue.IsEmpty() {
		err := queue.ExecuteNext()
		if err != nil {
			p.logger.Error(err.Error())
		}
	}
	p.logger.Verbose("Queue finished")
}
