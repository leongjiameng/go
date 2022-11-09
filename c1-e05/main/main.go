package main

import "fmt"

type custom int

var x custom
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Printf("\n%v\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
