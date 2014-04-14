package main

import (
	"github.com/Pursuit92/simulator"
)

func main() {
	sim := &simulator.Simulator{}
	ModelPrompt(sim)
	CustPrompt(sim)
	interDist,interMin,interMax,interExt1,interExt2 := RandPrompt("interarrival")
	servDist,servMin,servMax,servExt1,servExt2 := RandPrompt("work time")
	switch interDist {
	case RandUniform:
		sim.InterRand = simulator.Uniform(interMin,interMax)
	case RandNormal:
		sim.InterRand = simulator.Normal(interMin,interMax,interExt1,interExt2)
	case RandExponential:
		sim.InterRand = simulator.Exponential(interMin,interMax,interExt1)
	case RandPoisson:
		sim.InterRand = simulator.Poisson(interMin,interMax,interExt1)
	}
	switch servDist {
	case RandUniform:
		sim.ServRand = simulator.Uniform(servMin,servMax)
	case RandNormal:
		sim.ServRand = simulator.Normal(servMin,servMax,servExt1,servExt2)
	case RandExponential:
		sim.ServRand = simulator.Exponential(servMin,servMax,servExt1)
	case RandPoisson:
		sim.ServRand = simulator.Poisson(servMin,servMax,servExt1)
	}

	sim.Run()


	DisplayResults(sim)

}
