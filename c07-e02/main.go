package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James May",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p1 *person) {
	p1.name = "Hello Aloha"
	//(*p1).name = "Hello Kitty"
}
