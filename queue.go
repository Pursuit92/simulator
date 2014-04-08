package simulator

import (
	"container/list"
)

type QueueStatusEnum int

type Queue struct {
	Name string
	internal *list.List
	Status QueueStatus
	StatusHist *list.List
}

type QueueStatus struct {
	Status QueueStatusEnum
	Contents *list.List
}

func NewQueue() Queue {
	return Queue{internal: list.New()}
}

func (q Queue) Enqueue(i interface{}) {
	q.internal.PushFront(i)
}

func (q Queue) Dequeue() interface{} {
	front := q.internal.Front()
	if front != nil {
		return q.internal.Remove(front)
	} else {
		return nil
	}
}

func (q Queue) Size() int {
	return q.internal.Len()
}
