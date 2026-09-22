/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	copyMap := make(map[*Node]*Node)
	cursor := head

	for cursor != nil {
		copyMap[cursor] = &Node{Val: cursor.Val}
		cursor = cursor.Next
	}

	for key, val := range copyMap {
		val.Next = copyMap[key.Next]
		val.Random = copyMap[key.Random]
	}

	return copyMap[head]
}