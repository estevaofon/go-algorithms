package stack

// Stack represents a stack of integers
type Stack struct {
	items []int
}

// Pop removes and returns the last element of the stack.
// Panics if the stack is empty.
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Empty slice")
	}
	last := (s.items)[len(s.items)-1]
	s.items = (s.items)[0 : len(s.items)-1]
	return last
}

// Push adds a value to the stack.
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}
