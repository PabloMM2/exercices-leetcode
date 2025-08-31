func isPalindrome(s string) bool {

	rule := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	fixed := rule.ReplaceAllString(s, "")

	var reverse []byte
	for i := len(fixed) - 1; i >= 0; i-- {
		reverse = append(reverse, fixed[i])
	}

	if strings.ToLower(fixed) == strings.ToLower(string(reverse)) {
		return true
	}

	return false
}