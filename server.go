package simulator

import (
	"container/list"
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

func (s *Server) Update() {
	switch s.Status.Status {
	case Idle:
		s.TimeIdle++
	case Serving:
		s.TimeServing++
		s.cust.Update()
	default:
	}
	s.StatusHist.PushBack(s.Status)
	s.TimeLeft--
	if s.TimeLeft == 0 {
		s.sim.Done.PushFront(s.cust)
		s.Status = ServStatus{Status: Idle}
		s.cust.Done()
		s.cust = nil
	}
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
