func levelOrder(root *TreeNode) [][]int {
    var response [][]int

    if root == nil {
        return response
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        values := []int{}

        for _, node := range queue {
           values = append(values, node.Val)
           if node.Left != nil {
            queue = append(queue, node.Left)
           }
           if node.Right != nil {
            queue = append(queue, node.Right)
           }
           queue = queue[1:]
        }
        response = append(response, values)
    }

    return response
}