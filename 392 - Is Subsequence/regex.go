func isSubsequence(s string, t string) bool {

	regx := ".*"
	for _, val := range s {
		regx += fmt.Sprintf("%s.*", string(val))
	}

	match, _ := regexp.MatchString(regx, t)

	return match
}