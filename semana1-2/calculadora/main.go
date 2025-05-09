package main

import (
	cal "calculadora/mathutils"
	p "closures/prueba"
	"fmt"
)

func main() {
	cal.Add(5, 3)
	p := p.Persona{Nombre: "Juan", Edad: 20}
	fmt.Print(p)
}
