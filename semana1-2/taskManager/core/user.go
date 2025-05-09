package core

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Pais   string `json:"pais"`
	//No sé si tiene que haber aquí una tarea embebida
}
