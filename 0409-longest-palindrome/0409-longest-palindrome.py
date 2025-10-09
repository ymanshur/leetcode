class Solution:
    def longestPalindrome(self, s: str) -> int:
        freq = {}
        for c in s:
            if c in freq:
                freq[c] += 1
            else:
                freq[c] = 1
        
        odd = 0
        for cnt in freq.values():
            if cnt % 2 == 1:
                odd += 1
        
        res = len(s)
        if odd > 0:
            res -= (odd - 1)

        return res