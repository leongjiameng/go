package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first+" "+p.last, "said hi. I am", p.age, "years old.")
}

func (person) talk() {
	fmt.Println("Someone said hi")
}

func main() {
	p1 := person{
		first: "Hello",
		last:  "Guys",
		age:   27,
	}
	fmt.Println(p1)
	p1.speak()
	p1.talk()
}

/*
Hands-on exercise #4
● Create a user defined struct with
○ the identifier “person”
○ the fields:
■ first
■ last
■ age
● attach a method to type person with
○ the identifier “speak”
○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/
