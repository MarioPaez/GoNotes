package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	spanish             = "Spanish"
	french              = "French"
	galician            = "Galician"
	prefixHelloEnglish  = "Hello, "
	prefixHelloSpanish  = "Hola, "
	prefixHelloGalician = "Ola, "
	prefixHelloFrench   = "Bonjour, "
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = prefixHelloSpanish
	case french:
		prefix = prefixHelloFrench
	case galician:
		prefix = prefixHelloGalician
	default:
		prefix = prefixHelloEnglish
	}
	return
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Mariano")
}

func main() {
	fmt.Println(Hello("Mario", ""))
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

}
