package main

import "fmt"

type Node struct {
    valor   int
    proximo *Node
}

// Método no próprio struct
func (n *Node) Adicionar(value int) {
    if n.proximo == nil {
        n.proximo = &Node{
            valor:   value,
            proximo: nil,
        }
    } else {
        n.proximo.Adicionar(value)
    }
}

// Exibir toda a lista
func (n *Node) Exibir() {
    atual := n
    for atual != nil {
        fmt.Printf("%d -> ", atual.valor)
        atual = atual.proximo
    }
    fmt.Println("nil")
}

func main() {
    node := Node{valor: 10}
    
    node.Adicionar(20)
    node.Adicionar(30)
    node.Adicionar(40)
    
    node.Exibir()  // 10 -> 20 -> 30 -> 40 -> nil
}
