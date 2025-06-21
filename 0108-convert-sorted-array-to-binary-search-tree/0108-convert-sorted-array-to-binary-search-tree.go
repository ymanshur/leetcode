/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    root := BST(nums, 0, len(nums) - 1)
    //print(root)
    return root
}

func BST(nums []int, start, end int) *TreeNode {
    if start > end {
        return nil
    }

    mid := (end - start) / 2 + start

    //fmt.Printf("%d %d %d\n", start, mid, end)

    return &TreeNode{
        Val: nums[mid],
        Left: BST(nums, start, mid - 1),
        Right: BST(nums, mid + 1, end),
    }
}

func print(root *TreeNode) {
    if root == nil {
        fmt.Printf("null ")
        return
    }

    //fmt.Printf("%d ", root.Val)

    print(root.Left)
    print(root.Right)
}