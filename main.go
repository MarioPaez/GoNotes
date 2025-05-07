package main

import (
	"fmt"
	"unsafe"
)

func addElement(slice []int) {
	slice[0] = 8
	slice = append(slice, 60)
	fmt.Println("Valor del slice dentro:", slice)
}

func addElement2(slice *[]int) {
	*slice = append(*slice, 60)
	fmt.Println("Valor del slice dentro:", *slice)
}

func main() {
	var positivo int8 = 127  //El número entero sin signo se almacena en binario puro
	var negativo int8 = -128 //El número entero con signo se almacena en complemento a dos

	fmt.Println("Valor de positivo en bits:", fmt.Sprintf("%08b", positivo))
	fmt.Println("Valor de negativo en bits:", fmt.Sprintf("%08b", negativo))
	fmt.Println("Valor de negativo en bits (complemento a 2):", fmt.Sprintf("%08b", uint8(negativo)))

	var arq int = 64

	fmt.Println("Valor de arq en bits:", fmt.Sprintf("%064b", arq))

	var unsigned uint = 64
	fmt.Println("Valor de unsigned en bits:", fmt.Sprintf("%064b", unsigned))

	var t bool = true
	fmt.Println("Valor de t:", fmt.Sprintf("%t", t))
	var f bool = false
	fmt.Println("Valor de f:", fmt.Sprintf("%t", f))

	bVal := *(*byte)(unsafe.Pointer(&f))
	fmt.Println("Valor de t en bits:", fmt.Sprintf("%08b", bVal))

	var cadena1 string = "Holá"
	fmt.Println("Dirección de memoria antes:", len(cadena1), unsafe.Pointer(&cadena1))

	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Valor del slice antes:", slice)
	addElement(slice)
	fmt.Println("Valor del slice fuera:", slice)
	addElement2(&slice)
	fmt.Println("Valor del slice despues de pasarlo por referencia:", slice)

}
