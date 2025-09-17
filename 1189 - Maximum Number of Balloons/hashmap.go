func maxNumberOfBalloons(text string) int {
	mt := map[string]int{}

	for _, val := range text {
		mt[string(val)] += 1
	}

	response := 0
	for {
		if mt["b"] >= 1 && mt["a"] >= 1 && mt["l"] >= 2 && mt["o"] >= 2 && mt["n"] >= 1 {
			response++
			mt["b"] -= 1
			mt["a"] -= 1
			mt["l"] -= 2
			mt["o"] -= 2
			mt["n"] -= 1
		} else {
			break
		}
	}

	return response
}