package student

import "sync"

type Student struct {
	ID     int
	Name   string
	Grades []float64
}

type Key struct {
	Path, Country string
}




func prueba(){
	
	var counter = struct{
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.Lock()

}

func (s Student) Average() float64 {
	var sum float64
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}

func (s Student) Print() {
	println("ID:", s.ID)
	println("Name:", s.Name)
	println("Average Grade:", s.Average())
	for _, grade := range s.Grades {
		println("Grade:", grade)
	}
}

func (s Student) CambioIdPorValor(valor int) {
	s.ID = valor
	println("CambioIdPorValor ID:", s.ID)
}

func (s *Student) CambioIdPorReferencia(valor int) {
	s.ID = valor
	println("CambioIdPorValor ID:", s.ID)
}
