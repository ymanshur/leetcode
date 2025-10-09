# @param {Integer[]} nums
# @return {Integer}
def majority_element(nums)
    res = nums[0]
    cnt = 1
    nums[1..-1].each do |num|
        if cnt == 0
            res = num
            cnt = 1
            next
        end

        if num == res
            cnt += 1
        else
            cnt -= 1
        end
    end
    
    res
end