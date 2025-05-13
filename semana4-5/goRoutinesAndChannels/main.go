package main

import (
	"fmt"
	"math/rand"
	"time"
)

func productor(name string, out chan<- string) {
	for i := range 5 {
		time.Sleep(time.Duration(rand.Intn(1200)) * time.Millisecond)
		out <- fmt.Sprintf("%s: mensaje %d", name, i)
	}
	close(out)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go productor("ch1", ch1)
	go productor("ch2", ch2)
	go productor("ch3", ch3)

	for ch1 != nil || ch2 != nil || ch3 != nil {
		select {
		case res1, ok := <-ch1:
			if !ok {
				ch1 = nil
			}
			fmt.Println(res1)

		case res2, ok := <-ch2:
			if !ok {
				ch2 = nil
			}
			fmt.Println(res2)
		case res3, ok := <-ch3:
			if !ok {
				ch3 = nil
			}
			fmt.Println(res3)

		case <-time.After(1 * time.Second):
			fmt.Println("Nadie ha respondido en 1s")
		}

	}

	fmt.Println("Todos los canales han terminado")

}
