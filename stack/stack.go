package stack

type Stack struct {
	size int
	top  *StackElement
}

type StackElement struct {
	value interface{}
	next  *StackElement
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &StackElement{value: value, next: s.top}
	s.size++
}

func (s *Stack) Peek() interface{} {
	if s.size > 0 {
		return s.top.value
	} else {
		return nil
	}
}

func (s *Stack) Pop() (value interface{}) {
	if s.size == 0 {
		return nil
	}

	value, s.top = s.top.value, s.top.next
	s.size--
	return
}
