package main

import (
	"fmt"
	"sync"
)

func sum(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range in {
		fmt.Println("Recibiendo el valor de ", value)
		out <- (2 * value)
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	var wg, sumWg sync.WaitGroup

	wg.Add(1)
	go sum(in, out, &wg)
	sumWg.Add(1)
	go func() {
		defer sumWg.Done()
		for i := range out {
			fmt.Println("Valor de out: ", i)
		}
	}()

	for i := range 5 {
		in <- i
	}

	close(in)
	wg.Wait()
	close(out)
	sumWg.Wait()
	fmt.Println("Terminamos la ejecución")
}
