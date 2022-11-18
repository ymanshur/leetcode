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
    var prev *ListNode // point to new head
    var curr = head // point to head
    for curr != nil {
        // saving the next point
        next := curr.Next
        // changing it previous
        curr.Next = prev
        // setting previous to current
        prev = curr
        // setting next to current
        curr = next
 
        // list := prev
        // for list != nil {
        //     fmt.Printf("%+v->", list.Val)
        //     list = list.Next
        // }
        // fmt.Println()
    }
    return prev
}