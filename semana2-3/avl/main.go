package main

import "fmt"

/*
1. LL -> FE de z = +2 y FE hijoIzq (y) = +1 o 0 -> rotación hacia la derecha sobre z.
2. RR -> FE de z = -2 y FE hijoDer (y) = -1 o 0 -> rotación hacia la izquierda sobre z.
3. LR -> FE de z = +2 y FE hijoIzq (y) = -1 -> rotación hacia la izquierda sobre y y rotacion derecha sobre z.
4. RL -> FE de z = -2 y FE hijoDer (y) +1 -> rotación hacia la derecha sobre y y rotación izquierda sobre z.
*/

type AVL struct {
	root *Node
}

type Node struct {
	val         int
	left, right *Node
	height      int
}

func (avl *AVL) Insert(val int) {
	avl.root = auxInsert(avl.root, val)
}

func auxInsert(node *Node, val int) *Node {
	if node == nil {
		return &Node{val: val, height: 1}
	}

	if val < node.val {
		node.left = auxInsert(node.left, val)
	} else if val > node.val {
		node.right = auxInsert(node.right, val)
	} else {
		fmt.Println("Ya existe un nodo con el valor ", val)
		return node
	}

	updateHight(node)
	return balanceTree(node)
}

func balanceTree(node *Node) *Node {

	if node == nil {
		return node
	}

	feZ := FE(node)
	feYLeft := FE(node.left)
	feYRight := FE(node.right)

	//CASO LL
	if feZ == 2 && feYLeft >= 0 {
		node = rotateRight(node)
	}
	//CASO RR
	if feZ == -2 && feYRight <= 0 {
		node = rotateLeft(node)
	}
	//CASO LR
	if feZ == 2 && feYLeft == -1 {
		node.left = rotateLeft(node.left)
		node = rotateRight(node)
	}
	//CASO RL
	if feZ == -2 && feYRight == 1 {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
	}

	return node
}

func rotateRight(z *Node) *Node {
	y := z.left
	t2 := y.right

	y.right = z
	z.left = t2

	updateHight(z)
	updateHight(y)

	return y
}

func rotateLeft(z *Node) *Node {
	y := z.right
	t2 := y.left

	y.left = z
	z.right = t2

	updateHight(z)
	updateHight(y)

	return y
}

func FE(node *Node) int {
	if node == nil {
		return 0
	}
	return getNodeHight(node.left) - getNodeHight(node.right)
}

func updateHight(node *Node) {
	if node == nil {
		return
	}
	node.height = 1 + max(getNodeHight(node.left), getNodeHight(node.right))
}

func getNodeHight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (avl *AVL) BFS() {
	node := avl.root
	cola := make([]*Node, 0)
	cola = append(cola, node)
	var auxNode *Node
	for len(cola) > 0 {
		auxNode, cola = cola[0], cola[1:]
		fmt.Printf("(H= %d)", auxNode.height)
		fmt.Print(auxNode.val, ", ")
		if auxNode.left != nil {
			cola = append(cola, auxNode.left)
		}
		if auxNode.right != nil {
			cola = append(cola, auxNode.right)
		}
	}
	fmt.Println()
}

func (avl *AVL) isBalanced() bool {
	return auxIsBalanced(avl.root)
}

func auxIsBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	fe := FE(node)
	return (fe == -1 || fe == 0 || fe == 1) && auxIsBalanced(node.left) && auxIsBalanced(node.right)
}

func (avl *AVL) Remove(val int) {
	avl.root = auxRemove(avl.root, val)
}

func auxRemove(node *Node, val int) *Node {
	if node == nil {
		fmt.Println("No existe ningún nodo con el valor ", val)
		return node
	}

	if val < node.val {
		node.left = auxRemove(node.left, val)
	} else if val > node.val {
		node.right = auxRemove(node.right, val)
	} else {
		//sin hijos o con 1 hijo
		if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			//Tenemos 2 hijos -> sucesor in orden y eliminar este
			minNode := inOrderSucesor(node.right)
			node.val = minNode.val
			node.right = auxRemove(node.right, minNode.val)
		}
	}
	updateHight(node)
	return balanceTree(node)
}

func inOrderSucesor(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (avl *AVL) Height() int {
	return getNodeHight(avl.root)
}

func (avl *AVL) Size() int {
	return auxSize(avl.root)
}

func auxSize(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + auxSize(node.left) + auxSize(node.right)
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	preOrder(node.left)
	preOrder(node.right)
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Print(node.val, " ")
	inOrder(node.right)
}

func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.val, " ")
}

func main() {

	avl := AVL{}

	for i := -1; i <= 10; i++ {
		avl.Insert(i)
	}
	fmt.Print("BFS: ")
	avl.BFS()
	fmt.Println("El árbol está balanceado? ", avl.isBalanced())
	avl.Remove(14)
	avl.Remove(7)
	fmt.Print("BFS: ")
	avl.BFS()
	fmt.Println("ALtura del árbol", avl.Height())
	fmt.Println("Número de nodos del arbol", avl.Size())
	preOrder(avl.root)
	fmt.Println()
	inOrder(avl.root)
	fmt.Println()
	postOrder(avl.root)
	fmt.Println()
}
