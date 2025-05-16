func hasCycle(head *ListNode) bool {
	var values = make(map[*ListNode]bool)
	current := head
	for current != nil {
		if _, ok := values[current]; ok {
			return true
		}

		values[current] = true
		current = current.Next
	}

	return false
}