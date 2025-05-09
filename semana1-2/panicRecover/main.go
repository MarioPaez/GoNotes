package main

import (
	"fmt"
)

func safeDivide(a, b int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error recuperado:", r)
		}
	}()

	if b == 0 {
		panic("División por cero")
	}

	fmt.Println(a / b)
}

func Append(slice []int, value ...int) []int {
	n := len(slice)
	total := n + len(value)

	if total > cap(slice) {
		newSlice := make([]int, n, 2*n)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:total]
	copy(slice[n:], value)
	return slice
}

func generica[T any](a T) {
	fmt.Println(a)
}

func main() {
	safeDivide(10, 2) // → Resultado: 5.0
	safeDivide(8, 0)  // → Error: división por cero
	safeDivide(20, 4) // → Resultado: 5.0

	x := 10
	defer func() { fmt.Printf("Valor %#v", x) }()
	x = 20
}
