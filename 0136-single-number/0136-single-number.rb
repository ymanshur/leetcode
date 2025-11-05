# @param {Integer[]} nums
# @return {Integer}
def single_number(nums)
    res = nums[0]
    nums[1..-1].each do |num|
        res ^= num
    end
    res
end