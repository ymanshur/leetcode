class Solution {
public:
    int numSubarrayProductLessThanK(vector<int>& nums, int k) {
        if (k <= 1) {
            return 0;
        }

        int curr = 1;
        int left = 0;
        int ans = 0;

        for (int right = 0; right < nums.size(); right++) {
            curr *= nums[right];

            while (curr >= k) {
                curr /= nums[left];
                left++;
            }

            ans += (right - left + 1);
        }

        return ans;
    }
};