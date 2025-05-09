package main

import (
	"exercise/student"
	"fmt"
)

type Classroom []student.Student

func find(c Classroom, name string) *student.Student {
	for _, student := range c {
		if student.Name == name {
			return &student
		}
	}
	return nil

}

func main() {
	mario := student.Student{
		ID:     1,
		Name:   "Mario",
		Grades: []float64{8.5, 9.0, 7.5},
	}

	luigi := student.Student{
		ID:     2,
		Name:   "Luigi",
		Grades: []float64{7.0, 8.0, 9.5},
	}
	peach := student.Student{
		ID:     3,
		Name:   "Peach",
		Grades: []float64{9.0, 8.5, 9.5},
	}

	terceroA := Classroom{
		mario, luigi, peach}

	mario.Print()
	mario.Average()

	find(terceroA, "Mario").Print()

	fmt.Println("ID:", luigi.ID)
	luigi.CambioIdPorValor(10)
	fmt.Println("ID:", luigi.ID)
	luigi.CambioIdPorReferencia(20)
	fmt.Println("ID:", luigi.ID)
}
