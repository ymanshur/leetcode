/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var newHead *ListNode
    var helper func(*ListNode, *ListNode)
    helper = func(prev, curr *ListNode) {
        if curr == nil {
            newHead = prev
            return
        }
        
        helper(curr, curr.Next)
        curr.Next = prev
    }

    helper(nil, head)
    return newHead
}