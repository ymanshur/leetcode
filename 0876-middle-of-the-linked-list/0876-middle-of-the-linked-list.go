/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func lenNode(node ListNode) int {
    var length int
    if node.Next == nil {
        return length
    }
    length++
    for node.Next != nil {
        node = *node.Next
        length++
    }
    return length    
}

func middleNode(head *ListNode) *ListNode {
    var nodeLength int = lenNode(*head)
    
    for i := 0; i < nodeLength / 2; i++ {
        head = head.Next
    }
    
    return head
}