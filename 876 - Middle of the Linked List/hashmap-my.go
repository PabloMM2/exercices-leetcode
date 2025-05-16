func middleNode(head *ListNode) *ListNode {
	var values = make(map[int]*ListNode)
	count := 0
	current := head

	for current != nil {
		values[count] = current
		count++
		current = current.Next
	}

	middle := (count / 2)

	return values[middle]
}