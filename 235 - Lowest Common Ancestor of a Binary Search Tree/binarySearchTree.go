func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var response *TreeNode

    var GetValues func(root*TreeNode)
    GetValues = func(root *TreeNode) {
        if root == nil {
            return
        }

        response = root
        if p.Val == root.Val || q.Val == root.Val  {
            return 
        } else if root.Val < p.Val && root.Val < q.Val {
            GetValues(root.Right)
        } else if root.Val > p.Val && root.Val > q.Val {
            GetValues(root.Left)
        } else {
            return
        }
        
        return
    }

    GetValues(root)
    return response
}