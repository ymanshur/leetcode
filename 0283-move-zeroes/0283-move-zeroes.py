class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l = -1
        for r in range(len(nums)):
            if l == -1 and nums[r] == 0:
                l = r
            
            if l != -1 and nums[r] != 0:
                nums[l], nums[r] = nums[r], nums[l]
                l += 1