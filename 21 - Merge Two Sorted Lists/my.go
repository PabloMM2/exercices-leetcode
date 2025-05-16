func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var current *ListNode
	var head *ListNode
	c1 := list1
	c2 := list2

	if c1 == nil {
		return c2
	} else if c2 == nil {
		return c1
	}

	for c1 != nil && c2 != nil {
		if c1.Val <= c2.Val {
			if current == nil {
				current = c1
				head = current
			} else {
				current.Next = c1
				current = current.Next
			}
			c1 = c1.Next
		} else if c2.Val < c1.Val {
			if current == nil {
				current = c2
				head = current
			} else {
				current.Next = c2
				current = current.Next
			}
			c2 = c2.Next
		}
	}

	if c1 == nil {
		current.Next = c2
	} else if c2 == nil {
		current.Next = c1
	}

	return head
}