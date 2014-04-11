package simulator

import (
	"container/list"
	"fmt"
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

func NewQueue() *Queue {
	return &Queue{internal: list.New(),StatusHist: list.New()}
}

func (q *Queue) Enqueue(i interface{}) {
	q.internal.PushBack(i)
}

func (q *Queue) Dequeue() interface{} {
	front := q.internal.Front()
	if front != nil {
		return q.internal.Remove(front)
	} else {
		return nil
	}
}

func (q *Queue) Size() int {
	return q.internal.Len()
}

func (q *Queue) Update() {
	for e := q.internal.Front(); e != nil; e = e.Next() {
		e.Value.(*Customer).Update()
	}
	cpy := *q
	cpy.internal = list.New()
	cpy.internal.PushFrontList(q.internal)
	q.StatusHist.PushBack(&cpy)
}

func AllEmpty(lst *list.List) bool {
	return all(lst,func(i interface{}) bool {
		q := i.(*Queue)
		return q.Size() == 0
	})
}

func (q *Queue) StateStrings() *list.List {
	l := list.New()
	head := list.New()
	head.PushBack(fmt.Sprintf("%s Contents",q.Name))
	head.PushBack("Time in Queue")
	l.PushBack(head)
//	if q.internal.Len() != 0 {
//		panic("Hey!")
//	}
	for e := q.internal.Front(); e != nil; e = e.Next() {
		cust := e.Value.(*Customer)
		custEnt := list.New()
		custEnt.PushBack(cust.Name)
		custEnt.PushBack(fmt.Sprintf("%d",cust.TimeQueue))
		l.PushBack(custEnt)
	}
	return l
}
