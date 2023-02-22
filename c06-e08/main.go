package main

import "fmt"

func main() {
	a := foo()
	a()
}

func foo() func() {
	return func() {
		fmt.Println("I am the returned func")
	}
}

/*
Hands-on exercise #8
● Create a func which returns a func
● assign the returned func to a variable
● call the returned func
*/
