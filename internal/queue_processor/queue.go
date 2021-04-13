package queue_processor

type Queue struct {
	queue []Task
}

func NewQueue() *Queue {
	return &Queue{queue: []Task{}}
}

func (q *Queue) Add(task Task) {
	q.queue = append(q.queue, task)
}

func (q *Queue) ExecuteNext() error {
	task := q.queue[0]
	q.queue = q.queue[1:]
	err := task.Execute()
	return err
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}
