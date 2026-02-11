package main

import (
	"fmt"
	"net/http"
)

func Hello(name, language string) string {
	const spanish = "Spanish"
	const french = "French"
	const englishHelloPrefix = "Hello, "
	const spanishHelloprefix = "Hola, "
	const frenchHelloPrefix = "Bonjour, "
	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloprefix
	}
	if name == "" {
		name = "world"
	}
	return prefix + name
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	//fmt.Println(Hello("", ""))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	language := r.URL.Query().Get("language")

	message := Hello(name, language)

	fmt.Fprintln(w, message)
}
