func isSymmetric(root *TreeNode) bool {
	left := root.Left
	right := root.Right

	return Symmetric(left, right)
}

func Symmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l != nil && r == nil || l == nil && r != nil {
		return false
	}

	lr := Symmetric(l.Left, r.Right)
	rl := Symmetric(r.Left, l.Right)

	return (l.Val == r.Val) && lr && rl
}