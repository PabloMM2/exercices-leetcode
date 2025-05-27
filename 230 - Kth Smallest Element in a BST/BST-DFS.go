func kthSmallest(root *TreeNode, k int) int {
  response := 0
  count := 0
  var Smallest func(root *TreeNode)

  Smallest = func(root *TreeNode) {
    if root == nil || count >= k {
      return
    }

    Smallest(root.Left)
    count++
    if count == k {
      response = root.Val
    }
    
    Smallest(root.Right)
    return 
  }

  Smallest(root)
  return response
}