package main

import "fmt"

type Stack struct {
    numeros []int
}

func (s *Stack) Pop() int {
    if len(s.numeros) == 0 {
        panic("Empty slice")
    }
    ultimo := (s.numeros)[len(s.numeros)-1]
    s.numeros = (s.numeros)[0:len(s.numeros)-1]
    return ultimo
}

func (s *Stack) Push(valor int) {
    s.numeros = append(s.numeros, valor)
}

func (s *Stack) IsEmpty() bool {
    return  len(s.numeros) == 0
}


func (s *Stack) Size() int {
    return len(s.numeros)
}

func main() {
    stack := Stack{}
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    isEmpty := stack.IsEmpty()
    fmt.Println("Está vazia:", isEmpty)
    fmt.Println(stack.numeros)
    ultimo := stack.Pop()
    fmt.Println("Elemento removido:", ultimo)
    fmt.Println(stack.numeros)
    size := stack.Size()
    fmt.Println("Tamanho da lista", size)
}

