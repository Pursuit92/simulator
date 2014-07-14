package main

type FactorNode struct {
	Factor string
	Value int
	Children *FactorNode
	Next *FactorNode
}

func BuildTree(factors map[string][]int) *FactorNode {
	n := &FactorNode{}
	for i,v := range factors {
		n.AddFactor(i,v)
	}
	return n
}

func (n *FactorNode) AddFactor(factor string,levels []int) {
	if n.Children == nil {
		for _,v := range levels {
			n.AddChild(factor,v)
		}
	} else {
		for ch := n.Children; ch != nil; ch = ch.Next {
			ch.AddFactor(factor,levels)
		}
	}
}

func (n *FactorNode) AddChild(factor string,v int) {
	if n.Children == nil {
		n.Children = NewFactorNode(factor,v)
	} else {
		for ch := n.Children; ch != nil; ch = ch.Next {
			if ch.Next == nil {
				ch.Next = NewFactorNode(factor,v)
				break
			}
		}
	}
}

func NewFactorNode(factor string,v int) *FactorNode {
	return &FactorNode{Value: v,Factor: factor}
}

func GenRuns(n *FactorNode) <-chan SingleRun {
	runs := make(chan SingleRun)
	if n.Children == nil {
		return nil
	} else {
		go func() {
			run := SingleRun{}
			run.Factors = make(map[string]int)
			for ch := n.Children; ch != nil; ch = ch.Next {
				GenRunsRec(ch,run,runs)
			}
			close(runs)
		}()
		return runs
	}
}

func GenRunsRec(n *FactorNode,run SingleRun, runs chan<- SingleRun) {
	run.Factors[n.Factor] = n.Value
	if n.Children != nil {
		for ch := n.Children; ch != nil; ch = ch.Next {
			GenRunsRec(ch,run,runs)
		}
	} else {
		finalRun := SingleRun{}
		finalRun.Factors = make(map[string]int)
		for i,v := range run.Factors {
			finalRun.Factors[i] = v
		}
		runs <- finalRun
	}
}
