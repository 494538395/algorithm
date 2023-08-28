package memory

// 小顶堆，优化内存过期 key 淘汰效率
// 根据过期进行入堆排序
// 越是最近要过期的越会在堆顶

// MinHeap 小顶堆.
type MinHeap []*Item

// Len function.
func (h MinHeap) Len() int { return len(h) }

// Less function.
func (h MinHeap) Less(i, j int) bool { return h[i].Expiration.Before(h[j].Expiration) }

// Swap function.
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push into heap 入队.
func (h *MinHeap) Push(x interface{}) {
	item := x.(*Item)
	*h = append(*h, item)
}

// Pop from heap 出队.
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
