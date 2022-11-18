/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head = &ListNode{}
    var tail = head
    for list1 != nil && list2 != nil {
        if list1.Val > list2.Val {
            tail.Next = list2
            list2 = list2.Next
        } else {
            tail.Next = list1
            list1 = list1.Next
        }
        tail = tail.Next
    }
    
    if list1 != nil {
        tail.Next = list1
    } else if list2 != nil {
        tail.Next = list2
    }
    
    return head.Next
}