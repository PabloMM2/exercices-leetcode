func search(nums []int, target int) int {
    l, r := 0, len(nums) -1  
    for l <= r {
        m := int ((l + r) / 2)
        if target > nums[m] {
            l = m + 1
        }else if target < nums[m] {
            r = m -1 
        } else if nums[m] == target {
            return m
        }
    }
    return -1
}