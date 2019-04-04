package queue

import "errors"

type queueItem struct {
}

// Queue is an object that handles its own queue
type Queue struct {
	items []*queueItem
}

// NewQueue will create and initialize a new queue object
func NewQueue() *Queue {
	queue := Queue{}
	queue.items = make([]*queueItem, 0)

	return &queue
}

// Add will add a new item to the queue
func (o *Queue) Add(newItem *queueItem) error {
	if newItem == nil {
		return errors.New("queue: Unable to add item to queue, null reference detected")
	}

	o.items = append(o.items, newItem)

	return nil
}

// Process will prompt a loop that handles the entire queue, when there are no futher items left in the queue this function will return
func (o *Queue) Process() error {
	return nil
}
