package simulator

type NewCustStrat int
type QueueAllocStrat int
type ServAllocStrat int

const (
	CustShortest NewCustStrat = iota
	CustRand
)

const (
	QueueShortest QueueAllocStrat = iota
	QueueLongest
	QueueLinear
	QueueRand
)

const (
	ServLinear ServAllocStrat = iota
	ServRand
)

func (s *Simulator) Configure() {
	switch s.CustStrat {
	case CustShortest:
		s.NewCustAlloc = MakeQueueSelector(ShortestFirst,s.Queues)
	default:
		s.NewCustAlloc = MakeQueueSelector(UniformSel,s.Queues)
	}
	switch s.QueueStrat {
	case QueueShortest:
		s.QueueAlloc = MakeQueueSelector(ShortestFirst,s.Queues)
	case QueueLongest:
		s.QueueAlloc = MakeQueueSelector(LongestFirst,s.Queues)
	case QueueLinear:
		s.QueueAlloc = MakeQueueSelector(Linear,s.Queues)
	default:
		s.QueueAlloc = MakeQueueSelector(UniformSel,s.Queues)
	}
	switch s.ServerStrat {
	case ServLinear:
		s.ServerAlloc = MakeServerSelector(Linear,s.Servers)
	default:
		s.ServerAlloc = MakeServerSelector(UniformSel,s.Servers)
	}
}
