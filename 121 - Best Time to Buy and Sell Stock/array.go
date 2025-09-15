func maxProfit(prices []int) int {
	max := 0
	min := prices[0]

	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if (prices[i] - min) > max {
			max = (prices[i] - min)
		}
	}

	return max
}