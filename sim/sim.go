package main

import (
	"github.com/Pursuit92/simulator"
)

func main() {
	sim := &simulator.Simulator{}
	ModelPrompt(sim)
	CustPrompt(sim)
	RandPrompt(sim)
	/*
	sim.CustStrat = simulator.CustShortest
	sim.QueueStrat = simulator.QueueLongest
	sim.ServerStrat = simulator.ServRand
	sim.InterRand = simulator.Uniform(0,10)
	sim.ServRand = simulator.Uniform(0,10)
	sim.OneToOne = true
	sim.NumQueues = 1
	sim.NumServers = 1
	sim.NumCusts = 10
	*/

	sim.Run()

	println(sim)
}
