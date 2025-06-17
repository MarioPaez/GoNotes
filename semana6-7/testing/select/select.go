package selectTdd

import (
	"fmt"
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
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

/*
1. Vamos a realizar lo mínimo y lo más sencillo para que el test pase, luego viene el refactor y hacerlo más eficiente.
*/
