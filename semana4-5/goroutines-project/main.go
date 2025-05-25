package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	ID    int
	Input string
}

type Result struct {
	TaskID     int
	WorkerID   int
	Success    bool
	DurationMs int64
}

type Stats struct {
	sync.Mutex
	Successes int
	Failures  int
	ByWorker  map[int]int
}

func createNTasks(id int) Task {
	return Task{ID: id, Input: CreateRandomString()}
}

// Fan-out
func worker(id int, tasks <-chan Task, result chan<- Result, wg *sync.WaitGroup) {
	//Al final de la función, restamos 1 al contador interno.
	defer wg.Done()

	for task := range tasks {
		timeTask := time.Duration(rand.Intn(1000)) * time.Millisecond
		hasFailed := rand.Float32() > 0.2
		time.Sleep(timeTask)
		fmt.Printf("La tarea %d ha tardado %d ms\n", task.ID, timeTask.Milliseconds())
		result <- Result{
			TaskID:     task.ID,
			WorkerID:   id,
			Success:    hasFailed,
			DurationMs: timeTask.Milliseconds(),
		}
	}
}

func main() {
	numTasks := 10
	numWorkers := 5

	tasks := make(chan Task)
	results := make(chan Result)
	stats := Stats{
		ByWorker: make(map[int]int),
	}

	var wg sync.WaitGroup
	var statsWg sync.WaitGroup

	statsWg.Add(1)
	go func() {
		defer statsWg.Done()
		for result := range results {
			stats.Lock()
			if result.Success {
				stats.Successes++
			} else {
				stats.Failures++
			}
			stats.ByWorker[result.WorkerID]++
			stats.Unlock()
		}
	}()

	for i := range numWorkers {
		wg.Add(1) // Añadimos 1 al contador interno
		go worker(i, tasks, results, &wg)
	}

	for i := range numTasks {
		tasks <- createNTasks(i)
	}

	close(tasks) //Importante para que los bucles donde espero tasks se cierren.
	wg.Wait()    //Cuando contador interno = 0, se desbloquea y continúa

	close(results)
	statsWg.Wait()

	fmt.Println("Todas las tareas han sido procesadas")

	fmt.Println("Resumen:")
	fmt.Printf("  Tareas exitosas: %d\n", stats.Successes)
	fmt.Printf("  Tareas fallidas: %d\n", stats.Failures)
	fmt.Println("  Tareas por worker:")
	for id, count := range stats.ByWorker {
		fmt.Printf("    Worker %d: %d tareas\n", id, count)
	}
}
