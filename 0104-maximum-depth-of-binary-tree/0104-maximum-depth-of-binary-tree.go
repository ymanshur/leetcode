/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftMaxD, rightMaxD := maxDepth(root.Left), maxDepth(root.Right)
    if leftMaxD > rightMaxD {
        return 1 + leftMaxD
    }
    return 1 + rightMaxD
}