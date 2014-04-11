package simulator

import (
	"container/list"
	"fmt"
)

func all(lst *list.List, test func(interface{}) bool) bool {
	for e := lst.Front(); e != nil; e = e.Next() {
		if ! test(e.Value) {
			return false
		}
	}
	return true
}

// PrintTable takes a list of lists and "Pretty Prints" it in table form.
// The length of the first list is assumed to be the length of the rest.
func PrintTable(lst *list.List) {
	numFields := lst.Front().Value.(*list.List).Len()
	fieldLengths := make([]int,numFields)
	for i := 0; i < numFields; i++ {
		for e := lst.Front(); e != nil; e = e.Next() {
			ent := e.Value.(*list.List)
			j := 0
			var f *list.Element
			for f = ent.Front(); f != nil; f = f.Next() {
				if i == j {
					break
				} else {
					j++
				}
			}
			if i == j {
				thisLen := len(f.Value.(string))
				if fieldLengths[i] < thisLen {
					fieldLengths[i] = thisLen
				}
			}
		}
	}
	PrintSep(fieldLengths)
	for e := lst.Front(); e != nil; e = e.Next() {
		PrintFields(fieldLengths,e.Value.(*list.List))
		PrintSep(fieldLengths)
	}
}

func PrintSep(fieldLengths []int) {
	fmt.Print(" ")
	for _,v := range fieldLengths {
		fmt.Print("+")
		for i := 0; i < v+2; i++ {
			fmt.Print("-")
		}
	}
	fmt.Print("+\n")
}

func PrintFields(fieldLengths []int, lst *list.List) {
	fmt.Print(" ")
	e := lst.Front()
	for _,v := range fieldLengths {
		fmt.Print("| ")
		var padding int
		var str string
		if e == nil {
			padding = v
			str = ""
		} else {
			str = e.Value.(string)
			padding = v - len(str)
			e = e.Next()
		}

		fmt.Print(str)
		for i := 0; i < padding+1; i++ {
			fmt.Print(" ")
		}
	}
	fmt.Print("|\n")
}

func TableEntry(label string, val interface{}) *list.List {
	l := list.New()
	l.PushBack(label)
	l.PushBack(fmt.Sprintf("%v",val))
	return l
}

func funcSum(l *list.List, f func(interface{}) int) int {
	var sum int
	for e := l.Front(); e != nil; e = e.Next() {
		sum += f(e.Value)
	}
	return sum
}


