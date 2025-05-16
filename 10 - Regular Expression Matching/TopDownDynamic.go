func isMatch(s string, p string) bool {

	memo := make(map[[2]int]bool)

	var dp func(i, j int) bool

	dp = func(i, j int) bool {
		key := [2]int{i, j}
		if val, exists := memo[key]; exists {
			return val
		}

		var ans bool
		if j == len(p) {
			ans = i == len(s)
		} else {

			firstMatch := len(s) > i && (string(s[i]) == string(p[j]) || string(p[j]) == ".")

			if len(p) > j+1 && string(p[j+1]) == "*" {
				ans = dp(i, j+2) || (firstMatch && dp(i+1, j))
			} else {
				ans = firstMatch && dp(i+1, j+1)
			}

		}

		memo[key] = ans
		return ans
	}

	return dp(0, 0)
}