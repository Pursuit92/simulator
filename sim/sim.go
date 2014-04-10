package main

import (
	"github.com/Pursuit92/simulator"
)

func main() {
	sim := &simulator.Simulator{}
	/*
	sim.NewCustAlloc = simulator.MakeQueueSelector(simulator.UniformSel,sim.Queues)
	sim.ServerAlloc = simulator.MakeServerSelector(simulator.UniformSel,sim.Servers)
	sim.QueueAlloc = simulator.MakeQueueSelector(simulator.UniformSel,sim.Queues)
	sim.InterRand = simulator.Uniform(0,10)
	sim.ServRand = simulator.Uniform(0,10)
	sim.OneToOne = true
	sim.NumQueues = 1
	sim.NumServers = 1
	sim.NumCusts = 1

	sim.Run()
	*/

	println(sim)
}
