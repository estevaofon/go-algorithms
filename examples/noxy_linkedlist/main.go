package main

import "fmt"

type Node struct {
	valor   int
	proximo *Node // Ponteiro para Node
}

func adicionar(node *Node, value int) {
	if node.proximo == nil {
		node.proximo = &Node{
			valor:   value,
			proximo: nil,
		}
	} else {
		// Recursivo: adiciona no final da lista
		adicionar(node.proximo, value)
	}
}

func printList(node *Node) {
	for {
		if node != nil {
			fmt.Println(node.valor)
			node = node.proximo
		} else {
			break
		}
	}
}

func main() {
	node := Node{
		valor:   10,
		proximo: nil,
	}
	adicionar(&node, 20)
	adicionar(&node, 30)
	adicionar(&node, 40)
	fmt.Println(node)
	printList(&node)
}
