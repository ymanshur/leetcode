/**
 Do not return anything, modify nums in-place instead.
 */
function moveZeroes(nums: number[]): void {
    let l = -1;
    for (let r = 0; r < nums.length; r ++) {
        if (l == -1 && nums[r] == 0) {
            l = r;
        }

        if (l != -1 && nums[r] != 0) {
            [nums[l], nums[r]] = [nums[r], nums[l]];
            l++;
        }
    }
};