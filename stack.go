package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Empty slice")
	}
	last := (s.items)[len(s.items)-1]
	s.items = (s.items)[0 : len(s.items)-1]
	return last
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	isEmpty := stack.IsEmpty()
	fmt.Println("Is empty:", isEmpty)
	fmt.Println(stack.items)
	last := stack.Pop()
	fmt.Println("Removed element:", last)
	fmt.Println(stack.items)
	size := stack.Size()
	fmt.Println("Stack size:", size)
}
