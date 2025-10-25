# @param {Integer} n
# @return {Integer}
def hamming_weight(n)
    cnt = 0
    while n > 0 do
        remainder = n % 2
        if remainder == 1
            cnt += 1
        end
        n = n >> 1
    end
    cnt
end