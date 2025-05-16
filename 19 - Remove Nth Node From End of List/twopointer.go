func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{0, head}
	fast := node
	slow := node

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return node.Next
}