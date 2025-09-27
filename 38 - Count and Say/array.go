func countAndSay(n int) string {
	response := []string{}

	for i := 0; i < n; i++ {
		if len(response) == 0 {
			response = append(response, "1")
			continue
		}

		count := 1
		aux := []string{}
		for j := 0; j < len(response); j++ {
			if j == len(response)-1 || response[j] != response[j+1] {
				aux = append(aux, fmt.Sprintf("%d", count), response[j])
				count = 1
			} else {
				count += 1
			}
		}

		response = aux
	}

	return strings.Join(response, "")
}