/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev, curr *ListNode = nil, head
    for curr != nil {
        // saving the next point
        // next := curr.Next
        // changing it previous
        // curr.Next = prev
        // setting previous to current
        // prev = curr
        // setting next to current
        // curr = next
        
        prev, curr, curr.Next = curr, curr.Next, prev

        // list := prev
        // for list != nil {
        //     fmt.Printf("%+v->", list.Val)
        //     list = list.Next
        // }
        // fmt.Println()
    }
    return prev
}