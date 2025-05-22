package main

import (
	"fmt"
	"sync"
)

type Task struct {
	ID    int
	Input string
}

func createNTasks(id int) Task {
	return Task{ID: id, Input: CreateRandomString()}
}

// Fan-out
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0

	for range tasks {
		//fmt.Printf("Worker %d procesando tarea %d\n", id, task.ID)
		count++
		//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
	fmt.Printf("El worker con id %d ha consumido un total de %d tareas\n", id, count)
}

func main() {
	numTasks := 1000
	numWorkers := 5

	tasks := make(chan Task)
	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	for i := range numTasks {
		tasks <- createNTasks(i)
	}
	close(tasks)
	wg.Wait()

	fmt.Println("Todas las tareas han sido procesadas")
}
