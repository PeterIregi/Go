package main

import "errors"

var ErrDivideByZero = errors.New("Cannot divide by zero")

func main() {

}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

//lets see if it will fail now
//lets improve the test
