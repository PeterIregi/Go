package main

import (
	"fmt"
)

func Hello(s string) string {
	//defining a constant
	const englishHelloPrefix = "Hello, "
	return englishHelloPrefix + s
}

func main() {
	fmt.Println(Hello("world"))
}
