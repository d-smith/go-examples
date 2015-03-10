package stack

//Stack represents a stack data structure. Elements can be pushed on to and 
//popped off of the stack
type Stack struct {
	size int
	top  *Element
}

//Element holds an element on the stack
type Element struct {
	value interface{}
	next  *Element
}

//Size returns the number of elements on the stack
func (s *Stack) Size() int {
	return s.size
}

//Push pushes an element on the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value: value, next: s.top}
	s.size++
}

//Peek returns the data associated with the top element in the stack, without
//popping the element off the stack
func (s *Stack) Peek() interface{} {
	if s.size > 0 {
		return s.top.value
	}
	
	return nil
}

//Pop removes the top element from the stack, returning the data
//associated with it.
func (s *Stack) Pop() (value interface{}) {
	if s.size == 0 {
		return nil
	}

	value, s.top = s.top.value, s.top.next
	s.size--
	return
}
