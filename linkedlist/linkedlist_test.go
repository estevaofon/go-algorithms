package linkedlist_test

import (
	"github.com/estevaofon/go-algorithms/linkedlist"
)

func ExampleNode() {
	node := linkedlist.Node{Valor: 10}

	node.Adicionar(20)
	node.Adicionar(30)
	node.Adicionar(40)

	node.Exibir()

	// Output:
	// 10 -> 20 -> 30 -> 40 -> nil
}
