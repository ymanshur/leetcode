/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev, curr, next *ListNode
    prev, curr = nil, head
    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev, curr = curr, next
    }

    return prev
}