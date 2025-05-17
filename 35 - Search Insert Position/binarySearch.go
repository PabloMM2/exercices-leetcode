func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums) -1
    response :=0

    for l <= r {
        m := int((l+r)/2)
        if target > nums[m] {
            l = m +1
        } else if target < nums[m] {
            r = m -1
        } else if nums[m] == target {
            return m
        }
    }

    if l > r {
        return l
    }

    return response
}