func getMinimumDifference(root *TreeNode) int {
    var prev *int
    var GetValues func(root *TreeNode)
    min := math.MaxInt
    
    GetValues =  func(root *TreeNode) {
        if root == nil {
            return
        }

        GetValues(root.Left)
        if prev != nil {
            min = int(math.Min(float64(min), float64(root.Val) - float64(*prev)))
        }
        prev = &root.Val
        GetValues(root.Right)
        return
    }

    GetValues(root)
    return min
}