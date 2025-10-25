class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        unique = set(nums)
        return len(unique) != len(nums)