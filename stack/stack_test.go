package stack_test

import (
	"fmt"

	"github.com/estevaofon/go-algorithms/stack"
)

func ExampleStack() {
	s := stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	isEmpty := s.IsEmpty()
	fmt.Println("Is empty:", isEmpty)
	// Outputting items directly is not possible as items is private,
	// checking via Pop or Size is preferred for black-box testing.

	last := s.Pop()
	fmt.Println("Removed element:", last)

	size := s.Size()
	fmt.Println("Stack size:", size)

	// Output:
	// Is empty: false
	// Removed element: 3
	// Stack size: 2
}
