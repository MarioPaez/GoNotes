package formatter

import (
	"fmt"
)

type TextFormatter struct{}

func (f *TextFormatter) Export(tareas []any) string {

	for _, tarea := range tareas {
		fmt.Println("Tarea: ", tarea)
	}
	return ""
}
