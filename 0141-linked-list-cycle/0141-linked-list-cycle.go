/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    nodeMap := make(map[*ListNode]int)
    for head != nil {
        if _, ok := nodeMap[head]; !ok {
            nodeMap[head] = head.Val
        } else {
            return true
        }
        head = head.Next
    }
    return false
}