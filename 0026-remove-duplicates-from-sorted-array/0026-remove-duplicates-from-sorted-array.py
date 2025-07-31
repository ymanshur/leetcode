class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        # nums_unique = sorted(set(nums))
        #
        # nums_unique = []
        # curr = None
        # for i in range(0, len(nums)):
        #     if curr != nums[i]:
        #         curr = nums[i]
        #         nums_unique.append(curr)
        #
        # n = len(nums_unique)
        # nums[:n] = nums_unique
        # return n

        # 1, 1, 2

        l = 0
        for r in range(0, len(nums)):
            if nums[r] != nums[l]:
                l += 1
                nums[l], nums[r] = nums[r], nums[l]

        return l + 1