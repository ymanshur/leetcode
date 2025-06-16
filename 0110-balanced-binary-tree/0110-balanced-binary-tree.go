/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    left := depth(root.Left)
    right := depth(root.Right)

    //fmt.Printf("%d->%d ", root.Val, abs(left - right))
    return abs(left - right) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := depth(root.Left)
    right := depth(root.Right)

    //fmt.Printf("%d->%d ", root.Val, max(left, right) + 1)
    return max(left, right) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}