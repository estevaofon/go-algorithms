package main

import (
	"fmt"

	"github.com/estevaofon/go-algorithms/stack"
)

func main() {
	// Initialize a new stack
	// Note: Since Stack struct has private fields, we simply instantiate it.
	// If you wanted a constructor like NewStack(), you'd add that to the package.
	myStack := stack.Stack{}

	fmt.Println("Adding items to the stack...")
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)

	fmt.Printf("Current size: %d\n", myStack.Size())
	fmt.Printf("Is empty? %t\n", myStack.IsEmpty())

	fmt.Println("\nPopping items:")
	for !myStack.IsEmpty() {
		val := myStack.Pop()
		fmt.Printf("Popped: %d\n", val)
	}

	fmt.Printf("\nFinal size: %d\n", myStack.Size())
}
