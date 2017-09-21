package queue

import (
	"fmt"
	"time"
)

// implements TQueueItem
type BasicTimeoutQueueItem struct {
	value        interface{}
	enqueuedTime time.Time
	duration     int
}

func NewBasicTimeoutQueueItem(value interface{},
	timeout int) (*BasicTimeoutQueueItem, error) {
	return &BasicTimeoutQueueItem{
		value:        value,
		enqueuedTime: time.Unix(0, 0), // UNIX epoch (Jan 1, 1970 UTC)
		duration:     timeout,
	}, nil
}

func (this *BasicTimeoutQueueItem) Value() interface{} {
	return this.value
}

func (this *BasicTimeoutQueueItem) Set_value(v interface{}) {
	this.value = v
}

func (this *BasicTimeoutQueueItem) EnqueuedTime() time.Time {
	return this.enqueuedTime
}

func (this *BasicTimeoutQueueItem) Set_enqueuedTime(t time.Time) {
	this.enqueuedTime = t
}

func (this *BasicTimeoutQueueItem) Duration() int {
	return this.duration
}

func (this *BasicTimeoutQueueItem) Set_duration(duration int) {
	this.duration = duration
}

// implements TQueue
type BasicTimeoutQueue struct {
	dequeueChan chan TQueueItem
}

func NewBasicTimeoutQueue() TQueue {
	q := &BasicTimeoutQueue{
		dequeueChan: make(chan TQueueItem),
	}

	return q
}

func (this *BasicTimeoutQueue) Enqueue(item_ TQueueItem) error {
	//fmt.Println(item_)
	var item *BasicTimeoutQueueItem
	item, ok := item_.(*BasicTimeoutQueueItem)
	if !ok {
		return fmt.Errorf("bad item %s", item_)
	}

	item.enqueuedTime = time.Now()
	go func() {
		this.dequeueChan <- item
	}()
	//<-time.After(duration)
	return nil
}

func (this *BasicTimeoutQueue) UpdateEnqueue(item_ TQueueItem,
	enqueue_time time.Time,
	duration int) error {
	//fmt.Println(item_)
	var item *BasicTimeoutQueueItem
	item, ok := item_.(*BasicTimeoutQueueItem)
	if !ok {
		return fmt.Errorf("bad item %s", item_)
	}

	item.enqueuedTime = enqueue_time
	item.duration = duration
	go func() {
		this.dequeueChan <- item
	}()
	//<-time.After(duration)
	return nil
}

func (this *BasicTimeoutQueue) GetDequeueChan() chan TQueueItem {
	return this.dequeueChan
}
