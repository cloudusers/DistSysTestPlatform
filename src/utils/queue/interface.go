package queue

import (
	"time"
)

// item for TQueue
type TQueueItem interface {
	// item value
	Value() interface{}
	Set_value(v interface{})

	// enqueued time
	EnqueuedTime() time.Time
	Set_enqueuedTime(t time.Time)

	// duration (second timeout)
	Duration() int
	Set_duration(i int)
}

// concurrent-safe, time-bounded queue.
//
// designed for ExplorePolicies.
type TQueue interface {
	// enqueue
	Enqueue(TQueueItem) error
	UpdateEnqueue(TQueueItem,
		time.Time,
		int) error

	// get channel for dequeue
	GetDequeueChan() chan TQueueItem
}
