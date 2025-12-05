func isValid(s string) bool {
	var stack Stack

	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case ")":
			last := stack.Top()
			if last == "(" {
				stack.Pop()
			} else {
				stack.Push(string(s[i]))
			}
		case "}":
			last := stack.Top()
			if last == "{" {
				stack.Pop()
			} else {
				stack.Push(string(s[i]))
			}
		case "]":
			last := stack.Top()
			if last == "[" {
				stack.Pop()
			} else {
				stack.Push(string(s[i]))
			}
		default:
			stack.Push(string(s[i]))
		}
	}

	return stack.IsEmpty()

}

type Stack struct {
	values []string
}

func (s *Stack) Push(st string) {
	s.values = append(s.values, st)
}

func (s *Stack) Pop() string {
	last := s.values[s.Size()-1]
	s.values = s.values[:s.Size()-1]
	return last
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	return s.values[s.Size()-1]
}

func (s *Stack) Size() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}
