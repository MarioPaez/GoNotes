package core

import "time"

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "TODO"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN_PROGRESS"
	TASK_STATUS_DONE        TaskStatus = "DONE"
)

type Tarea struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	//Estado No sé si es mejor que vaya esto
	Estado        TaskStatus `json:"estado"`
	FechaCreacion time.Time  `json:"fecha_creacion"`
	Usuario       Usuario    `json:"usuario"`
}
