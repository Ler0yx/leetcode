package nAryTreePreorderTraversal

func preorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    var ans = make([]int, 0)
    ans = append(ans, root.Val)    
    for i:=0; i < len(root.Children); i ++ {
        ans = append(ans, preorder(root.Children[i])...)
    }
    return ans
}