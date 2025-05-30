/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    slowP, fastP := head, head // slow pointer, fast pointer are start at same node
    for fastP != nil && fastP.Next != nil {
        slowP = slowP.Next
        fastP = fastP.Next.Next
        if slowP == fastP {
            return true
        }
    }
    return false
}

/**
 * The O(n) time and space complexity is using map or set to store the visited node,
 * and check if the next node was visited, return true, otheriwise, return false (at the end).
 * 
 * But there is O(n) time and O(1) space complexity, I am going to use two pointer with diff pace,
 * so if the linked-list is cyclic, the pointers will meet.
 */