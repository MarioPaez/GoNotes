package core

import (
	"fmt"
	"time"

	formatter "formatter.local"
	"github.com/google/uuid"
)

type TaskManager []Tarea

func (manager *TaskManager) Crear(nombre, descripcion string, user Usuario) {

	task := Tarea{
		ID:            uuid.NewString(),
		Nombre:        nombre,
		Descripcion:   descripcion,
		Estado:        TASK_STATUS_TODO,
		FechaCreacion: time.Now(),
		Usuario:       user,
	}

	*manager = append(*manager, task)
}

func (manager TaskManager) Listar() {
	formatter := formatter.JSONFormatter{}
	fmt.Println(formatter.Export([]any{manager}))
}

func (manager *TaskManager) CompletarTarea(id string) {
	defer func() {
		fmt.Println("Se ha completado la tarea con ID:", id)
	}()
	for i, task := range *manager {
		if task.ID == id {
			(*manager)[i].Estado = TASK_STATUS_DONE
			break
		}

	}
}

func (manager *TaskManager) EliminarTarea(id string) {
	for i, task := range *manager {
		if task.ID == id {
			*manager = append((*manager)[:i], (*manager)[i+1:]...)
		}
	}
}
