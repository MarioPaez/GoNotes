package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string {

	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%d\n", i)
			time.Sleep(100000)
		}
	}()
	return c
}

func main() {
	c := boring("Boring!")

	for range 5 {
		fmt.Printf("You say: %s\n", <-c)
	}

	fmt.Println("Me piro!")
}
