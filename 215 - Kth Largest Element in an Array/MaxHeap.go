func findKthLargest(nums []int, k int) int {
    h := &MaxHeap{}
    heap.Init(h)

    for _, num := range nums {
        heap.Push(h, num)
    }

    count := 1
    for count < k {
        heap.Pop(h)
        count ++
    }

    return heap.Pop(h).(int)
}

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}