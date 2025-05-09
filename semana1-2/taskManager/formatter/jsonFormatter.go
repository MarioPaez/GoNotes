package formatter

import (
	"bytes"
	"encoding/json"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Export(tareas []any) string {
	var prettyJSON bytes.Buffer

	for _, tarea := range tareas {
		tareaByte, marshalErr := json.Marshal(tarea)
		if marshalErr != nil {
			panic("No se ha podido convertir la estructura Tarea a JSON: " + marshalErr.Error())
		}
		indentErr := json.Indent(&prettyJSON, tareaByte, "", "  ")
		if indentErr != nil {
			panic("No se ha podido indentar el JSON: " + indentErr.Error())
		}

	}

	return prettyJSON.String()

}
