package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

type Temperatura interface {
	ToCelsius() Celsius
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) ToCelsius() Celsius {
	return c
}

func main() {

	a, b := 4, 5
	fmt.Print(a, b)

	fmt.Println("Hola, ", a, b, "lo que digas")
	fmt.Println("mundo")

	fmt.Printf("Probando v %v y T %T", a, a)

	var temperaturas []Temperatura

	var temperaturaF Fahrenheit = 100
	var temperaturaK Kelvin = 373.15
	var temperaturaC Celsius = 22

	temperaturas = append(temperaturas, temperaturaF, temperaturaC, temperaturaK)

	for _, temperatura := range temperaturas {
		switch t := temperatura.(type) {
		case Fahrenheit:
			fmt.Printf("Original %T. Celsius: %.2f\n", t, t.ToCelsius())
		case Kelvin:
			fmt.Printf("Original %T. Celsius: %.2f\n", t, t.ToCelsius())
		case Celsius:
			fmt.Printf("Original %T. Celsius: %.2f\n", t, t.ToCelsius())
		}
	}

}
