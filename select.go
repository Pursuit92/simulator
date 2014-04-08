package simulator

import (
	"container/list"
	"math/rand"
)

type Selector func(*list.List) interface{}

func compFirst(lst *list.List, comp func(interface{}, interface{}) bool) interface{} {
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
	return compFirst(lst,
		func(v, curr interface{}) bool {
			return v.(Queue).Size() < curr.(Queue).Size()
		})
}

func LongestFirst(lst *list.List) interface{} {
	return compFirst(lst,
		func(v, curr interface{}) bool {
			return v.(Queue).Size() > curr.(Queue).Size()
		})
}

func Linear(lst *list.List) interface{} {
	e := lst.Front()
	if e != nil {
		lst.Remove(e)
		return e.Value
	}
	return nil
}

func funcSelector(lst *list.List, f func(int) int) interface{} {
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

func Uniform(lst *list.List) interface{} {
	return funcSelector(lst,
		func(s int) int {
			return rand.Int() % s
		})
}

