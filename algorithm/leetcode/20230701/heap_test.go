package _0230701

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type medianFinder struct {
	queMin, queMax hp
}

func constructor() medianFinder {
	return medianFinder{}
}

func (mf *medianFinder) addNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *medianFinder) findMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func TestHeap(t *testing.T) {

	finder := constructor()

	finder.addNum(1)
	finder.addNum(4)
	finder.addNum(3)

	fmt.Println(finder.findMedian())

}

func TestArr(t *testing.T) {
	var a sort.IntSlice

	a = append(a, 1)
	a = append(a, 9)
	a = append(a, 4)

	fmt.Println(a)

}
