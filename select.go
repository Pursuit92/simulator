package simulator

import (
	"container/list"
	"math/rand"
)

type Selector func(*list.List) interface{}

func compFirst(lst *list.List, comp func(interface{}, interface{}) bool) interface{} {
	println("Performing a select by comparison")
	curr := lst.Front()
	for e := curr.Next(); e != nil; e = e.Next() {
		if comp(e.Value, curr.Value) {
			curr = e
		}
	}
	if curr != nil {
		lst.Remove(curr)
		return curr.Value
	}
	return nil
}

func ShortestFirst(lst *list.List) interface{} {
	println("Creating a Shortest First selector")
	return compFirst(lst,
		func(v, curr interface{}) bool {
			return v.(*Queue).Size() < curr.(*Queue).Size()
		})
}

func LongestFirst(lst *list.List) interface{} {
	println("Creating a Longest First selector")
	return compFirst(lst,
		func(v, curr interface{}) bool {
			return v.(*Queue).Size() > curr.(*Queue).Size()
		})
}

func Linear(lst *list.List) interface{} {
	println("Performing Linear Selection")
	e := lst.Front()
	if e != nil {
		lst.Remove(e)
		return e.Value
	}
	return nil
}

func funcSelector(lst *list.List, f func(int) int) interface{} {
	println("Performing a Function selection")
	print("List is ")
	println(lst)
	r := f(lst.Len())
	e := lst.Front()
	for i := 0; i < r; i++ {
		e = e.Next()
	}
	if e != nil {
		lst.Remove(e)
		return e.Value
	}
	return nil
}

func UniformSel(lst *list.List) interface{} {
	println("Creating a Uniform selector")
	return funcSelector(lst,
		func(s int) int {
			return rand.Int() % s
		})
}

type QueueSelector func() *Queue
type ServerSelector func() *Server

func MakeQueueSelector(sel Selector, l *list.List) QueueSelector {
	println("Converting generic to Queue selector")
	return func() *Queue {
		return sel(l).(*Queue)
	}
}

func MakeServerSelector(sel Selector, l *list.List) ServerSelector {
	println("Converting generic to Server selector")
	return func() *Server {
		return sel(l).(*Server)
	}
}
