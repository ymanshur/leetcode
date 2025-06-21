class Solution {
public:
    int climbStairs(int n) {
        if (n < 4) {
            return n;
        }

        int prev = 2, curr = 3;
        int tmp;
        for (int i = 4; i <= n; i++) {
            tmp = curr;
            curr = curr + prev;
            prev = tmp;
        }

        return curr;
    }
};