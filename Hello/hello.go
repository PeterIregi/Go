// writing a Hello World go program
package main

import "fmt"

// lets create the function
func Swap(fp, sp string) string {
	return sp + fp
}
func main() {
	//today I am going to try and see if I can swap the words in the greeting using a function i will create called swap
	// lets test it out
	fmt.Println(Swap("Hello ", "World! "))
	//should print out World! Hello
	//should print out Hello World to the console when I run it
	fmt.Println("Hello world!")
	//was sucessfull
}
