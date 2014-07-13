package main

import (
	"github.com/Pursuit92/simulator"
	"fmt"
)

type SingleRun struct {
	Model string
	Factors map[string]int
	Results map[string]int
	Reps int
}

func buildSims(conf FlagConf) []*SingleRun {
	numSims := len(conf.Models)
	for _,v := range conf.Factors {
		numSims *= len(v)
	}
	runs := make([]*SingleRun, numSims)
	/* Figure this out!
	 * Needs to generate records with all combinations
	 * of factor levels

	i := 0
	for _,v := range conf.Models {
		for j,w := range conf.Factors {
			for k,x := range w {
			}
		}
	}
	*/

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
	switch r.Model {
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
		s.OneToOne = true
	}
	s.Run()
}

func (r *SingleRun) Strings() []string {
	fields := 1 + len(r.Factors) + len(r.Results)
	out := make([]string,fields)
	out[0] = r.Model
	i := 1
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

func AllStrings(runs []*SingleRun) [][]string {
	numRuns := len(runs)
	outs := make([][]string,numRuns)
	for i,v := range runs {
		outs[i] = v.Strings()
	}
	return outs
}
