var searchInsert = function (nums, target) {
    let l = 0
    let r = nums.length - 1
    while (l <= r) {
        let m = parseInt((l + r) / 2)

        if (nums[m] < target) {
            l = m + 1
        } else if (nums[m] > target) {
            r = m - 1
        } else if (nums[m] == target) {
            return m
        }
    }

    if(l > r) {
        return l
    }
    return 0
};