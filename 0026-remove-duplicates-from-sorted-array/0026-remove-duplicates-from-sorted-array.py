class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        nums_set = set()
        for i in range(0, len(nums)):
            nums_set.add(nums[i])
        
        n = len(nums_set)
        nums[:n] = sorted([num for num in nums_set])
        return n