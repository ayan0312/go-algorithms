package circularqueue

import "errors"

type QueueItem struct {
	item interface{}
}

type CircularQueue struct {
	array []*QueueItem
	front int
	rear  int
}

func CreateQueueItem(item interface{}) *QueueItem {
	return &QueueItem{item: item}
}

func New(capacity int) *CircularQueue {
	cq := new(CircularQueue)
	cq.array = make([]*QueueItem, capacity)
	cq.front = 0
	cq.rear = 0
	return cq
}

func (cq *CircularQueue) EnQueue(element interface{}) error {
	if (cq.rear+1)%len(cq.array) == cq.front {
		return errors.New("")
	}
	cq.array[cq.rear] = CreateQueueItem(element)
	cq.rear = (cq.rear + 1) % len(cq.array)
	return nil
}

func (cq *CircularQueue) DeQueue() (*QueueItem, error) {
	if cq.rear == cq.front {
		return nil, errors.New("")
	}
	dqElement := cq.array[cq.front]
	cq.array[cq.front] = nil
	cq.front = (cq.front + 1) % len(cq.array)
	return dqElement, nil
}
