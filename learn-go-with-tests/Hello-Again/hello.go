package main

import (
	"fmt"
)

func Hello(s string) string {
	const englishHelloPrefix = "Hello, "
	if s == "" {
		return englishHelloPrefix + "world"
	}
	return englishHelloPrefix + s
}

func main() {
	fmt.Println(Hello(""))
}
