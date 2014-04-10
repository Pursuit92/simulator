package simulator

import (
	"container/list"
	"fmt"
)

type randGen func() int

type Simulator struct {
	NumQueues,NumServers,NumCusts int

	NewCustAlloc QueueSelector
	ServerAlloc ServerSelector
	QueueAlloc QueueSelector

	Customers *list.List
	Queues *list.List
	Servers *list.List

	tmpQueues *list.List
	tmpServers *list.List

	InterRand randGen
	ServRand randGen

	// Customers that are finished
	Done *list.List

	OneToOne bool
}

func (s *Simulator) GenCusts() {
	s.Customers = list.New()
	for i := 0; i < s.NumCusts; i++ {
		c := NewCust(s.InterRand())
		c.Name = fmt.Sprintf("Customer %d",i)
		c.Interarrival = s.InterRand()
		s.Customers.PushBack(c)
	}
}

func (s *Simulator) GenServers() {
	s.Servers = list.New()
	s.tmpServers = list.New()
	for i := 0; i < s.NumServers; i++ {
		srv := NewServ()
		srv.Name = fmt.Sprintf("Server %d",i)
		s.Servers.PushBack(srv)
	}
}

func (s *Simulator) GenQueues() {
	s.Queues = list.New()
	s.tmpQueues = list.New()
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

func (s *Simulator) ProcArrivals() {
	for fst := s.Customers.Front();
		fst.Value.(*Customer).Interarrival == 0;
		fst = s.Customers.Front() {
			// Remove a queue from the list of queues
			q := s.NewCustAlloc()
			// remove the customer and enqueue
			cust := s.Customers.Remove(fst).(*Customer)
			q.Enqueue(cust)

			// add the queue back
			s.Queues.PushFront(q)
		}
	// update the frontmost customer
	s.Customers.Front().Value.(*Customer).Update()
}

func (s *Simulator) QueuesToServers() {
	if s.OneToOne {
		for srv := s.ServerAlloc(); s.Servers.Len() != 0; srv = s.ServerAlloc() {
			s.tmpServers.PushBack(srv)
			q := s.QueueAlloc()
			s.tmpQueues.PushBack(q)
			srv.StartServing(q.Dequeue().(*Customer),s.ServRand())
		}
		s.Queues.PushFrontList(s.tmpQueues)
		s.Servers.PushFrontList(s.tmpServers)
		s.tmpQueues = list.New()
		s.tmpServers = list.New()
	} else {
		for srv := s.ServerAlloc(); s.Servers.Len() != 0; srv = s.ServerAlloc() {
			s.tmpServers.PushBack(srv)

			found := false

			// locate a queue with a customer
			var q *Queue
			for q = s.QueueAlloc(); s.Queues.Len() != 0; q = s.QueueAlloc() {
				s.tmpQueues.PushBack(q)
				if q.Size() != 0 {
					found = true
					break
				}
			}
			// Return them to the pool
			s.Queues.PushFrontList(s.tmpQueues)
			s.tmpQueues = list.New()

			if found {
				srv.StartServing(q.Dequeue().(*Customer), s.ServRand())
			} else {
				// if no suitable queue found, no other servers can serve either
				break
			}
		}
		s.Servers.PushFrontList(s.tmpServers)
		s.tmpServers = list.New()
	}
}

func (s *Simulator) UpdateServers() {
	all(s.Servers,func(i interface{}) bool {
		srv := i.(*Server)
		srv.Update()
		return true
	})
}

func (s *Simulator) UpdateQueues() {
	all(s.Queues,func(i interface{}) bool {
		q := i.(*Queue)
		q.Update()
		return true
	})
}

func (s *Simulator) ConfigRands() {
}

func (s *Simulator) ConfigRest() {
}

func (s *Simulator) ConfigReset() {
}

func (s *Simulator) Run() {
	s.ConfigRands()

	s.Init()

	s.ConfigRest()

	for {
		s.ProcArrivals()

		s.QueuesToServers()

		s.UpdateServers()
		s.UpdateQueues()

		if AllIdle(s.Servers) && AllEmpty(s.Queues) && s.Customers.Len() == 0 {
			break
		}
	}
}
