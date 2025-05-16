func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current := &ListNode{}
	head := current
	c1 := list1
	c2 := list2

	for c1 != nil && c2 != nil {
		if c1.Val <= c2.Val {
			current.Next = c1
			c1 = c1.Next
		} else {
			current.Next = c2
			c2 = c2.Next
		}
		current = current.Next
	}

	if c1 != nil {
		current.Next = c1
	} else if c2 != nil {
		current.Next = c2
	}

	return head.Next
}