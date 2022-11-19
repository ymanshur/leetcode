/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    // reverse the children
    root.Left, root.Right = root.Right, root.Left
    // invert the left sub tree
    invertTree(root.Left)
    // invert the right sub tree
    invertTree(root.Right)

    return root
}