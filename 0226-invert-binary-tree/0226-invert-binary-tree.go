/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseChild(root *TreeNode) {
    if root == nil {
        return
    }
    root.Left, root.Right = root.Right, root.Left
    reverseChild(root.Left)
    reverseChild(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
    reverseChild(root)
    return root
}