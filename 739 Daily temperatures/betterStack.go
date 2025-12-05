func dailyTemperatures(temperatures []int) []int {
	response := make([]int, len(temperatures))
	stack := Stack{}

	for key, val := range temperatures {

		for !stack.IsEmpty() && val > temperatures[stack.Top()] {
			response[stack.Top()] = key - stack.Top()
			stack.Pop()
		}

		stack.Push(key)
	}

	return response
}

type Stack struct {
	Values []int
}

func (s *Stack) Push(value int) {
	s.Values = append(s.Values, value)
}

func (s *Stack) Pop() {
	s.Values = s.Values[:s.Size()-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Values) == 0
}

func (s *Stack) Size() int {
	return len(s.Values)
}

func (s *Stack) Top() int {
	return s.Values[s.Size()-1]
}