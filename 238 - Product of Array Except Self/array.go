func productExceptSelf(nums []int) []int {

	response := []int{}
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	result := 1

	for i := 0; i < len(nums); i++ {
		result *= nums[i]
		left[i] = result
	}

	result = 1

	for j := len(nums) - 1; j >= 0; j-- {
		result *= nums[j]
		right[j] = result
	}

	for j := 0; j < len(nums); j++ {
		res := 0
		if j == 0 {
			res = right[j+1]
		} else if j == len(nums)-1 {
			res = left[j-1]
		} else {
			res = left[j-1] * right[j+1]
		}
		response = append(response, res)
	}

	return response
}