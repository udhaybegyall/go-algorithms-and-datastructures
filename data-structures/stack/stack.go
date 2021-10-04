package stack

type Stack struct {
	values     []int
	valueCount int
}

func (s *Stack) isEmpty() bool {

	return len(s.values) == 0
}

func (s *Stack) push(value int) {

	s.values = append(s.values[:s.valueCount], value)
	s.valueCount += 1
}

func (s *Stack) pop() int {

	if s.valueCount == 0 {
		return 0
	}

	var value = s.values[len(s.values)-1]

	if len(s.values) > 1 {
		s.values = s.values[:len(s.values)-1]
	} else {
		s.values = s.values[0:]
	}
	s.valueCount = len(s.values)

	return value
}
