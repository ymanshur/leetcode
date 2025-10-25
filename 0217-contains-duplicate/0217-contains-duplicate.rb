# @param {Integer[]} nums
# @return {Boolean}
def contains_duplicate(nums)
    unique = nums.to_set
    unique.length != nums.length
end