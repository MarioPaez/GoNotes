package main

import "fmt"

func main() {
	abs := ABS{}
	fmt.Println("El árbol está vacío? ", abs.IsEmpty())
	abs.root = abs.createTree()
	PreOrder(abs.root)
	fmt.Println()
	InOrder(abs.root)
	fmt.Println()
	PostOrder(abs.root)
	fmt.Println()
	abs.BFS(abs.root)
	fmt.Println()
	fmt.Println("La altura del arbol es de:", abs.Altura())

	fmt.Println(abs.BuscarValorDFS(1))
	fmt.Println(abs.BuscarValorDFS(2))
	fmt.Println(abs.BuscarValorDFS(3))
	fmt.Println(abs.BuscarValorDFS(4))
	fmt.Println(abs.BuscarValorDFS(5))
	fmt.Println(abs.BuscarValorDFS(6))
	fmt.Println(abs.BuscarValorDFS(7))
	fmt.Println(abs.BuscarValorDFS(8))

	fmt.Println(abs.BuscarValorBFS(1))
	fmt.Println(abs.BuscarValorBFS(2))
	fmt.Println(abs.BuscarValorBFS(3))
	fmt.Println(abs.BuscarValorBFS(4))
	fmt.Println(abs.BuscarValorBFS(5))
	fmt.Println(abs.BuscarValorBFS(6))
	fmt.Println(abs.BuscarValorBFS(7))
	fmt.Println(abs.BuscarValorBFS(8))

	fmt.Println("Empezamos con BST")
	bst := BST{root: nil}
	bst.Insert(5)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(8)
	bst.Insert(4)

	bst.BFS(bst.root)
	fmt.Println()
	fmt.Println(bst.BuscarBST(1))
	fmt.Println(bst.BuscarBST(2))
	fmt.Println(bst.BuscarBST(3))
	fmt.Println(bst.BuscarBST(4))
	fmt.Println(bst.BuscarBST(5))
	fmt.Println(bst.BuscarBST(6))
	fmt.Println(bst.BuscarBST(7))
	fmt.Println(bst.BuscarBST(8))

	PreOrder(bst.root)
	fmt.Println()
	InOrder(bst.root)
	fmt.Println()
	PostOrder(bst.root)
	fmt.Println()
	bst.BFS(bst.root)
	fmt.Println()

	bst.EliminarBST(5)
	fmt.Println("Número de nodos: ", bst.ContarNodos(bst.root))
	fmt.Println("Suma de los nodos: ", bst.SumarNodos(bst.root))
	bst.BFS(bst.root)
	pr := bst.BuscarBST(6)
	fmt.Printf("Dirección de memoria del nodo: %p\n", pr)
}
