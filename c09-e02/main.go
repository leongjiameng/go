package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

// Method set
func main() {
	p1 := person{
		first: "James",
	}
	saySomething(&p1)
	//saySomething(p1) cannot use p1 (variable of type person) as human value in argument to saySomething: person does not implement human (method speak has pointer receiver)

	p1.speak()
}
