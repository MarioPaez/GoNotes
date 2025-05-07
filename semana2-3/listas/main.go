package main

import "fmt"

type ListaSlice struct {
	list []int
}

func (l *ListaSlice) Append(val ...int) {
	n := len(l.list)
	total := n + len(val)

	if total > cap(l.list) { //Ampliar la capacidad
		newSlice := make([]int, total, total*2)
		copy(newSlice, l.list)
		l.list = newSlice
	}

	fmt.Println("La longitud de slice es: ", len(l.list), " y la capaciad es: ", cap(l.list))
	l.list = l.list[:total]
	copy(l.list[n:], val)
}

func main() {
	var lista ListaSlice
	lista.Append(1, 2, 3)
	lista.Append(1, 2, 3)
	fmt.Println(lista.list)
}
