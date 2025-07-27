class Solution:
    def search(self, nums: List[int], target: int) -> int:
        low, high = 0, len(nums) - 1

        while low < high:
            mid = low + ((high - low) >> 1)

            if nums[low] <= nums[mid]:
                if nums[low] <= target <= nums[mid]: # pivot part
                    high = mid
                else:
                    low = mid + 1
            else:
                if nums[mid] < target <= nums[high]:
                    low = mid + 1
                else:
                    high = mid
        
        if nums[low] == target:
            return low
        
        return -1