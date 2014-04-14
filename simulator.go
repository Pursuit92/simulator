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
	Steps int
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
				cust.Status = CustStatus{Name: q.Name, Status: InQueue}
				q.Enqueue(cust)

				// add the queue back
				s.Queues.PushFront(q)
			} else {
				break
			}
		}
	// update the frontmost customer
	if s.Customers.Len() >= 1 {
		s.Customers.Front().Value.(*Customer).IsFront = true
	}
}

func (s *Simulator) QueuesToServers() {
	//fmt.Println("QueuesToServers Start")
	//s.simpleState()
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

			if srv.Status.Status == Idle {
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
		}
		s.Servers.PushFrontList(s.tmpServers)
		s.tmpServers = list.New()
	}
	//s.simpleState()
	//fmt.Println("QueuesToServers Done")
}

func (s *Simulator) UpdateServers() {
	all(s.Servers,func(i interface{}) bool {
		srv := i.(*Server)
		srv.Update()
		//srv.PrintState()
		return true
	})
}

func (s *Simulator) UpdateUnarrived() {
	all(s.Customers,func(i interface{}) bool {
		c := i.(*Customer)
		c.Update()
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
	fmt.Printf("Running Simulator... ")
	s.Init()

	s.Configure()

	for {
		s.Steps++

/*
		//println("Step")
		s.PrintNYA()
		s.PrintQueues()
		s.PrintServers()
		s.PrintDone()
*/

		s.ProcArrivals()

		s.QueuesToServers()

		s.UpdateUnarrived()
		s.UpdateServers()
		s.UpdateQueues()

		if AllIdle(s.Servers) && AllEmpty(s.Queues) && s.Customers.Len() == 0 {
			break
		}
	}
	//s.simpleState()
	fmt.Printf("Done! Ran for %d steps.\n",s.Steps)
/*
	s.PrintNYA()
	s.PrintQueues()
	s.PrintServers()
	s.PrintDone()
*/
}

func (s *Simulator) PrintQueues() {
	for e := s.Queues.Front(); e != nil; e = e.Next() {
		PrintTable(e.Value.(*Queue).StateStrings())
		if e.Next() != nil {
			fmt.Println()
		}
	}
}

func (s *Simulator) PrintQueuesAt(time int) {
	for e := s.Queues.Front(); e != nil; e = e.Next() {
		f := e.Value.(*Queue).StatusHist.Front()
		for i := 0; i < time && f != nil; i++ {
			f = f.Next()
		}
		if f != nil {
			PrintTable(f.Value.(*Queue).StateStrings())
		}
		if e.Next() != nil {
			fmt.Println()
		}
	}
}

func (s *Simulator) PrintServers() {
	PrintTable(ServerTable(s.Servers))
}

func (s *Simulator) PrintServersAt(time int) {
	servs := list.New()
	for e := s.Servers.Front(); e != nil; e = e.Next() {
		f := e.Value.(*Server).StatusHist.Front()
		for i := 0; i < time && f != nil; i++ {
			f = f.Next()
		}
		if f != nil {
			servs.PushBack(f.Value.(*Server))
		}
	}

	PrintTable(ServerTable(servs))
}

func (s *Simulator) PrintNYA() {
	PrintTable(NYATable(s.Customers))
}

func (s *Simulator) PrintNYAAt(time int) {
	custs := list.New()
	for e := s.Customers.Front(); e != nil; e = e.Next() {
		f := e.Value.(*Customer).StatusHist.Front()
		for i := 0; i < time && f != nil; i++ {
			f = f.Next()
		}
		if f != nil {
			custs.PushBack(f.Value.(*Customer))
		}
	}

	PrintTable(NYATable(custs))
}


func (s *Simulator) PrintDone() {
	PrintTable(DoneTable(s.Done))
}

func (s *Simulator) PrintDoneAt(time int) {
	custs := list.New()
	for e := s.Done.Front(); e != nil; e = e.Next() {
		f := e.Value.(*Customer).StatusHist.Front()
		for i := 0; i < time && f != nil; i++ {
			f = f.Next()
		}
		if f != nil {
			custs.PushBack(f.Value.(*Customer))
		}
	}

	PrintTable(DoneTable(custs))
}

func (s *Simulator) PrintStep(step int) {
	step -= 1
	s.PrintQueuesAt(step)
	fmt.Println()
	s.PrintServersAt(step)
}

func (s *Simulator) PrintResults() {
	tab := list.New()
	tab.PushFront(TableEntry("Average Wait Time",s.AvgWait()))
	tab.PushFront(TableEntry("Probability of Waiting",s.ProbWait()))
	tab.PushFront(TableEntry("Probability of Server Idle",s.ProbIdle()))
	tab.PushFront(TableEntry("Average Server Utilization",s.ServerUt()))
	tab.PushFront(TableEntry("Average Service Time",s.AvgServ()))
	tab.PushFront(TableEntry("Average Interarrival Time",s.AvgInter()))
	tab.PushFront(TableEntry("Average Time Spent in System",s.AvgSys()))
	PrintTable(tab)
}

func (s *Simulator) AvgWait() string {
	totWait := funcSum(s.Done,func(i interface{}) int {
		cust := i.(*Customer)
		return cust.TimeQueue
	})

	return fmt.Sprintf("%0.2f",float64(totWait) / float64(s.NumCusts))
}

func (s *Simulator) ProbWait() string {
	totWait := funcSum(s.Done,func(i interface{}) int {
		cust := i.(*Customer)
		if cust.TimeQueue > 0 {
			return 1
		}
		return 0
	})

	return fmt.Sprintf("%0.2f",float64(totWait) / float64(s.NumCusts))
}

func (s *Simulator) ProbIdle() string {
	idling := funcSum(s.Servers, func(i interface{}) int {
		srv := i.(*Server)
		return srv.TimeIdle
	})

	return fmt.Sprintf("%0.2f",float64(idling) / float64(s.Steps))
}

func (s *Simulator) ServerUt() string {
	used := funcSum(s.Servers, func(i interface{}) int {
		srv := i.(*Server)
		return srv.TimeServing
	})

	return fmt.Sprintf("%0.2f",float64(used) / float64(s.Steps))
}

func (s *Simulator) AvgServ() string {
	servTime := funcSum(s.Done, func(i interface{}) int {
		cust := i.(*Customer)
		return cust.TimeServed
	})

	return fmt.Sprintf("%0.2f",float64(servTime) / float64(s.NumCusts))
}

func (s *Simulator) AvgSys() string {
	sysTime := funcSum(s.Done, func(i interface{}) int {
		cust := i.(*Customer)
		return cust.TimeServed + cust.TimeQueue
	})

	return fmt.Sprintf("%0.2f",float64(sysTime) / float64(s.NumCusts))
}

func (s *Simulator) AvgInter() string {
	inter := funcSum(s.Done, func(i interface{}) int {
		cust := i.(*Customer)
		return cust.InterOrig
	})
	return fmt.Sprintf("%0.2f",float64(inter) / float64(s.NumCusts))
}

