package queue

import (
	"errors"
	"sync"
)

// NewQueue : returns new instance of queue
func NewQueue() *Queue {
	return new(Queue)
}

// Queue holds data
// this queue build by using slice
type Queue struct {
	data []string
	lock sync.Mutex
}

// Enqueue : push the data to top of queue
func (q *Queue) Enqueue(value string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	// add data to tail of queue
	q.data = append(q.data, value)
}

// Dequeue : pop the data from queue. It will remove data from front
func (q *Queue) Dequeue() (string, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// if queue is empty return error
	if q.Empty() {
		return "", errors.New("dequeue error : Queue is empty")
	}

	dequeuedData := q.Front()

	// pop top
	q.data = q.data[1:]

	return dequeuedData, nil
}

// Front returns front value of queue
func (q *Queue) Front() string {
	if !q.Empty() {
		return q.data[0]
	}

	return ""
}

// Size returns the size of queue
func (q *Queue) Size() int {
	return len(q.data)
}

// Empty : returns if queue is empty or not
func (q *Queue) Empty() bool {
	return q.Size() == 0
}
