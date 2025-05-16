func copyRandomList(head *Node) *Node {
	current := head
	values := make(map[*Node]*Node)

	for current != nil {
		n := &Node{current.Val, nil, nil}
		values[current] = n
		current = current.Next
	}

	current = head
	for current != nil {
		newNode := values[current]
		newNode.Random = values[current.Random]
		newNode.Next = values[current.Next]
		current = current.Next
	}

	return values[head]
}