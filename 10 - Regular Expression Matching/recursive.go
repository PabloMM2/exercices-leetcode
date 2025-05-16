import "fmt"

func main() {
	isMatch()
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := len(s) > 0 && (string(s[0]) == string(p[0]) || string(p[0]) == ".")

	if len(p) >= 2 && string(p[1]) == "*" {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}

	return firstMatch && isMatch(s[1:], p[1:])
}