func calPoints(operations []string) int {

	var stack Stack
	response := 0

	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			last := stack.Pop()
			peak := stack.Peak()
			stack.Push(last)
			stack.Push(last + peak)
		case "D":
			peak := stack.Peak()
			stack.Push(peak * 2)
		case "C":
			stack.Pop()
		default:
			val, _ := strconv.Atoi(operations[i])
			stack.Push(val)
		}
	}

	for !stack.IsEmpty() {
		response += stack.Pop()
	}

	return response
}

type Stack struct {
	items []int
}

func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) Pop() int {
	last := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return last
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peak() int {
	return s.items[s.Size()-1]
}