package heap2

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type myHp struct {
	sort.IntSlice
}

func (h *myHp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *myHp) Pop() interface{} {
	a := h.IntSlice
	last := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return last
}

type Jerry struct {
	h myHp // 存放小于中位数的
}

func Constructor() Jerry {
	return Jerry{}
}

func TestHeap(t *testing.T) {

	jerry := Constructor()
	jerry.add(2)
	jerry.add(4)
	jerry.add(1)

	fmt.Println(jerry.h.IntSlice)
}
func (j *Jerry) add(num int) {
	heap.Push(&j.h, num)
}
