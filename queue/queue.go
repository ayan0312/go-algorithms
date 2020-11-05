package queue

type QueueItem struct {
	item interface{}
	prev *QueueItem
}

type Queue struct {
	current *QueueItem
	last    *QueueItem
	depth   uint64
}

func New() *Queue {
	return &Queue{depth: 0}
}

func (q *Queue) Enqueue(item interface{}) {
	if q.depth == 0 {
		q.current = &QueueItem{item: item, prev: nil}
		q.last = q.current
		q.depth++
		return
	}

	newQueueItem := &QueueItem{item: item, prev: nil}
	q.last.prev = newQueueItem
	q.last = newQueueItem
	q.depth++
}

func (q *Queue) Dequeue() interface{} {
	if q.depth > 0 {
		item := q.current.item
		q.current = q.current.prev
		q.depth--

		return item
	}

	return nil
}
