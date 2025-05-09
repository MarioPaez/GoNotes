package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
}

type Estudiante struct {
	Persona
	Carrera string
	Nombre  string
	Edad    int
}

func (p Persona) Saludar() {
	fmt.Printf("Hola, soy %s y tengo %d años\n", p.Nombre, p.Edad)
}

func Counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	c := Counter()

	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

	estudiante := Estudiante{Persona{"Juan", 20}, "Ingeniería", "Rodrigo", 20}
	estudiante.Saludar()
}
