package main

import (
	"fmt"
	"github.com/Pursuit92/simulator"
)

func ModelPrompt(s *simulator.Simulator) {
	fmt.Print("Select Simulation Model:\n\n")
	fmt.Print("\t1. Single Queue, Single Server\n\t2. Single Queue, Multi Server\n\t3. Multi Queue, Multi Server\n\t4. Multi Server, Per-Server Queues\n\n")
	var i int
	var err error

	for {
		fmt.Print("Input 1-4 [1]: ")
		_,err = fmt.Scan(&i)
		if err != nil || i < 1 || i > 4 {
			fmt.Println("Invalid Input")
		} else {
			fmt.Println()
			break
		}
	}
	switch i {
	case 1:
		s.NumServers = 1
		s.NumQueues = 1
		s.CustStrat = simulator.CustShortest
		s.QueueStrat = simulator.QueueLinear
		s.ServerStrat = simulator.ServLinear
		s.OneToOne = true
	case 2:
		s.NumQueues = 1
		s.CustStrat = simulator.CustShortest
		s.QueueStrat = simulator.QueueLinear
		ServersPrompt(s)
	case 3:
		QueuesPrompt(s)
		ServersPrompt(s)
	case 4:
		s.OneToOne = true
		ServersPrompt(s)
		QueuesPrompt(s)
		s.NumQueues = s.NumServers
	}
}

func QueuesPrompt(s *simulator.Simulator) {
	var i int
	if ! s.OneToOne {
		for {
			fmt.Print("Input Number of Queues: ")
			_,err := fmt.Scan(&i)
			if err != nil || i < 1 {
				fmt.Println("Invalid Input")
			} else {
				s.NumQueues = i
			fmt.Println()
				break
			}
		}
	}

	fmt.Print("Select new customer queue selection strategy:\n\n")
	fmt.Print("\t1. Shortest First\n\t2. Random\n\n")

	for {
		fmt.Print("Input 1-2 [1]: ")
		_,err := fmt.Scan(&i)
		if err != nil || i > 2 || i < 1 {
			fmt.Println("Invalid Input")
		} else {
			fmt.Println()
			break
		}
	}

	switch i {
	case 1:
		 s.CustStrat = simulator.CustShortest
	case 2:
		s.CustStrat = simulator.CustRand
	}

}

func ServersPrompt(s *simulator.Simulator) {
	var i int
	for {
		fmt.Print("Input Number of Servers: ")
		_,err := fmt.Scan(&i)
		if err != nil || i < 1 {
			fmt.Println("Invalid Input")
		} else {
			s.NumServers = i
			fmt.Println()
			break
		}
	}


	if ! s.OneToOne {
		fmt.Print("Select server allocation strategy:\n\n")
		fmt.Print("\t1. Linear\n\t2. Random\n\n")

		for {
			fmt.Print("Input 1-2 [1]: ")
			_,err := fmt.Scan(&i)
			if err != nil || i > 2 || i < 1 {
				fmt.Println("Invalid Input")
			} else {
				fmt.Println()
				break
			}
		}

		switch i {
		case 1:
			 s.ServerStrat = simulator.ServLinear
		case 2:
			s.ServerStrat = simulator.ServRand
		}

		if s.NumQueues > 1 {
			fmt.Print("Select Queue -> Server allocation strategy:\n\n")
			fmt.Print("\t1. Linear\n\t2. Longest First\n\t3. Random\n\n")

			for {
				fmt.Print("Input 1-3 [1]: ")
				_,err := fmt.Scan(&i)
				if err != nil || i > 3 || i < 1 {
					fmt.Println("Invalid Input")
				} else {
					fmt.Println()
					break
				}
			}

			switch i {
			case 1:
				 s.QueueStrat = simulator.QueueLinear
			case 2:
				s.QueueStrat = simulator.QueueLongest
			case 3:
				s.QueueStrat = simulator.QueueRand
			}
		}
	}

}

func CustPrompt(s *simulator.Simulator) {
	var i int
	for {
		fmt.Print("Input Number of Customers: ")
		_,err := fmt.Scan(&i)
		if err != nil || i < 1 {
			fmt.Println("Invalid Input")
		} else {
			s.NumCusts = i
			fmt.Println()
			break
		}
	}
}

type RandDistEnum int
const (
	RandUniform RandDistEnum = 1
	RandNormal = 2
	RandExponential = 3
	RandPoisson = 4
)

func RandPrompt(which string) (RandDistEnum,int,int,int,int) {
	fmt.Print("Select ",which," random distribution:\n\n")
	fmt.Print("\t1. Uniform\n\t2. Normal\n\t3. Exponential\n\t4. Poisson\n\n")
	var i RandDistEnum
	for {
		fmt.Print("Input 1-4 [1]: ")
		_,err := fmt.Scan(&i)
		if err != nil || i < 1 || i > 4 {
			fmt.Println("Invalid Input")
		} else {
			fmt.Println()
			break
		}
	}

	min,max := MinMaxPrompt()

	var ext1,ext2 int
	switch i {
	case RandUniform:
		return i,min,max,0,0
	case RandNormal:
		ext1,ext2 = RandExtraPrompt("mean","deviation")
	default:
		ext1,_ = RandExtraPrompt("mean","")
	}

	return i,min,max,ext1,ext2
}

func prompt(str string, i *int) {
	for {
		fmt.Print("Input ",str,": ")
		_,err := fmt.Scan(i)
		if err != nil || *i < 0 {
			fmt.Println("Invalid Input")
		} else {
			break
		}
	}
}

func MinMaxPrompt() (int,int) {
	var min,max int
	prompt("minimum",&min)
	prompt("maximum",&max)

	return min,max
}

func RandExtraPrompt(str1,str2 string) (int,int) {
	var ext1,ext2 int
	prompt(str1,&ext1)
	if str2 != "" {
		prompt(str2,&ext2)
	}
	return ext1,ext2
}

func DisplayResults(s *simulator.Simulator) {
	for {
		var step int
		prompt(fmt.Sprintf("Select time to view 1-%d (0 to quit): ",s.Steps),&step)
		if step != 0 {
			s.PrintStep(step)
		} else {
			s.PrintResults()
			break
		}
	}
}

