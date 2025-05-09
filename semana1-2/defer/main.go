package main

import (
	"fmt"
)

func operacion(saldo float64, ops ...func(float64) float64) float64 {
	// Función para capturar cualquier panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de un error:", r)
		}
	}()

	for _, op := range ops {
		// Captura el valor actual antes de la operación
		saldoAntes := saldo

		// defer para imprimir el resultado de cada operación (pero se ejecutará después del return o panic)

		defer fmt.Printf("Operación completada con saldo: %.2f\n", saldoAntes)

		// Ejecutar la operación
		saldo = op(saldo)
	}

	return saldo

}

func main() {
	ingreso := func(s float64) float64 {
		return s + 100
	}

	retiro := func(s float64) float64 {
		return s - 50
	}

	panico := func(s float64) float64 {
		panic("¡Error en operación!")
	}

	saldoFinal := operacion(200, ingreso, retiro, panico, ingreso)

	fmt.Println("Saldo final:", saldoFinal)
}

/*
Crea una función que simule varias operaciones bancarias (ingresos y retiros) sobre una cuenta. Usa defer para asegurarte de que se imprimen los logs de cada operación al salir de la función, incluso si ocurre un error.

💡 Requisitos:
Define una función operacion que reciba:

Un saldo inicial (float64).

Una lista de operaciones, donde cada operación sea una función que reciba y devuelva un float64.

Dentro de operacion, ejecuta cada operación, y usa defer para registrar lo que ocurrió (antes de que termine la función).

Agrega una operación que cause un panic, y usa recover para capturarla en un defer.
*/
