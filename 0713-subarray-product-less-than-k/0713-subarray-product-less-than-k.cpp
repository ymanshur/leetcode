class Solution {
public:
    int numSubarrayProductLessThanK(vector<int>& nums, int k) {
        int ans = 0;

        if (k <= 1) {
            return ans;
        }

        for (int i = 0; i < nums.size(); i++) {
            int curr = 1;
            for (int j = i; j >= 0; j--) {
                curr *= nums[j];

                if (curr >= k) {
                    break;
                }

                ans++;
            }
        }

        return ans;
    }
};