# @param {Integer[]} nums
# @return {Void} Do not return anything, modify nums in-place instead.
def move_zeroes(nums)
    # [2, 1, 0, 0, 12, 0, 9]
    #        l      r
    # [2, 1, 12, 0, 0, 0, 9]
    #            l        r
    l = -1
    (0..nums.length-1).each do |r|
        if l == -1 && nums[r] == 0
            l = r
        end

        if l != -1 && nums[r] != 0
            nums[l], nums[r] = nums[r], nums[l]
            l += 1
        end
    end
    nums
end