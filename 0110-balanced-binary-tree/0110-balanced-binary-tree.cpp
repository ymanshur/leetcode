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
        if (!root) {
            return true;
        }

        int height = 0;
        return isNodeBalanced(root, height);
    }

private:
    bool isNodeBalanced(TreeNode* root, int& height) {
        if (!root) {
            return true;
        }

        height++;

        int leftH = 0, rightH = 0;
        if (!isNodeBalanced(root->left, leftH)) {
            return false;
        }

        if (!isNodeBalanced(root->right, rightH)) {
            return false;
        }

        if (leftH > rightH) {
            height += leftH;
        } else {
            height += rightH;
        }

        //cout << root->val << "->" << abs(leftH - rightH) << " ";

        return abs(leftH - rightH) <= 1;
    }
};