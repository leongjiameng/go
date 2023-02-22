package main

import "fmt"

func main() {
	defer foo()
	bar()
	fmt.Println("Check")
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	defer func() {
		fmt.Println("Defer func run")
	}()
	fmt.Println("bar")
}

/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
