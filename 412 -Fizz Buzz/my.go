func fizzBuzz(n int) []string {
	response := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			response = append(response, "FizzBuzz")
		} else if i%5 == 0 {
			response = append(response, "Buzz")
		} else if i%3 == 0 {
			response = append(response, "Fizz")
		} else {
			response = append(response, fmt.Sprintf("%d", i))
		}
	}
	return response
}