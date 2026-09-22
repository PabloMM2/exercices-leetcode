func hasCycle(head *ListNode) bool {
	var nodes map[*ListNode]bool = make(map[*ListNode]bool)
	for head != nil {
		_, ok := nodes[head]
		if ok {
			return true
		}
		nodes[head] = true
		head = head.Next
	}
	return false
}