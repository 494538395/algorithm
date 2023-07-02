package heap

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type MedianFinder struct {
	minH iMinHeap // 小顶堆，保存较大的一半
	maxH iMaxHeap // 大顶堆，保存较小的一半
}

func Constructor() MedianFinder {
	maxH, minH := iMaxHeap{}, iMinHeap{}
	heap.Init(&maxH)
	heap.Init(&minH)
	finder := MedianFinder{
		minH: minH,
		maxH: maxH,
	}
	return finder
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.minH.Len() != mf.maxH.Len() {
		heap.Push(&mf.minH, num)
		heap.Push(&mf.maxH, heap.Pop(&mf.minH))
	} else {
		heap.Push(&mf.maxH, num)
		heap.Push(&mf.minH, heap.Pop(&mf.maxH))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.minH.Len() == mf.maxH.Len() {
		return float64(mf.minH.Peek()+mf.maxH.Peek()) * 0.5
	}
	return float64(mf.minH.Peek())
}

// ---------------- 大顶堆/小顶堆 定义 ---------------- //
// int类型小顶堆，sort.IntSlice 默认升序，即小顶堆
type iMinHeap struct{ sort.IntSlice }

// 实现 container/Heap 接口的 Push,Pop 方法
func (h *iMinHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *iMinHeap) Pop() interface{} {
	n := h.IntSlice.Len()
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

// 查看堆顶元素，不改变结构
func (h iMinHeap) Peek() int {
	return h.IntSlice[0]
}

// int类型大顶堆，大顶堆需要重新 Less() 比较器
type iMaxHeap struct{ sort.IntSlice }

func (h iMaxHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

// 实现 container/Heap 接口的 Push,Pop 方法
func (h *iMaxHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *iMaxHeap) Pop() interface{} {
	n := h.IntSlice.Len()
	x := h.IntSlice[n-1]
	h.IntSlice = h.IntSlice[:n-1]
	return x
}

func (h iMaxHeap) Peek() int {
	return h.IntSlice[0]
}

func TestHeap(t *testing.T) {
	finder := Constructor()
	finder.AddNum(-1)
	fmt.Println(finder.FindMedian())

	finder.AddNum(2)
	fmt.Println(finder.FindMedian())
	finder.AddNum(3)
	fmt.Println(finder.FindMedian())
}
