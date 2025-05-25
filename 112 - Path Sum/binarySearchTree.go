func hasPathSum(root *TreeNode, targetSum int) bool {

	return Sum(root, targetSum, 0)
}

func Sum(root *TreeNode, targetSum, sum int) bool {
	if root == nil {
		return false
	}

	sum += root.Val

	if root.Left == nil && root.Right == nil {
		return sum == targetSum
	}

	return Sum(root.Left, targetSum, sum) || Sum(root.Right, targetSum, sum)
}