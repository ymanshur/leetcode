/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    // return nil node if the input is nil
    if head == nil {
        return head
    }
    stack := make([]int, 0)
    for head != nil {
        stack = append(stack, head.Val)
        head = head.Next
    }
    topIndex := len(stack)-1
    // init newHead with top off stack
    newHead := &ListNode{Val:stack[topIndex]}
    for i := topIndex-1; i >= 0; i-- {
        newNode := &ListNode{
            Val: stack[i],
        }
        // get tail pointer, then link to newNode
        tail := newHead
        for tail.Next != nil {
            tail = tail.Next
        }
        tail.Next = newNode
    }
    return newHead
}

/**
 * 1. Store the s-linked link into stack
 * 2. Create new s-linked link from the stack
 */
 