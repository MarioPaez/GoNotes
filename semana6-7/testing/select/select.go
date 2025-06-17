package selectTdd

import (
	"net/http"
	"time"
)

var Urls = []string{
	"https://httpbin.org/get",
	"https://httpbin.org/json",
	"https://httpbin.org/status/200",
	"https://httpbin.org/status/404",
	"https://httpbin.org/status/500",
	"https://httpbin.org/delay/2",
	"https://postman-echo.com/get?test=1",
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://api.github.com",
	"https://www.google.com",
}

// Probar a realizar un Racer de un array de urls de entrada
func Racer(url1, url2 string) string {
	timeA := time.Now()
	http.Get(url1)
	totalA := time.Since(timeA)

	timeB := time.Now()
	http.Get(url2)
	totalB := time.Since(timeB)
	if totalA > totalB {
		return url2
	}
	return url1
}

/*
1. Vamos a realizar lo mínimo y lo más sencillo para que el test pase, luego viene el refactor y hacerlo más rápido.
*/
