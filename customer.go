package simulator

import (
	"container/list"
	"fmt"
)

type CustStatus struct {
	Name string
	Status CustStatusEnum
}

type CustStatusEnum int
const (
	Unarrived CustStatusEnum = iota
	InQueue
	BeingServed
	Done
)

type Customer struct {
	Name string
	Status CustStatus
	StatusHist *list.List
	TimeQueue int
	TimeServed int
	Interarrival int
	InterOrig int
	IsFront bool
	queue *Queue
	server *Server
}

func (c *Customer) Update() {
	switch c.Status.Status {
	case InQueue:
		c.TimeQueue++
	case BeingServed:
		c.TimeServed++
	case Unarrived:
		if c.IsFront {
			c.Interarrival--
		}
	default:
	}
	cpy := *c
	c.StatusHist.PushBack(&cpy)
}

func (c *Customer) Done() {
	c.Status = CustStatus{Status: Done}
}

func NewCust(iat int) *Customer {
		c := Customer{}
		c.Status = CustStatus{Status: Unarrived}
		c.StatusHist = list.New()
		c.Interarrival = iat
		c.InterOrig = iat
		return &c
}

func NYATable(lst *list.List) *list.List {
	l := list.New()
	head := list.New()
	head.PushBack("Name")
	head.PushBack("Time")
	l.PushBack(head)
	for e := lst.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Value.(*Customer).ArrivalStrings())
	}
	return l
}

func DoneTable(lst *list.List) *list.List {
	l := list.New()
	head := list.New()
	head.PushBack("Name")
	head.PushBack("Queue Time")
	head.PushBack("Server Time")
	l.PushBack(head)
	for e := lst.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Value.(*Customer).DoneStrings())
	}
	return l
}

func (c *Customer) ArrivalStrings() *list.List {
	l := list.New()
	l.PushBack(c.Name)
	l.PushBack(fmt.Sprintf("%d",c.Interarrival))
	return l
}

func (c *Customer) DoneStrings() *list.List {
	l := list.New()
	l.PushBack(c.Name)
	l.PushBack(fmt.Sprintf("%d",c.TimeQueue))
	l.PushBack(fmt.Sprintf("%d",c.TimeServed))
	return l
}
