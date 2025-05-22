package main

import "fmt"

type Task struct {
	ID    int
	Input string
}

func createNTasks(n int) []Task {
	res := make([]Task, 0)
	for i := range n {
		res = append(res, Task{ID: i, Input: CreateRandomString()})
	}
	return res
}

func main() {

	tasks := createNTasks(1000)

	fmt.Println("Tareas que hay que hacer: ", tasks)

}
