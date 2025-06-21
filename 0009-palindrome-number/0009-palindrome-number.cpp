class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0 || (x > 0 && x % 10 == 0)) {
            return false;
        }

        int reserved = 0;
        while (reserved < x) {
            reserved = reserved * 10 + x % 10;
            x /= 10 ;
        }

        return (x == reserved) || (x == reserved / 10);
    }
};