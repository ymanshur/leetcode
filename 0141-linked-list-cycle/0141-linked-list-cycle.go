/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    nodeMap := make(map[*ListNode]bool)
    for head != nil {
        if nodeMap[head] {
            return true
        } else {
            nodeMap[head] = true
        }
        head = head.Next
    }
    return false
}