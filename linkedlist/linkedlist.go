package linkedlist

import "fmt"

type Node struct {
	Valor   int
	Proximo *Node
}

// Adicionar insere um valor no final da lista.
func (n *Node) Adicionar(value int) {
	if n.Proximo == nil {
		n.Proximo = &Node{
			Valor:   value,
			Proximo: nil,
		}
	} else {
		n.Proximo.Adicionar(value)
	}
}

// Exibir imprime todos os valores da lista.
func (n *Node) Exibir() {
	atual := n
	for atual != nil {
		fmt.Printf("%d -> ", atual.Valor)
		atual = atual.Proximo
	}
	fmt.Println("nil")
}
