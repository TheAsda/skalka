package queue_processor

import (
	"github.com/TheAsda/skalka/internal"
	"github.com/TheAsda/skalka/internal/progress_logger"
	"github.com/TheAsda/skalka/pkg/config"
	"os"
	"sync"
)

type QueueProcessor struct {
	queues      map[string]Queue
	wg          sync.WaitGroup
	errors      map[string]error
	logger      progress_logger.ProgressLogger
	stepLoggers map[string]*progress_logger.StepLogger
	workdir     string
}

func (p *QueueProcessor) FillQueues(jobs config.Jobs) {
	for name, job := range jobs {
		queue := NewQueue()
		logger := p.logger.GetStepLogger(name, len(job.Steps))
		p.stepLoggers[name] = logger
		for _, step := range job.Steps {
			env := append(os.Environ(), step.Env...)
			queue.Add(*NewTask(step.Name, step.Run, logger.GetStdout(), logger.GetStderr(), env, p.workdir))
		}
		p.queues[name] = *queue
	}
}

func (p *QueueProcessor) Start() {
	for job, queue := range p.queues {
		p.wg.Add(1)
		go func(queue Queue, job string) {
			defer p.wg.Done()
			err := p.RunQueue(job, queue)
			if err != nil {
				p.errors[job] = err
			}
		}(queue, job)
	}
}

func (p *QueueProcessor) Wait() map[string]error {
	p.wg.Wait()
	return p.errors
}

func (p *QueueProcessor) RunQueue(name string, queue Queue) error {
	logger := p.stepLoggers[name]
	if logger == nil {
		return internal.NewError("Step logger is not defined")
	}
	for queue.IsEmpty() {
		err := logger.LogStep("Next step")
		if err != nil {
			return err
		}
		err = queue.ExecuteNext()
		if err != nil {
			return err
		}
	}
	return nil
}
