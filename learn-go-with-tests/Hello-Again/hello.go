package main

import (
	"fmt"
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
	fmt.Println(Hello("", ""))
}
