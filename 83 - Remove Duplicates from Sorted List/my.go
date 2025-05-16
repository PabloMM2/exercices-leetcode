func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	current := head
	next := current.Next

	for next != nil {
		if current.Val == next.Val {
			current.Next = next.Next
			next = current.Next
		} else {
			current = next
			next = current.Next
		}
	}
	return head
}