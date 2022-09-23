package vhost

import (
	"sync"
)

type Queue struct {
	msgs []Message
	mu   sync.Mutex
}

func CreateQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Empty() {
	q.mu.Lock()
	q.msgs = nil
	q.mu.Unlock()
}

func (q *Queue) IsErrorEvent() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, m := range q.msgs {
		if ErrorEvent == m.Event {
			return true
		}
	}
	return false
}

func (q *Queue) Exists(uri string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, m := range q.msgs {
		if uri == m.Sender {
			return true
		}
	}
	return false
}

func (q *Queue) Count() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.msgs == nil {
		return 0
	}
	return len(q.msgs)
}

func (q *Queue) Enqueue(msg Message) {
	q.mu.Lock()
	q.msgs = append(q.msgs, msg)
	q.mu.Unlock()
}

/*
func (q *queue) dequeue() *Message {
	if q.isEmpty() {
		return nil
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	msg := q.msgs[0]
	q.msgs = q.msgs[1:]
	return msg
}
*/
