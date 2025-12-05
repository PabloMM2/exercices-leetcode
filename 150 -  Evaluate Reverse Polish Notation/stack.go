func evalRPN(tokens []string) int {
	s := Stack{}

	for _, token := range tokens {

		switch token {
		case "+":
			fir := s.Pop()
			sec := s.Pop()
			val := sec + fir
			s.Push(val)
		case "-":
			fir := s.Pop()
			sec := s.Pop()
			val := sec - fir
			s.Push(val)
		case "*":
			fir := s.Pop()
			sec := s.Pop()
			val := sec * fir
			s.Push(val)
		case "/":
			fir := s.Pop()
			sec := s.Pop()
			val := sec / fir
			s.Push(val)
		default:
			val, _ := strconv.Atoi(token)
			s.Push(val)
		}
	}

	return s.values[0]
}

type Stack struct {
	values []int
}

func (s *Stack) Push(x int) {
	s.values = append(s.values, x)
}

func (s *Stack) Pop() int {
	last := s.values[s.Size()-1]
	s.values = s.values[:s.Size()-1]
	return last
}

func (s *Stack) Top() int {
	return s.values[s.Size()-1]
}

func (s *Stack) Size() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}