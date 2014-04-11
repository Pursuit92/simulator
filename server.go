package simulator

import (
	"container/list"
	"fmt"
)

type ServStatus struct {
	Name string
	Status ServStatusEnum
}

type ServStatusEnum int
const (
	Idle ServStatusEnum = iota
	Serving
)

type Server struct {
	Name string
	Status ServStatus
	StatusHist *list.List
	TimeLeft int
	TimeIdle int
	TimeServing int
	sim *Simulator
	cust *Customer
}

func (s *Server) PrintState() {
	if s.cust != nil {
		fmt.Printf("Serving %s\n",s.cust.Name)
		fmt.Printf("Time Remaining: %d\n",s.TimeLeft)
	} else {
		fmt.Println("Idle")
	}
}

func (s *Server) Update() {
	switch s.Status.Status {
	case Idle:
		s.TimeIdle++
	case Serving:
		s.TimeServing++
		s.cust.Update()
		s.TimeLeft--
		if s.TimeLeft <= 0 {
			s.sim.Done.PushFront(s.cust)
			s.Status = ServStatus{Status: Idle}
			s.cust.Done()
			s.cust = nil
		}
	default:
	}
	cpy := *s
	s.StatusHist.PushBack(&cpy)
}

func (s *Server) StartServing(c *Customer, time int) {
	s.Status = ServStatus{Name: c.Name, Status: Serving}
	s.TimeLeft = time
	c.Status = CustStatus{Name: s.Name, Status: BeingServed}
	c.queue = nil
	c.server = s
	s.cust = c
}

func NewServ() *Server {
		s := Server{}
		s.Status = ServStatus{Status: Idle}
		s.StatusHist = list.New()
		return &s
}

func AllIdle(lst *list.List) bool {
	return all(lst,func(i interface{}) bool {
		s := i.(*Server)
		return s.Status.Status == Idle
	})
}

func (s *Server) StateStrings() *list.List {
	l := list.New()
	l.PushBack(s.Name)
	if s.Status.Status != Idle {
		l.PushBack(s.cust.Name)
		l.PushBack(fmt.Sprintf("%d",s.TimeLeft))
	} else {
		l.PushBack("Idle")
		l.PushBack("")
	}

	return l
}

func ServerTable(srvs *list.List) *list.List {
	head := list.New()
	l := list.New()
	head.PushBack("Server")
	head.PushBack("Status")
	head.PushBack("Time Remaining")
	l.PushBack(head)
	for e := srvs.Front(); e != nil; e = e.Next() {
		srv := e.Value.(*Server)
		l.PushBack(srv.StateStrings())
	}
	return l
}
