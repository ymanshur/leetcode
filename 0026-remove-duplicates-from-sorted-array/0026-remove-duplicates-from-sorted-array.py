class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        nums_set = sorted(set(nums))
        n = len(nums_set)
        nums[:n] = nums_set
        return n