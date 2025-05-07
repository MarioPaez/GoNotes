package main

import (
	"fmt"
)

type Node struct {
	val         int
	left, right *Node
}

type ABS struct {
	root *Node
}

func (a *ABS) IsEmpty() bool {
	return a == nil || a.root == nil
}

func PreOrder(nodo *Node) {
	//root, left, right
	if nodo == nil {
		return
	}
	fmt.Print(nodo.val, ", ")
	PreOrder(nodo.left)
	PreOrder(nodo.right)
}

// DFS
func InOrder(nodo *Node) {
	//root, left, right
	if nodo == nil {
		return
	}

	InOrder(nodo.left)
	fmt.Print(nodo.val, ", ")
	InOrder(nodo.right)
}

func PostOrder(nodo *Node) {
	//root, left, right
	if nodo == nil {
		return
	}

	PostOrder(nodo.left)
	PostOrder(nodo.right)
	fmt.Print(nodo.val, ", ")
}

func (a *ABS) createTree() *Node {
	n8 := &Node{val: 8}
	n4 := &Node{val: 4, left: n8}
	n5 := &Node{val: 5}
	n6 := &Node{val: 6}
	n7 := &Node{val: 7}
	n2 := &Node{val: 2, left: n4, right: n5}
	n3 := &Node{val: 3, left: n6, right: n7}
	root := &Node{val: 1, left: n2, right: n3}
	return root
}

// BFS
func (a *ABS) BFS(nodo *Node) {
	cola := make([]Node, 0)
	cola = append(cola, *nodo)
	for len(cola) != 0 {
		if cola[0].left != nil {
			cola = append(cola, *cola[0].left)
		}
		if cola[0].right != nil {
			cola = append(cola, *cola[0].right)
		}
		fmt.Print(cola[0].val, ", ")
		cola = cola[1:]
	}
}

// Me apoyo con una cola
func (a *ABS) BuscarValorBFS(valor int) *Node {
	if a == nil || a.root == nil {
		return nil
	}
	var nodoAux *Node
	cola := make([]*Node, 0)
	cola = append(cola, a.root)
	for len(cola) > 0 {
		nodoAux = cola[0]
		if nodoAux.val == valor {
			return nodoAux
		}

		if nodoAux.left != nil {
			cola = append(cola, nodoAux.left)
		}
		if nodoAux.right != nil {
			cola = append(cola, nodoAux.right)
		}
		cola = cola[1:]
	}

	return nil
}

func (a *ABS) BuscarValorDFS(valor int) *Node {
	if a == nil || a.root == nil {
		return nil
	}
	return BuscarValorDFSAux(a.root, valor)
}

func BuscarValorDFSAux(nodo *Node, valor int) *Node {

	if nodo == nil {
		return nil
	}

	if nodo.val == valor {
		return nodo
	}

	leftTree := BuscarValorDFSAux(nodo.left, valor)

	if leftTree != nil {
		return leftTree
	}
	return BuscarValorDFSAux(nodo.right, valor)
}

func (a *ABS) Altura() int {
	if a == nil {
		return 0
	}
	return AlturaAux(a.root)
}

func AlturaAux(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + max(AlturaAux(node.left), AlturaAux(node.right))
}
