package simulator

import (
	"container/list"
)

func all(lst *list.List, test func(interface{}) bool) bool {
	for e := lst.Front(); e != nil; e = e.Next() {
		if ! test(e.Value) {
			return false
		}
	}
	return true
}
