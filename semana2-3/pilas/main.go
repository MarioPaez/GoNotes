package main

import "fmt"

type Element[T any] struct {
	next, prev *Element[T]
	val        T
}

type Stack[T any] struct {
	len  int
	head *Element[T]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack[T]) Init(val T) {
	s.head = &Element[T]{next: nil, prev: nil, val: val}
	s.len++
}

func (s *Stack[T]) Push(val T) {
	if s.IsEmpty() {
		s.Init(val)
	} else {
		newElement := &Element[T]{next: nil, prev: s.head, val: val}
		s.head = newElement
		fmt.Println("He pusheado el ", val, " y el anterior vale: ", newElement.prev)
		s.len++
	}
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.head == nil {
		var zero T
		return zero, false
	}
	return s.head.val, true
}

func (s *Stack[T]) Pop() (T, bool) {
	if !s.IsEmpty() {

		aux := s.head.val
		s.head = s.head.prev
		fmt.Println("He popeado el ", aux, " y ahora head vale: ", s.head)
		if s.head != nil {
			s.head.next = nil
		}
		s.len--
		return aux, true
	} else {
		var zero T
		return zero, false
	}
}

func (s *Stack[T]) Size() int {
	return s.len
}

func (s *Stack[T]) String() string {
	if s.IsEmpty() {
		return ""
	}
	result := "Top->"
	for current := s.head; current != nil; current = current.prev {
		result += fmt.Sprint(current.val)
	}

	result += "<- Bottom"
	return result
}

func main() {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Peek())
	fmt.Println("Tamaño: ", stack.Size())
	stack.Push(3)
	fmt.Print(stack.String())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println("Tamaño: ", stack.Size())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
