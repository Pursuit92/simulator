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

	CustStrat NewCustStrat
	QueueStrat QueueAllocStrat
	ServerStrat ServAllocStrat
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
		srv.sim = s
		s.Servers.PushBack(srv)
	}
}

func (s *Simulator) GenQueues() {
	s.Queues = list.New()
	s.tmpQueues = list.New()
	for i := 0; i < s.NumQueues; i++ {
		q := NewQueue()
		q.Name = fmt.Sprintf("Queue %d",i)
		s.Queues.PushBack(q)
	}
}

func (s *Simulator) Init() {
	s.GenCusts()
	s.GenServers()
	s.GenQueues()
	s.Done = list.New()
}

func (s *Simulator) ProcArrivals() {
	for fst := s.Customers.Front();
		fst != nil;
		fst = s.Customers.Front() {
		    if fst.Value.(*Customer).Interarrival == 0 {
				// Remove a queue from the list of queues
				q := s.NewCustAlloc()
				// remove the customer and enqueue
				cust := s.Customers.Remove(fst).(*Customer)
				q.Enqueue(cust)

				// add the queue back
				s.Queues.PushFront(q)
			} else {
				break
			}
		}
	// update the frontmost customer
	if s.Customers.Len() >= 1 {
		s.Customers.Front().Value.(*Customer).Update()
	}
}

func (s *Simulator) QueuesToServers() {
	fmt.Println("QueuesToServers Start")
	s.simpleState()
	if s.OneToOne {
		for ; s.Servers.Len() != 0 ; {
			srv := s.ServerAlloc()
			s.tmpServers.PushBack(srv)
			q := s.QueueAlloc()
			s.tmpQueues.PushBack(q)
			if q.Size() > 0 && srv.Status.Status == Idle {
				srv.StartServing(q.Dequeue().(*Customer),s.ServRand())
			}
		}
		s.Queues.PushFrontList(s.tmpQueues)
		s.Servers.PushFrontList(s.tmpServers)
		s.tmpQueues = list.New()
		s.tmpServers = list.New()
	} else {
		for ; s.Servers.Len() != 0; {
			srv := s.ServerAlloc()
			s.tmpServers.PushBack(srv)

			found := false

			// locate a queue with a customer
			var q *Queue
			for ; s.Queues.Len() != 0; {
				q = s.QueueAlloc()
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
	s.simpleState()
	fmt.Println("QueuesToServers Done")
}

func (s *Simulator) UpdateServers() {
	all(s.Servers,func(i interface{}) bool {
		srv := i.(*Server)
		srv.Update()
		srv.PrintState()
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

func (s *Simulator) simpleState() {
	fmt.Printf("%d Servers\n%d Queues\n%d Customers to arrive\n%d Customers Done\n",
		s.Servers.Len(),s.Queues.Len(),s.Customers.Len(),s.Done.Len())
}

func (s *Simulator) Run() {
	s.Init()

	s.Configure()

	for {
		println("Step")
		s.ProcArrivals()

		s.QueuesToServers()

		s.UpdateServers()
		s.UpdateQueues()

		if AllIdle(s.Servers) && AllEmpty(s.Queues) && s.Customers.Len() == 0 {
			break
		}
	}
	s.simpleState()
}
