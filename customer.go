package simulator

import (
	"container/list"
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
		c.Interarrival--
		return
	default:
	}
	c.StatusHist.PushBack(c.Status)
}

func (c *Customer) Done() {
	c.Status = CustStatus{Status: Done}
}

func NewCust(iat int) *Customer {
		c := Customer{}
		c.Status = CustStatus{Status: Unarrived}
		c.StatusHist = list.New()
		c.Interarrival = iat
		return &c
}
