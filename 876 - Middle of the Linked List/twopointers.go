package twopointers

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var start *ListNode = &ListNode{0, head}
	slow := start
	fast := start

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			fast = fast.Next
		} else {
			fast = fast.Next.Next
		}
	}

	return slow
}
