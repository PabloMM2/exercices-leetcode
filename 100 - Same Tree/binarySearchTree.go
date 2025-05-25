func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q == nil {
		return true
	}

	response := p.Val == q.Val

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return response && left && right
}