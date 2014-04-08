package simulator

import (
	"container/list"
	"fmt"
)

type Simulator struct {
	NumQueues,NumServers,NumCusts int

	InterDist func() int
	NewCustAlloc Selector
	ServerAlloc Selector
	QueueAlloc Selector

	Customers *list.List
	Queues *list.List
	Servers *list.List

	tmpQueues *list.List
	tmpServers *list.List

	// Customers that are finished
	Done *list.List
}

func (s *Simulator) GenCusts() {
	for i := 0; i < s.NumCusts; i++ {
		c := NewCust(s.InterDist())
		c.Name = fmt.Sprintf("Customer %d",i)
		s.Customers.PushBack(c)
	}
}

func (s *Simulator) GenServers() {
	for i := 0; i < s.NumServers; i++ {
		srv := NewServ()
		srv.Name = fmt.Sprintf("Server %d",i)
		s.Servers.PushBack(srv)
	}
}

func (s *Simulator) GenQueues() {
	for i := 0; i < s.NumQueues; i++ {
		q := NewQueue()
		q.Name = fmt.Sprintf("Queue %d",i)
		s.Servers.PushBack(q)
	}
}

func (s *Simulator) Init() {
	s.GenCusts()
	s.GenServers()
	s.GenQueues()
}
