import "container/heap"
func lastStoneWeight(stones []int) int {
    h := &MaxHeap{}
    heap.Init(h)

    for _, stone := range stones {
        heap.Push(h, stone)
    }

    for h.Len() > 1 {
        s1 := heap.Pop(h).(int)
        s2 := heap.Pop(h).(int)
        if s1 != s2 {
            heap.Push(h, s1 - s2)
        }
    }

    if len(*h) == 1 {
        return heap.Pop(h).(int)
    }
    return 0
}

type MaxHeap []int

func (h MaxHeap) Len() int {return len(h) }
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