package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	len  int
}

func (dll DoubleLinkedList) ToSlice() []int {
	slice := make([]int, 0, dll.len)

	current := dll.head

	for current != nil {
		slice = append(slice, current.val)
		current = current.next
	}

	return slice
}

func (dll *DoubleLinkedList) IsEmpty() bool {
	return dll.len == 0
}

func (dll *DoubleLinkedList) Append(val int) {

	if dll.IsEmpty() {
		dll.head = &Node{val: val, next: nil, prev: nil}
		dll.tail = dll.head
		dll.len++
		return
	}

	dll.tail.next = &Node{val: val, next: nil, prev: dll.tail}
	dll.tail = dll.tail.next
	dll.len++
}

func (dll *DoubleLinkedList) Prepend(val int) {
	if dll.IsEmpty() {
		dll.head = &Node{val: val, next: nil, prev: nil}
		dll.tail = dll.head
		dll.len++
		return
	}

	aux := dll.head
	dll.head = &Node{val: val, next: aux, prev: nil}
	aux.prev = dll.head
	dll.len++
}

func (list *DoubleLinkedList) PrintForward() {
	current := list.head
	for current != nil {
		fmt.Printf("%d <-> ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *DoubleLinkedList) PrintBackward() {
	current := list.tail
	for current != nil {
		fmt.Printf("%d <-> ", current.val)
		current = current.prev
	}
	fmt.Println("nil")
}

func (list *DoubleLinkedList) InserAt(val, index int) {

	if index == 0 {
		list.Prepend(val)
		return
	}

	if index == list.len {
		list.Append(val)
		return
	}

	if index > 0 && index < list.len {
		current := list.head
		aux := 0
		for aux+1 < index {
			current = current.next
			aux++
		}

		next := current.next
		current.next = &Node{val: val, next: next, prev: current}
		next.prev = current.next
		list.len++
		return
	}

	fmt.Println("el indice debe estar entre 0 y la longitud de la lista doblemente enlazada")
}

func (dll *DoubleLinkedList) DeleteFirst() {
	if dll.len == 1 {
		dll.head = nil
		dll.tail = nil
		dll.len--
		return
	}

	if !dll.IsEmpty() {
		aux := dll.head
		dll.head = aux.next
		dll.head.prev = nil
		dll.len--
	}
}

func (dll *DoubleLinkedList) DeleteLast() {
	if dll.len == 1 {
		dll.head = nil
		dll.tail = nil
		dll.len--
	}
	if !dll.IsEmpty() {
		aux := dll.tail
		dll.tail = aux.prev
		dll.tail.next = nil
		dll.len--
	}
}

func (dll *DoubleLinkedList) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.len = 0
}

func (dll *DoubleLinkedList) Length() int {
	return dll.len
}

func (dll *DoubleLinkedList) Reverse() {

	current := dll.head

	for current != nil {
		auxNext := current.next
		current.next = current.prev
		current.prev = auxNext
		current = auxNext
	}

	auxHead := dll.head
	dll.head = dll.tail
	dll.tail = auxHead

}

func main() {
	dll := DoubleLinkedList{head: nil, tail: nil, len: 0}
	dll.DeleteFirst()
	dll.DeleteLast()
	dll.Append(2)
	dll.DeleteFirst()
	dll.PrintForward()
	dll.DeleteLast()
	dll.Append(3)
	dll.Append(5)
	dll.Append(2)
	dll.PrintForward()
	dll.Prepend(1)
	dll.PrintForward()
	fmt.Println("Longitud de la LDE: ", dll.len)
	dll.InserAt(69, 2)
	dll.InserAt(69, 4)
	dll.InserAt(69, 6)
	dll.Clear()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	dll.PrintForward()

	dll.PrintForward()  // 1 -> 2 -> 3 -> nil
	dll.PrintBackward() // 3 -> 2 -> 1 -> nil

	dll.Reverse()

	dll.PrintForward() // 3 -> 2 -> 1 -> nil
	dll.PrintBackward()

	slice := dll.ToSlice()
	fmt.Println(slice) // [1 2 3]

}
