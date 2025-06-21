/*
0  -> 0000 -> 0
2  -> 0010 -> 1
4  -> 0100 -> 1
6  -> 0110 -> 2
8  -> 1000 -> 1
13 -> 1101 -> 3
*/

class Solution {
public:
    vector<int> countBits(int n) {
        vector<int> ans(n + 1);
        for (int i = 0; i <= n; i++) {
            ans[i] = ans[i / 2] + i % 2;
        }

        return ans;
    }
};