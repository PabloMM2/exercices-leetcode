func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	current := head
	next := current

	for next != nil {
		next = current.Next
		current.Next = prev
		prev = current
		if next != nil {
			current = next
		}
	}

	return current
}