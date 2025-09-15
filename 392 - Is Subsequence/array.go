func isSubsequence(s string, t string) bool {

	counter := 0
	last := 0

	for _, val := range s {
		for i := last; i < len(t); i++ {
			last++
			if string(val) == string(t[i]) {
				counter++
				break
			}
		}
	}

	return counter == len(s)
}