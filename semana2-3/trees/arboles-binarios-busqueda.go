package main

import "fmt"

type BST struct {
	root *Node
}

func (a *BST) IsEmpty() bool {
	return a == nil || a.root == nil
}

func (a *BST) Insert(valor int) {
	a.root = auxInsert(a.root, valor)
}

func auxInsert(node *Node, valor int) *Node {
	if node == nil {
		return &Node{val: valor}
	}

	if valor < node.val {
		node.left = auxInsert(node.left, valor)
	} else {
		node.right = auxInsert(node.right, valor)
	}
	return node
}

func (a *BST) BuscarBST(valor int) *Node {
	if a == nil || a.root == nil {
		return nil
	}
	return auxBuscarBST(a.root, valor)
}

func auxBuscarBST(node *Node, valor int) *Node {
	if node == nil {
		return nil
	}
	if node.val == valor {
		return node
	}
	if valor < node.val {
		return auxBuscarBST(node.left, valor)
	}

	return auxBuscarBST(node.right, valor)
}

func (a *BST) BFS(nodo *Node) {
	cola := make([]*Node, 0)
	cola = append(cola, nodo)
	for len(cola) != 0 {
		if cola[0].left != nil {
			cola = append(cola, cola[0].left)
		}
		if cola[0].right != nil {
			cola = append(cola, cola[0].right)
		}
		fmt.Print(cola[0].val, ", ")
		cola = cola[1:]
	}
}

func (a *BST) EliminarBST(valor int) {
	if a == nil || a.root == nil {
		fmt.Println("No se puede eliminar porque no hay valores en árbol")
	}
	a.root = auxEliminarBST(a.root, valor)
}

func auxEliminarBST(node *Node, valor int) *Node {
	if node == nil {
		fmt.Println("El valor no se encuentra en el árbol")
		return nil
	}

	if valor < node.val {
		node.left = auxEliminarBST(node.left, valor)
	} else if valor > node.val {
		node.right = auxEliminarBST(node.right, valor)
	} else {
		//Caso 1 y 2
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			fmt.Printf("Dirección de memoria del nodo a borrar: %p\n", node)
			minNode := findMin(node.right)
			node.val = minNode.val
			node.right = auxEliminarBST(node.right, minNode.val)
		}
	}
	return node
}

func findMin(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (a *BST) ContarNodos(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + a.ContarNodos(node.left) + a.ContarNodos(node.right)
}

func (a *BST) SumarNodos(nodo *Node) int {
	if nodo == nil {
		return 0
	}
	return nodo.val + a.SumarNodos(nodo.left) + a.SumarNodos(nodo.right)
}
