package main

import (
	"github.com/Pursuit92/simulator"
	"fmt"
)

type SingleRun struct {
	Factors map[string]int
	Results map[string]string
}

func buildSims(conf FlagConf) []*SingleRun {
	numSims := conf.Reps
	for _,v := range conf.Factors {
		numSims *= len(v)
	}
	runs := make([]*SingleRun, numSims)

	tree := BuildTree(conf.Factors)

	i := 0
	for run := range GenRuns(tree) {
		for j := 0; j < conf.Reps; j++ {
			cpy := &SingleRun{}
			cpy.Factors = make(map[string]int)
			cpy.Results = make(map[string]string)
			for k,v := range run.Factors {
				cpy.Factors[k] = v
			}
			for _,v := range conf.Results {
				cpy.Results[v] = ""
			}
			runs[i] = cpy
			i++
		}
	}

	return runs
}

func (r *SingleRun) Run() {
	s := &simulator.Simulator{}
	s.InterRand = simulator.Poisson(1,100,r.Factors[IATime])
	s.ServRand = simulator.Poisson(1,100,r.Factors[STime])
	s.NumQueues = r.Factors[Queues]
	s.NumServers = r.Factors[Servers]
	s.NumCusts = r.Factors[Customers]
	// Not configurable :P
	s.CustStrat = simulator.CustRand
	s.QueueStrat = simulator.QueueRand
	s.ServerStrat = simulator.ServRand
	// reconfigure for model - this might change stuff
	switch r.Factors[Model] {
	case SQSS:
		s.NumQueues = 1
		s.NumServers = 1
		r.Factors[Servers] = 1
		r.Factors[Queues] = 1
		s.OneToOne = true
	case SQMS:
		s.NumQueues = 1
		r.Factors[Queues] = 1
	case PSQ:
		r.Factors[Servers] = r.Factors[Queues]
		s.NumServers = s.NumQueues
		s.OneToOne = true
	default:
	}
	s.Run()

	for i,_ := range r.Results {
		switch i {
		case Wait:
			r.Results[i] = s.AvgWait()
		case ProbWait:
			r.Results[i] = s.ProbWait()
		}
	}


}

func (r *SingleRun) Strings() []string {
	fields := len(r.Factors) + len(r.Results)
	out := make([]string,fields)
	i := 0
	for _,v := range r.Factors {
		out[i] = fmt.Sprint(v)
		i++
	}
	for _,v := range r.Results {
		out[i] = fmt.Sprint(v)
		i++
	}
	return out
}

func (r *SingleRun) Header() []string {
	fields := len(r.Factors) + len(r.Results)
	out := make([]string,fields)
	i := 0
	for v,_ := range r.Factors {
		out[i] = fmt.Sprint(v)
		i++
	}
	for v,_ := range r.Results {
		out[i] = fmt.Sprint(v)
		i++
	}
	return out
}

func AllStrings(runs []*SingleRun) [][]string {
	numRuns := len(runs)
	outs := make([][]string,1+numRuns)
	outs[0] = runs[0].Header()
	for i,v := range runs {
		outs[i+1] = v.Strings()
	}
	return outs
}
