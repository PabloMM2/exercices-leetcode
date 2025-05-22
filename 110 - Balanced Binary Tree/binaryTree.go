func isBalanced(root *TreeNode) bool {
    balance, _ := Height(root)
    return balance
}

func Height(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }

    lb, left := Height(root.Left)
    rb, right := Height(root.Right)

    balanced := lb && rb && math.Abs(float64(left-right)) <= 1 

    return balanced, (1 + int(math.Max(float64(left), float64(right))))
}