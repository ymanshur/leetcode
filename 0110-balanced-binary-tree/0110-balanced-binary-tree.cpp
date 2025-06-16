/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    bool isBalanced(TreeNode* root) {
        return nodeHeight(root) != -1;
    }

private:
    int nodeHeight(TreeNode* root) {
        if (!root) {
            return 0;
        }
        
        int leftHeight = nodeHeight(root->left);
        if (leftHeight == -1) {
            return -1;
        }

        int rightHeight = nodeHeight(root->right);
        if (rightHeight == -1) {
            return -1;
        }

        //cout << root->val << "->" << abs(leftHeight - rightHeight) << " ";

        if (abs(leftHeight - rightHeight) > 1) {
            return -1;
        }

        return max(leftHeight, rightHeight) + 1;
    }
};