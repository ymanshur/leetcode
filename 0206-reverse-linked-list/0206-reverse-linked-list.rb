# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val = 0, _next = nil)
#         @val = val
#         @next = _next
#     end
# end
# @param {ListNode} head
# @return {ListNode}
def reverse_list(head)
    return if head.nil?
    return head if head.next.nil?

    stack = []

    curr = head
    while curr
        stack.push(curr)
        curr = curr.next
    end

    head = stack.pop
    curr = head

    while !stack.empty?
        node = stack.pop
        curr.next = node
        curr = node
    end
    curr.next = nil

    head
end