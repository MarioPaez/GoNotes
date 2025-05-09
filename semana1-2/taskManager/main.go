package main

import (
	"core"
)

func main() {

	sliceTareas := make(core.TaskManager, 0)

	sliceTareas.Crear("Tarea 1", "Descripcion de la tarea 1", Usuario1)
	sliceTareas.Crear("Tarea 2", "Descripcion de la tarea 2", Usuario1)

	sliceTareas.CompletarTarea(sliceTareas[0].ID)
	sliceTareas.EliminarTarea(sliceTareas[0].ID)

	sliceTareas.Listar()

}
