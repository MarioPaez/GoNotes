package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type CircleList struct {
	head, tail *Node
	len        int
}

func (cl *CircleList) IsEmpty() bool {
	return cl.len == 0
}

func (cl *CircleList) Length() int {
	return cl.len
}

func (cl *CircleList) Append(val int) {

	newNode := &Node{val: val}
	if cl.head == nil {
		cl.head = newNode
		cl.tail = newNode
		newNode.next = newNode
		cl.len++
		return
	}

	newNode.next = cl.head
	cl.tail.next = newNode
	cl.tail = newNode
	cl.len++
}

func (cl *CircleList) Preppend(val int) {

	newNode := &Node{val: val}
	if cl.head == nil {
		cl.Append(val)
		return
	}

	aux := cl.head
	newNode.next = aux
	cl.tail.next = newNode
	cl.head = newNode

	cl.len++
}

func (cl *CircleList) Print() {
	if cl.IsEmpty() {
		fmt.Println("Lista vacía")
		return
	}

	current := cl.head

	for current.next != cl.head {
		fmt.Printf("%d -> ", current.val)
		current = current.next
	}
	fmt.Printf("%d -> ", current.val)
	fmt.Println("CicloDeNuevo")
}

func (cl *CircleList) Reverse() {
	current := cl.head
	auxHead := cl.head
	var ant, aux *Node

	for current != cl.tail {
		aux = current.next
		current.next = ant
		ant = current
		current = aux
	}

	current.next = ant
	cl.head = current
	cl.tail = auxHead
	cl.tail.next = cl.head
}

func main() {
	cl := CircleList{}
	cl.Print()
	cl.Append(1)
	cl.Append(2)
	cl.Append(3)
	cl.Preppend(0)
	cl.Print()
	cl.Reverse()
	cl.Print()
}
