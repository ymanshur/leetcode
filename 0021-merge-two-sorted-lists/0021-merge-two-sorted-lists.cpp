/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode* curr1 = list1;
        ListNode* curr2 = list2;
        ListNode* ans = nullptr;

        while (curr1 && curr2) {
            int val;
            if (curr1->val < curr2->val) {
                val = curr1->val;
                curr1 = curr1->next;
            } else {
                val = curr2->val;
                curr2 = curr2->next;
            }

            ListNode* x = new ListNode(val);
            x->next = ans;
            ans = x;
        }

        while (curr1) {
            ListNode* x = new ListNode(curr1->val);
            x->next = ans;
            ans = x;
            curr1 = curr1->next;
        }

        while (curr2) {
            ListNode* x = new ListNode(curr2->val);
            x->next = ans;
            ans = x;
            curr2 = curr2->next;
        }

        ListNode* ans2 = nullptr;
        while (ans) {
            ListNode* x = new ListNode(ans->val);
            x->next = ans2;
            ans2 = x;
            ans = ans->next;
        }
        
        return ans2;
    }
};