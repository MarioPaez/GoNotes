package main

import "fmt"

// Aquí hay un problema si quiero extender un slice cuya longitud sea igual a su capacidad
func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func ExtendFixed(slice []int, element int) []int {
	n := len(slice)
	c := cap(slice)
	if n == c {
		newSlice := make([]int, n, c*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func InitProblem() {
	var iBuffer [10]int
	slice := iBuffer[0:0]
	fmt.Printf("longitud: ", len(slice))
	for i := 0; i < 20; i++ {
		slice = ExtendFixed(slice, i)
		fmt.Printf("Dirección de memoria del array[0]", &slice[0])
		fmt.Println(slice)
	}
	slice2 := []int{0, 1, 2, 3, 4}
	fmt.Println(slice2)
	slice2 = Append(slice2, 5, 6, 7, 8)
	slice3 := Append(slice2, slice...)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

// func Append(slice []int, items ...int) []int {
// 	for _, item := range items {
// 		slice = ExtendFixed(slice, item)
// 	}
// 	return slice
// }

func Append(slice []int, items ...int) []int {
	n := len(slice)
	total := n + len(items)

	if total > cap(slice) {
		newSize := total*3/2 + 1 //Garantiza eficiencia de memoria
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total] //Extendemos
	copy(slice[n:], items)
	return slice
}
