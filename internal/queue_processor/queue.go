package queue_processor

import "github.com/TheAsda/skalka/internal/progress_logger"

type Queue struct {
	queue   []Task
	current int
	logger  progress_logger.ProgressLogger
}

func NewQueue(logger progress_logger.ProgressLogger) *Queue {
	return &Queue{queue: []Task{}, logger: logger, current: 0}
}

func (q *Queue) Add(task Task) {
	q.queue = append(q.queue, task)
}

func (q *Queue) ExecuteNext() error {
	task := q.queue[q.current]
	q.current++
	err := q.logger.LogStep(q.current, len(q.queue), task.name)
	if err != nil {
		return err
	}
	err = task.Execute()
	return err
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}
