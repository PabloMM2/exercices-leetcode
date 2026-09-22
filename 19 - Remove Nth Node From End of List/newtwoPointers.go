/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * Time - 0,0ms
 * Memo - 4,07MB
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil && n == 1 {
		return nil
	}

	left := &ListNode{Val: 0, Next: head}
	right := left

	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right.Next != nil {
		left = left.Next
		right = right.Next
	}

	if *left.Next == *head {
		return head.Next
	}

	left.Next = left.Next.Next

	return head
}