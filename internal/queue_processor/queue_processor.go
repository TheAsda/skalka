package queue_processor

import (
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/pkg/config"
	"os"
	"sync"
)

type QueueProcessor struct {
	queues  map[string]Queue
	wg      sync.WaitGroup
	errors  map[string]error
	logger  progress_logger.ProgressLogger
	workdir string
}

func (p *QueueProcessor) FillQueues(jobs config.Jobs) {
	for name, job := range jobs {
		queue := NewQueue(p.logger)
		for _, step := range job.Steps {
			env := append(os.Environ(), step.Env...)
			queue.Add(*NewTask(step.Name, step.Run, p.logger.GetStdout(), p.logger.GetStderr(), env, p.workdir))
		}
		p.queues[name] = *queue
	}
}

func (p *QueueProcessor) Start() {
	for job, queue := range p.queues {
		p.wg.Add(1)
		go func(queue Queue, jobName string) {
			defer p.wg.Done()
			err := p.logger.LogJob(jobName)
			if err != nil {
				p.errors[jobName] = err
				return
			}
			err = p.RunQueue(queue)
			if err != nil {
				p.errors[jobName] = err
			}
		}(queue, job)
	}
}

func (p *QueueProcessor) Wait() map[string]error {
	p.wg.Wait()
	return p.errors
}

func (p *QueueProcessor) RunQueue(queue Queue) error {
	for queue.IsEmpty() {
		err := queue.ExecuteNext()
		if err != nil {
			return err
		}
	}
	return nil
}
