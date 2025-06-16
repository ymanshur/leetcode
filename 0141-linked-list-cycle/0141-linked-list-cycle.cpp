/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode *curr1 = head, *curr2 = head;

        while (curr1 && curr2 && curr2->next) {
            if (curr1 == curr2->next) {
                return true;
            }

            curr1 = curr1->next;
            curr2 = curr2->next->next;
        }

        return false;
    }
};