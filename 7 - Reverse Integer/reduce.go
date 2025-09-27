func reverse(x int) int {
	num := x
	aux := ""
	limit := int(math.Pow(2, 31))
	//limit := 1<<31
	if x < 0 {
		num *= -1
	}

	for num > 0 {
		aux += fmt.Sprintf("%d", num%10)
		num = num / 10
	}

	res, _ := strconv.Atoi(aux)
	if res >= limit {
		return 0
	}

	if x < 0 {
		res *= -1
	}

	return res
}