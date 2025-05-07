package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (list *LinkedList) Append(val int) {
	newNode := &Node{val: val, next: nil}

	if list.head == nil {
		list.head = newNode
		list.len++
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	list.len++
}

func (list *LinkedList) Prepend(val int) {
	aux := list.head
	list.head = &Node{val: val, next: aux}
	list.len++
}

func (list *LinkedList) InsertAt(val, position int) {
	if position >= 0 && position <= list.len {
		if position == 0 {
			list.Prepend(val)
			return
		}
		if position == list.len {
			list.Append(val)
			return
		}
		aux := 0
		current := list.head
		for aux+1 != position {
			current = current.next
			aux++
		}
		auxNode := current.next
		current.next = &Node{val: val, next: auxNode}
		list.len++
	}
}

func (list *LinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.val)
		current = current.next
	}

	fmt.Println("nil")
}

func (list LinkedList) Longitud() int {
	return list.len
}

func (list *LinkedList) BorrarUltimo() {
	current := list.head

	if current != nil {
		if current.next == nil {
			list.head = nil
			list.len--
			return
		}

		for current.next.next != nil {
			current = current.next
		}

		current.next = nil
		list.len--
	}

}

func (list *LinkedList) BorrarPrimero() {
	current := list.head
	if current == nil {
		return
	}

	list.head = current.next
	list.len--
}

func (list *LinkedList) BorrarEn(val int) {
	aux := 0
	if val >= 0 && val < list.len {
		if val == 0 {
			list.BorrarPrimero()
			return
		}
		current := list.head
		for aux+1 != val {
			current = current.next
			aux++
		}
		current.next = current.next.next
		list.len--
	}
}

func (list *LinkedList) FindNode(val int) (*Node, error) {
	current := list.head

	for current != nil {
		if current.val == val {
			return current, nil
		}
		current = current.next
	}
	return nil, fmt.Errorf("no se encuentra ningún nodo con ese valor")
}

func (list *LinkedList) GetAt(index int) (*Node, error) {
	if index < 0 || index >= list.len {
		return nil, fmt.Errorf("no hay nodos en esa posicion")
	}

	current := list.head
	aux := 0

	for aux+1 < index {
		current = current.next
		aux++
	}

	return current, nil
}

func (list *LinkedList) ToSlice() []int {
	slice := make([]int, list.len)
	current := list.head
	aux := 0
	for current != nil {
		slice[aux] = current.val
		current = current.next
		aux++
	}

	return slice
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Reverse() {
	var ant *Node
	var aux *Node
	current := list.head

	for current != nil {
		aux = current.next
		current.next = ant
		ant = current
		current = aux
	}

	list.head = ant

}

func main() {
	listaPrueba := LinkedList{}
	fmt.Println(listaPrueba.IsEmpty())

	listaPrueba.Append(1)
	listaPrueba.BorrarUltimo()
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())

	listaPrueba.Append(2)
	listaPrueba.Append(3)
	listaPrueba.Append(4)
	listaPrueba.Append(5)
	listaPrueba.Append(6)

	listaPrueba.Print()

	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())

	listaPrueba.BorrarUltimo()
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())

	listaPrueba.BorrarPrimero()
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())

	listaPrueba.BorrarEn(2)
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())
	listaPrueba.BorrarEn(2)
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())
	listaPrueba.BorrarEn(0)
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())

	listaPrueba.Prepend(1)
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())

	listaPrueba.InsertAt(2, 1)
	listaPrueba.Print()
	fmt.Println("Longitud de la lista ", listaPrueba.Longitud())

	fmt.Println(listaPrueba.FindNode(2))
	fmt.Println(listaPrueba.FindNode(3))

	fmt.Println(listaPrueba.GetAt(2))
	fmt.Println(listaPrueba.GetAt(3))

	fmt.Println(listaPrueba.ToSlice())
	fmt.Println(listaPrueba.IsEmpty())

	listaPrueba.Reverse()
	listaPrueba.Print()

	slice := make([]int, 1)
	slice[0] = 1

	slice = Append(slice, 3)
	fmt.Println(slice)

}

func Append(slice []int, values ...int) []int {
	longitud := len(slice)
	total := longitud + len(values)

	if total > cap(slice) {
		newSlice := make([]int, longitud, total)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:total]
	copy(slice[longitud:], values)

	return slice

}
