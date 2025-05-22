func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0
    var Diameter func (root *TreeNode) int

    Diameter = func (root *TreeNode) int {
        if root == nil {
            return 0
        }

        left := Diameter(root.Left)
        right := Diameter(root.Right)

        diameter = int( math.Max( float64(diameter), float64(left+right) ) )

        return 1 + int(math.Max(float64(left), float64(right)))
    }

    Diameter(root)
    return diameter
}

