func flatten(root *TreeNode)  {
    for root != nil {
        left, right := root.Left, root.Right
        root.Left = nil
        if left != nil {
            root.Right = left
            for left.Right != nil {
                left = left.Right
            }
            left.Right = right
        }
        root = root.Right
    }
}
