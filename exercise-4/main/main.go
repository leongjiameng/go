package main

import "fmt"

type custom int

var x custom

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Printf("\n%v", x)
}
