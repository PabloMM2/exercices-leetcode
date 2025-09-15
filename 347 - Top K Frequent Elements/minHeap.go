import "container/heap"

func topKFrequent(nums []int, k int) []int {
	var values = make(map[int]int)

	for _, num := range nums {
		values[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for key, val := range values {
		node := Node{key, val}
		heap.Push(h, node)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	response := []int{}

	for h.Len() > 0 {
		node := heap.Pop(h).(Node)
		response = append(response, node.num)
	}

	return response
}

type Node struct {
	num       int
	frecuency int
}

type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frecuency < h[j].frecuency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}